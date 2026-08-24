package app

import (
	"context"

	"github.com/lacsar712/cokeoven/internal/clock"
)

// SegmentPlan describes inter-segment vent staging for drying batches.

type SegmentPlan struct {
	VentSteps int
}

func (a *App) ExecutePlan(ctx context.Context, plan SegmentPlan) error {
	if a.scheduler == nil {
		return nil
	}
	// Propagate ctx so a window revocation (cancel) issued at the screen layer
	// reaches the plan coordination layer; otherwise InstallVentPlanCtx receives
	// an uncancellable context and keeps appending heating steps after the
	// carbonization window has been withdrawn.
	return a.scheduler.InstallVentPlanCtx(ctx, clock.VentPlan{VentSteps: plan.VentSteps}, "segment-plan")
}

func (a *App) SegmentVentStepsDone() int {
	if a.scheduler == nil {
		return 0
	}
	return a.scheduler.VentStepsDone()
}
