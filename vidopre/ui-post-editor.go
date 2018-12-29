package vidopre

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

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

func buildLastMsg(msg string) {
	tt := time.Now()
	lastMessageInEditor = fmt.Sprintf("[%s] %s", tt.Format("2006-01-02 15:04:05"), msg)
}

func createPageIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Crea le pagine index asemblando tutti i vari post")
	CreateIndexPostPages(Conf.PostSourceDir, Conf.OutDirPage, Conf.PostPerPage)

	buildLastMsg(fmt.Sprintf("Index pages created in  %s", Conf.OutDirPage))

	http.Redirect(w, r, "/", http.StatusFound)
}

func viewWebgenOut(w http.ResponseWriter, r *http.Request) {
	log.Println("Navigate to webgen out")
	go openBrowser(Conf.WebgenOutIndexFile)

	http.Redirect(w, r, "/", http.StatusFound)
}

func createSite(w http.ResponseWriter, r *http.Request) {
	log.Println("Lancia webgen")

	go execWebgen()

	buildLastMsg("Webgen lanciato in una command console, per favore controlla lì il risultato.")
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
	buildLastMsg(fmt.Sprintf("Messaggio salvato su %s", selectedFileName))
	http.Redirect(w, r, "/", http.StatusFound)
}

func editPostRoot(w http.ResponseWriter, req *http.Request) {
	log.Println("HTTP: Edit post with title", selectedTitle)
	switch req.Method {
	case "GET":
		handleIndexGet(w, req)
	case "POST":
		log.Println("POST", req.RequestURI)
		handleIndexPost(w, req)
	}

}
func handleIndexGet(w http.ResponseWriter, req *http.Request) {
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

	err := t.ExecuteTemplate(w, "base", pagectx)
	if err != nil {
		log.Fatal(err)
	}
}

func handleIndexPost(w http.ResponseWriter, req *http.Request) {
	u, err := url.Parse(req.RequestURI)
	if err != nil {
		log.Println("Error uri: ", err)
		return
	}
	if ok := checkDoRquest(req.RequestURI); !ok {
		log.Println("Command invalid", req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}
	q := u.Query()
	log.Println(q) // interface: do?clear=preprocessor

	if val, ok := q["clear"]; ok {
		log.Println("DO clear last message", val)
		if val[0] == "preprocessor" {
			lastMessageInEditor = ""
			return
		}
	}

	log.Println("Command invalid", req.RequestURI)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Bad Request"))
}

func checkDoRquest(reqURI string) bool {
	aa := strings.Split(reqURI, "/")
	cmd := aa[len(aa)-1]
	//fmt.Println(cmd)
	return strings.HasPrefix(cmd, "do")
}

func startEditor(title string, content string, openNewPage bool) {

	selectedContent = content
	selectedTitle = title
	lastMessageInEditor = ""

	surl := Conf.UiServerUrl
	urlInbrowser := fmt.Sprintf("http://%s", surl)
	http.HandleFunc("/", editPostRoot)
	http.HandleFunc("/save-post/", savePost)
	http.HandleFunc("/create-page-index/", createPageIndex)
	http.HandleFunc("/exec-webgen/", createSite)
	http.HandleFunc("/open-webgen-out/", viewWebgenOut)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	log.Println("Starting http server at ", urlInbrowser)
	if openNewPage {
		go openBrowser(urlInbrowser)
	}

	log.Fatal(http.ListenAndServe(surl, nil))
}

func execWebgen() error {
	var cmd string
	var args []string
	if runtime.GOOS == "windows" {
		cmd = "cmd"
		args = []string{"/c", "start", fmt.Sprintf("%s\\webgen", Conf.WebGenLocation), "-d", Conf.WebGenWebPageDir}
	} else {
		log.Fatal("OS not recognized")
		return fmt.Errorf("OS not supported %s", runtime.GOOS)
	}
	log.Printf("Exec webgen (source %s) in %s", Conf.WebGenLocation, Conf.WebGenWebPageDir)
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Printf("Error on executing webgen: %v", err)
		return err
	}
	log.Printf("Webgen executed %s\n", out)
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

func EditLastPost(dirIn string, openNewPage bool) {
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

	startEditor(itemName, string(src), openNewPage)
}
