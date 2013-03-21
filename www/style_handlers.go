package www


import (
	 "bytes"
	 "fmt"
	 "net/http"
)

const FAM_FAM_URL = "http://static.freeflightsim.org"
// Represent a Character in the Alphabet with json encoding for output
//type Ico struct{
//	Ico string `json:"char"`
//	File string `json:"word"`
//}/

/*
static_icons['icoHelp'] = "help.png"
static_icons['icoExecute'] = "accept.png"
static_icons['icoHtml'] = "html.png"

static_icons['icoDev'] = "shape_align_bottom.png"
static_icons['icoDatabase'] = "database.png"

static_icons['icoSelectStyle'] = "color_swatch.png"

static_icons['icoLogin'] = "key.png"

static_icons['icoRefresh'] = "refresh.gif"

static_icons['icoOn'] = "bullet_pink.png"
static_icons['icoOff'] = "bullet_black.png"
*/

// A map of the Alphabet definition
var FamFamIcons = map[string]string {

	"icoRefresh": "refresh.gif",
	"icoOff": "bullet_black.png",
	"icoOn": "bullet_pink.png",
	
	"icoConnect": "connect.png",
	"icoDisconnect": "disconnect.png",
	
	"icoDev": "shape_align_bottom.png",
}

var LocalIcons = map[string]string {
	"icoFgx": "fgx-cap-16.png",
	"icoFlightGear": "flightgear_icon.png",
}

func Style_dynamic_css(w http.ResponseWriter, r *http.Request){
	
	var buffer bytes.Buffer
	buffer.WriteString("/* Fam fam icons */\n") 
	
    for ico, file := range FamFamIcons {
		s := fmt.Sprintf(".%s{background-image: url('%s/icons/famfam_silk/%s') !important; background-repeat: no-repeat;}\n", ico, FAM_FAM_URL, file)
		buffer.WriteString(s) 
    }
    buffer.WriteString("\n\n/* Local icons */\n")
    for ico, file := range LocalIcons {
		s := fmt.Sprintf(".%s{background-image: url('/static/icons/%s') !important; background-repeat: no-repeat;}\n", ico, file)
		buffer.WriteString(s) 
    }
    
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprint(w, buffer.String())
}