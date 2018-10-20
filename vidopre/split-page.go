package vidopre

import (
	"log"
	"path/filepath"

	"github.com/spf13/afero"
)

func SplitPages(dirIn string, dirOut string) {
	dir, err := filepath.Abs(dirIn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Split all pages (*.page) in %s\n", dir)
	appfs := afero.NewOsFs() // Wrapper del file os molto interessante
	items, err := afero.ReadDir(appfs, dirIn)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		log.Println("Processing: ", item.Name())
		splitSinglePage(dirIn, item.Name(), dirOut)
	}
}

func splitSinglePage(dirIn string, fname string, dirOut string) {

}
