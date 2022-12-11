package interfaces

import (
	"context"

	"passwordcheck/internal/structJson"
)

/*Interfaces for Services */
type PWService interface {
	CheckPW(ctx context.Context, jsonStructure structJson.PSReceiveStructure) (bool, []string)
}
