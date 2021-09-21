package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Variable and Constant

	// deklarasi dengan var
	var text = "Hello World"
	fmt.Println(text)

	text = "Hello World diubah"
	fmt.Println(text)

	// deklarasi dengan var dan juga tipe data
	var text1 string
	text1 = "ini teks 1"
	fmt.Println(text1)

	var text2 string = "ini teks 2"
	text2 = "ini teks 1 diubah"
	fmt.Println(text2)

	// deklarasi dengan menggunakan perantara ":="
	text3 := "ini teks 3"
	text3 = "ini teks 3 diubah"
	fmt.Println(text3)

	text4 := "ini teks 4"
	fmt.Println(text4)

	// contant
	const judul = "ini judul"
	fmt.Println(judul)

	// Tipe Data

	// Tipe Data Numerik Non-Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// Tipe Data Numerik Desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// Tipe Data bool (Boolean)
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// Tipe Data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var message2 = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message2)

	// Nilai nil & Zero Value
	// Semua tipe d`ata yang sudah dibahas di atas memiliki zero value (nilai default tipe data). Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, tetap akan ada nilai default-nya.
	// - Zero value dari string adalah "" (string kosong).
	// - Zero value dari bool adalah false.
	// - Zero value dari tipe numerik non-desimal adalah 0.
	// - Zero value dari tipe numerik desimal adalah 0.0.

	// Zero value berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong. nil tidak bisa digunakan pada tipe data yang sudah dibahas di atas. Ada beberapa tipe data yang bisa di-set nilainya dengan nil, diantaranya:
	// - pointer
	// - tipe data fungsi
	// - slice
	// - map
	// - channel
	// - interface kosong atau interface{}`

	// Konversi Data Menggunakan Teknik Casting
	var a float64 = float64(24)
	fmt.Println(a) // 24

	var b int32 = int32(24.00)
	fmt.Println(b) // 24

	var str = "Halo"
	var c string = string(str[0])
	fmt.Println(c) // H

	// Packages strings dan strconv

	// Package strings
	// 1. strings.Index()
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2

	// 2. strings.Replace()
	var text5 = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text5, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	// 3.strings.Repeat()
	var str2 = strings.Repeat("na", 4)
	fmt.Println(str2) // "nananana"

	// 4. strings.ToLower()
	var str3 = strings.ToLower("aLAy")
	fmt.Println(str3) // "alay"

	// 5. strings.ToUpper()
	var str4 = strings.ToUpper("eat!")
	fmt.Println(str4) // "EAT!"

	// more https://pkg.go.dev/strings

	// Package strconv

	// Fungsi strconv.Atoi()
	var str5 = "124"
	var num, err = strconv.Atoi(str5)

	if err == nil {
		fmt.Println(num) // 124
	}

	// Fungsi strconv.Itoa()
	var num2 = 124
	str, _ := strconv.Itoa(num2)

	fmt.Println(str) // "124"

	// Fungsi strconv.ParseInt()
	var str = "124"
	num, _ := strconv.ParseInt(str, 10, 64)

	fmt.Println(num) // 124

	var str = "1010"
	num, _ := strconv.ParseInt(str, 2, 8)

	fmt.Println(num) // 10

	// Fungsi strconv.FormatInt()
	var num = int64(24)
	str := strconv.FormatInt(num, 8)

	fmt.Println(str) // 30

	// Fungsi strconv.ParseFloat()
	var str = "24.12"
	num, _ := strconv.ParseFloat(str, 32)

	fmt.Println(num) // 24.1200008392334

	// Fungsi strconv.FormatFloat()
	var num = float64(24.12)
	str := strconv.FormatFloat(num, 'f', 6, 64)

	fmt.Println(str) // 24.120000

	// Fungsi strconv.ParseBool()
	var str = "true"
	var bul, _ = strconv.ParseBool(str)

	fmt.Println(bul) // true

	// Fungsi strconv.FormatBool()
	var bul = true
	var str = strconv.FormatBool(bul)

	fmt.Println(str) // true

	// morea https://pkg.go.dev/strconv
}
