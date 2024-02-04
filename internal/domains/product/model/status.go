package model

import (
	"fmt"
)

type Status uint8

const (
	StatusDisabled Status = iota
	StatusEnabled
	StatusEnabledPartial
	StatusDeleted
)

func (r Status) String() string {
	switch r {
	case StatusDisabled:
		return "disabled"

	case StatusEnabled:
		return "enabled"

	case StatusEnabledPartial:
		return "enabled_partial"

	case StatusDeleted:
		return "deleted"

	default:
		return fmt.Sprintf("%d", r)
	}
}
