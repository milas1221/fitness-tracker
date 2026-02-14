package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)


type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}


func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат строки: должно быть два поля")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %w", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %w", err)
	}
	ds.Duration = duration

	return nil
}


func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps, ds.Weight, ds.Height, ds.Duration,
	)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
