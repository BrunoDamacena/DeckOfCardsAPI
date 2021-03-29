package main

import (
	"math/rand"
	"strings"
	"time"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// create a new deck
func new_deck(card_str string) []Card {
	if card_str == "" {
		return complete_deck()
	}
	return custom_deck(card_str)

}

// create a complete deck on the correct order of suits and values
func complete_deck() []Card {
	suits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
	values := []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
	deck := []Card{}
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{value, suit, value[0:1] + suit[0:1]})
		}
	}
	return deck
}

/*
	create a custom deck based on the cards you provide on the card_str
	the card_str must have every card Code separated by commas
	if there is any invalid card on the structure, the method returns nil
*/
func custom_deck(card_str string) []Card {
	deck := []Card{}
	for _, card_code := range strings.Split(card_str, ",") {
		var suit, value string

		switch card_code[1:2] {
		case "S":
			suit = "SPADES"
		case "D":
			suit = "DIAMONDS"
		case "C":
			suit = "CLUBS"
		case "H":
			suit = "HEARTS"
		default:
			return nil
		}
		switch card_code[0:1] {
		case "A":
			value = "ACE"
		case "J":
			value = "JACK"
		case "Q":
			value = "QUEEN"
		case "K":
			value = "KING"
		case "1":
			value = "10"
		case "2", "3", "4", "5", "6", "7", "8", "9":
			value = card_code[0:1]
		default:
			return nil
		}

		deck = append(deck, Card{value, suit, card_code})
	}
	return deck
}

// sets a random seed and shuffle the deck
func shuffle(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

// draw {{amount}} cards and removes it from the original deck
func draw_cards(deck *[]Card, amount int) []Card {
	cards_drawned := (*deck)[0:amount]
	*deck = (*deck)[amount:len(*deck)]
	return cards_drawned
}
