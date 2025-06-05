package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"Desktop\Go-Lang\GOlang\Dev\final4-sprint-ex\cmd\tracker\daysteps"
)

const (
	stepLength = 0.65 // длина шага в метрах
	mInKm      = 1000 // метров в километре
)

// постарался переработать корректность ошибок и их вывод
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("input must contain exactly one comma")
	}

	stepsStr := strings.TrimSpace(parts[0])
	if stepsStr == "" {
		return 0, 0, errors.New("steps value is empty")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format: %v", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("steps must be positive")
	}

	durationStr := strings.TrimSpace(parts[1])
	if durationStr == "" {
		return 0, 0, errors.New("duration is empty")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration format: %v", err)
	}
	if duration <= 0 {
		return 0, 0, errors.New("duration must be positive")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) (string, error) {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return "", fmt.Errorf("invalid input '%s': %v", data, err)
	}

	if weight <= 0 {
		return "", errors.New("weight must be positive")
	}
	if height <= 0 {
		return "", errors.New("height must be positive")
	}

	distance := float64(steps) * walkingStepLength / mInKm
	calories := weight * distance * 0.029 // Simplified walking calories calculation

	result := fmt.Sprintf(
		"Steps: %d\nDistance: %.2f km\nCalories burned: %.2f",
		steps, distance, calories,
	)
	return result, nil
}
