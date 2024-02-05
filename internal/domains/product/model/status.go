package model

type Status uint8

const (
	StatusDisabled Status = iota
	StatusEnabled
	StatusEnabledPartial
	StatusDeleted
)
