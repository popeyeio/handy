package handy

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func GetUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return StrEmpty, err
	}
	return hex.EncodeToString(u[:]), nil
}
