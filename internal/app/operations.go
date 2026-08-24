package app

import (
	"context"
	"fmt"
	"time"

	"github.com/lacsar712/cokeoven/internal/model"
)

func (a *App) ValidateGasDrift(ctx context.Context, moistPct float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	limit := a.cfg.TargetMoistPct + a.cfg.MaxGradientDeltaPct
	if moistPct <= limit {
		return nil
	}
	return fmt.Errorf("moisture: %v", model.ErrGasDrift)
}

func (a *App) ConfirmCarbonHold(ctx context.Context, anchor time.Time) error {
	if a.avgWindow == nil {
		return model.Wrap("app", "window", model.ErrCarbonHold)
	}
	if err := a.avgWindow.Require(anchor); err != nil {
		return fmt.Errorf("gradient hold: %w", err)
	}
	return nil
}
