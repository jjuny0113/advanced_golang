package cli

import (
	"advancedGolang/rest"
	"flag"
	"fmt"
)

func Start() {
	//rest.Start(4000)
	port := flag.Int("port", 4000, "port to listen on")
	flag.Parse()
	fmt.Println(*port)
	rest.Start(*port)
	
}
