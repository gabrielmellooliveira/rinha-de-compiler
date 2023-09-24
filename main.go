package main

import (
	"encoding/json"
	interpreter "github.com/gabrielmellooliveira/rinha-de-compiler/src"
	"github.com/gabrielmellooliveira/rinha-de-compiler/src/models"
	"os"
	"path"
)

func main() {
	// Rodar esse comando para executar os testes
	// test.RunTests()

	// Capturar o nome do arquivo
	fileName := GetFileName()

	// Ler o arquivo
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic("Erro ao ler arquivo da rinha: " + err.Error())
	}

	// Mapear o json
	var ast models.AST
	err = json.Unmarshal(content, &ast)

	if err != nil {
		panic("Erro ao converter o conteúdo do arquivo da rinha para JSON: " + err.Error())
	}

	// Executar o código
	environment := make(map[string]interface{})
	interpreter.Execute(ast.Expression, environment)
}

const RINHA_PATH string = "/var/rinha/"
const RINHA_DEFAULT_FILE string = "run-print.rinha.json"

func GetFileName() string {
	args := os.Args

	if len(args) < 2 {
		return RINHA_PATH + RINHA_DEFAULT_FILE
	} else {
		fileName := args[1]
		ext := path.Ext(fileName)

		if ext == ".json" {
			return RINHA_PATH + fileName
		} else if ext == ".rinha" {
			return RINHA_PATH + fileName + ".json"
		} else {
			return RINHA_PATH + fileName + ".rinha.json"
		}
	}
}
