package main

import "fmt"

type shikaku struct {
	tate int
	yoko int
}

type sum struct{}

func (x sum) keisan(tate1 int, yoko1 int, tate2 int, yoko2 int) {
	menseki1 := tate1 * yoko1
	menseki2 := tate2 * yoko2
	fmt.Println(menseki1 + menseki2)
}

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "fizzbuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		}
	}

	chouhoukei1 := new(shikaku)
	chouhoukei1.tate = 5
	chouhoukei1.yoko = 5
	chouhoukei2 := new(shikaku)
	chouhoukei2.tate = 10
	chouhoukei2.yoko = 10
	kekka := new(sum)
	kekka.keisan(chouhoukei1.tate, chouhoukei1.yoko, chouhoukei2.tate, chouhoukei2.yoko)
}
