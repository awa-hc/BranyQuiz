package auth

import (
	"context"
)

type AuthRepository interface {
	LoginWithEmail(ctx context.Context, email string, password string) (string, error)
}
