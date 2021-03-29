# Deck Of Cards API
A simple Go application to handle a deck of cards.

## How to run
### Install your dependencies
First, you need to [download and install Go](https://golang.org/dl/).

With Go installed, you need to get the UUID package. Just execute on your command line:

`go get github.com/google/uuid`

### Serve the API
To serve the API on your host, just execute on the project's root:
`go run .`
Then, the API will be up and running on your localhost!

## Methods
### Create a new Deck
With the POST route `/create`, you can create a new Deck! On the body, you can specify if the new Deck is coming shuffled or not (default is not), like so:
```
{
	"shuffled": false
}
```
On the Query String you can specify the cards that are gonna be present on the deck, like so:

`/create?cards=AS,KH,8C`

In that way, the new deck will have only an Ace of Spades, a King of Hearts, and an 8 of Clubs.

This route returns a JSON containing the Deck UUID, if it's shuffled or not, and the total of cards remaining on it, like so:

`localhost:8080/create?cards=AS,KH,8C`
```
{
    "deck_id": "6d1e1342-81d7-4703-9c25-bda87b9ec30a",
    "shuffled": false,
    "remaining": 3
}
```

### Open a Deck
With the GET route `/open/${deck_UUID}`, you can open an existing Deck! 

This route returns a JSON containing the Deck UUID, if it's shuffled or not, the total of cards remaining on it, and all the cards on it, in order, like so:

`localhost:8080/open/6d1e1342-81d7-4703-9c25-bda87b9ec30a`
```
{
    "deck_id": "6d1e1342-81d7-4703-9c25-bda87b9ec30a",
    "shuffled": false,
    "remaining": 3,
    "cards": [
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
        {
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        },
        {
            "value": "8",
            "suit": "CLUBS",
            "code": "8C"
        }
    ]
}
``` 

### Draw a Card
With the PUT route `/draw/${deck_UUID}/${number_of_cards}`, you can draw the first N cards of this deck! 

This route returns a JSON containing the drawn cards from the Deck, like so:

`localhost:8080/draw/6d1e1342-81d7-4703-9c25-bda87b9ec30a/2`
```
{
    "cards": [
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
        {
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        }
    ]
}
```