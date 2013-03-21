
package main

import (
    "fmt"
   // "time"
    "net/http"
)
import (   
  //	"code.google.com/p/go.net/websocket"
	
  	"github.com/gorilla/mux"
  
    "github.com/fgx/fgxgobot/www"
    "github.com/fgx/fgxgobot/xstate"
   	
   	//"xwebsocket"
)


func main() {

	//---------------------------------------------------
	//= Initialize the xtate.GStateMachine :-)
	gsm := xstate.InitializeStateMachine()
	
	// TODO setup config with file ? cmd ? help ? best practises
	//  > set option to false, check master config and load, check last config and load
	//  > check command line to override all previous
	gsm.Debug = true
	
	gsm.Start()


	

	router := mux.NewRouter()
	//router.Methods("GET", "POST")
	
	//= HELP ? this is not working, even touch full path
	// Am expection localhost:9999/static/REAME.txt to appear
	
	
	//router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("/home/gogo/src/github.com/fgx/fgxgobot/www/static")) ))
	router.PathPrefix("/static/").Handler( http.StripPrefix("/static", http.FileServer(http.Dir("/home/gogo/src/github.com/fgx/fgxgobot/www/static"))) )

	router.HandleFunc("/flights", www.Ajax_flights)
	router.HandleFunc("/flight/{callsign}", www.Ajax_flight)
	
	router.HandleFunc("/mpservers", www.Ajax_mpservers)
	
	
	router.HandleFunc("/radio/alphabet", www.Ajax_radio_alphabet)
	router.HandleFunc("/radio/callsign2words", www.Ajax_radio_callsign2words)
	
	//http.Handle("/ws", websocket.Handler(xwebsocket.WsHandler))
	router.HandleFunc("/dynamic.css", www.Style_dynamic_css)
	http.Handle("/", router)
	
	//http.HandleFunc("/crossfeed", crossfeed_handler)
    if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
	
	
	fmt.Println("end main")
}
