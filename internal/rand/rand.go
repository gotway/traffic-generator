package rand

import "math/rand"

func Item(items ...string) string {
	return items[Int(0, len(items)-1)]
}

func Int(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Bool() bool {
	return rand.Intn(2) == 0
}
