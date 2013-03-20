package www


import (
    "fmt"
    "strings"
   // "io/ioutil"
    "net/http"
	//"net/url"
    "encoding/json"
)
import (

	"github.com/gorilla/mux"
	
	"github.com/fgx/fgxgobot/radio"
	"github.com/fgx/fgxgobot/xstate"
)

//-------------------------------------------------------------------------
// == Ajax handlers ==

// Returns a payload with  {error: "My Error", success: true}
func GetErrorPayload(err string) string{

	payload := map[string]interface{}{
        "error": err,
        "success":  true, //This is a quirk for extjs
    }
	s, _ := json.MarshalIndent(payload, "" , "  ")
	return string(s)
}

func SetAjaxHeaders(w http.ResponseWriter){
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Origin", "*") // cache control ?
	//w.Header().Set("Access-Control-Allow-Origin", "*") // no cache ?
}

// Return flights data at ajax url = /flights 
func Ajax_flights(w http.ResponseWriter, r *http.Request) {
	
	s := xstate.GStateMachine.Flights.GetAjaxFlightsPayload()
	fmt.Fprint(w, s)
}

// Return flight data at ajax url = /flight/{callsign}
func Ajax_flight(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	callsign := vars["callsign"]
	s := xstate.GStateMachine.Flights.GetAjaxFlightPayload(callsign)
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

// Returns callsign as a words at ajax url= /radio/callsign2words?callsign=Foo123
func Ajax_radio_callsign2words(w http.ResponseWriter, r *http.Request) {
	
	
	callsign := r.URL.Query().Get("callsign")
	
	callsign = strings.TrimSpace(callsign)
	
	if callsign == "" {
		fmt.Fprint(w, GetErrorPayload("Need a ?callsign=") )
		return
	}
	//TODO - Throw error is its blank
	
	payload := radio.GetAjaxCallsign2Words(callsign)
	
	SetAjaxHeaders(w)
	fmt.Fprint(w, payload)
}


