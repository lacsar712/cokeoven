package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lacsar712/cokeoven/internal/config"
	"github.com/lacsar712/cokeoven/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	anchor := a.clk.Now().Add(-3 * time.Minute)
	err = a.ConfirmCarbonHold(context.Background(), anchor)
	if err == nil {
		t.Fatal("expected gradient hold error")
	}
	if !errors.Is(err, model.ErrCarbonHold) {
		t.Fatalf("expected ErrCarbonHold, got %v", err)
	}
	_ = time.Second
}
