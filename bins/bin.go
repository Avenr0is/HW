package bins

import "time"


type Bin struct {
	Id        string `json:"id"`
	Private   bool   `json:"private"`
	CreatedAt time.Time `json:"time"`
	Name      string   `json:name`
}

type Binlist struct {
	items []Bin
}