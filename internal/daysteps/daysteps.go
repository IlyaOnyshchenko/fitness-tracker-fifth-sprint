package daysteps

import (
	"fitness-tracker-fifth-sprint/internal/personaldata"
	"fitness-tracker-fifth-sprint/internal/spentenergy"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataInput := strings.Split(datastring, ",")
	// Так как формат вводимой строки "6000,1h00m" (разделённый одной запятыой), то проверяем длину получившегося слайса
	// Возвращаем возможную ошибку в виде сообщения "недопустимый формат данных"
	if len(dataInput) != 2 {
		return fmt.Errorf("недопустимый формат данных")
	}
	ds.Steps, err = strconv.Atoi(dataInput[0])
	if err != nil {
		return fmt.Errorf("ошибочный формат данных количества шагов: %w", err)
	}
	// Значение шагов не может быть отрицательным или равным нулю
	if ds.Steps <= 0 {
		return fmt.Errorf("нулевые шаги")
	}
	// Находим продолжительность прогулки walkDur (третий элемент dataInput). Парсим строку в переменную time.Duration, обрабатывая возможную ошибку
	ds.Duration, err = time.ParseDuration(dataInput[1])
	if err != nil {
		return fmt.Errorf("ошибка преобразования времени в часы: %w", err)
	}
	// Проверяем, чтобы продолжительность была положительна и отлична от нуля
	if ds.Duration <= 0 {
		return fmt.Errorf("нулевая продолжительность")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Считаем количество пройденного пути dist в км, используя константы stepLength и mInKm
	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	// Получаем значение количества сожженных калорий и присваеиваем её переменной calories через функцию WalkingSpentCalories из пакета spentcalories
	// Также обрабатываем возможные ошибки и выводим их на экран
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка входных данных: %w", err)
	}
	// Итоговый результат присваиваем переменной res в виде форматированного вывода и возвращаем её значение из функции
	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories)
	return res, nil
}
