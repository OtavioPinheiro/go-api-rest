package models

type Personalidade struct {
	Nome     string `josn:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
