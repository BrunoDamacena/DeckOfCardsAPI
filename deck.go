package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func new_deck(card_str string) []Card {
	if card_str == "" {
		return complete_deck()
	}
	return custom_deck(card_str)

}

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

func custom_deck(card_str string) []Card {
	deck := []Card{}
	for _, card := range strings.Split(card_str, ",") {
		var suit, value string

		switch card[1:2] {
		case "S":
			suit = "SPADES"
		case "D":
			suit = "DIAMONDS"
		case "C":
			suit = "CLUBS"
		case "H":
			suit = "HEARTS"
		default:
			fmt.Println("ERROR: invalid suit", card[1:2])
			return nil
		}
		switch card[0:1] {
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
			value = card[0:1]
		default:
			fmt.Println("ERROR: invalid value", card[0:1])
			return nil
		}

		deck = append(deck, Card{value, suit, card})
	}
	return deck
}

func shuffle(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func draw_cards(deck *[]Card, amount int) []Card {
	cards_drawned := (*deck)[0:amount]
	*deck = (*deck)[amount:len(*deck)]
	return cards_drawned
}
