package serverLib

import (
	"encoding/json"
)

type CmdOutYourTurn struct {
	Turn bool `json:"turn"`
}

func (c *CmdOutYourTurn) String() string {
	s, err := json.Marshal(c)
	if err != nil {
		return "yourTurn {}"
	}
	return "yourTurn " + string(s)
}

type CmdOutGameOver struct {
	CoverCard []string `json:"coverCard"`
	Score     []int    `json:"score"`
}

func (c *CmdOutGameOver) String() string {
	s, err := json.Marshal(c)
	if err != nil {
		return "gameOver {}"
	}
	return "gameOver " + string(s)
}

type CmdOutBackToRoom struct {
}

func (c *CmdOutBackToRoom) String() string {
	return "backToRoom"
}

type CmdOutGameTable struct {
	CardLow         []*Card      `json:"cardLow"`
	IsDeckLowEmpty  bool         `json:"isDeckLowEmpty"`
	CardMid         []*Card      `json:"cardMid"`
	IsDeckMidEmpty  bool         `json:"isDeckMidEmpty"`
	CardHigh        []*Card      `json:"cardHigh"`
	IsDeckHighEmpty bool         `json:"isDeckHighEmpty"`
	Noble           []*Noble     `json:"noble"`
	Gems            map[Gems]int `json:"gems"`
}

func (c *CmdOutGameTable) String() string {
	s, err := json.Marshal(c)
	if err != nil {
		return "gameTable {}"
	}
	return "gameTable " + string(s)
}
