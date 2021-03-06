package d2s

import (
	"encoding/json"
)

// quest struct will represent any quest in game.
type quest [2]byte

// isCompleted will return a bool, telling us if the quest is completed or not.
// We'll do this by looking at bit 0, if it's set or not.
func (q quest) isCompleted() bool {
	return ((q[0] >> 0) & 1) > 0
}

func (q *quest) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		IsCompleted bool `json:"is_completed"`
	}{
		IsCompleted: q.isCompleted(),
	})
}

// Prison Of Ice quest, this has a certain bit we're interested in to calculate resistances.
// If you have consumed the scroll of resistance you'll get +10 @ res.
type prisonOfIce [2]byte

func (q *prisonOfIce) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		IsCompleted    bool `json:"is_completed"`
		ConsumedScroll bool `json:"consumed_scroll"`
	}{
		IsCompleted:    ((q[0] >> 0) & 1) > 0,
		ConsumedScroll: ((q[0] >> 7) & 1) > 0,
	})
}
