package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Так как является экспортируемой, необходимо проверить корректность аргументов функции:
	// Количество шагов steps не должно быть нулевым или отрицательным
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов отрицательным быть не может")
	}
	// Вес weight и рост heght не должны быть нулевыми или отрицательными. Если хотя бы одна из величин отрицательна, это уже ошибка
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("рост и вес не могут быть отрицательными или нулевыми")
	}
	// Продолжительность активности не должна быть нулевой или отрицательной
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность бега - величина не отрицательная")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := float64(duration) / float64(time.Minute)
	// Итоговое значение функции выводим по формуле с поправкой на коэффициент для расчета калорий при ходьбе walkingCaloriesCoefficient
	// Также возвращаем значение отсутствия ошибки
	return ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Необходимость проверки входных аргументов функции та же, что и у WalkingSpentCalories()
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов отрицательным быть не может")
	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("рост и вес не могут быть отрицательными или нулевыми")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность бега - величина не отрицательная")
	}
	// Для расчета калорий нам нужно значение средней скорости meanSpeed
	meanSpeed := MeanSpeed(steps, height, duration)
	// Продолжительность активности необходимо перевести в минуты. Присваиваем переменной durationInMinutes
	durationInMinutes := float64(duration) / float64(time.Minute)
	// Итоговое значение функции выводим по формуле с использованием константы minInH. Также возвращаем значение отсутствия ошибки
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Здесь обработаем ошибку ввода отрицательных шагов и продолжительности активности
	if steps <= 0 || duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	// Величина скорости нам необходима выраженная в км/ч. Для этого воспользуемся методом hour из пакета time
	return dist / (float64(duration) / float64(time.Hour))
}

func Distance(steps int, height float64) float64 {
	// Так как длина шага и их количество выражены в метрах, используем константу mInKm для перевода значения в километры
	return (float64(steps) * stepLengthCoefficient * height) / mInKm
}
