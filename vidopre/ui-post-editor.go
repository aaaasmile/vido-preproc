package vidopre

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/spf13/afero"
)

type ctxPostEdit struct {
	TitlePost   string
	ContentPost string
	Buildnr     string
}

var (
	selectedTitle   string
	selectedContent string
)

func editPost(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("EditPost").Parse(tempHtmlBase))
	t = template.Must(t.Parse(tempHtmlIndex))
	t = template.Must(t.Parse(tempHtmlEditPost))

	pagectx := &ctxPostEdit{
		Buildnr:     BuildNr,
		ContentPost: selectedContent,
		TitlePost:   selectedTitle,
	}
	fmt.Println(pagectx)

	err := t.ExecuteTemplate(w, "base", pagectx)
	if err != nil {
		log.Fatal(err)
	}
}

func startEditor(title string, content string) {

	selectedContent = content
	selectedTitle = title
	surl := "localhost:4200"
	http.HandleFunc("/", editPost)
	log.Println("Starting http server at ", fmt.Sprintf("http://%s", surl))
	log.Fatal(http.ListenAndServe(surl, nil))
}

func EditNewPost() {
	startEditor("", "")
}

func EditLastPost(dirIn string) {
	items, err := afero.ReadDir(appfs, dirIn) // sorted by name as default, order is acending. Oldest first.
	if err != nil {
		log.Fatal(err)
	}
	if len(items) == 0 {
		log.Fatalln("Post source directory is empty. Command not available")
	}
	dir, _ := filepath.Abs(dirIn)
	sourceFname := filepath.Join(dir, items[len(items)-1].Name())
	afs := &afero.Afero{Fs: appfs}
	src, err := afs.ReadFile(sourceFname)
	if err != nil {
		log.Fatal("Error reading file ", err)
	}
	log.Printf("Editing post %s\n", sourceFname)

	startEditor("", string(src))
}
