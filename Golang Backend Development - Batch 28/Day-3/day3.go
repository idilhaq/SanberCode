package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hari 3 - Operator dan Conditional")

	// 1. Operator Aritmatika
	// operator penjumlahan
	jumlah := 8 + 3
	fmt.Println(jumlah) // hasilnya 13

	// operator pengurangan
	kurang := 8 - 3
	fmt.Println(kurang) // hasilnya 5

	// operator perkalian
	kali := 8 * 3
	fmt.Println(kali) // hasilnya 24

	// operator pembagian
	bagi := 8 / 4
	fmt.Println(bagi) // hasilnya 2

	// operator modulus
	modulus := 8 % 3
	fmt.Println(modulus) // hasilnya 2

	// augmented
	var angka = 8
	fmt.Println(angka) // 8
	angka += 10
	fmt.Println(angka) // 18

	var angka2 = 5
	fmt.Println(angka2) // 5
	angka2 += 5
	fmt.Println(angka2) // 10

	// unary
	angka3 := 8
	fmt.Println(angka3) // 8
	angka3++
	fmt.Println(angka3) // 11

	angka4 := 5
	fmt.Println(angka4) // 5
	angka4--
	fmt.Println(angka4) // 4

	// 2. Operator Perbandingan
	var angka5 = 8

	fmt.Println(angka5 > 5) // true

	fmt.Println(angka5 < 5) // false

	fmt.Println(angka5 >= 5) // true

	fmt.Println(angka5 <= 5) // false

	fmt.Println(angka5 == 5) // false

	fmt.Println(angka5 != 5) // true

	// 3. Operator Logika
	var a = true
	var b = false
	var c = true
	var d = false

	fmt.Println(a && c) // true

	fmt.Println(a && b) // false

	fmt.Println(a || b) // true

	fmt.Println(b || d) // false

	fmt.Println(!b && !d) // true

	fmt.Println(!a || b) // false
}
