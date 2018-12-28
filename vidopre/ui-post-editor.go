package vidopre

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/afero"
)

type ctxPostEdit struct {
	TitlePost          string
	ContentPost        string
	Buildnr            string
	LastMessage        string
	WebgenOutIndexFile string
}

var (
	selectedTitle       string
	selectedContent     string
	selectedFileName    string
	lastMessageInEditor string
)

func createPageIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Crea le pagine index asemblando tutti i vari post")
	CreateIndexPostPages(Conf.PostSourceDir, Conf.OutDirPage, Conf.PostPerPage)

	lastMessageInEditor = fmt.Sprintf("Pagine index create in  %s", Conf.OutDirPage)
	http.Redirect(w, r, "/", http.StatusFound)
}

func createSite(w http.ResponseWriter, r *http.Request) {
	log.Println("Lancia webgen")
	if err := execWebgen(); err != nil {
		log.Printf("Error on executing webgen: %v", err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func savePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Save post in ", selectedFileName)
	selectedContent = r.FormValue("contentpost")
	afs := &afero.Afero{Fs: appfs}
	f, err := afs.Create(selectedFileName) // Nota che con open non riesco a scrivere
	if err != nil {
		log.Fatalln("Error open file", err)
	}
	defer f.Close()
	_, err = f.WriteString(selectedContent)
	if err != nil {
		log.Fatal("Unable to save the file:", err)
	}
	lastMessageInEditor = fmt.Sprintf("Messaggio salvato su %s", selectedFileName)
	http.Redirect(w, r, "/", http.StatusFound)
}

func editPost(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP: Edit post with title", selectedTitle)

	// t := template.Must(template.New("EditPost").Parse(tempHtmlBase))
	// t = template.Must(t.Parse(tempHtmlIndex))
	// t = template.Must(t.Parse(tempHtmlEditPost))

	pagectx := &ctxPostEdit{
		Buildnr:            BuildNr,
		ContentPost:        selectedContent,
		TitlePost:          selectedTitle,
		LastMessage:        lastMessageInEditor,
		WebgenOutIndexFile: Conf.WebgenOutIndexFile,
	}

	templName := "templates/index.html"
	log.Println("Load the template and reload on request")
	t := template.Must(template.New("EditPost").ParseFiles(templName))

	//fmt.Println(pagectx)

	err := t.ExecuteTemplate(w, "base", pagectx)
	if err != nil {
		log.Fatal(err)
	}
}

func startEditor(title string, content string) {

	selectedContent = content
	selectedTitle = title
	lastMessageInEditor = ""

	surl := Conf.UiServerUrl
	urlInbrowser := fmt.Sprintf("http://%s", surl)
	http.HandleFunc("/", editPost)
	http.HandleFunc("/save-post/", savePost)
	http.HandleFunc("/create-page-index/", createPageIndex)
	http.HandleFunc("/exec-webgen/", createSite)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	log.Println("Starting http server at ", urlInbrowser)
	go openBrowser(urlInbrowser)
	log.Fatal(http.ListenAndServe(surl, nil))
}

func execWebgen() error {
	var cmd string
	var args []string
	if runtime.GOOS == "windows" {
		cmd = "cmd"
		args = []string{"/c", "start", "-d", Conf.WebGenWebPageDir}
	} else {
		log.Fatal("OS not recognized")
		return fmt.Errorf("OS not supported %s", runtime.GOOS)
	}
	log.Printf("Exec webgen (source %s) in %s", Conf.WebGenLocation, Conf.WebGenWebPageDir)
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return err
	}
	log.Printf("Webgen %s\n", out)
	return nil
}

func openBrowser(url string) error {
	var cmd string
	var args []string
	if runtime.GOOS == "windows" {
		cmd = "cmd"
		args = []string{"/c", "start"}
	} else {
		log.Fatal("OS not recognized")
		return fmt.Errorf("OS not supported %s", runtime.GOOS)
	}
	args = append(args, url)
	log.Println("open a browser url ", url)
	return exec.Command(cmd, args...).Start()

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
	itemName := items[len(items)-1].Name()
	selectedFileName = filepath.Join(dir, itemName)
	afs := &afero.Afero{Fs: appfs}
	src, err := afs.ReadFile(selectedFileName)
	if err != nil {
		log.Fatal("Error reading file ", err)
	}
	log.Printf("Editing post %s\n", selectedFileName)

	startEditor(itemName, string(src))
}
