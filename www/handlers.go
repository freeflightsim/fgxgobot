/*
Package www - Deals with all web front end http site

<ul>
<li>ajax and json response</li>
<li>html pages</li>
<li>static stuff eg js</li>
</ul>

*/
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
//-------------------------------------------------------------------------


// Handles ajax url = /flights 
func Ajax_flights(w http.ResponseWriter, r *http.Request) {
	
	s := xstate.GStateMachine.Flights.GetAjaxPayload()
	fmt.Fprint(w, s)
}

// Handles ajax url = /mpservers 
func Ajax_mpservers(w http.ResponseWriter, r *http.Request) {
	
	s := xstate.GStateMachine.MpServers.GetAjaxPayload()
	fmt.Fprint(w, s)
}

// Handles ajax url = /radio/dialect 
func Ajax_radio_alphabet(w http.ResponseWriter, r *http.Request) {
	
	s := radio.GetAjaxAlphabet()
	fmt.Println("YES",s)
	fmt.Fprint(w, s)
}
// url= /radio/callsign
func Ajax_radio_callsign(w http.ResponseWriter, r *http.Request) {
	
	req_callsign := "unix09" // extract from request
	words := radio.Callsign2Words(req_callsign)
	fmt.Println("YES",words)
	fmt.Fprint(w, words)
}

//-------------------------------------------------------------------------
// == Html Handling ==
//-------------------------------------------------------------------------
type TemplateContext struct{
	Name string
	Addr string
}


// HomePage - Rattles out the index.html template 
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