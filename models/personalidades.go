package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `josn:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
