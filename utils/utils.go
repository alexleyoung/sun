package utils

import "time"

func IsBefore(t1 time.Time , t2 time.Time) bool {
	return t1.Before(t2)
}