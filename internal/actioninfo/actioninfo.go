// Пакет реализует вывод общей информации обо всех тренировках и прогулках
package actioninfo

import (
	"fmt"
	"log"
)

// Создаём интерфейс DataParser с методами Parse(string) и ActionInfo().
// Эти методы реализованы для типов Training и DaySteps, а значит реализуют интерфейс
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Проверяем, реализуют ли интерфейс методы Parse(string) и ActionInfo() для типов Training и DaySteps
func Info(dataset []string, dp DataParser) {
	// Перебираем все значения слайса dataset в цикле.
	for _, v := range dataset {
		// Распарсиваем каждое значение с помощью метода Parse()
		err := dp.Parse(v)
		// Обрабатываем ошибку парсинга
		if err != nil {
			// Логируем ошибку и переходим к следующей итерации цикла.
			log.Println(err)
			continue
		}
		// Сформировываем строку с информацией об активности с помощью метода ActionInfo()
		info, err := dp.ActionInfo()
		if err != nil {
			// Логируем ошибку при её возникновении
			log.Println(err)
		}
		// Выводим строку с информацией об активности
		fmt.Println(info)
	}
}
