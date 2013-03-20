package www


import (
    "fmt"
    //"time"
   // "io/ioutil"
    "net/http"
	//"net/url"
    "html/template"
   // "encoding/json"
)
import (
	//"github.com/fgx/fgxgobot/radio"
	//"github.com/fgx/fgxgobot/xstate"
)


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