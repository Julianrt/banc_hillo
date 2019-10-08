package models

import "errors"

type ValidationError error

var (
	SaldoInsuficiente = ValidationError(errors.New("Saldo insuficiente"))
)