package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/aaaasmile/vido-preproc/conf"
	"github.com/aaaasmile/vido-preproc/util"
	"github.com/aaaasmile/vido-preproc/web/srv"
	"github.com/kardianos/service"
)

func RunService(cr <-chan struct{}, logger service.Logger, configfile string) error {
	if logger == nil {
		logger = service.ConsoleLogger
	}

	conf.ReadConfig(util.GetFullPath(configfile))
	log.Println("Configuration is read")
	if conf.Current.RootURLPattern == "" {
		log.Fatal("RootURLPattern is not defined")
	}
	if err := srv.InitFromConfig(conf.Current.DebugVerbose); err != nil {
		log.Println("Error from InitFromConfig: ", err)
		return err
	}

	log.Println("Prepare service init")
	var wait time.Duration
	serverurl := conf.Current.ServiceURL
	protoHtt := "http"

	serviceURL := fmt.Sprintf("%s://%s%s", protoHtt, strings.Replace(serverurl, "0.0.0.0", "localhost", 1), conf.Current.RootURLPattern)
	serviceURL = strings.Replace(serviceURL, "127.0.0.1", "localhost", 1)
	logger.Infof("Server started with URL %s", serverurl)
	log.Println("Try this url: ", serviceURL)

	http.Handle(conf.Current.RootURLPattern+"static/", http.StripPrefix(conf.Current.RootURLPattern+"static", http.FileServer(http.Dir(util.GetFullPath("static")))))

	log.Println("Root pattern ", conf.Current.RootURLPattern)
	http.HandleFunc(conf.Current.RootURLPattern, srv.HandleIndex)

	srv := &http.Server{
		Addr: serverurl,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      nil,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Server is not listening anymore: ", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt) //We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	log.Println("Enter in server loop")
loop:
	for {
		select {
		case <-sig:
			log.Println("stop because interrupt")
			break loop
		case <-cr:
			log.Println("stop because service shutdown")
			break loop
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Bye, service")

	return nil
}

func init() {
	//message.SetString(language.German, "Build redirect with JWT", "Umleitung mit JWT bauen")
}