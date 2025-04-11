package main

type LLM interface {
	Prompt(string) (string, error)
}
