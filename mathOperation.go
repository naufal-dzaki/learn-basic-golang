package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var c int = a + b

	fmt.Println("Nilai a:", a)
	fmt.Println("Nilai b:", b)
	fmt.Println("Nilai c:", c)

	// Operasi Aritmatika
	fmt.Println("Penjumlahan a + b =", a+b)
	fmt.Println("Pengurangan c - b =", c-b)
	fmt.Println("Perkalian a * b =", a*b)
	fmt.Println("Pembagian c / a =", c/a)
	fmt.Println("Modulus c % b =", c%b)

	// augmented assignment
	var i = 10

	i += 5
	fmt.Println(i)

	i *= 2
	fmt.Println(i)

	i -= 2
	fmt.Println(i)

	i /= 2
	fmt.Println(i)

	i %= 4
	fmt.Println(i)

	// unary operator
	i++
	fmt.Println(i)

	i--
	fmt.Println(i)
}
