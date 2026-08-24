package app

import (
	"context"
	"fmt"
	"time"

	"github.com/lacsar712/cokeoven/internal/model"
)

// CalibrateProbe allows acceptance tests to inject feed calibration faults.
var CalibrateProbe func(ctx context.Context) error

const feedTempLimitC = 55.0

func (a *App) CalibrateFeed(ctx context.Context, tower model.TowerID, holder string) error {
	if err := a.feedLeases.Require(tower, holder, 30*time.Second); err != nil {
		return err
	}
	// Release the lease on every return path: a probe fault that aborts
	// calibration must not leave the tower position occupied, otherwise the
	// next charging attempt finds a stale lease and reports the slot busy.
	defer a.feedLeases.ReleaseHolder(tower, holder)
	if CalibrateProbe != nil {
		if err := CalibrateProbe(ctx); err != nil {
			return fmt.Errorf("calibrate: %w", err)
		}
	}
	return nil
}
