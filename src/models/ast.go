package models

import "errors"

// Types

type AST struct {
	Name       string      `json:"name"`
	Expression interface{} `json:"expression"`
	Location   Location    `json:"location"`
}

type Expression struct {
	Kind     string      `json:"kind"`
	Name     Name        `json:"name"`
	Value    interface{} `json:"value"`
	Next     interface{} `json:"next"`
	Location Location    `json:"location"`
}

type Name struct {
	Text     string   `json:"text"`
	Location Location `json:"location"`
}

type Location struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Filename string `json:"filename"`
}

// Methods

func (a AST) GetField(fieldName string) string {
	return a.Expression.(map[string]interface{})[fieldName].(string)
}

func GetField(content interface{}, fieldName string) any {
	return content.(map[string]interface{})[fieldName]
}

func GetFieldWithError(content interface{}, fieldName string) (interface{}, error) {
	cont, err := content.(map[string]interface{})

	if err == true {
		return nil, errors.New("1 - Erro ao obter o campo " + fieldName + " do conteúdo")
	}

	response, err := cont[fieldName]

	if err == true {
		return nil, errors.New("2 - Erro ao obter o campo " + fieldName + " do conteúdo")
	}

	return response, nil
}
