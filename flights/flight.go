package flights

import (
    "time"
    //"fmt"
)
import (
    "github.com/fgx/fgxgobot/crossfeed"
)
//= XFlight - representaion of a flight
//= XPos    - represents a position in time
//= AjaxFlight - object to spool out ajax record
//= AjaxFlightsPayload - object to spool out ajax Payload

//= No of positions to store
//= TODO: this need to be a configurable online ? how
//        so for now a constant
const HISTORY_MAX_POSITIONS = 5 //180  // 3 mins at 1 req per second


//--------------------------------------------------------------------
//= XFlight is an entry in the FlightsStore and contains data on a flight
// The flight track is stores in the  Positions in a slice
// A flight will accumulate positions until HISTORY_MAX_POSITIONS
// Current position should be XFlight.Positions[0] unless there is a better way
// TODO: The flight need some calculations based on positions
//       eg speed trend, vertical speed trend, distance traveled, lookahead etc

type XFlight struct {
	Callsign string 
	Model string  
	Positions []*XPos 
}

//= Position
type XPos struct {
	Lat float32 
	Lon float32 
	HdgT int
	AltFt int
	SpdKt int
	Ts time.Time 
}


//= UpdatePosition inserts an item into the list of positions
func (me *XFlight) UpdatePosition(fly crossfeed.CF_Flight, ts time.Time){
	
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
	p := new(XPos)
	p.Ts = ts 
	p.Lat = fly.Lat
	p.Lon = fly.Lon
	p.HdgT = fly.HdgT
	p.SpdKt = fly.SpdKt
	p.AltFt = fly.AltFt
	
	// append out position to end of slice
	me.Positions = append(me.Positions, p)
	
	//= Remove first item if up to Max Positions
	if len(me.Positions) == HISTORY_MAX_POSITIONS + 1 {
		 me.Positions = me.Positions[1:]
	}
}
//= Return current position (last entry in slice)
func (me *XFlight) Position() *XPos{
	return me.Positions[ len(me.Positions) - 1 ]
}

//--------------------------------------------------------------------



//--------------------------------------------------------------------
//= NewXFlights from Crossfeed flight
func NewXFlight(cffly crossfeed.CF_Flight) *XFlight{
 	xfly := new(XFlight)
    xfly.Callsign = cffly.Callsign
    xfly.Model = cffly.Model
    xfly.Positions = make([]*XPos,0, HISTORY_MAX_POSITIONS)
	return xfly
}



//--------------------------------------------------------------------
//= Structs for spooling out Ajax

//= The Complete payload sent
type AjaxFlightsPayload struct {
	Success bool `json:"success"`
	Ts string `json:"ts"`
	Flights []*AjaxFlight `json:"flights"`
}

//= A Flight Row
type AjaxFlight struct {
	Callsign string `json:"callsign"`
	Model string `json:"model"`
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	AltFt int `json:"alt_ft"`
	SpdKt int `json:"spd_kt"`
	HdgT int  `json:"hdg_t"`
	Ts string `json:"ts"`  // ISO eg 2015-12-25 00:00:01.123 << Yes we may need milisec
}

//= Creates a new Flight from poisiton etc
// TODO This will need to change to maybe returning a list of last x positions
//      or other tricks required by a map or a radar interface, overview etc
func NewAjaxFlight(xfly *XFlight) *AjaxFlight{
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
	return rec
}


