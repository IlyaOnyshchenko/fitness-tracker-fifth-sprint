package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Println("Имя:", p.Name)
	fmt.Println("Вес:", p.Weight)
	fmt.Println("Рост:", p.Height)
}
