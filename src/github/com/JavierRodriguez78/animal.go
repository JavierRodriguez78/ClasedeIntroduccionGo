package main

import (
	"fmt"
)

func main() {
	a := animal{"Leon", "Mamifero"}
	a.Name()
	a.Species()

	l := &leon{"amarillo", a}
	l.Name()
	l.Species()
	l.Color()


}

type animal struct {
	name   string
	specie string
}

type leon struct {
	color string
	animal
}

func (a *animal) Name() {
	fmt.Println(a.name)
}

func (a *animal) Species() {
	fmt.Printf("%s belongs to %s species\n", a.name, a.specie)
}

func (l *leon) Color() {
	fmt.Printf("%s is the color: %s\n", l.name, l.color)
}


type Reader interface {
	Read (p []byte)( n int, err error)
}

type Writer interface {
	Write (p []byte)(n int, err, error)
}

func ( l *leon) Write(p []byte (n int, err error)){
	return 0, nil
}

func ( l *leon )Read ( p []byte ( n int, err error)){
	return 0 , nil
}