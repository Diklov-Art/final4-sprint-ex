package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	stepLength = 0.65 // длина шага в метрах
	mInKm      = 1000 // метров в километре
)

// ParsePackage принимаем данные о кол-во шагов + время прогулки
func parsePackage(data string) (int, time.Duration, error) {
	// Проверяем наличие пробелов в данных
	if strings.Contains(data, " ") {
		return 0, 0, fmt.Errorf("incorrect data format")
	}

	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("incorrect data format")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("incorrect format of the number of steps")
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("the number of steps must be positive")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("incorrect duration format")
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("duration must be greater than zero")
	}

	return steps, duration, nil
}

// DayActionInfo парсим строку ---> дистанция в км + кол-во кал.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Ошибка обработки данных: %v (ввод: %s)", err, data)
		return "" // возвращаем пустую строку как ожидают тесты
	}

	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("Ошибка расчета калорий: %v (ввод: %s)", err, data)
		return "" // возвращаем пустую строку
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distance, calories)

	return result
}
