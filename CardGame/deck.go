package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// we are going to make a new datatype deck which is a slcie of string
type deck []string
func newDeck() deck{
	cards:=deck{}
	cardSuits := []string{"Spades", "Ace","Diamonds", "Hearts"}
	cardValues := []string{"one","two","three","king"}
	for _,suit:=range cardSuits {
		for _,value:=range cardValues{
			cards = append(cards,value+" of "+suit)
		}
	}
	return cards
}
//receiver function *better refer it to a single letter/ 2-letter starting with datatype like d for deck.
func (d deck) print(){
	for i,card :=range d{
		fmt.Println(i,card)
	}
}
// d is the copy of deck object
func deal(d deck, handSize int) (deck,deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()),0666)
}
func readDeckFromFile(filename string) deck{
	bs, err:= ioutil.ReadFile(filename)     //receiving byte slice and error from the file.
	if err!=nil{
		//if error, log the error and quit the program
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	s :=strings.Split(string(bs),",")
	return deck(s)
}