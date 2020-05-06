package pointgenerate

import "github.com/go-zs/calculation/check"

const (
	defaultMaxNumber = 13
)

type (
	Option struct {
		maxNumber int
	}
	GenOption func(o *Option)
)

func SetMaxNumber(num int) GenOption {
	return func(o *Option) {
		o.maxNumber = num
	}
}

func Generate24(f func(arr []int), options ...GenOption) {
	option := Option{maxNumber: defaultMaxNumber}
	for _, o := range options {
		o(&option)
	}
	for i := 1; i <= option.maxNumber; i++ {
		for j := 1; j <= i; j++ {
			for m := 1; m <= j; m++ {
				for n := 1; n <= m; n++ {
					if check.Check24([]int{i, j, m, n}) {
						f([]int{i, j, m, n})
					}
				}
			}
		}
	}
}
