package clock

import (
	"time"

	"github.com/lacsar712/cokeoven/internal/model"
)

type CarbonWindow struct {
	clk      Clock
	duration time.Duration
}

func NewCarbonWindow(clk Clock, duration time.Duration) *CarbonWindow {
	if duration <= 0 {
		duration = 2 * time.Minute
	}
	return &CarbonWindow{clk: clk, duration: duration}
}

func (w *CarbonWindow) Active(anchor time.Time) bool {
	// Carbonization window cadence must follow the process clock, not the wall
	// clock: a drill standstill that freezes the oven-group beat must also freeze
	// the closure window, otherwise the closure percentage creeps up on the
	// duty-room wall clock instead of the carbonization process rhythm.
	now := w.clk.Now()
	return !now.Before(anchor) && now.Sub(anchor) < w.duration
}

func (w *CarbonWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrCarbonHold
}
