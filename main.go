package main

import (
	"fmt"

	"github.com/mrk21/go-web-config-sample/config"
)

func main() {
	fmt.Printf("%+v\n", *config.Get())
}
