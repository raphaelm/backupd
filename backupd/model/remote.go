package model

type Remote struct {
	ID       int64  `json:"id"`
	Driver   string `json:"driver"`
	Location string `json:"location"`
}

type Remotes []Remote
