
package main

import (
    "fmt"
   // "time"
    "net/http"
)
import (   
	//"os"
  	"github.com/gorilla/mux"
  
   	//"code.google.com/p/go.net/websocket" 	
   	
    "github.com/fgx/fgxgobot/www"
    "github.com/fgx/fgxgobot/xstate"
   // "github.com/fgx/fgxgobot/xwebsocket"
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

	//go xwebsocket.WsHubb.Run()
	

	router := mux.NewRouter()
	
	//router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("/home/gogo/src/github.com/fgx/fgxgobot/www/static")) ))
	router.PathPrefix("/static/").Handler( http.StripPrefix("/static",  http.FileServer(http.Dir("/home/gogo/src/github.com/fgx/fgxgobot/static"))) )
    
    //pwd, _ := os.Getwd()
    //router.Static("/static", pwd)
    //router.PathPrefix("/static").Handler(http.FileServer(http.Dir("/home/gogo/src/github.com/fgx/fgxgobot/static/")))
    //router.PathPrefix("/js").Handler(http.FileServer(http.Dir("./static/js")))

	//http.Handle("/ws", websocket.Handler(xwebsocket.WsHandler))
    
    
	router.HandleFunc("/flights", www.Ajax_flights)
	router.HandleFunc("/flight/{callsign}", www.Ajax_flight)
	
    router.HandleFunc("/mpservers.json", www.Ajax_mpservers)
	
	
	router.HandleFunc("/radio/alphabet", www.Ajax_radio_alphabet)
	router.HandleFunc("/radio/callsign2words", www.Ajax_radio_callsign2words)
	
	
	router.HandleFunc("/dynamic.01.css", www.Style_dynamic_css)
    
    router.HandleFunc("/", www.Html_home_page)
	http.Handle("/", router)
	
	//http.HandleFunc("/crossfeed", crossfeed_handler)
    if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
	
	
	fmt.Println("end main")
}
