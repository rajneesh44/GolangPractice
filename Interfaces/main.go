/*
	INTREFACES ARE LIKE TEMPLATES IN CPP.
	i.e. make things General for different data types.

*/
package main

import (
	"fmt"
)

type bot interface{      // making an interface (custom type)
	getGreeting() string
}

type englishBot struct{

}

type spanishBot struct{

}

func main(){
	eb:=englishBot{}
	sp:=spanishBot{}

	printGreeting(eb)
	printGreeting(sp)

}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string{
	// very custom logic to generate english logic 
	return "hi there!"
}
func (sp spanishBot) getGreeting() string{
	// very custom logic to generate english logic 
	return "Hola!"	
}
// func printGreeting(eb englishBot){
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sp spanishBot){
// 	fmt.Println(sp.getGreeting())
// }



/*
	1. Interfaces are not generic types
	2. Interfaces are implicit   (auto linking of data and interfaces)
	3. Intrefaces are contract to help us manage types.
	4. Intrefaces are tough.
*/