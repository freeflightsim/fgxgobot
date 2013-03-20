
package flights

import (
    "time"
    //"fmt"
)
import (
    "github.com/fgx/fgxgobot/crossfeed"
)


// No of positions to store
// 
// 	TODO: this need to be a configurable online ? how
//  	  so for now a constant
//        Defaut is 180 = 3 min track at 1 req per second
const HISTORY_MAX_POSITIONS = 180  



// Flight is an entry in the FlightsStore and contains data on a flight
// The flight track is stored in the XFlight.Positions slice
// A flight will accumulate positions until HISTORY_MAX_POSITIONS
// Current position should be XFlight.Positions[0] unless there is a better way
// 
//		TODO: The flight need some calculations based on positions
//	          eg speed trend, vertical speed trend, distance traveled, lookahead etc
type Flight struct {
	Callsign string 
	Model string  
	Positions []*Pos 
}

// A Position
type Pos struct {
	Lat float32 
	Lon float32 
	HdgT int
	AltFt int
	SpdKt int
	Ts time.Time 
}


// UpdatePosition() -  inserts an item into the list of XFlights.Positions
// but only if the position has changed ie not parked
func (me *Flight) UpdatePosition(fly crossfeed.CF_Flight, ts time.Time){
	
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


//--------------------------------------------------------------------
// NewXFlight() - constructs and returns a new Flight (created from a crossfeed.CF_Flight)
func NewFlight(cffly crossfeed.CF_Flight) *Flight{
 	xfly := new(Flight)
    xfly.Callsign = cffly.Callsign
    xfly.Model = cffly.Model
    xfly.Positions = make([]*Pos,0, HISTORY_MAX_POSITIONS)
	return xfly
}




// Represents a json encoder spooling out complete ajax payload for  /flights
type AjaxFlightsPayload struct {
	Success bool `json:"success"`
	Ts string `json:"ts"`
	Flights []*AjaxFlight `json:"flights"`
}

// Represents a json encoder row of a Flight
type AjaxFlight struct {
	Callsign string `json:"callsign"`
	Model string `json:"model"`
	Aero string `json:"aero"`
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	AltFt int `json:"alt_ft"`
	SpdKt int `json:"spd_kt"`
	HdgT int  `json:"hdg_t"`
	Ts string `json:"ts"`  // ISO eg 2015-12-25 00:00:01.123 << Yes we may need milisec
	PositionsCount int `json:"positions_count"`
}

//= Creates a new Flight from poisiton etc
// TODO This will need to change to maybe returning a list of last x positions
//      or other tricks required by a map or a radar interface, overview etc
func NewAjaxFlight(xfly *Flight) *AjaxFlight{
	rec := new(AjaxFlight)
	rec.Callsign = xfly.Callsign
	rec.Model = xfly.Model
	
	lp := xfly.Position() // last position 
	rec.Lat = lp.Lat
	rec.Lon = lp.Lon
	rec.AltFt = lp.AltFt
	rec.SpdKt = lp.SpdKt
	rec.HdgT = lp.HdgT
	rec.Ts = lp.Ts.Format("2006-01-02 15:04:05")
	rec.PositionsCount = len(xfly.Positions)
	return rec
}


