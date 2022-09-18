package model

type Capitals struct {
	Id      int    `form:"id" json:"id"`
	Country string `form:"country" json:"country", binding: "required"`
	Capital string `form:"capital" json:"capital", binding: "required"`
}
