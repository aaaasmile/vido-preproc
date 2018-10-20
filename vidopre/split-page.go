package vidopre

import (
	"log"
	"path/filepath"
)

func SplitPages(dirname string) {
	dir, err := filepath.Abs(dirname)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Split all pages (*.page) in %s\n", dir)
	// TODO iterate dir
}
