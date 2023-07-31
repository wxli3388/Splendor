package serverLib

type Gems int

const (
	Diamond Gems = iota
	Sapphire
	Emerald
	Ruby
	Onyx
	Gold
)

const (
	CardMidStartNumber  int = 41
	CardHighStartNumber int = 71
)

type Card struct {
	Cost   map[Gems]int `json:"cost"`
	Score  int          `json:"score"`
	Deduct Gems         `json:"deduct"`
	Number int          `json:"number"`
}

type Noble struct {
	Cost   map[Gems]int `json:"cost"`
	Score  int          `json:"score"`
	Number int          `json:"number"`
}

var nobleSetting = []*Noble{
	&Noble{
		Cost:   map[Gems]int{Ruby: 4, Onyx: 4},
		Score:  3,
		Number: 1,
	},
	&Noble{
		Cost:   map[Gems]int{Diamond: 4, Sapphire: 4},
		Score:  3,
		Number: 2,
	},
	&Noble{
		Cost:   map[Gems]int{Diamond: 3, Ruby: 3, Onyx: 3},
		Score:  3,
		Number: 3,
	},
	&Noble{
		Cost:   map[Gems]int{Diamond: 3, Sapphire: 3, Onyx: 3},
		Score:  3,
		Number: 4,
	},
	&Noble{
		Cost:   map[Gems]int{Diamond: 4, Onyx: 4},
		Score:  3,
		Number: 5,
	},
	&Noble{
		Cost:   map[Gems]int{Sapphire: 4, Emerald: 4},
		Score:  3,
		Number: 6,
	},
	&Noble{
		Cost:   map[Gems]int{Diamond: 3, Sapphire: 3, Emerald: 3},
		Score:  3,
		Number: 7,
	},
	&Noble{
		Cost:   map[Gems]int{Emerald: 3, Ruby: 3, Onyx: 3},
		Score:  3,
		Number: 8,
	},
	&Noble{
		Cost:   map[Gems]int{Emerald: 4, Ruby: 4},
		Score:  3,
		Number: 9,
	},
	&Noble{
		Cost:   map[Gems]int{Sapphire: 3, Emerald: 3, Ruby: 3},
		Score:  3,
		Number: 10,
	},
}

var cardLowSetting = []*Card{
	&Card{
		Cost:   map[Gems]int{Diamond: 3},
		Score:  0,
		Deduct: Ruby,
		Number: 1,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Sapphire: 1, Emerald: 1, Onyx: 1},
		Score:  0,
		Deduct: Ruby,
		Number: 2,
	},
	&Card{
		Cost:   map[Gems]int{Ruby: 2, Onyx: 1},
		Score:  0,
		Deduct: Diamond,
		Number: 3,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 1, Emerald: 1, Ruby: 1, Onyx: 1},
		Score:  0,
		Deduct: Diamond,
		Number: 4,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Sapphire: 1, Onyx: 1},
		Score:  0,
		Deduct: Diamond,
		Number: 5,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Sapphire: 1, Ruby: 1, Onyx: 2},
		Score:  0,
		Deduct: Emerald,
		Number: 6,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Sapphire: 1},
		Score:  0,
		Deduct: Emerald,
		Number: 7,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Onyx: 2},
		Score:  0,
		Deduct: Sapphire,
		Number: 8,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 3},
		Score:  0,
		Deduct: Onyx,
		Number: 9,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Sapphire: 2, Emerald: 1, Ruby: 1},
		Score:  0,
		Deduct: Onyx,
		Number: 10,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 1, Emerald: 2, Ruby: 1, Onyx: 1},
		Score:  0,
		Deduct: Diamond,
		Number: 11,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 2, Onyx: 2},
		Score:  0,
		Deduct: Diamond,
		Number: 12,
	},
	&Card{
		Cost:   map[Gems]int{Ruby: 4},
		Score:  1,
		Deduct: Sapphire,
		Number: 13,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 2, Onyx: 2},
		Score:  0,
		Deduct: Sapphire,
		Number: 14,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Emerald: 1, Ruby: 1, Onyx: 1},
		Score:  0,
		Deduct: Sapphire,
		Number: 15,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Sapphire: 2, Ruby: 1},
		Score:  0,
		Deduct: Onyx,
		Number: 16,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 4},
		Score:  1,
		Deduct: Onyx,
		Number: 17,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 2, Ruby: 1},
		Score:  0,
		Deduct: Onyx,
		Number: 18,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 4},
		Score:  1,
		Deduct: Ruby,
		Number: 19,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 1, Ruby: 2, Onyx: 2},
		Score:  0,
		Deduct: Emerald,
		Number: 20,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Sapphire: 3, Emerald: 1},
		Score:  0,
		Deduct: Emerald,
		Number: 21,
	},
	&Card{
		Cost:   map[Gems]int{Ruby: 3},
		Score:  0,
		Deduct: Emerald,
		Number: 22,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 4},
		Score:  1,
		Deduct: Diamond,
		Number: 23,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Sapphire: 1, Emerald: 1, Onyx: 1},
		Score:  0,
		Deduct: Ruby,
		Number: 24,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Emerald: 1, Ruby: 2, Onyx: 1},
		Score:  0,
		Deduct: Sapphire,
		Number: 25,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 1, Emerald: 3, Ruby: 1},
		Score:  0,
		Deduct: Sapphire,
		Number: 26,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Sapphire: 1, Ruby: 1, Onyx: 1},
		Score:  0,
		Deduct: Emerald,
		Number: 27,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 2, Ruby: 2},
		Score:  0,
		Deduct: Emerald,
		Number: 28,
	},
	&Card{
		Cost:   map[Gems]int{Onyx: 4},
		Score:  1,
		Deduct: Emerald,
		Number: 29,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Emerald: 2, Ruby: 2},
		Score:  0,
		Deduct: Sapphire,
		Number: 30,
	},
	&Card{
		Cost:   map[Gems]int{Onyx: 3},
		Score:  0,
		Deduct: Sapphire,
		Number: 31,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 3},
		Score:  0,
		Deduct: Diamond,
		Number: 32,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 2, Emerald: 2, Onyx: 1},
		Score:  0,
		Deduct: Diamond,
		Number: 33,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Ruby: 1, Onyx: 3},
		Score:  0,
		Deduct: Ruby,
		Number: 34,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 2, Emerald: 1},
		Score:  0,
		Deduct: Ruby,
		Number: 35,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Emerald: 1, Onyx: 2},
		Score:  0,
		Deduct: Ruby,
		Number: 36,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Emerald: 2},
		Score:  0,
		Deduct: Onyx,
		Number: 37,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Sapphire: 1, Emerald: 1, Ruby: 1},
		Score:  0,
		Deduct: Onyx,
		Number: 38,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 1, Ruby: 3, Onyx: 1},
		Score:  0,
		Deduct: Onyx,
		Number: 39,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Ruby: 2},
		Score:  0,
		Deduct: Ruby,
		Number: 40,
	},
}

var cardMidSetting = []*Card{
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Onyx: 5},
		Score:  2,
		Deduct: Ruby,
		Number: 41,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 6},
		Score:  3,
		Deduct: Sapphire,
		Number: 42,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 6},
		Score:  3,
		Deduct: Emerald,
		Number: 43,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 6},
		Score:  3,
		Deduct: Diamond,
		Number: 44,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Emerald: 3, Onyx: 2},
		Score:  1,
		Deduct: Onyx,
		Number: 45,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 5},
		Score:  2,
		Deduct: Emerald,
		Number: 46,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 5, Ruby: 3},
		Score:  2,
		Deduct: Onyx,
		Number: 47,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Ruby: 2, Onyx: 3},
		Score:  1,
		Deduct: Ruby,
		Number: 48,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 1, Emerald: 4, Ruby: 2},
		Score:  2,
		Deduct: Onyx,
		Number: 49,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Emerald: 2, Ruby: 3},
		Score:  1,
		Deduct: Emerald,
		Number: 50,
	},
	&Card{
		Cost:   map[Gems]int{Ruby: 5},
		Score:  2,
		Deduct: Diamond,
		Number: 51,
	},
	&Card{
		Cost:   map[Gems]int{Ruby: 6},
		Score:  3,
		Deduct: Ruby,
		Number: 52,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 1, Ruby: 4, Onyx: 2},
		Score:  2,
		Deduct: Diamond,
		Number: 53,
	},
	&Card{
		Cost:   map[Gems]int{Onyx: 5},
		Score:  2,
		Deduct: Ruby,
		Number: 54,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 3, Ruby: 2, Onyx: 2},
		Score:  1,
		Deduct: Diamond,
		Number: 55,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 4, Sapphire: 2, Onyx: 1},
		Score:  2,
		Deduct: Emerald,
		Number: 56,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 5, Sapphire: 3},
		Score:  2,
		Deduct: Sapphire,
		Number: 57,
	},
	&Card{
		Cost:   map[Gems]int{Onyx: 6},
		Score:  3,
		Deduct: Onyx,
		Number: 58,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 5, Emerald: 3},
		Score:  2,
		Deduct: Emerald,
		Number: 59,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 2, Emerald: 3, Onyx: 3},
		Score:  1,
		Deduct: Emerald,
		Number: 60,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 5},
		Score:  2,
		Deduct: Onyx,
		Number: 61,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Ruby: 1, Onyx: 4},
		Score:  2,
		Deduct: Sapphire,
		Number: 62,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 2, Emerald: 2, Ruby: 2},
		Score:  1,
		Deduct: Sapphire,
		Number: 63,
	},
	&Card{
		Cost:   map[Gems]int{Ruby: 5, Onyx: 3},
		Score:  2,
		Deduct: Diamond,
		Number: 64,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Sapphire: 3, Ruby: 3},
		Score:  1,
		Deduct: Diamond,
		Number: 65,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 1, Sapphire: 4, Emerald: 2},
		Score:  2,
		Deduct: Ruby,
		Number: 66,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Sapphire: 2, Emerald: 2},
		Score:  1,
		Deduct: Onyx,
		Number: 67,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 3, Ruby: 2, Onyx: 3},
		Score:  1,
		Deduct: Ruby,
		Number: 68,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 5},
		Score:  2,
		Deduct: Sapphire,
		Number: 69,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 2, Sapphire: 3, Onyx: 2},
		Score:  1,
		Deduct: Emerald,
		Number: 70,
	},
}

var cardHighSetting = []*Card{
	&Card{
		Cost:   map[Gems]int{Ruby: 7},
		Score:  4,
		Deduct: Onyx,
		Number: 71,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 7, Ruby: 3},
		Score:  5,
		Deduct: Ruby,
		Number: 72,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 7},
		Score:  4,
		Deduct: Emerald,
		Number: 73,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Ruby: 3, Onyx: 6},
		Score:  4,
		Deduct: Diamond,
		Number: 74,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Sapphire: 3, Emerald: 5, Ruby: 3},
		Score:  3,
		Deduct: Onyx,
		Number: 75,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 3, Ruby: 6, Onyx: 3},
		Score:  4,
		Deduct: Onyx,
		Number: 76,
	},
	&Card{
		Cost:   map[Gems]int{Onyx: 7},
		Score:  4,
		Deduct: Diamond,
		Number: 77,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Sapphire: 6, Emerald: 3},
		Score:  4,
		Deduct: Emerald,
		Number: 78,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 7, Sapphire: 3},
		Score:  5,
		Deduct: Sapphire,
		Number: 79,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 3, Emerald: 6, Ruby: 3},
		Score:  4,
		Deduct: Ruby,
		Number: 80,
	},
	&Card{
		Cost:   map[Gems]int{Emerald: 7},
		Score:  4,
		Deduct: Ruby,
		Number: 81,
	},
	&Card{
		Cost:   map[Gems]int{Ruby: 7, Onyx: 3},
		Score:  5,
		Deduct: Onyx,
		Number: 82,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 7, Emerald: 3},
		Score:  5,
		Deduct: Emerald,
		Number: 83,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 5, Sapphire: 3, Ruby: 3, Onyx: 3},
		Score:  3,
		Deduct: Emerald,
		Number: 84,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 6, Sapphire: 3, Onyx: 3},
		Score:  4,
		Deduct: Sapphire,
		Number: 85,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Onyx: 7},
		Score:  5,
		Deduct: Diamond,
		Number: 86,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 7},
		Score:  4,
		Deduct: Sapphire,
		Number: 87,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Sapphire: 5, Emerald: 3, Onyx: 3},
		Score:  3,
		Deduct: Ruby,
		Number: 88,
	},
	&Card{
		Cost:   map[Gems]int{Sapphire: 3, Emerald: 3, Ruby: 5, Onyx: 3},
		Score:  3,
		Deduct: Diamond,
		Number: 89,
	},
	&Card{
		Cost:   map[Gems]int{Diamond: 3, Emerald: 3, Ruby: 3, Onyx: 5},
		Score:  3,
		Deduct: Sapphire,
		Number: 90,
	},
}
