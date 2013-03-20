// FGx-go - Fun with golang and FlightGear network

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
    "github.com/FreeFlightSim/ffsgobot/www"
    "github.com/FreeFlightSim/ffsgobot/xstate"
   	//"xwebsocket"
)


//== Lets Gooooooooooooooooo! ;-)
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
	http.Handle("/static/", http.FileServer(http.Dir("/home/daffodil/gogo/src/www/static")))
	
	
	//= Mad but initially trying to use the naming convention of
	// www.Ajax_foo where foo == http://server.com/foo
	// www.Ajax_foo_bar would be == http://server.com/foo_bar
	http.HandleFunc("/flights", www.Ajax_flights)
	http.HandleFunc("/mpservers", www.Ajax_mpservers)
	
	//http.HandleFunc("/radio/callsign", www.Ajax_radio_callsign)
	http.HandleFunc("/radio/alphabet", www.Ajax_radio_alphabet)
	
	
	//http.Handle("/ws", websocket.Handler(xwebsocket.WsHandler))
	http.HandleFunc("/", www.HandleHomePage)
	
	//http.HandleFunc("/crossfeed", crossfeed_handler)
    if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
	
	
	fmt.Println("end main")
}
