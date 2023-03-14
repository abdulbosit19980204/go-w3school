//Go Maps
//Maps are used to store data values in key:value pairs.

//Each element in a map is a key:value pair.

//A map is an unordered and changeable collection that does not allow duplicates.

// The length of a map is the number of its elements. You can find it using the len() function.

// The default value of a map is nil.

// Maps hold references to an underlying hash table.

// Go has multiple ways for creating maps.

//Creating Maps Using var and :=
/*
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}
*/

package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	//Creating Maps Using Using make()Function:
	/*
		var a = make(map[KeyType]ValueType)
		b := make(map[KeyType]ValueType)
	*/
	var c = make(map[string]string) // The map is empty now
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"
	// a is no longer empty
	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)

	//Creating an Empty Map
	//var a map[KeyType]ValueType
	/*
		Note: The make()function is the right way to create an empty map.
		If you make an empty map in a different way and write to it,
		it will causes a ***runtime panic***.
	*/
	var e = make(map[string]string)
	var f map[string]string

	fmt.Println(e == nil)
	fmt.Println(f == nil)

	//Accessing Map Elements
	var h = make(map[string]string)
	h["brand"] = "Ford"
	h["model"] = "Mustang"
	h["year"] = "1964"

	fmt.Printf(h["brand"])

	//Updating and Adding Map Elements
	var g = make(map[string]string)
	g["brand"] = "Ford"
	g["model"] = "Mustang"
	g["year"] = "1964"

	fmt.Println(a)

	g["year"] = "1970" // Updating an element
	g["color"] = "red" // Adding an element

	fmt.Println(g)

	//Remove Element from Map
	var i = make(map[string]string)
	i["brand"] = "Ford"
	i["model"] = "Mustang"
	i["year"] = "1964"

	fmt.Println(i)

	delete(i, "year")

	fmt.Println(i)
}
