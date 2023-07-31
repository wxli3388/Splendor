package serverLib

import (
	"sync"
)

type PlayerAssets struct {
	Score                int
	Deduct               map[Gems]int
	Gems                 map[Gems]int
	Cards                []*Card
	ReservedCard         []*Card
	ReservedCardPosition map[int]int
	Noble                []*Noble
	mutex                sync.RWMutex
}

func NewPlayerAssets() *PlayerAssets {
	return &PlayerAssets{
		Deduct:               map[Gems]int{Diamond: 0, Sapphire: 0, Emerald: 0, Ruby: 0, Onyx: 0, Gold: 0},
		Gems:                 map[Gems]int{Diamond: 0, Sapphire: 0, Emerald: 0, Ruby: 0, Onyx: 0, Gold: 0},
		Cards:                []*Card{},
		ReservedCard:         []*Card{},
		ReservedCardPosition: map[int]int{},
		Noble:                []*Noble{},
	}
}
