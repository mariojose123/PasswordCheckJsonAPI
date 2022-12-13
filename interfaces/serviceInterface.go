/*Interfaces for Password Checker API*/
package interfaces

import (
	"context"

	"passwordcheck/internal/structJson"
)

/*Interfaces for Password Check Service */
type PWService interface {
	CheckPW(ctx context.Context, jsonStructure structJson.PSReceiveStructure) (bool, []string)
}
