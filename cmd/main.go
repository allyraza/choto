package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/allyraza/choto"
)

func main() {
	config := &choto.Config{}

	flag.StringVar(&config.Address, "address", ":8080", "Address to listen on.")
	flag.Parse()

	app := choto.New(config)

	tmpl := make(map[string]*template.Template)

	tmpl["index.html"] = template.Must(template.ParseFiles("tpls/index.html", "tpls/base.html"))

	tmpl["about.html"] = template.Must(template.ParseFiles("tpls/about.html", "tpls/base.html"))

	app.Mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		tmpl["index.html"].ExecuteTemplate(w, "base.html", struct{}{})
	})

	app.Mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		tmpl["about.html"].ExecuteTemplate(w, "base.html", struct{}{})
	})

	log.Printf("Starting choto on %v...\n", config.Address)

	server := http.Server{Addr: config.Address, Handler: app}
	go func() {
		server.ListenAndServe()
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	log.Println("Shutdown signal received, quiting...")
	server.Shutdown(context.Background())
}
