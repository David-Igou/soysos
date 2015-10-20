package main

type CatFact struct {
	Species string `json:"species"`
	Fact    string `json:"fact"`
}

type CatFacts []CatFact
