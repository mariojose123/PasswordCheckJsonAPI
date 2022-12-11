package interfaces

import "passwordcheck/internal/structJson"

type PWChecker interface {
	AddMatch(jsonStruct structJson.PSReceiveStructure, noMatch []string) []string
}
