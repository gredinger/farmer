package farmer

import (
	"log"
	"net/http"

	"gopkg.in/ini.v1"

	"github.com/gorilla/mux"
)

//WebApp contains the webserver application
type WebApp struct {
	Router   *mux.Router
	settings *Settings
}

//Settings contains db info/application settings
type Settings struct {
	BindAddress string `ini:"BindAddress"`
}

//LoadSettings pulls settings for the webapp.
func (wa *WebApp) LoadSettings(fn string) {
	fr, err := ini.Load(fn)
	if err != nil {
		log.Fatal(err)
	}
	s := new(Settings)
	err = fr.MapTo(s)
	if err != nil {
		log.Fatal(err)
	}
	wa.settings = s
}

//Run executes a webserver
func (wa *WebApp) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe(wa.settings.BindAddress, r)
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Derp"))
}
