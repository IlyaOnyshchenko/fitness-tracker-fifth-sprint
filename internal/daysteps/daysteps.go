// Пакет daysteps. В нём представлено описание обычного вида активности с последующим её расчетом и выводом
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/IlyaOnyshchenko/fitness_tracker_fifth_sprint/internal/personaldata"
	"github.com/IlyaOnyshchenko/fitness_tracker_fifth_sprint/internal/spentenergy"
)

// Создаём структуру для хранения данных о пользователе, количестве пройденных шагов и времени активности
type DaySteps struct {
	Steps                 int           // Количество шагов
	Duration              time.Duration // Время активности
	personaldata.Personal               // Встраиваем структуру Personal, содержащую информацию о пользователе из пакета personaldata
}

// Метод Parse необходим для преобразования входных данных формата "6000,1h00m" в данные о количестве шагов и времени активности
func (ds *DaySteps) Parse(datastring string) (err error) {
	// Используем функцию Split из пакета strings для получения двух строк с информацией с разделителем ","
	dataInput := strings.Split(datastring, ",")
	// Проверяем длину получившегося слайса, обрабатываем возможную ошибку
	if len(dataInput) != 2 {
		return fmt.Errorf("недопустимый формат данных")
	}
	// Получаем количество шагов, преобразуя строку в число и также обрабатывая возможную ошибку преобразования
	ds.Steps, err = strconv.Atoi(dataInput[0])
	if err != nil {
		return fmt.Errorf("ошибочный формат данных количества шагов: %w", err)
	}
	// Значение шагов не может быть отрицательным или равным нулю
	if ds.Steps <= 0 {
		return fmt.Errorf("нулевые шаги")
	}
	// Находим продолжительность прогулки Duration структуры DaySteps(ds). Парсим строку в переменную time.Duration, обрабатывая возможную ошибку
	ds.Duration, err = time.ParseDuration(dataInput[1])
	if err != nil {
		return fmt.Errorf("ошибка преобразования времени в часы: %w", err)
	}
	// Проверяем, чтобы продолжительность была положительна и отлична от нуля
	if ds.Duration <= 0 {
		return fmt.Errorf("нулевая продолжительность")
	}
	// В случае успешных преобразований данных, возвращаем значение ошибки, равное nil
	return nil
}

// Метод ActionInfo необходим для вывода информации об активности (шаги, расстояние, сожжённые калории)
func (ds DaySteps) ActionInfo() (string, error) {
	// Считаем количество пройденного пути dist в км, передавая в аргументы метода Distance из пакета spentenergy шаги и рост пользователя,
	// взятый из структуры Personal
	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	// Получаем значение количества сожженных калорий и присваеиваем её переменной calories через функцию WalkingSpentCalories из пакета spentenergy
	// Также обрабатываем возможные ошибки
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка входных данных: %w", err)
	}
	// Итоговый результат присваиваем переменной res в виде форматированного вывода и возвращаем её значение из функции
	res := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories)
	return res, nil
}
