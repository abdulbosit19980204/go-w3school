package main

import (
	"fmt"
)

// oddiy funkisya
func myMessage() {
	fmt.Println("I just got executed!")
}

// parametrli funkisya fnmae string tipidagi parametr
func familyName(fname string) {
	fmt.Println("Hello", fname, "welcome")
}

// multiple parametrlar bilan
func familyNameM(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

// qiymat qaytaruvchi funkisyalar x va y qabul qilinuvchi parametrlar int tipida qiymat qaytariladi
func myFunction(x int, y int) int {
	return x + y
}

// qaytaruluvchi qiymatlarni nomlashimiz mumkin
func myFunctionN(x int, y int) (result int) {
	result = x + y
	return

	// yoki quydagicha yozishimiz mumkin
	//return result
}

// Bundan tashqari func bir nechta qiymat qaytarishi mumkin Bu yerda result int va txt1 string qaytarilyapti
func myFunctionM(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

//Recursion Functions

// Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.
// Funktsiya o'zini chaqirsa va to'xtash shartiga erishsa, rekursiv hisoblanadi.
// x=11 bolganda funkisya 0 qiymat qaytaradi va ishini to'xtatadi. Bolmasam x di qiymatini 1 ga oshirib beradi
func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

// factorial ni topish funkisyasi
func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	myMessage() // call the function
	familyName("Tuychiev")
	familyNameM("Tuychiev", 25)
	fmt.Println("Bizning funkisyamiz myFunc = ", myFunction(5, 11), "qiymat qaytaryapti")
	fmt.Println(myFunctionN(7, 8))

	//Store the Return Value in a Variable
	returnedVal := myFunctionN(88, 11)
	fmt.Println("returnedVal = ", returnedVal)

	// Multiple return func
	fmt.Println(myFunctionM(5, "Hello"))

	//recursiv func
	testcount(5)

	//factorial func
	fmt.Println("8 faktarila = ", factorial_recursion(8))

}
