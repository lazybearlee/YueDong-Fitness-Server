package appmodel

type UserStepRank struct {
	UserID    uint   `json:"userId"`
	HeaderImg string `json:"headerImg"`
	Nickname  string `json:"nickname"`
	Step      uint   `json:"step"`
	Rank      uint   `json:"rank"`
}

type UserDistanceRank struct {
	UserID    uint    `json:"userId"`
	HeaderImg string  `json:"headerImg"`
	Nickname  string  `json:"nickname"`
	Distance  float64 `json:"distance"`
	Rank      uint    `json:"rank"`
}
