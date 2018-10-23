package vidopre

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/afero"
)

type CtxNewPost struct {
	Date    string
	Title   string
	Content string
}

var (
	mesi = map[time.Month]string{
		time.January:   "Gennaio",
		time.February:  "Febbraio",
		time.March:     "Marzo",
		time.April:     "Aprile",
		time.May:       "Maggio",
		time.June:      "Giugno",
		time.July:      "Luglio",
		time.August:    "Agosto",
		time.September: "Settembre",
		time.October:   "Ottobre",
		time.November:  "Novembre",
		time.December:  "Dicembre",
	}
	giorniSett = map[time.Weekday]string{
		time.Sunday:    "Domenica",
		time.Monday:    "Lunedì",
		time.Tuesday:   "Martedì",
		time.Wednesday: "Mercoledì",
		time.Thursday:  "Giovedì",
		time.Friday:    "Venerdì",
		time.Saturday:  "Sabato",
	}
)

func NewPost(dirOut string, title string, content string) {
	dirOutAbs, err := filepath.Abs(dirOut)
	if err != nil {
		log.Fatal(err)
	}

	afs := &afero.Afero{Fs: appfs}
	ct := time.Now()
	postFileName := fmt.Sprintf("%d-%d-%d-_%s", ct.Year(), ct.Month(), ct.Day(), title)
	if len(postFileName) > 25 {
		postFileName = postFileName[0:25]
	}
	postFileName = strings.Replace(postFileName, " ", "_", -1)
	postFileName += ".txt"
	postFileName = filepath.Join(dirOutAbs, postFileName)
	f, err := afs.Create(postFileName)
	if err != nil {
		log.Fatalln("Error create file", err)
	}
	defer f.Close()

	wd := giorniSett[ct.Weekday()]
	mm := mesi[ct.Month()]
	ctx := CtxNewPost{
		Title:   title,
		Content: content,
		Date:    fmt.Sprintf("%s, %d %s %d", wd, ct.Day(), mm, ct.Year()),
	}
	var t = template.Must(template.New("Page").Parse(tempNewPost))
	err = t.Execute(f, ctx)
	if err != nil {
		log.Fatal("Template error: ", err)
	}

	log.Println("Created new post file", postFileName)

}
