package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"Desktop\Go-Lang\GOlang\Dev\final4-sprint-ex\cmd\tracker\spentcalories"
)

const (
	stepLengthCoefficient      = 0.414 // коэффициент длины шага
	mInKm                      = 1000  // метров в километре
	minInH                     = 60    // минут в часе
	walkingCaloriesCoefficient = 0.029 // коэффициент для ходьбы
)
// постарался сделать корректный вывод ошибок
func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("expected 3 comma-separated values, got %d", len(parts))
	}

	stepsStr := strings.TrimSpace(parts[0])
	if stepsStr == "" {
		return 0, "", 0, errors.New("steps value is empty")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps format: %v", err)
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("steps must be positive")
	}

	activity := strings.TrimSpace(parts[1])
	if activity != "Walking" && activity != "Running" {
		return 0, "", 0, fmt.Errorf("invalid activity type: %s (expected Walking or Running)", activity)
	}

	durationStr := strings.TrimSpace(parts[2])
	if durationStr == "" {
		return 0, "", 0, errors.New("duration is empty")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration format: %v", err)
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("duration must be positive")
	}

	return steps, activity, duration, nil
}
func distance(steps int, stepLength float64) float64 {
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(distance float64, duration time.Duration) float64 {
	return distance / duration.Hours()
}

func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	dist := distance(steps, runningStepLength)
	speed := meanSpeed(dist, duration)
	durationInMinutes := duration.Minutes()
	return weight * speed * durationInMinutes / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	stepLength := height * stepLengthCoefficient
	dist := distance(steps, stepLength)
	speed := meanSpeed(dist, duration)
	durationInMinutes := duration.Minutes()
	return weight * speed * durationInMinutes / minInH * walkingCaloriesCoefficient, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", fmt.Errorf("parsing error: %v", err)
	}

	var dist, speed, calories float64
	var stepLength float64

	switch activity {
	case "Walking":
		stepLength = height * stepLengthCoefficient
		dist = distance(steps, stepLength)
		speed = meanSpeed(dist, duration)
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	case "Running":
		stepLength = runningStepLength
		dist = distance(steps, stepLength)
		speed = meanSpeed(dist, duration)
		calories, err = RunningSpentCalories(steps, weight, duration)             // не понял как исключить дублирование кода :(
	default:
		return "", fmt.Errorf("unsupported activity type: %s", activity)
	}

	if err != nil {
		return "", fmt.Errorf("calculation error: %v", err)
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity, duration.Hours(), dist, speed, calories)
	
	return result, nil
}