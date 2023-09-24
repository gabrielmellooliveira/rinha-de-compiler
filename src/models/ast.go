package models

import "errors"

type AST struct {
	Name       string      `json:"name"`
	Expression interface{} `json:"expression"`
	Location   interface{} `json:"location"`
}

func GetField(content interface{}, fieldName string) any {
	return content.(map[string]interface{})[fieldName]
}

func GetFieldWithError(content interface{}, fieldName string) (interface{}, error) {
	cont, ok := content.(map[string]interface{})

	if ok == false {
		return nil, errors.New("Erro ao obter o campo " + fieldName + " do conteúdo")
	}

	response, ok := cont[fieldName]

	if ok == false {
		return nil, errors.New("Erro ao obter o campo " + fieldName + " do conteúdo")
	}

	return response, nil
}
