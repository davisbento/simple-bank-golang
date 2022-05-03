package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/goombaio/namegenerator"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer in the range [min, max].
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random length n string.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	return nameGenerator.Generate()
}

func RandomBalance() int64 {
	return RandomInt(0, 100)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
