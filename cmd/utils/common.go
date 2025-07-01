package utils

import (
	"hash/fnv"
	"time"
)

func hash(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func NowInVietnam() time.Time {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return time.Now().UTC()
	}
	return time.Now().In(loc)
}
