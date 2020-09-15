package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	// var msg string = "Hello"
// 	// msg += " World"

// 	// fmt.Println(msg)

// 	// msg = " Hola Xavi"
// 	// fmt.Println(msg)

// 	// var tipo1 = "casa"
// 	// var tipo2 = 5

// 	// fmt.Println(tipo1)
// 	// fmt.Println(tipo2)
// 	// fmt.Println(tipo1 + msg)

// 	// var altura, peso int = 100, 50

// 	// fmt.Println("Altura", altura, " Peso", peso)

// 	// var edad int
// 	// fmt.Println("La edad es ", edad)

// 	// var nombre string
// 	// fmt.Println("El nombre es" + nombre)

// 	// var peso1, peso2 = 50, 65
// 	// fmt.Println("El peso 1", peso1, "El peso 2", peso2)

// 	// var (
// 	// 	nombre1   = "Javier"
// 	// 	apellido1 = "Rodriguez"
// 	// 	edad1     int
// 	// )
// 	// fmt.Println("Nombre"+nombre1, "Apellido"+apellido1, "edad", edad1)

// 	// //var count = 10
// 	// count := 10

// 	// fmt.Println(count)

// 	// const count1 = 20

// 	// fmt.Println(count1)

// 	// color, altura := "rojo", 20

// 	// fmt.Println("color "+color, "altura", altura)

// 	// c1 := complex(4, 3)
// 	// c2 := 5 + 43i
// 	// cadd := c1 + c2
// 	// fmt.Println("sum: ", cadd)
// 	// cmul := c1 * c2
// 	// fmt.Println("product: ", cmul)

// 	// if age := 10; age <= 4 {
// 	// 	fmt.Println("Junior")
// 	// } else if age > 5 && age < 8 {
// 	// 	fmt.Println("Senior")
// 	// } else {
// 	// 	fmt.Println("Cross")
// 	// }

// 	// finger := 6
// 	// fmt.Printf("Finger %d is ", finger)
// 	// switch finger {
// 	// case 3:
// 	// 	fmt.Println("Middle")
// 	// case 4:
// 	// 	fmt.Println("Ring")
// 	// default:
// 	// 	fmt.Println("Pinky")
// 	// }

// 	// finger := 3
// 	// fmt.Printf("Finger %d is ", finger)
// 	// switch finger {
// 	// case 3:
// 	// 	fmt.Println("Middle")
// 	// 	fallthrough
// 	// case 4:
// 	// 	fmt.Println("Ring")
// 	// default:
// 	// 	fmt.Println("Pinky")
// 	// }

// 	// a := 3.1
// 	// fmt.Println(suma(int(a), 5))

// 	// //area1, perimeter := rectPer(10.8, 5.6)
// 	// //fmt.Printf("Area %f Perimetro %f", area1, perimeter)

// 	// fmt.Println(rectangle(4, 8))

// 	// area, _ := rectPer(10.8, 5.6)
// 	// fmt.Printf("Area %f", area)

// 	//variadicEx(5, "red", "blue", "green", "yellow")

// 	// type Employee struct {
// 	// 	firstName, lastName string
// 	// 	age                 int
// 	// }

// 	// emp1 := Employee{
// 	// 	firstName: "Geeks",
// 	// 	age:       25,
// 	// 	lastName:  "Hubs",
// 	// }
// 	// fmt.Println("Employee 1", emp1)

// 	// emp2 := Employee{"Geeks", "Hubs", 25}
// 	// fmt.Println("Employee 2: ", emp2)

// 	// emp3 := struct {
// 	// 	firstName, lastName string
// 	// 	age                 int
// 	// }{
// 	// 	firstName: "Geeks", lastName: "Hubs",
// 	// 	age: 31}

// 	// fmt.Println("Employee 3", emp3)

// 	// var emp4 Employee

// 	// fmt.Println("First Name:", emp4.firstName)
// 	// fmt.Println("Last Name:", emp4.lastName)
// 	// fmt.Println("Age:", emp4.age)

// 	// type Vehiculo struct {
// 	// 	string
// 	// 	int
// 	// }

// 	// p1 := Vehiculo{
// 	// 	string: "BMW",
// 	// 	int:    1999,
// 	// }
// 	// fmt.Println(p1.string)
// 	// fmt.Println(p1.int)

// 	// type Vehiculo struct {
// 	// 	string
// 	// 	int
// 	// }

// 	// type Person struct {
// 	// 	name     string
// 	// 	age      int
// 	// 	vehiculo Vehiculo

// 	// }

// 	// p := Person{
// 	// 	name: "vicboma1",
// 	// 	age:  24,
// 	// 	vehiculo: Vehiculo{
// 	// 		string: "BMW",
// 	// 		int:    1999,
// 	// 	},
// 	// }

// 	// fmt.Println(p)

// 	// b := 0xFF
// 	// var a *int = &b
// 	// fmt.Printf("Type of a is %T\n", a)
// 	// fmt.Println("Address b:", a)
// 	// fmt.Println("Value b:", b)

// 	// a := 34
// 	// var b *int
// 	// if b == nil {
// 	// 	fmt.Println("b: ", b)
// 	// 	b = &a
// 	// 	fmt.Println("b after init:", b)
// 	// }

// 	// a := 10

// 	// fmt.Println(a)
// 	// size := new(int)
// 	// fmt.Printf("Size value:  %d, typy: %T, address:  %v\n", *size, size, size)
// 	// *size = 50
// 	// fmt.Println("New size value:", *size)
// 	// *size = 100
// 	// fmt.Println("New size value:", *size)

// 	// b := 255
// 	// a := &b
// 	// fmt.Println("Address b: ", a)
// 	// fmt.Println("Value b:", *a)

// 	// b := 5
// 	// a := &b
// 	// fmt.Println("Address b: ", a, " Value b:", *a)
// 	// *a = 5000
// 	// fmt.Println("Address b: ", a, " New Value b:", *a)

// 	// a := 50
// 	// fmt.Println("Before function call:", a)
// 	// b := &a
// 	// change(b)
// 	// fmt.Println("After function call:", *b)

// 	// str := "GeeksHubs"
// 	// d := hello(&str)
// 	// fmt.Println("Value d: ", *d)

// 	// go hello()
// 	// time.Sleep(1 * time.Second)
// 	// fmt.Println("Proceso main")

// 	// go numeros()
// 	// go letras()
// 	// time.Sleep(3 * time.Second)
// 	// fmt.Println("main function")

// 	// var a chan int
// 	// if a == nil {
// 	// 	fmt.Println("Channel ‘a’ is null, creating...")
// 	// 	a = make(chan int)
// 	// 	fmt.Printf("Type of a is %T\n", a)
// 	// }

// 	//deadlock()

// 	ch := make(chan int)
// 	go producer(ch)
// 	for {
// 		v, ok := <-ch
// 		if ok == false {
// 			fmt.Println("Salta el break")
// 			break
// 		}
// 		fmt.Println("Received ", v, ok)
// 	}

// }

// // func change(val *int) {
// // 	(*val)--
// // }

// // func hello(val *string) *string {
// // 	i := "Hello " + (*val)
// // 	return &i
// // }

// // func hello() {
// // 	fmt.Println("Hola mundo")
// // }

// func numeros() {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Printf("%d ", i)
// 	}
// }

// func letras() {
// 	for i := 'a'; i <= 'e'; i++ {
// 		time.Sleep(50 * time.Millisecond)
// 		fmt.Printf("%c ", i)
// 	}
// }

// func hello(done chan bool) {
// 	fmt.Println("Hello world channel-goroutine")
// 	done <- false
// }

// func channel() {
// 	done := make(chan bool)
// 	go hello(done)
// 	res := <-done
// 	fmt.Println("channel hello function", res)
// }

// func deadlock() {
// 	ch := make(chan int)
// 	ch <- 5
// }

// func producer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }
