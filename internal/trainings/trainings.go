package trainings

import (
	"fitness-tracker-fifth-sprint/internal/personaldata"
	"fitness-tracker-fifth-sprint/internal/spentenergy"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataInput := strings.Split(datastring, ",")
	// Так как формат вводимой строки "6000,Ходьба,1h00m" (разделённый двумя запятыми), то проверяем длину получившегося слайса
	// Возвращаем возможную ошибку в виде сообщения "недопустимый формат данных"
	if len(dataInput) != 3 {
		return fmt.Errorf("недопустимый формат данных")
	}
	// Находим количество шагов steps путём преобразования строки в число. Количество шагов - первый элемент dataInput
	// Также не забываем обработать ошибку преобразования
	t.Steps, err = strconv.Atoi(dataInput[0])
	if err != nil {
		return err
	}
	// Значение шагов не может быть отрицательным или равным нулю
	if t.Steps <= 0 {
		return fmt.Errorf("нулевые шаги")
	}
	// Вид активности - переменная типа string. Элемент слайса dataInput (второй) по типу совпадает со string. Присваиваем значение переменной activType
	t.TrainingType = dataInput[1]
	// Находим продолжительность прогулки walkDur (третий элемент dataInput). Парсим строку в переменную time.Duration, обрабатывая возможную ошибку
	t.Duration, err = time.ParseDuration(dataInput[2])
	if err != nil {
		return err
	}
	// Проверяем, чтобы продолжительность была положительна и отлична от нуля
	if t.Duration <= 0 {
		return fmt.Errorf("нулевая продолжительность")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Получив входные данные для расчета пройденного пути Distance и средней скорости MeanSpeed,
	// вызываем вспомогательные функции и присваиваем названным переменным значения функций
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	// конвертируем полученную переменную walkDur типа time.Duration в количество часов типа float64
	convWalkDur := t.Duration.Hours()
	// Необходимо вывести результат в зависимости от вида выполняемой физической активности activType
	switch t.TrainingType {
	// В случае бега рассчитываем сожжённые калории с помощью функции RunningSpentCalories() и присваиваем значение переменной runningSpentCalories
	case "Бег":
		runningSpentCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		// Выводим лог ошибки в случае возникновения
		if err != nil {
			log.Println(err)
		}
		// Итоговый результат присваиваем переменной res в виде форматированного вывода и возвращаем её значение из функции и значение отсутствия ошибки nil
		res := fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", convWalkDur, distance, meanSpeed, runningSpentCalories)
		return res, nil
	// В случае ходьбы рассчитываем сожжённые калории с помощью функции WalkingSpentCalories() и присваиваем значение переменной walkingSpentCalories
	case "Ходьба":
		walkingSpentCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		// Выводим лог ошибки в случае возникновения
		if err != nil {
			log.Println(err)
		}
		// Итоговый результат присваиваем переменной res в виде форматированного вывода и возвращаем её значение из функции и значение отсутствия ошибки nil
		res := fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", convWalkDur, distance, meanSpeed, walkingSpentCalories)
		return res, nil
	default:
		// Если при вводе было введено ни "Бег", ни "Ходьба", выводим ошибку "неизвестный тип тренировки"
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
