package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/sinakeshmiri/simpleGoBNB/pkg/config"
	"github.com/sinakeshmiri/simpleGoBNB/pkg/handlers"
	"github.com/sinakeshmiri/simpleGoBNB/pkg/render"
	"log"
	"net/http"
	"time"
)

var AppConfig config.Config
var session *scs.SessionManager

func main() {

	AppConfig.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = AppConfig.InProduction
	AppConfig.Session = session
	AppConfig.TemplatesCache, _ = render.CreateTemplateCache()
	AppConfig.UseCache = false
	repo := handlers.NewRepo(&AppConfig)
	handlers.NewHandlers(repo)
	render.NewTemplates(&AppConfig)
	//http.HandleFunc("/", handlers.Repo.HomePage)
	//http.HandleFunc("/about", handlers.Repo.AboutPage)
	//_ = http.ListenAndServe(":8484", nil)
	srv := &http.Server{
		Addr:    ":8181",
		Handler: route(&AppConfig),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

//e:=errors.New("F")
