package handlers

import (
	"github.com/sinakeshmiri/simpleGoBNB/pkg/config"
	"github.com/sinakeshmiri/simpleGoBNB/pkg/models"
	"github.com/sinakeshmiri/simpleGoBNB/pkg/render"

	"net/http"
)

type Repository struct {
	App *config.Config
}

var Repo *Repository

func NewRepo(app *config.Config) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomePage(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_IP", remoteIP)
	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) AboutPage(writer http.ResponseWriter, request *http.Request) {

	var aboutTemplateDataMap = make(map[string]string)
	aboutTemplateDataMap["content"] = "test"
	remoteIP := m.App.Session.GetString(request.Context(), "remote_IP")
	aboutTemplateDataMap["remote_IP"] = remoteIP
	var aboutTemplateData = models.TemplateData{}
	aboutTemplateData.TextData = aboutTemplateDataMap
	render.RenderTemplate(writer, "about.page.tmpl", &aboutTemplateData)

}

/*
func AboutPage(writer http.ResponseWriter, request *http.Request) {
	n, err := fmt.Fprintf(writer, "it's so simple in GO ! , it's the about page !")
	if err != nil {
		fmt.Println("something went wrong! 500")
	} else {
		fmt.Println(fmt.Sprintf("whoopsie goopsi here is number of bytes written: %d", n))
	}
}
*/
