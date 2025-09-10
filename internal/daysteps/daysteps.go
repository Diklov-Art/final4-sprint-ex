package daysteps

import (
	"fmt"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := spentcalories.ParsePackage(data)
	if err != nil {
		return fmt.Sprintf("Ошибка: %v", err)
	}

	distance := float64(steps) * spentcalories.StepLength / spentcalories.MInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return fmt.Sprintf("Ошибка расчета калорий: %v", err)
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distance, calories)

	return result
}
