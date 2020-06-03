package main 

import (
	"fmt"
	)

type info struct{
		email string
		zipCode int
}

type person struct{
	firstName string
	lastName string
	contact info
}


func main(){
	
	// //first way 
	// alex:=person{"rajneesh", "sharma"}
	// fmt.Println(alex.firstName, alex.lastName)
	// //second way
	// raj:= person{firstName:"rajneesh", lastName:"sharma"}
	// fmt.Println(raj)
	// // third way 
	// var rajjo person
	// rajjo.firstName = "rajneesh"
	// rajjo.lastName = "sharma"
	// fmt.Println(rajjo)

	// using multiple structs
	raj:=person{
		firstName:"rajneesh",
		lastName:"sharma",
		contact:info{
			email:"rajneesh@gmail",
			zipCode:202001,
		},
	}
	// fmt.Printf("%+v",raj)

	//making a reciever function and using it
	
	raj.updateName("rajjo")    //this did not update the name  //basically passing by value and passiing the data to a copy of data.
	raj.print()

	rajPointer:=&raj   // pointer passing data by reference
	rajPointer.updateName("rajjo")    //this updates the name  //basically passing by value and passiing the data to a copy of data.
	raj.print()
}

func (p person) print(){
	fmt.Printf("%+v",p)
}
func (pointerToPerson *person) updateName (newFirstName string){
	(*pointerToPerson).firstName=newFirstName
	// fmt.Println(p) this prints the changed name
}