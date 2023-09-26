package operations

import (
	"fmt"
	"reflect"
)

func BinaryOperation(lhs int32, rhs int32, op string) interface{} {
	switch op {
	case "Add":
		return lhs + rhs
	case "Sub":
		return lhs - rhs
	case "Mul":
		return lhs * rhs
	case "Div":
		if rhs == 0 {
			panic("Não é possivel realizar divisão por zero")
		}

		return lhs / rhs
	case "Rem":
		return lhs % rhs
	case "Eq":
		return lhs == rhs
	case "Neq":
		return lhs != rhs
	case "Lt":
		return lhs < rhs
	case "Gt":
		return lhs > rhs
	case "Lte":
		return lhs <= rhs
	case "Gte":
		return lhs >= rhs
	default:
		panic("Operação não identificada: " + op)
	}
}

func BinaryOperationFloat64(lhs float64, rhs float64, op string) interface{} {
	switch op {
	case "Add":
		return lhs + rhs
	case "Sub":
		return lhs - rhs
	case "Mul":
		return lhs * rhs
	case "Div":
		if rhs == 0.0 {
			panic("Não é possivel realizar divisão por zero")
		}

		return lhs / rhs
	case "Rem":
		return int32(lhs) % int32(rhs)
	case "Eq":
		return lhs == rhs
	case "Neq":
		return lhs != rhs
	case "Lt":
		return lhs < rhs
	case "Gt":
		return lhs > rhs
	case "Lte":
		return lhs <= rhs
	case "Gte":
		return lhs >= rhs
	default:
		panic("Operação não identificada: " + op)
	}
}

func BinaryEqOperation(lhs interface{}, rhs interface{}) bool {
	if reflect.TypeOf(lhs).Kind() == reflect.Float64 {
		lhs = int32(lhs.(float64))
	}

	if reflect.TypeOf(rhs).Kind() == reflect.Float64 {
		rhs = int32(rhs.(float64))
	}

	if reflect.TypeOf(lhs).Kind() == reflect.String || reflect.TypeOf(rhs).Kind() == reflect.String {
		return lhs == rhs
	}

	return (lhs).(int32) == (rhs).(int32)
}

func BinaryNeqOperation(lhs interface{}, rhs interface{}) bool {
	if reflect.TypeOf(lhs).Kind() == reflect.Float64 {
		lhs = int32(lhs.(float64))
	}

	if reflect.TypeOf(rhs).Kind() == reflect.Float64 {
		rhs = int32(rhs.(float64))
	}

	if reflect.TypeOf(lhs).Kind() == reflect.String || reflect.TypeOf(rhs).Kind() == reflect.String {
		return lhs != rhs
	}

	return (lhs).(int32) != (rhs).(int32)
}

func BinaryAddOperation(lhs interface{}, rhs interface{}) interface{} {
	if intLhs, ok := lhs.(int32); ok {
		if intRhs, ok := rhs.(int32); ok {
			return intRhs + intLhs
		} else if strRhs, ok := rhs.(string); ok {
			// fmt.Printf("%d%s", intLhs, strRhs)
			return fmt.Sprintf("%d%s", intLhs, strRhs)
		}
	} else if strLhs, ok := lhs.(string); ok {
		if intRhs, ok := rhs.(int32); ok {
			return fmt.Sprintf("%s%d", strLhs, intRhs)
		} else if strRhs, ok := rhs.(string); ok {
			return fmt.Sprintf("%s%s", strLhs, strRhs)
		}
	}

	return nil
}

func BinaryBoolOperation(lhs bool, rhs bool, op string) bool {
	switch op {
	case "And":
		return lhs && rhs
	case "Or":
		return lhs || rhs
	default:
		panic("Operação não identificada: " + op)
	}
}
