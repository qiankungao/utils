package rand

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(RandInt(1000,9999))
	}
}
