
package main

import (
    "fmt"
   // "time"
    "net/http"
)
import (   
  //	"code.google.com/p/go.net/websocket"
	
   // "flights"
  //  "mpserver"
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

	
	//= HELP ? this is not working, even touch full path
	// Am expection localhost:9999/static/REAME.txt to appear
	http.Handle("/", http.StripPrefix("/static/", http.FileServer(http.Dir("/home/gogo/src/github.com/fgx/fgxgobot/www/static")) ))
	
	

	http.HandleFunc("/flights", www.Ajax_flights)
	http.HandleFunc("/mpservers", www.Ajax_mpservers)
	
	
	http.HandleFunc("/radio/alphabet", www.Ajax_radio_alphabet)
	http.HandleFunc("/radio/callsign2words", www.Ajax_radio_callsign2words)
	
	//http.Handle("/ws", websocket.Handler(xwebsocket.WsHandler))
	//http.HandleFunc("/", www.HandleHomePage)
	
	//http.HandleFunc("/crossfeed", crossfeed_handler)
    if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
	
	
	fmt.Println("end main")
}
