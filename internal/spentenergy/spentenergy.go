package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()
	if hours <= 0 {
		return 0
	}
	return distance / hours
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные данные")
	}

	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	return (weight * speed * minutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные данные")
	}

	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := (weight * speed * minutes) / minInH
	return calories * walkingCaloriesCoefficient, nil
}
