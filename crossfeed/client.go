
package crossfeed

import (
    //"fmt"
    //"sort"
   // "time"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

// FGX_CF_FLIGHTS_URL - the upstream url for flights data
const FGX_CF_FLIGHTS_URL = "http://crsossfeed.fgx.ch/data"


//= A Flight Row from ajax crossfeed
type CF_Flight struct {
	Fid uint64 `json:"fid"`
	Callsign string `json:"callsign"`
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	HdgT int `json:"hdg"`
	AltFt int `json:"alt_ft"`
	SpdKt int `json:"spd_kts"`
	Model string `json:"model"`
}

//= Represents the whole Blob from fgx-crossfeed	
type CF_Reply struct {
	Success bool `json:"success"`
	LastUpdated string `json:"last_updated"`
	Source string `json:"source"`
	Flights []CF_Flight `json:"flights"`
	Error string `json:"flights"`
}



//= Make a requests to remote CF server and reply to channel
// - TODO: time the request/respose time
func MakeRequest(c chan CF_Reply) {

	//t_start := time.Now()
	//fmt.Println("CrossFeedClient->make_request() >>")
	
	//= Make http request to remote server
	resp, ereq := http.Get(FGX_CF_FLIGHTS_URL)
	if ereq != nil {
		panic(ereq)
	}
	
	//= Convert the request body, expected json string
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println("body: ",string(body))
	
	//= Now Decode the json into our objects defined
	var cf CF_Reply
	e2 := json.Unmarshal(body, &cf)
	if e2 != nil {
		panic(e2)
	}
	//= Phew that worked
	//fmt.Println("OK", "Got Data",  cf.LastUpdated ) 
	
	//end_t := time.Now()
	//fmt.Println("OK", "Timed",  "start_end", start_t, end_t )
	c <- cf
	
}

