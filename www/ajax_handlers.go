package www


import (
    "fmt"
    //"time"
   // "io/ioutil"
    "net/http"
	//"net/url"
   // "encoding/json"
)
import (
	"github.com/fgx/fgxgobot/radio"
	"github.com/fgx/fgxgobot/xstate"
)

//-------------------------------------------------------------------------
// == Ajax handlers ==

func SetAjaxHeaders(w http.ResponseWriter){
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-FOO", "BAR")
}

// Return flights data at ajax url = /flights 
func Ajax_flights(w http.ResponseWriter, r *http.Request) {
	
	s := xstate.GStateMachine.Flights.GetAjaxPayload()
	fmt.Fprint(w, s)
}

// Returns mpservers data ajax url = /mpservers 
func Ajax_mpservers(w http.ResponseWriter, r *http.Request) {
	
	s := xstate.GStateMachine.MpServers.GetAjaxPayload()
	fmt.Fprint(w, s)
}

// Returns radio alphabet data ajax url = /radio/alphabet
func Ajax_radio_alphabet(w http.ResponseWriter, r *http.Request) {
	
	s := radio.GetAjaxAlphabet()
	fmt.Println("YES",s)
	fmt.Fprint(w, s)
}

// Returns callsign as a words at ajax url= /radio/callsign2words
func Ajax_radio_callsign2words(w http.ResponseWriter, r *http.Request) {
	
	
	callsign := r.URL.Query().Get("callsign")
	//TODO - Throw error is its blank
	
	payload := radio.GetAjaxCallsign2Words(callsign)
	
	fmt.Fprint(w, payload)
}


