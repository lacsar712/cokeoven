package app

import (
	"context"
	"time"

	"github.com/lacsar712/cokeoven/internal/clock"
)

type CokeRamp struct {
	clk   clock.Clock
	tick  time.Duration
	steps int
}

func NewCokeRamp(clk clock.Clock, tick time.Duration, steps int) *CokeRamp {
	if steps <= 0 {
		steps = 40
	}
	return &CokeRamp{clk: clk, tick: tick, steps: steps}
}

func (r *CokeRamp) Ramp(ctx context.Context, target float64, apply func(float64)) error {
	step := target / float64(r.steps)
	if step <= 0 {
		step = 0.5
	}
	cur := 0.0
	for cur < target {
		cur += step
		if cur > target {
			cur = target
		}
		apply(cur)
		if pc, ok := r.clk.(*clock.ProcessClock); ok {
			pc.Step()
		}
		time.Sleep(2 * time.Millisecond)
	}
	return nil
}

func (a *App) RunCokeRamp(ctx context.Context, target float64) error {
	return a.dryRamp.Ramp(ctx, target, func(v float64) { _ = v })
}
