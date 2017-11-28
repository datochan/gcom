package utils

import (
	"testing"
	"fmt"
)

func TestRandomMacAddress(t *testing.T) {
	uuid := RandomMacAddress()
	fmt.Println(uuid)
}
