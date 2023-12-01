package main

import (
	"fmt"
	"testing"

	"github.com/mrk21/go-web-config-sample/config"
)

func TestEnv(t *testing.T) {
	fmt.Printf("%+v\n", *config.Get())
}
