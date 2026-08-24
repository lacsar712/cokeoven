package store

import (
	"testing"

	"github.com/lacsar712/cokeoven/internal/model"
)

func TestCase(t *testing.T) {
	orig := BatterySnapshot{
		Tower: model.TowerID("tower-a1"),
		Segments: []SegmentSnapshot{
			{Zone: model.ZoneID("tower-a1-zone-00"), TempC: 42.0},
		},
	}
	clone := CloneBatterySnapshot(orig)
	clone.Segments[0].TempC = 99.0
	if orig.Segments[0].TempC == 99.0 {
		t.Fatal("clone mutated original batch Segments backing array")
	}
}
