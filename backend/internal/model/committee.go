package model

type Committee struct {
	ID        int64  `json:"id"`
	ShortName string `json:"short_name"`
	Name      string `json:"name"`
}
