package spentenergy

import (
	"errors"
	"time"
)


const (
	stepLengthCoefficient       = 0.45 // коэффициент длины шага
	mInKm                       = 1000.0
	minInH                      = 60.0
	walkingCaloriesCoefficient  = 0.5  // корректирующий коэффициент для ходьбы
)


func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceM := float64(steps) * stepLength
	return distanceM / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return distance / hours
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры: steps, weight, height и duration должны быть положительными")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed == 0 {
		return 0, errors.New("не удалось вычислить среднюю скорость")
	}
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные параметры: steps, weight, height и duration должны быть положительными")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	if meanSpeed == 0 {
		return 0, errors.New("не удалось вычислить среднюю скорость")
	}
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}