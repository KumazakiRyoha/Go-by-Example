package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	src string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		src: "some name",
	}
	fmt.Printf("co={num: %v,ste: %v}\n", co.num, co.src)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
