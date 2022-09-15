package lib

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	once sync.Once
)

// SeedMathRand Credit: https://github.com/hashicorp/consul/blob/main/lib/rand.go
func SeedMathRand() error {
	var (
		n   *big.Int
		err error
	)

	once.Do(func() {
		n, err = crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			log.Errorf("cannot seed math/rand: %s", err)
		} else {
			log.Debugf("seeding math/rand %+v", n.Int64())
			rand.Seed(n.Int64())
		}
	})

	return err
}

func RandInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func MapRand[K comparable, V any](m map[K]V) V {
	n := rand.Intn(len(m))
	i := 0
	for _, v := range m {
		if i == n {
			return v
		}
		i++
	}
	panic("unreachable")
}

func MapRandKey[K comparable, V any](m map[K]V) K {
	keys := MapKeys(m)
	n := rand.Intn(len(m))
	return keys[n]
}
