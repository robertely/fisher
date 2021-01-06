package crazyivan

import (
	"math/rand"
	"time"
)

type Crazyivan struct {
	Name string
}

// New - Chonkers
func New() Crazyivan {
	return Crazyivan{}
}

// Check - Chonks
func Check(c Crazyivan) (bool, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1, nil
}
