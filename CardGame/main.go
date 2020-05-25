package main

import "fmt"

func main(){
	//var card string = "Ace of Spades"
	card := "Ace of spades"   //use colon= while *initializing a variable*
	card2:= newCard()
	fmt.Println(card)
	fmt.Println(card2)
	//card3:=newCrad()     //calling int function
	//fmt.Println(card3)
}

func newCard() string{
	return "Five of Diamonds"
}
func newCrad() int {
	return 2
}

// we have to specially tell the data type as it is a statically typed language
// basic data types : bool , string , int . float64