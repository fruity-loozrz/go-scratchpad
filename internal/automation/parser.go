package automation

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	fractionalRegex  = regexp.MustCompile(`^([+-]?\d+)/(\d+)`)
	floatRegex       = regexp.MustCompile(`^([+-]?\d+\.\d+)$`)
	intRegex         = regexp.MustCompile(`^([+-]?\d+)$`)
	singlePlusRegex  = regexp.MustCompile(`^\+$`)
	singleMinusRegex = regexp.MustCompile(`^\-$`)
)

type actionType int

const (
	actionTypeNone actionType = iota
	actionTypeMove
	actionTypeBpm
	actionTypeInterpolation
)

type interpolationType string

const (
	interpolationTypeCubic  interpolationType = "cubic"
	interpolationTypeLinear interpolationType = "linear"
)

const (
	bpmToken                 = "bpm"
	defaultBpm               = 140.0
	equalToken               = "="
	interpolateToken         = "interpolate"
	defaultInterpolationType = interpolationTypeCubic
)

type action struct {
	actionType        actionType
	bpm               float64
	move              *Move
	interpolationType interpolationType
}

func parseReal(field string) (float64, bool) {
	if matches := fractionalRegex.FindStringSubmatch(field); len(matches) == 3 {
		num, _ := strconv.Atoi(matches[1])
		den, _ := strconv.Atoi(matches[2])
		return float64(num) / float64(den), true
	}

	if matches := intRegex.FindStringSubmatch(field); len(matches) == 2 {
		num, _ := strconv.Atoi(matches[1])
		return float64(num), true
	}

	if matches := floatRegex.FindStringSubmatch(field); len(matches) == 2 {
		num, _ := strconv.ParseFloat(matches[1], 64)
		return num, true
	}

	if matches := singlePlusRegex.FindStringSubmatch(field); len(matches) == 1 {
		return 1.0, true
	}

	if matches := singleMinusRegex.FindStringSubmatch(field); len(matches) == 1 {
		return -1.0, true
	}

	return 0.0, false
}

func parseMove(fields []string) (*Move, bool) {
	if len(fields) == 0 {
		return nil, false
	}
	dh, ok := parseReal(fields[0])
	if !ok {
		return nil, false
	}

	if len(fields) == 1 {
		return &Move{Dh: dh, Dt: 1.0}, true
	}

	// Handle equal sign: Dt = Abs(Dh)
	if fields[1] == equalToken {
		dt := dh
		if dt < 0 {
			dt = -dt
		}
		return &Move{Dh: dh, Dt: dt}, true
	}

	dt, ok := parseReal(fields[1])
	if !ok {
		return nil, false
	}
	return &Move{Dt: dt, Dh: dh}, true
}

func parseBpm(fields []string) (float64, bool) {
	if len(fields) < 2 {
		return 0.0, false
	}
	if fields[0] != bpmToken {
		return 0.0, false
	}
	return parseReal(fields[1])
}

func parseInterpolation(fields []string) (interpolationType, bool) {
	if len(fields) < 2 {
		return "", false
	}
	if fields[0] != interpolateToken {
		return "", false
	}
	return interpolationType(fields[1]), true
}

func isValidInterpolationType(interpolationType interpolationType) bool {
	return interpolationType == interpolationTypeCubic || interpolationType == interpolationTypeLinear
}

func parseLine(line string) (*action, error) {
	if idx := strings.Index(line, "#"); idx != -1 {
		line = line[:idx]
	}

	fields := strings.Fields(line)

	if len(fields) == 0 {
		return &action{actionType: actionTypeNone}, nil
	}

	if move, ok := parseMove(fields); ok {
		return &action{
			actionType: actionTypeMove,
			move:       move,
		}, nil
	}

	if bpm, ok := parseBpm(fields); ok {
		return &action{
			actionType: actionTypeBpm,
			bpm:        bpm,
		}, nil
	}

	if interpolationType, ok := parseInterpolation(fields); ok {
		return &action{
			actionType:        actionTypeInterpolation,
			interpolationType: interpolationType,
		}, nil
	}

	return nil, fmt.Errorf("could not parse line %q", line)
}

func Parse(input string) (*Program, error) {
	program := &Program{
		Bpm:   defaultBpm,
		Moves: []Move{},
	}
	program.SetInterpolationType(defaultInterpolationType)

	bpmIsSet := false
	interpolationIsSet := false

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		action, err := parseLine(line)
		if err != nil {
			return nil, err
		}

		switch action.actionType {
		case actionTypeMove:
			{
				program.Moves = append(program.Moves, *action.move)
			}
		case actionTypeBpm:
			{
				if bpmIsSet {
					return nil, fmt.Errorf("duplicate bpm set: %f", action.bpm)
				}
				program.Bpm = action.bpm
				bpmIsSet = true
			}
		case actionTypeInterpolation:
			{
				if interpolationIsSet {
					return nil, fmt.Errorf("duplicate interpolation set: %v", action.interpolationType)
				}
				if !isValidInterpolationType(action.interpolationType) {
					return nil, fmt.Errorf("invalid interpolation type: %v", action.interpolationType)
				}
				program.SetInterpolationType(action.interpolationType)
				interpolationIsSet = true
			}
		case actionTypeNone:
			{
				continue
			}
		default:
			{
				return nil, fmt.Errorf("unknown action type: %v", action.actionType)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return program, nil
}
