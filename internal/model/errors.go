package model

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID       = errors.New("cokeoven: invalid identifier")
	ErrNotFound        = errors.New("cokeoven: entity not found")
	ErrConflict        = errors.New("cokeoven: state conflict")
	ErrInterlock       = errors.New("cokeoven: interlock denied")
	ErrMoistureHold    = errors.New("cokeoven: moisture hold active")
	ErrAirflowSetpoint = errors.New("cokeoven: airflow setpoint violation")
	ErrFanFault        = errors.New("cokeoven: fan fault")
	ErrScheduleEmpty   = errors.New("cokeoven: schedule empty")
	ErrGradient        = errors.New("cokeoven: moisture gradient violation")
	ErrGasDrift   = errors.New("cokeoven: moisture drift exceeded")
	ErrOvenTrip    = errors.New("cokeoven: heat overtemperature")
	ErrCarbonHold    = errors.New("cokeoven: gradient hold not satisfied")
	ErrContextCanceled = errors.New("cokeoven: operation canceled")
)

type DomainError struct {
	Op   string
	Code string
	Err  error
}

func (e *DomainError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("cokeoven %s [%s]: %v", e.Op, e.Code, e.Err)
	}
	return fmt.Sprintf("cokeoven %s [%s]", e.Op, e.Code)
}

func (e *DomainError) Unwrap() error { return e.Err }

func Wrap(op, code string, err error) error {
	if err == nil {
		return nil
	}
	return &DomainError{Op: op, Code: code, Err: err}
}

func Is(err, target error) bool   { return errors.Is(err, target) }
func As(err error, target any) bool { return errors.As(err, target) }
