// Пакет personaldata реализует хранение и вывод данных о пользователе
package personaldata

import "fmt"

// В структуре Personal хранятся данные об имени, росте и весе пользователя
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Функция для вывода имени, веса и роста пользователя в соответствии с ТЗ
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f кг.\n", p.Weight)
	fmt.Printf("Рост: %.2f м.\n", p.Height)
}
