package main

import (
	"encoding/json"
	interpreter "github.com/gabrielmellooliveira/rinha-de-compiler/src"
	"github.com/gabrielmellooliveira/rinha-de-compiler/src/models"
	"os"
)

const RINHA_SOURCE_PATH string = "./var/rinha/source.rinha.json"

func main() {
	// Rodar esse comando para executar os testes
	// test.RunTests()

	// Ler o arquivo
	content, err := os.ReadFile(RINHA_SOURCE_PATH)
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
