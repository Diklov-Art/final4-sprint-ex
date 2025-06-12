package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"Desktop\final4-sprint-ex\internal\daysteps"
)

const (
	stepLength = 0.65 // длина шага в метрах
	mInKm      = 1000 // метров в километре
)

// принимаем данные о кол-во шагов + время прогулки
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("incorrect data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("incorrect format of the number of steps")
	}

	if steps <= 0 {
		return 0, 0, errors.New("the number of steps must be positive")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, errors.New("incorrect duration format")
	}

	return steps, duration, nil // возвращаем данные или ошибку
}

// парсим строку ---> дистанция в км + кол-во кал.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return "", fmt.Errorf("failed to parse data: %v", err)
	}

	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return "", fmt.Errorf("failed to calculate calories: %v", err)
	}

	result :=  fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distance, calories)

	return result, nil	
}