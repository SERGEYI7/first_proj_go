package first_proj_go

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Add Функция для складывания чисел и возврата результата [Какаятоссылка]
//
// [Какаятоссылка]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
