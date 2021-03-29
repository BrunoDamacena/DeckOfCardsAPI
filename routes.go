package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Deck struct {
	Uuid     string
	Cards    *[]Card
	Shuffled bool `json:"shuffled"`
}

var all_decks = []Deck{}

/*
	search and return a deck from the UUID
	if not exists, return nil
*/
func getDeckFromUUID(uuid string) *Deck {
	for _, deck := range all_decks {
		if deck.Uuid == uuid {
			return &deck
		}
	}
	return nil
}

// handles default error message for the API route
func handleErrorMessage(message string, statusCode int, w *http.ResponseWriter) {
	(*w).WriteHeader(statusCode)
	(*w).Write([]byte(`{"message": "Error: ` + message + `"}`))
}

// route to creates a new deck
func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// only allows POST requests
	if r.Method != "POST" {
		handleErrorMessage("method not allowed", http.StatusInternalServerError, &w)
		return
	}
	url := r.URL
	var body Deck
	json.NewDecoder(r.Body).Decode(&body)

	cards_str := url.Query().Get("cards")
	new_deck_of_cards := new_deck(cards_str)

	if new_deck_of_cards == nil {
		handleErrorMessage("could not create new deck of cards. please check out the requested cards", http.StatusInternalServerError, &w)
		return
	}

	if body.Shuffled {
		shuffle(new_deck_of_cards)
	}

	deck_uuid := uuid.NewString()

	new_deck := Deck{deck_uuid, &new_deck_of_cards, body.Shuffled}
	all_decks = append(all_decks, new_deck)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"deck_id": "` + deck_uuid + `", "shuffled": ` + strconv.FormatBool(new_deck.Shuffled) +
		`, "remaining": ` + strconv.Itoa(len(*new_deck.Cards)) + `}`))
}

// route to open a new deck
func open(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// only allows GET requests
	if r.Method != "GET" {
		handleErrorMessage("method not allowed", http.StatusInternalServerError, &w)
		return
	}
	url := r.URL
	deck_uuid := strings.Split(url.String(), "/")[2]

	deck := getDeckFromUUID(deck_uuid)

	if deck == nil {
		handleErrorMessage("deck not found!", http.StatusNotFound, &w)
		return
	}

	deck_json, err := json.Marshal(deck.Cards)

	if err != nil {
		handleErrorMessage(err.Error(), http.StatusInternalServerError, &w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"deck_id": "` + deck_uuid + `", "shuffled": ` + strconv.FormatBool(deck.Shuffled) +
		`,"remaining": ` + strconv.Itoa(len(*deck.Cards)) + `, "cards": ` +
		string(deck_json) + `}`))

}

// route to draw {{amount}} cards from the deck
func draw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// only allows PUT requests
	if r.Method != "PUT" {
		handleErrorMessage("method not allowed", http.StatusInternalServerError, &w)
		return
	}
	url := r.URL
	deck_uuid := strings.Split(url.String(), "/")[2]
	cards_to_draw, err := strconv.Atoi(strings.Split(url.String(), "/")[3])

	if err != nil {
		handleErrorMessage(err.Error(), http.StatusInternalServerError, &w)
		return
	}

	deck := getDeckFromUUID(deck_uuid)

	if deck == nil {
		handleErrorMessage("deck not found!", http.StatusNotFound, &w)
		return
	}

	if cards_to_draw > len(*deck.Cards) {
		handleErrorMessage("cannot draw that many cards! deck only have "+strconv.Itoa(len(*deck.Cards)), http.StatusInternalServerError, &w)
		return
	}

	cards_drawned := draw_cards(deck.Cards, cards_to_draw)

	deck_json, err := json.Marshal(cards_drawned)

	if err != nil {
		handleErrorMessage(err.Error(), http.StatusInternalServerError, &w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"cards": ` + string(deck_json) + `}`))
}
