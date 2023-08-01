package main

import (
	"fmt"
	"math/rand"

	geo "github.com/kellydunn/golang-geo"
	"gonum.org/v1/gonum/mat"
)

func main() {
	p1 := geo.NewPoint(37.4, 126.434)
	p2 := geo.NewPoint(35, 130.32)

	dist := p1.GreatCircleDistance(p2)
	fmt.Println(dist)
	zero := mat.NewDense(2, 2, nil) // zero 변수 삭제

	fmt.Println(zero) // zero

	data := make([]float64, 36)
	fmt.Println(data)

	for i := range data {
		data[i] = rand.NormFloat64()
	}

	a := mat.NewDense(6, 6, data)

	fmt.Println(mat.Trace(a)) // 대각요소들의 합
	zero.Copy(a)

	fmt.Println(zero)
	var c mat.Dense // 값이 0인 새로운 행렬을 구성합니다.
	c.Mul(a, a)     // a * a(행렬)
	fmt.Println(c)
}
