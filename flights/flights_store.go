package flights

import (
    "fmt"
    "time"
    "encoding/json"
    "sort"
)
import (
    "github.com/fgx/fgxgobot/crossfeed"
)

//--------------------------------------------------------------------
//= FlightsStore is the database for flights
type FlightsStore struct {
	n int 
	Flights map[string]*Flight  
	CfReplyChan chan crossfeed.CF_Reply // ?? help  not sure
}



// NewFlightsStore creates an instance of FlightsStore
func NewFlightsStore() *FlightsStore {
	ob := new(FlightsStore)
	ob.Flights =  make(map[string]*Flight)
	ob.CfReplyChan = make(chan crossfeed.CF_Reply) // HELP ?
	return ob
}


//= Starts timer to fetch flights data from crossfeed every few seconds
func (me *FlightsStore) StartCrossfeedTimer() {

	//= To do Make the interval configurable via web interface or websocket etc
	ticker := time.NewTicker(time.Millisecond * 2000)
	
    go func() {
		for t := range ticker.C {
			_ = t // Silence error
			
			fmt.Println(">> CF Make request >>")
			
			//= Make a remote request
			go crossfeed.MakeRequest(me.CfReplyChan)
			
			//= The reply is the json decoded object crossfeed.CF_Reply
			reply := <- me.CfReplyChan
			fmt.Println("  << Reply from cf: flights =", len(reply.Flights) ) 
			
			// get the date (source = "last_updated":"2013-03-18 21:33:43")
			ts, _ := time.Parse("2006-01-02 15:04:05", reply.LastUpdated )
			
			for _, cffly := range reply.Flights{
				
				//= Check if this Flight is in the map already
				xfly, ok := me.Flights[cffly.Callsign]
				//fmt.Println("xfly=", xfly, ok)
				
				if !ok{
					//= Flight not in map so create a new one
					xfly = NewFlight(cffly)
    				//xfly.Callsign = cffly.Callsign
    				//xfly.Positions = make([]*XPos,0, HISTORY_MAX_POSITIONS)
					me.Flights[cffly.Callsign] = xfly
				}
				
				xfly.UpdatePosition(cffly, ts)  // Update this flights position
				if cffly.Callsign == "Hermi" {
					fmt.Println(xfly.Positions)
					fmt.Println(xfly.Position())
				}
				// TODO << Send out message (help)
			}
		}
	}()
}


//= GetAjaxPayload - spools out the flights as json string 
//  this is send to client whether ajax request or websocket
func (me *FlightsStore) GetAjaxPayload() string {

	var pay = new(AjaxFlightsPayload)
    pay.Success = true // for extjs
    pay.Flights = make([]*AjaxFlight,0)
    
	// TODO Would be nice to sort properly by map Key But LONG
	sort_idx := make([]string, 0, len(me.Flights))
    for i := range me.Flights {
        sort_idx = append(sort_idx, i)
    }
    sort.Strings(sort_idx)
    for _, i := range sort_idx {
    	var ajF = NewAjaxFlight(me.Flights[i])
    	pay.Flights = append(pay.Flights, ajF)
    }
    
    s, _ := json.MarshalIndent(pay, "" , "  ")
    //s, _ := json.Marshal(pay)
    return string(s)
}


//= This was a Previous attempt using the http.Handle("/flights", this)
/*
func (me *FlightsStore) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    
    var pay = new(AjaxFlightsPayload)
    pay.Success = true
    pay.Flights = make([]*AjaxFlight,0)
    
    for _, ele := range me.Flights {
    	var ajF = NewAjaxFlight(ele)
    	pay.Flights = append(pay.Flights, ajF)
    }
    
    s, _ := json.Marshal(pay)
    
  	
  	w.Header().Set("Content-Type", "application/json")
  	fmt.Fprint(w, string(s))
}
*/
