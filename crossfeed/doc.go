/*
Package crossfeed - Remote Client for crossfeed.fgx.ch

About FGx Crossfeed

The client aqquires remote flight data from Geoff's CrossFeed fgx-cf (see the mpserver package http://godoc.org/github.com/fgx/fgxgobot/mpserver)
The crossfeed is a listener to the USP network and converts 
a snapshot to json serves via ajax requests 


	Feed:     http://crossfeed.fgx.ch/data
	Project:  http://fgx.ch/projects/fgx-cf
	Code:     https://gitorious.org/fgtools/crossfeed

Note: FGx replies with the Same Origin Policy * ie disabled

*/
package crossfeed