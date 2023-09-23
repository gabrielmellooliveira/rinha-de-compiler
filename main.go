package main

import (
	"encoding/json"
	interpreter "github.com/gabrielmellooliveira/rinha-de-compiler/src"
	"github.com/gabrielmellooliveira/rinha-de-compiler/src/models"
	"os"
)

const RINHA_SOURCE_PATH string = "./var/rinha/source.rinha.json"

var total = 0

func fibonacci(n int) int {
	total++
	println("n", n)

	if n < 2 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	println("Rinha de compiladores")

	/*println("fibonacci: ", fibonacci(10))
	print("total: ", total)*/

	/*println("\n######## Código de print #########")
	RunFile("./var/rinha/run-print.rinha.json")

	println("\n########## Código de if ##########")
	RunFile("./var/rinha/run-if.rinha.json")

	println("\n####### Código de function #######")
	RunFile("./var/rinha/run-fn.rinha.json")

	println("\n#### Código de Hello function ####")
	RunFile("./var/rinha/run-hello-fn.rinha.json")

	println("\n######## Código de Tuple #########")
	RunFile("./var/rinha/run-tuple.rinha.json")

	println("\n####### Código de Tuple 2 ########")
	RunFile("./var/rinha/run-tuple-2.rinha.json")*/

	/*println("\n####### Código de combination ########")
	RunFile("./var/rinha/run-combination-fn.rinha.json")*/

	/*println("\n####### Código de concate ########")
	RunFile("./var/rinha/run-concate.rinha.json")*/

	/*println("\n####### Código de division ########")
	RunFile("./var/rinha/run-div.rinha.json")*/

	/*println("\n####### Código de division by zero ########")
	RunFile("./var/rinha/run-div-by-zero.rinha.json")*/

	/*println("\n####### Código de multiplication ########")
	RunFile("./var/rinha/run-mult.rinha.json")*/

	/*println("\n####### Código de print to function ########")
	RunFile("./var/rinha/run-print-fn.rinha.json")*/

	/*println("\n####### Código de subtration ########")
	RunFile("./var/rinha/run-sub.rinha.json")*/

	/*println("########## Código de fib ##########")
	RunFile("./var/rinha/run-fib.rinha.json")*/

	/*println("########## Código de sum ##########")
	RunFile("./var/rinha/run-sum-fn.rinha.json")*/

	// Ler o arquivo
	content, err := os.ReadFile("")
	if err != nil {
		panic("Erro ao ler arquivo da rinha: " + err.Error())
	}

	// Transformar em JSON
	var ast models.AST
	err = json.Unmarshal(content, &ast)

	if err != nil {
		panic("Erro ao converter o conteúdo do arquivo da rinha para JSON: " + err.Error())
	}

	// Executar o código
	environment := make(map[string]interface{})
	interpreter.Execute(ast.Expression, environment)
}

func RunFile(fileName string) {
	// Ler o arquivo
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic("Erro ao ler arquivo da rinha: " + err.Error())
	}

	// Transformar em JSON
	var ast models.AST
	err = json.Unmarshal(content, &ast)

	if err != nil {
		panic("Erro ao converter o conteúdo do arquivo da rinha para JSON: " + err.Error())
	}

	// Executar o código
	environment := make(map[string]interface{})
	interpreter.Execute(ast.Expression, environment)
}
