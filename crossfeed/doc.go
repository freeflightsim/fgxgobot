/*
Package crossfeed - Remote Client for crossfeed.fgx.ch

About FGx Corssfeed

The client aqquires remote flight data from Geoff's CrossFeed fgx-cf

The crossfeed is a listener to the USP network and converts 
a snapshot to json serves via ajax requests (see ../mpserver)

	Feed:     http://crossfeed.fgx.ch/data
	Project:  http://fgx.ch/projects/fgx-cf
	Code:     https://gitorious.org/fgtools/crossfeed

Note: FGx replies with the Same Origin Policy * ie disabled

*/
package crossfeed