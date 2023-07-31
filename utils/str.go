package utils

import (
	"github.com/google/uuid"
	"math"
	"strconv"
	"strings"
)

const (
	base    = 10
	bitSize = 32
)

func Str2Uint(str string) uint {
	num, err := strconv.ParseUint(str, base, bitSize)
	if err != nil || num == math.MaxUint {
		return 0
	}
	return uint(num)
}

func Str2UintArr(str string) (ids []uint) {
	idArr := strings.Split(str, ",")
	for _, v := range idArr {
		ids = append(ids, Str2Uint(v))
	}
	return
}

func Str2UUIDArr(str string) (ids []uuid.UUID) {
	idArr := strings.Split(str, ",")
	for _, v := range idArr {
		ids = append(ids, uuid.MustParse(v))
	}
	return
}
