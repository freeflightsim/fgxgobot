
package flights

import (
    //"time"
    //"fmt"
    "path/filepath"
    "strings"
)
import (
    "github.com/fgx/fgxgobot/crossfeed"
    "github.com/fgx/fgxgobot/radio"
)


// No of positions to store
// 
// TODO: this need to be a configurable online ? how ?
// so for now a constant. 
//
// Current Defaut is 180 = 3 min track at 1 req per second
const HISTORY_MAX_POSITIONS = 100  



// Flight is an entry in the FlightsStore and contains data on a flight
//
// The flight track is stored in the Flight.Positions slice and a
// flight will accumulate positions until HISTORY_MAX_POSITIONS. 
// Positions are simply appended on each update.
//
// The current position is the last ie Flight.Positions[len - 1] 
// (unless there is a better way)
// 
// TODO: The flight need some calculations based on positions 
// eg speed trend, vertical speed trend, distance traveled, lookahead etc
type Flight struct {
	Callsign string 
	Model string
	Aero string    
	Positions []*Pos 
}
func (me *Flight) CallsignWords() string {
	return radio.Callsign2Words(me.Callsign) 
}

// A Position with its Ts timestamp
type Pos struct {
	Lat float32   `json:"lat"`
	Lon float32   `json:"lon"`
	HdgT int  `json:"hdg_t"`
	AltFt int  `json:"alt_ft"`
	SpdKt int  `json:"spd_kt"`
	Ts int64 `json:"ts"`
	Elapsed int64   `json:"elapsed"`
}


// UpdatePosition() -inserts an item into the list of Flight.Positions
// but only if the position has changed ie not parked
func (me *Flight) UpdatePosition(fly crossfeed.CF_Flight, ts int64){
	
	// Check if this position same as last position
	if len(me.Positions) > 1 {
		last := me.Position()
		if  last.Lat == fly.Lat  && 
			last.Lon == fly.Lon && 
			last.AltFt == fly.AltFt &&
			last.HdgT == fly.HdgT {
			//fmt.Println("dupe position", fly.Callsign)
			return
		}
	}
	
	//fmt.Println("Not Dupe", fly.Callsign)
	p := new(Pos)
	p.Ts = ts 
	p.Lat = fly.Lat
	p.Lon = fly.Lon 
	p.HdgT = fly.HdgT 
	p.SpdKt = fly.SpdKt 
	p.AltFt = fly.AltFt
	
	// append out position to end of slice
	me.Positions = append(me.Positions, p)
	
	// remove first item if up to Max Positions
	if len(me.Positions) == HISTORY_MAX_POSITIONS + 1 {
		 me.Positions = me.Positions[1:]
	}
}

// Returns *Pos pointer of the current position;  ie last entry in slice 
func (me *Flight) Position() *Pos{
	return me.Positions[ len(me.Positions) - 1 ]
}


// NewFlight() - constructs and returns a new Flight (created from the crossfeed.CF_Flight)
func NewFlight(cffly crossfeed.CF_Flight) *Flight{
 	fly := new(Flight)
    fly.Callsign = cffly.Callsign
    fly.Model = cffly.Model
    aero := filepath.Base(cffly.Model)
    if strings.HasSuffix(aero, ".xml") {
    	fly.Aero = aero[0:len(aero)-4]
    }else{
    	fly.Aero = aero
    }
    
    fly.Positions = make([]*Pos,0, HISTORY_MAX_POSITIONS)
	return fly
}




// Represents a json encoder spooling out complete ajax payload for  /flights
type AjaxFlightsPayload struct {
	Success bool `json:"success"`
	Ts string `json:"ts"`
	Flights []*AjaxFlightRow `json:"flights"`
}



// Represents a json encoder row of a Flight
type AjaxFlightRow struct {
	Callsign string `json:"callsign"`
	CallsignWords string `json:"callsign_words"`
	Model string `json:"model"`
	Aero string `json:"aero"`
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	AltFt int `json:"alt_ft"`
	SpdKt int `json:"spd_kt"`
	HdgT int  `json:"hdg_t"`
	Ts int64 `json:"ts"`  
	PositionsCount int `json:"positions_count"`
}

// Represents a json encoder spooling out complete ajax payload for  /flight/{callsign}
type AjaxFlightPayload struct {
	Success bool `json:"success"`
	Err bool `json:"err"`
	ErrMsg string `json:"err_msg"`
	Flight *AjaxFlightPositions `json:"flight"`
}

type AjaxFlightPositions struct {
	Callsign string `json:"callsign"`
	CallsignWords string `json:"callsign_words"`
	Model string `json:"model"`
	Aero string `json:"aero"`  
	Positions[] *Pos `json:"positions"`
}

//Creates a new AjaxFlight for encoding
func NewAjaxFlightRow(fly *Flight) *AjaxFlightRow{
	rec := new(AjaxFlightRow)
	rec.Callsign = fly.Callsign
	rec.CallsignWords = fly.CallsignWords()
	rec.Model = fly.Model
	rec.Aero = fly.Aero
	lp := fly.Position() // last position 
	rec.Lat = lp.Lat
	rec.Lon = lp.Lon
	rec.AltFt = lp.AltFt
	rec.SpdKt = lp.SpdKt
	rec.HdgT = lp.HdgT
	rec.Ts = lp.Ts
	rec.PositionsCount = len(fly.Positions)
	return rec
}


