package fsm

import (
	"context"
	"fmt"

	"github.com/lacsar712/cokeoven/internal/model"
)

var ErrIllegalDryTransition = fmt.Errorf("illegal dry transition")

type CokingFSM struct {
	id    model.TowerID
	state model.DryState
	hooks *DryHookChain
}

func NewCokingFSM(id model.TowerID, effect func(context.Context, model.TowerID, model.DryState, model.DryState) error) *CokingFSM {
	_ = effect
	return &CokingFSM{id: id, state: model.DryIdle, hooks: NewDryHookChain()}
}

func (f *CokingFSM) Hooks() *DryHookChain { return f.hooks }

func (f *CokingFSM) State() model.DryState { return f.state }

func (f *CokingFSM) Dispatch(ctx context.Context, event string) (model.DryState, error) {
	next, ok := allowedDry(f.state, event)
	if !ok {
		if f.hooks != nil {
			_ = f.hooks.RunAfter(ctx, f.state, f.state, event)
		}
		return f.state, fmt.Errorf("%s from %s: %w", event, f.state, ErrIllegalDryTransition)
	}
	from := f.state
	if f.hooks != nil {
		if err := f.hooks.RunBefore(ctx, from, next, event); err != nil {
			return f.state, err
		}
	}
	f.state = next
	if f.hooks != nil {
		if err := f.hooks.RunAfter(ctx, from, next, event); err != nil {
			return f.state, err
		}
	}
	return f.state, nil
}

func allowedDry(from model.DryState, event string) (model.DryState, bool) {
	switch from {
	case model.DryIdle:
		if event == "arm_heat" {
			return model.DryHeating, true
		}
	case model.DryHeating:
		if event == "hold" {
			return model.DryHold, true
		}
	case model.DryHold:
		if event == "cool" {
			return model.DryCool, true
		}
	case model.DryCool:
		if event == "done" {
			return model.DryIdle, true
		}
	}
	return from, false
}
