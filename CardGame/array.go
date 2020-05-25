package main

//import "fmt"

// Array and Slice
// array- fixed length
// variable sized array + other features.
// both must be of identical type i.e. either int or string
func main(){
	//fmt.Println("Hello")
	////cards:=  []string{newCard(), newCard(),"Ace of diamonds"}
	//cards:=  deck{newCard(), newCard(),"Ace of diamonds"}
	//fmt.Println(cards)
	//cards = append(cards, "siz of spades")
	//fmt.Println(cards)
	//cards.print()

	//for i, card:=range cards{
	//	if i < 3 {
	//		fmt.Println(i, card)
	//	}
		//else{
		//	fmt.Println("gaid")
		//}
	//}
//	Generating a newDeck through deck.go
//	cards:=newDeck()
//	hand, remainingDeck:=deal(cards, 5)
//	hand.print()
//	remainingDeck.print()
//	cards.print()

//converting into bytes
	//cards:=newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards:=readDeckFromFile("my_cards")
	cards.print()
}
func newCard() string{
	return "Five of diamonds"
}
