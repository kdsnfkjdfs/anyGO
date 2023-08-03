package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func main() {

	data := make([]float64, 36)
	fmt.Println(data)

	for i := range data {
		data[i] = rand.NormFloat64()
	}

	a := mat.NewDense(6, 6, data)

	fmt.Println(mat.Trace(a)) // 대각요소들의 합

	var c mat.Dense // 값이 0인 새로운 행렬을 구성합니다.
	c.Mul(a, a)     // a * a(행렬)
	fmt.Println(c)

}
