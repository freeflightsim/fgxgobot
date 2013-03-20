package www


import (
    "fmt"
    //"time"
   // "io/ioutil"
    "net/http"
    "html/template"
   // "encoding/json"
)
import (
	"github.com/fgx/fgxgobot/radio"
	"github.com/fgx/fgxgobot/xstate"
)

//-------------------------------------------------------------------------
// == Ajax handlers ==



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
// Returns callsign as a wors at ajax url= /radio/callsign2words
func Ajax_radio_callsign2words(w http.ResponseWriter, r *http.Request) {
	
	req_callsign := "unix09" // extract from request
	words := radio.Callsign2Words(req_callsign)
	fmt.Println("YES",words)
	fmt.Fprint(w, words)
}




//-------------------------------------------------------------------------
// == Html Handling ==
//-------------------------------------------------------------------------

// A stuct for the templates
// BUG(perdo) - need to get extjs and other configs
type TemplateContext struct{
	Name string
	Addr string
}


// Handles HomePage - returns out the index.html template 
func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	
	context := &TemplateContext{Name: "Foobar", Addr: "localhost:9999"}
	fmt.Println("HOMe_handler()")

	 t := template.New("index.html")
	 t, err := t.ParseFiles("/home/daffodil/gogo/src/www/index.html")
	 if err != nil {
	 	http.Error(w, err.Error(), http.StatusInternalServerError)
	 	return
	 }
	 w.Header().Set("Content-Type", "text/html")
	//index, _ := ioutil.ReadFile("www/index.html")
	//w.Write([]byte(index))
	t.Execute(w, context)

}