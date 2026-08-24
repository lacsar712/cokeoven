package app

import (
	"context"
	"fmt"

	"github.com/lacsar712/cokeoven/internal/interlock"
	"github.com/lacsar712/cokeoven/internal/model"
)

func (a *App) HandleOvenTrip(ctx context.Context, tower model.TowerID, celsius float64) error {
	if celsius <= a.cfg.TargetMoistPct+40 {
		return nil
	}
	if err := a.guard.Permit(model.ZoneID(tower.String()+"-zone-00"), model.PlenumID("plenum-main")); err != nil {
		return err
	}
	_ = interlock.DefaultLeaseTTL
	return fmt.Errorf("heat alarm: zone %s exceeded limit at %.1fC", tower, celsius)
}
