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
	selectedIsNew   bool
)

func savePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Save post - TO DO...")

}

func editPost(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP: Edit post with title", selectedTitle)

	t := template.Must(template.New("EditPost").Parse(tempHtmlBase))
	t = template.Must(t.Parse(tempHtmlIndex))
	t = template.Must(t.Parse(tempHtmlEditPost))

	pagectx := &ctxPostEdit{
		Buildnr:     BuildNr,
		ContentPost: selectedContent,
		TitlePost:   selectedTitle,
	}
	//fmt.Println(pagectx)

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
	http.HandleFunc("/save-post/", savePost)
	log.Println("Starting http server at ", fmt.Sprintf("http://%s", surl))
	log.Fatal(http.ListenAndServe(surl, nil))
}

func EditLastPost(dirIn string) {
	selectedIsNew = false
	items, err := afero.ReadDir(appfs, dirIn) // sorted by name as default, order is acending. Oldest first.
	if err != nil {
		log.Fatal(err)
	}
	if len(items) == 0 {
		log.Fatalln("Post source directory is empty. Command not available")
	}
	dir, _ := filepath.Abs(dirIn)
	itemName := items[len(items)-1].Name()
	sourceFname := filepath.Join(dir, itemName)
	afs := &afero.Afero{Fs: appfs}
	src, err := afs.ReadFile(sourceFname)
	if err != nil {
		log.Fatal("Error reading file ", err)
	}
	log.Printf("Editing post %s\n", sourceFname)

	startEditor(itemName, string(src))
}
