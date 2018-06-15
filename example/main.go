package main

import (
	"fmt"

	compressible "github.com/sniperkit/compressible/pkg"
)

func main() {

	fmt.Println("`text/html` compressible=", compressible.Is("text/html"))
	// -> true

	fmt.Println("`image/png` compressible=", compressible.Is("image/png"))
	// -> false

	var wt compressible.WithThreshold = 1024

	fmt.Println("`text/html`: compressible.WithThreshold=", wt, ", compressible=", wt.Compressible("text/html", 1023))
	// -> false

}
