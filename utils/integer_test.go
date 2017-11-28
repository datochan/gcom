package utils

import (
	"testing"
	"fmt"
	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

func TestGenerateIndex(t *testing.T) {
	result := GenerateIndex(0, 1, 10)
	fmt.Println(result)

	result = GenerateIndex(10, 1, 15)
	fmt.Println(result)

	result = GenerateIndex(10, -2, 0)
	fmt.Println(result)
}
