package main

import (
	"fmt"

	"github.com/mitchallen/go-monorepo-demo/pkg/alpha"
)

func main() {
	fmt.Println(alpha.CoinCount(200))

	alpha.Hello()
}
