package algorithm

import (
	"encoding/base64"
	"github.com/google/uuid"
)

func GetRandomStr() string {
	binary1, _ := uuid.New().MarshalBinary()
	binary2, _ := uuid.New().MarshalBinary()
	binary1 = append(binary1, binary2...)
	return base64.StdEncoding.EncodeToString(binary1)
}
