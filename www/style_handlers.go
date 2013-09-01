package www


import (
	 "bytes"
	 "fmt"
	 "net/http"
)

const FAM_FAM_URL = "http://static.freeflightsim.org"




var LocalIcons = map[string]string {
    "icoFgx": "fgx-cap-16.png",
    "icoFlightGear": "flightgear_icon.png",
    
    "icoAirport": "apt.png",
    "icoFix": "vfr_fix.png",
    "icoNdb": "ndb.16.png",
    "icoVor": "vor.png",
    "icoClr": "go.gif",
}

var FamFamIcons = map[string]string {
    
    "icoAirways": "chart_line.png",
    "icoFlightPlans": "page_white_actionscript.png",
    
    "icoHelp": "help.png",
    "icoExecute": "accept.png",
    "icoHtml": "html.png",
    
    "icoDev": "shape_align_bottom.png",
    "icoDatabase": "database.png",
    
    "icoSelectStyle": "color_swatch.png",
    
    "icoLogin": "key.png",
    
    "icoRefresh": "refresh.gif",
    
    "icoOn": "bullet_pink.png",
    "icoOff": "bullet_black.png",
    
    
    "icoBookMarkAdd": "book_add.png",
    
    "icoSettings": "cog.png",
    
    "icoCallSign": "page_white_c.png",
    
    
    "icoFlights": "text_horizontalrule.png",
    
    "icoMapCore": "map.png",
    "icoMap": "map.png",
    "icoMapAdd": "map_add.png",
    "icoMapGo": "map_go.png",
    
    "icoMpServers": "server_database.png",
    
    "icoBlue": "bullet_blue.png",
    "icoOrange": "bullet_orange.png",
    "icoPink": "bullet_pink.png",
    "icoGreen": "bullet_green.png",
    "icoRed": "bullet_red.png",
    "icoWhite": "bullet_white.png",
    "icoYellow": "bullet_yellow.png",
    
    
    "icoUsers": "group.png",
    "icoUser": "user.png",
    "icoUserAdd": "user_add.png",
    "icoUserEdit": "user_edit.png",
    "icoUserDelete": "user_delete.png",
    
    
    
    "icoCancel": "bullet_black.png",
    "icoSave": "accept.png",
    
    
    
    
    "icoRefreshStop": "clock_stop.png",
    "icoRefreshRun": "clock_run.png",
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