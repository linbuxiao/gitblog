package main

import (
	"os"
	"testing"
)

func TestReadMe(t *testing.T) {
	_, err := os.Stat("README.md")
	if err != nil {
		if os.IsNotExist(err) {
			t.Log("not exist")
		}
	}
	t.Log("exist")
}
