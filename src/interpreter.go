package interpreter

import (
	"fmt"
	"github.com/gabrielmellooliveira/rinha-de-compiler/src/models"
	"github.com/gabrielmellooliveira/rinha-de-compiler/src/operations"
	"reflect"
)

func Execute(expression interface{}, environment map[string]interface{}) interface{} {
	switch models.GetField(expression, "kind") {
	case "Call":
		callee := Execute(models.GetField(expression, "callee"), environment)
		calleeKind := models.GetField(callee, "kind").(string)

		if calleeKind != "Function" && calleeKind != "Var" {
			return nil
		}

		arguments := models.GetField(expression, "arguments").([]interface{})
		params := models.GetField(callee, "parameters").([]interface{})

		for i := 0; i < len(arguments); i++ {
			if _, ok := arguments[i].(map[string]interface{}); ok {
				if models.GetField(arguments[i], "kind") == nil {
					arguments[i] = models.GetField(arguments[i], "value")
				} else {
					arguments[i].(map[string]interface{})["value"] = Execute(arguments[i], environment)
				}
			}
		}

		for i := 0; i < len(params); i++ {
			if _, ok := params[i].(map[string]interface{}); ok {
				if models.GetField(params[i], "kind") == nil {
					params[i] = models.GetField(params[i], "text")
				} else {
					params[i] = Execute(params[i], environment)
				}
			}
		}

		if len(arguments) != len(params) {
			panic(fmt.Sprintf("Expected %d arguments, but got %d", len(params), len(arguments)))
		}

		newEnvironment := CopyEnvironment(environment)

		for i := 0; i < len(params); i++ {
			newEnvironment[params[i].(string)] = models.GetField(arguments[i], "value")
		}

		return Execute(models.GetField(callee, "value"), newEnvironment)

	case "Int", "Str", "Bool":
		return models.GetField(expression, "value")

	case "Binary":
		lhs := Execute(models.GetField(expression, "lhs"), environment)
		rhs := Execute(models.GetField(expression, "rhs"), environment)
		op := models.GetField(expression, "op").(string)

		if op == "Add" {
			if reflect.TypeOf(lhs).Kind() == reflect.Float64 {
				lhs = int32(lhs.(float64))
			}

			if reflect.TypeOf(rhs).Kind() == reflect.Float64 {
				rhs = int32(rhs.(float64))
			}

			/*if reflect.TypeOf(lhs).Kind() == reflect.Float64 && reflect.TypeOf(rhs).Kind() == reflect.Float64 {
				return operations.BinaryOperationFloat64(lhs.(float64), rhs.(float64), op)
			}*/

			if reflect.TypeOf(lhs).Kind() == reflect.String || reflect.TypeOf(rhs).Kind() == reflect.String {
				return operations.BinaryAddOperation(lhs, rhs)
			}

			return operations.BinaryOperation((lhs).(int32), (rhs).(int32), op)
		}

		if op == "Eq" {
			return operations.BinaryEqOperation(lhs, rhs)
		}

		if op == "Neq" {
			return operations.BinaryNeqOperation(lhs, rhs)
		}

		if op == "And" || op == "Or" {
			return operations.BinaryBoolOperation((lhs).(bool), (rhs).(bool), op)
		}

		// return operations.BinaryOperation((lhs).(int32), (rhs).(int32), op)

		if reflect.TypeOf(lhs).Kind() == reflect.Float64 {
			lhs = int32(lhs.(float64))
		}

		if reflect.TypeOf(rhs).Kind() == reflect.Float64 {
			rhs = int32(rhs.(float64))
		}

		return operations.BinaryOperation((lhs).(int32), (rhs).(int32), op)

	case "Let":
		value := Execute(models.GetField(expression, "value"), environment)
		name := models.GetField(expression, "name")

		switch value.(type) {
		case int:
			environment[models.GetField(name, "text").(string)] = value.(int)
		case float64:
			environment[models.GetField(name, "text").(string)] = value.(float64)
		case bool:
			environment[models.GetField(name, "text").(string)] = value.(bool)
		case string:
			environment[models.GetField(name, "text").(string)] = value.(string)
		default:
			environment[models.GetField(name, "text").(string)] = Execute(value, environment)
		}

		return Execute(models.GetField(expression, "next"), environment)

	case "Var":
		return environment[models.GetField(expression, "text").(string)]

	case "If":
		condition := Execute(models.GetField(expression, "condition"), environment)
		if condition.(bool) {
			return Execute(models.GetField(expression, "then"), environment)
		} else {
			return Execute(models.GetField(expression, "otherwise"), environment)
		}

	case "Print":
		if models.GetField(models.GetField(expression, "value"), "kind") == "Var" {
			value := models.GetField(expression, "value")
			_, err := models.GetFieldWithError(value, "kind")

			if err != nil {
				tuple := Execute(value, environment)
				first := Execute(models.GetField(tuple, "first"), environment)
				second := Execute(models.GetField(tuple, "second"), environment)

				fmt.Println(first, "|", second)
				return expression
			}

			if models.GetField(value, "kind") == "Var" {
				variable := Execute(value, environment)
				variableKind, err := models.GetFieldWithError(variable, "kind")

				if err != nil {
					fmt.Println(variable)
					return expression
				}

				if variableKind == "Tuple" {
					tuple := Execute(variable, environment)
					first := Execute(models.GetField(tuple, "first"), environment)
					second := Execute(models.GetField(tuple, "second"), environment)

					fmt.Println(first, "|", second)
					return expression
				}

				if variableKind == "Function" {
					fmt.Println("<#closure>")
					return expression
				}

				fmt.Println("variável com o nome:", models.GetField(value, "text"))
				return expression
			}

			fmt.Println(value)
		} else {
			fmt.Println(Execute(models.GetField(expression, "value"), environment))
			return expression
		}

	case "Tuple":
		return expression

	case "First":
		top := models.GetField(expression, "value")
		if _, ok := top.(map[string]interface{}); ok {
			return Execute(models.GetField(Execute(top, environment), "first"), environment)
		} else {
			return Execute(models.GetField(top, "first"), environment)
		}

	case "Second":
		top := models.GetField(expression, "value")
		if _, ok := top.(map[string]interface{}); ok {
			return Execute(models.GetField(Execute(top, environment), "second"), environment)
		} else {
			return Execute(models.GetField(top, "second"), environment)
		}

	case "Function":
		return expression
	}

	return nil
}

func CopyEnvironment(environment map[string]interface{}) map[string]interface{} {
	newEnvironment := make(map[string]interface{})

	for key, value := range environment {
		newEnvironment[key] = value
	}

	return newEnvironment
}
