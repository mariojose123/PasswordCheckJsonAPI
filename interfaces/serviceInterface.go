/*Interfaces for Password Checker API*/
package interfaces

import (
	"context"

	"github.com/mariojose123/PasswordCheckJsonAPI/internal/structJson"
)

/*Interfaces for Password Check Service */
type PWService interface {
	CheckPW(ctx context.Context, jsonStructure structJson.PSReceiveStructure) (bool, []string)
}
