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
	return time.Since(anchor) < w.duration
}

func (w *CarbonWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrCarbonHold
}
