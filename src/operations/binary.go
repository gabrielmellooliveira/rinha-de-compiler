package operations

import (
	"fmt"
	"math/big"
)

func BinaryOperation(lhs int32, rhs int32, op string) interface{} {
	switch op {
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
		if rhs == 0 {
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

func BinaryAddOperation(lhs interface{}, rhs interface{}) interface{} {
	if intLhs, ok := lhs.(*big.Int); ok {
		if intRhs, ok := rhs.(*big.Int); ok {
			fmt.Println("BinaryAddOperation 1:", new(big.Int).Add(intRhs, intLhs))
			return new(big.Int).Add(intRhs, intLhs)
		} else if strRhs, ok := rhs.(string); ok {
			fmt.Printf("%d%s", intLhs, strRhs)
			return fmt.Sprintf("%d%s", intLhs, strRhs)
		}
	} else if strLhs, ok := lhs.(string); ok {
		if intRhs, ok := rhs.(*big.Int); ok {
			fmt.Printf("%s%d", strLhs, intRhs)
			return fmt.Sprintf("%s%d", strLhs, intRhs)
		} else if strRhs, ok := rhs.(string); ok {
			fmt.Printf("%s%s", strLhs, strRhs)
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
