package serverLib

import (
	"math/rand"
	"reflect"
	"sync"
	"time"
)

type GameTable struct {
	Gems                map[Gems]int
	cardNumberPosition  map[int]int
	deckLow             []*Card
	cardLow             []*Card
	deckMid             []*Card
	cardMid             []*Card
	deckHigh            []*Card
	cardHigh            []*Card
	noble               []*Noble
	nobleNumberPosition map[int]int
	mutex               sync.RWMutex
	roundRemain         int
}

func NewGameTable(maxPlayer int) *GameTable {
	gemsCnt := map[int]int{2: 4, 3: 5, 4: 7}[maxPlayer]
	cardNumberPosition := map[int]int{}
	nobleNumberPosition := map[int]int{}

	cardTemp := make([]*Card, len(cardLowSetting))
	copy(cardTemp, cardLowSetting)
	ShuffleDeck(cardTemp)
	cardLow := []*Card{}
	for i := 0; i < 4; i++ {
		cardLow = append(cardLow, cardTemp[i])
		cardNumberPosition[cardTemp[i].Number] = i
	}
	deckLow := cardTemp[4:]

	cardTemp = make([]*Card, len(cardMidSetting))
	copy(cardTemp, cardMidSetting)
	ShuffleDeck(cardTemp)
	cardMid := []*Card{}
	for i := 0; i < 4; i++ {
		cardMid = append(cardMid, cardTemp[i])
		cardNumberPosition[cardTemp[i].Number] = i
	}
	deckMid := cardTemp[4:]

	cardTemp = make([]*Card, len(cardHighSetting))
	copy(cardTemp, cardHighSetting)
	ShuffleDeck(cardTemp)
	cardHigh := []*Card{}
	for i := 0; i < 4; i++ {
		cardHigh = append(cardHigh, cardTemp[i])
		cardNumberPosition[cardTemp[i].Number] = i
	}
	deckHigh := cardTemp[4:]

	nobleTemp := make([]*Noble, len(nobleSetting))
	copy(nobleTemp, nobleSetting)
	ShuffleDeck(nobleTemp)
	noble := []*Noble{}
	for i := 0; i < maxPlayer+1; i++ {
		noble = append(noble, nobleTemp[i])
		nobleNumberPosition[nobleTemp[i].Number] = i
	}

	gems := map[Gems]int{
		Diamond:  gemsCnt,
		Sapphire: gemsCnt, //藍
		Emerald:  gemsCnt, //綠
		Ruby:     gemsCnt,
		Onyx:     gemsCnt,
		Gold:     5,
	}
	return &GameTable{
		Gems:                gems,
		noble:               noble,
		nobleNumberPosition: nobleNumberPosition,
		cardNumberPosition:  cardNumberPosition,
		deckLow:             deckLow,
		cardLow:             cardLow,
		deckMid:             deckMid,
		cardMid:             cardMid,
		deckHigh:            deckHigh,
		cardHigh:            cardHigh,
		roundRemain:         4,
	}
}

func ShuffleDeck(list interface{}) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := reflect.ValueOf(list)

	swap := reflect.Swapper(list)
	length := v.Len()
	for i := length - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		swap(i, j)
	}
}

func (gt *GameTable) reportDeck() *CmdOutGameTable {
	// gt.mutex.RLock()
	// defer gt.mutex.RUnlock()
	return &CmdOutGameTable{
		CardLow:         gt.cardLow,
		IsDeckLowEmpty:  len(gt.deckLow) == 0,
		CardMid:         gt.cardMid,
		IsDeckMidEmpty:  len(gt.deckMid) == 0,
		CardHigh:        gt.cardHigh,
		IsDeckHighEmpty: len(gt.deckHigh) == 0,
		Noble:           gt.noble,
		Gems:            gt.Gems,
	}
}

func (gt *GameTable) DrawCard(deck *[]*Card, cardLevel []*Card, index int) {
	if len(*deck) > 0 {
		newCard := (*deck)[len(*deck)-1]
		gt.cardNumberPosition[newCard.Number] = index
		(*deck) = (*deck)[:len(*deck)-1]
		cardLevel[index] = newCard
	} else {
		cardLevel[index] = nil //todo: nil maybe a bad practice
	}
}

// func (gt *GameTable) IsGoldEnough() bool {
// 	gt.mutex.RLock()
// 	defer gt.mutex.RUnlock()
// 	return gt.Gems[Gold] > 0
// }

func (gt *GameTable) getGold() bool {
	if gt.Gems[Gold] == 0 {
		return false
	}
	gt.Gems[Gold] -= 1
	return true
}

// func (gt *GameTable) getCanBuyCard() []*Card{

// }
