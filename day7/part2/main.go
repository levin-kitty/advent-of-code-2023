package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var strength map[byte]int = map[byte]int{
	'J': -1,
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'Q': 10,
	'K': 11,
	'A': 12,
}

const NumCardTypes = 7

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards string
	Bid   int
	Type  int
	Rank  int
}

type Hands []Hand

func (h Hands) Len() int      { return len(h) }
func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Hands) Less(i, j int) bool {
	for k := 0; k < len(h[i].Cards); k++ {
		if strength[h[i].Cards[k]] != strength[h[j].Cards[k]] {
			return strength[h[i].Cards[k]] < strength[h[j].Cards[k]]
		}
	}
	return true
}

type CamelCards []Hands

func (c CamelCards) Winnings() (winnings int) {
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			winnings += (c[i][j].Rank * c[i][j].Bid)
		}
	}
	return
}

func main() {
	camelCards := GetCamelCards("../input")
	fmt.Println(camelCards.Winnings())

}

func GetCamelCards(fileName string) CamelCards {
	hands := make(Hands, 0)
	for _, line := range FileToLines(fileName) {
		tokens := strings.Split(line, " ")
		hands = append(hands, Hand{
			Cards: tokens[0],
			Bid:   MustAtoi(tokens[1]),
			Type:  CardsType(tokens[0]),
		})
	}

	camelCards := make([]Hands, NumCardTypes)
	for i := 0; i < NumCardTypes; i++ {
		camelCards[i] = make(Hands, 0)
	}
	for _, hand := range hands {
		camelCards[hand.Type] = append(camelCards[hand.Type], hand)
	}

	// Sort & Rank
	rank := 1
	for i := 0; i < len(camelCards); i++ {
		sort.Sort(camelCards[i])
		for j := 0; j < len(camelCards[i]); j++ {
			camelCards[i][j].Rank = rank
			rank++
		}
	}

	return camelCards
}

func CardsType(cards string) int {
	// count
	countByLetter := make(map[byte]int)
	for l := 0; l < len(cards); l++ {
		countByLetter[cards[l]]++
	}

	// maximum cards: 5
	maxFreqs := 0
	countFreqs := make([]int, 6)
	for k, v := range countByLetter {
		if k != 'J' {
			countFreqs[v]++
			if v > maxFreqs {
				maxFreqs = v
			}
		}
	}
	for i := 0; i < countByLetter['J']; i++ {
		countFreqs[maxFreqs], countFreqs[maxFreqs+1] = countFreqs[maxFreqs]-1, countFreqs[maxFreqs+1]+1
		maxFreqs++
	}

	switch {
	case countFreqs[5] > 0:
		return FiveOfAKind
	case countFreqs[4] > 0:
		return FourOfAKind
	case countFreqs[3] > 0:
		if countFreqs[2] > 0 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	case countFreqs[2] > 1:
		return TwoPair
	case countFreqs[2] > 0:
		return OnePair
	default:
		return HighCard
	}
}

/* Utility */

func MustAtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func FileToLines(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
