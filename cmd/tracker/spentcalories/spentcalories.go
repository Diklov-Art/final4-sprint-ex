package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	stepLengthCoefficient      = 0.414 // коэффициент длины шага
	mInKm                      = 1000  // метров в километре
	minInH                     = 60    // минут в часе
	walkingCaloriesCoefficient = 0.029 // коэффициент для ходьбы
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("неверный формат данных")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, "", 0, errors.New("неверный формат количества шагов")
	}

	activity := strings.TrimSpace(parts[1])
	if activity != "Ходьба" && activity != "Бег" {
		return 0, "", 0, errors.New("неизвестный тип тренировки")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return 0, "", 0, errors.New("неверный формат продолжительности")
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры")
	}

	speed := meanSpeed(steps, 0, duration) // рост не учитывается для бега
	durationInMinutes := duration.Minutes()
	calories := weight * speed * durationInMinutes / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры")
	}

	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * speed * durationInMinutes / minInH * walkingCaloriesCoefficient
	return calories, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	var dist, speed, calories float64
	var errCal error

	switch activity {
	case "Ходьба":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, errCal = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, errCal = RunningSpentCalories(steps, weight, duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if errCal != nil {
		return "", errCal
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity, duration.Hours(), dist, speed, calories), nil
}
