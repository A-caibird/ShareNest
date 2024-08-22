package main

import (
	"ShareNest/internal/util"
	"fmt"
	"testing"
)

func TestReadEnv(t *testing.T) {
	envMap, err := util.LoadEnvFile("../../env/.dev.env")
	if err != nil {
		fmt.Printf("%s", err)
	}
	for key, value := range envMap {
		fmt.Printf("%s = %s\n", key, value)
	}
}
