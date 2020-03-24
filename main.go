package main

import "fmt"
import "time"


func main(){
	// ini coment 1 baris
	fmt.Println("Hello World")

	var first_num int = 290
	fmt.Println("Thr first number is", first_num)

	// //or
	// var second_num float64 = 12.12
	// var thrid_num float64 = 42.22
	// fmt.Println("the numbers are", second_num, thrid_num)

	// or
	// var second_num, thrid_num float64 = 12.12, 42.22
	// fmt.Println("The numbers are", second_num, thrid_num)

	//or
	second_num, thrid_num := 12.22, 42.22
	fmt.Println("the numbers are", second_num, thrid_num)

	fourthnUm := first_num
	fourthnUm = 100
	fmt.Println(fourthnUm)

	fmt.Println("time now is", time.Now())

	// var constanta
	const var1 int = 2
	const var2 string = "hello"

	fmt.Println(var1)

	// deklarasi bolean
	var varBolean bool = false
	
	fmt.Println(varBolean)

	// var a int = 1
	// var b string = "bbbb"
	// var c bool = a == b
	// fmt.Println("nilai c adalah",c)
	//if statemnt
	i := 10

	if i == 10{
		fmt.Println("sama..")
	}else{
		fmt.Println("Tidak sama")
	}

	
		var nama string = "ismail"
		var umur int = 23
		var suhu int = 26
		var farenhite int = (suhu *9/5)+32
		var tahun int = 2020 - umur

		fmt.Println("Data Diri")
		fmt.Println("Nama : ", nama)
		fmt.Println("Umur : ", umur)
		fmt.Println("Celcius : ",suhu)
		fmt.Println("Farehite : ", farenhite)
		fmt.Println("Tahun Lahir : ", tahun)
		if farenhite >= 73{
			fmt.Println("Suhu : Panas")
		}else{
			fmt.Println("Suhu : Dingin")
		}
	
	// nama umur thn lahir celcius far
	// far >= 73 panas     >73 dingin
	fmt.Println("tambah apa aja")

}
