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

type SiteInfoStruct struct{
    SiteLong string
    SiteShort string
    APP_VER string
    CDN_URL string    
}

var GSiteInfo = &SiteInfoStruct{SiteLong: "FlightGear Multiplayer Map - golang experimental", 
    SiteShort: "fg-gomap", APP_VER: "1.0",
    CDN_URL: "http://static.freeflightsim.org"}
    
    
type Page struct{
    Title string
    Foo string
    SiteInfo *SiteInfoStruct
}
    
    
// A stuct for the templates
// BUG(perdo) - need to get extjs and other configs
type TemplateContext struct{
	Name string
	Addr string
}

var HomeTemplates = template.Must(template.ParseFiles("templates/_base.html", "templates/map-extjs4.html"))

// Handles HomePage - returns out the index.html template 
func Html_home_page(w http.ResponseWriter, r *http.Request) {
	
	//context := &TemplateContext{Name: "Foobar", Addr: "localhost:9999"}
	fmt.Println("HOMe_handler()")

    page := &Page{Title: "Welcome", Foo: "Bar", SiteInfo: GSiteInfo}
    if err := HomeTemplates.Execute(w, page); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

