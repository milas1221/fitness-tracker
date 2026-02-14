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

	// Не используем TrimSpace – пробелы должны вызывать ошибку
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("шаги должны быть положительным числом")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительным значением")
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