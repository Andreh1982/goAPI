package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade

type ErrorMessage struct {
	Msg  string `json:"Msg"`
	Name string `json:"Name"`
}

var ErrorMessages []ErrorMessage
