package handy

import (
	"encoding/hex"

	"github.com/satori/go.uuid"
)

func GetUUID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return StrEmpty, err
	}
	return hex.EncodeToString(u.Bytes()), nil
}
