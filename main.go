package main

import (
	"flag"
	"fmt"

	"./vidopre"
)

func main() {
	cmd := flag.String("cmd", "", "Comnads: splitpages, newpost, createindex")
	flag.Parse()
	switch *cmd {
	case "splitpages":
		vidopre.SplitPages("./data/page-in", "./data/post-src")
		break
	case "newpost":
		break
	case "createindex":
		vidopre.CreateIndexPostPages("./data/post-src", "./data/page-out", 13)
		break
	default:
		fmt.Println("Vido site pre-processor (before webgen). Please use --usage to see all options.")
	}

}
