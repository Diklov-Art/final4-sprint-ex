package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/cmd/tracker/spentcalories"
)

const (
	stepLength = 0.65 // длина шага в метрах
	mInKm      = 1000 // метров в километре
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("неверный формат данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("неверный формат количества шагов")
	}

	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, errors.New("неверный формат продолжительности")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distance, calories)
}
