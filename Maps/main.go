// Maps are a key value pair 

//maps vs structs ? 
// keys & values must be of same type respectively.
// 


package main

import "fmt"

func main(){
	//making a string-string map 
	// method-1
	colors:=map[string]string{
		"red":"#ff0000",
		"green":"#ffcddj",
	}
	fmt.Println(colors)
	//method-2
	var colors2 map[string]string
	fmt.Println(colors2)

	// method 3
	colors3:=make(map[string]string)
	fmt.Println(colors3)
	
	colors["white"]="#000000"
	fmt.Println(colors)

	//deleting data 
	delete(colors,"white")
	fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string){
	for color,hex:=range c{
		fmt.Println(color,">",hex)
	}
}


/*
	Maps 
	1.	Keys must be ofsame type.& values must be of same type
	2. 	All differnet keys are indexed.
	3. 	Reference Type		(On passing gives the refenece to original data)
	4.	Use to represent a collection of related data 
	5. 	Dont need to know all the keys at compile time 
	
	Structs
	1.	Values can be any type.
	2.	No indexing in keys
	3.	Value type. 		(On passing gievs, the copy of data)
	4.	You need to know all different fields at compile time.
	5.	User data type
*/