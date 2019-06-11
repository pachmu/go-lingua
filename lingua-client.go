package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type LinguaClient interface {
	GetTranslation(s string) (string, error)
}

func NewLinguaClient(url string) LinguaClient {
	return &lingua{
		url: url,
	}
}

type result struct {
	Translation   []translation `json:"translate"`
	Transcription string        `json:"transcription"`
}

type translation struct {
	Value string `json:"value"`
}

type lingua struct {
	url string
}

func (l *lingua) GetTranslation(s string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?word=%s", l.url, s))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	values := []string{s, "", fmt.Sprintf("[%s]", result.Transcription)}
	for _, t := range result.Translation {
		values = append(values, t.Value)
	}
	return strings.Join(values, "\n"), nil
}
