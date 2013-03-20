/*
Package crossfeed - Remote Client for crossfeed.fgx.ch

	================================================================
	About FGx Crossfeed
	================================================================

	The client aqquires remote flight data from Geoff's CrossFeed fgx-cf 
	- The crossfeed is at http://crossfeed.fgx.ch/data
	- This listens on the UDP wire to network traffic on the UDP crossfeed
	- It then converts these to ajax/json packets on request
	- Note: 
		FGx replies with the Same Origin Policy * ie disabled
	
	Links: 
		Feed:     http://crossfeed.fgx.ch/data
		Project:  http://fgx.ch/projects/fgx-cf
		Code:     https://gitorious.org/fgtools/crossfeed
		More info in mpserver package http://godoc.org/github.com/fgx/fgxgobot/mpserver

	

*/
package crossfeed