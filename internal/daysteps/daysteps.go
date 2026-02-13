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
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("некорректный формат строки")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil || steps <= 0 {
		return errors.New("некорректное количество шагов")
	}

	duration, err := time.ParseDuration(data[1])
	if err != nil || duration <= 0 {
		return errors.New("некорректная продолжительность")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
