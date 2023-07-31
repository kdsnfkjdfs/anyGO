package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	xys, err := readData("Data.txt")

	if err != nil {
		log.Fatalf("Could Not Data(NIGA): %v", err)
	}

	for _, xy := range xys {
		fmt.Println(xy.x, xy.y)

	}

}

type xy struct{ x, y float64 }

func readData(path string) ([]xy, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys []xy
	s := bufio.NewScanner(f)

	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscan(s.Text(), "%f,&f", &x, &y)

		if err != nil {
			log.Println("ERROR %q :%v", s.Text(), err)
		}

		xys = append(xys, xy{x, y})

	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("Hey You  could not Scan: %v", err)
	}

	return xys, nil
}
