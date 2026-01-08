package vnl

import (
	"errors"
)

var ErrEmptyActions = errors.New("actions list empty")
var ErrSmallBeatDuration = errors.New("beatDurationInSeconds <= 0")
var ErrSmallPlatterDuration = errors.New("platterRevolutionDurationInSeconds <= 0")
