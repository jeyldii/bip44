package bip44

import (
	"encoding/json"
)

type Constants []constant

type constant struct {
	Constant   string `json:"constant"`
	CoinSymbol string `json:"coinSymbol"`
	CoinName   string `json:"coinName"`
}

func convert(sl string) Constants {
	slip := Constants{}
	if err := json.Unmarshal([]byte(sl), &slip); err != nil {
		return nil
	}

	return slip
}

// Create is starting point. You get constants.
func Create() Constants {
	slip := convert(Slip0044)
	return slip
}


func(c Constants) Get(coinSymbol ...string) []constant {
	constants := make([]constant, 0)

	for _, symbol := range coinSymbol {
		for _, v := range c {
			if v.CoinSymbol == symbol {
				constants = append(constants, v)
			}
		}
	}

	return constants
}