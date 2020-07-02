package farmer

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/ini.v1"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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
	r.HandleFunc("/ws", wa.WebSocketHandler)
	err := http.ListenAndServe(wa.settings.BindAddress, r)
	if err != nil {
		log.Fatal(err)
	}
}

//WebSocketHandler does most of the communication for the game.
func (wa WebApp) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	wu := websocket.Upgrader{}
	c, err := wu.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	pc := PlayerConnection{conn: c}
	go pc.readPump()
	go pc.writePump()
}

//IndexHandler serves the game client
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Derp"))
}
