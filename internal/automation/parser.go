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
)

const (
	bpmToken   = "bpm"
	defaultBpm = 140.0
)

type action struct {
	actionType actionType
	bpm        float64
	move       *Move
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

func parseLine(line string) (*action, error) {
	fields := strings.Fields(line)

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

	return nil, fmt.Errorf("could not parse line %q", line)
}

func Parse(input string) (*Program, error) {
	program := &Program{
		Bpm:   defaultBpm,
		Moves: []Move{},
	}
	bpmIsSet := false

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
