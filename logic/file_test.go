package logic

import (
	"fmt"
	"os"
	"testing"
)

func TestSaveUpload(t *testing.T) {
	_, err := os.Open("ssss")
	fmt.Println(string(WrapError(err).stack))
}
