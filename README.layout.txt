=====================
Basic Structure
=====================

/crossfeed/ 
	Remote ajax client for data from "crossfeed" (ie UDP to Ajax)
	http://crossfeed.fgx.ch/data << returns all data of flights on mpnet
	
	
/flights/
	Stuff to do with Flights, in Flight Object and the 
	"FlightsStore" which is the "object" DB to hold current sessions
	
/mpnet/
	Stuff to so with the MultiPlyer Network, Inc the MpServer Object
	MPServersStore is an "object" that stores server data in memory
	MpServers are discovred via dns lookup 
	ie mpserver01.flightgear.org .. mpserver02, 03 .. to MAX_DNS_NO currently 50
	
/radio/
	Stuff to do with radio telephonony (callsigns etc).. WIP
	
/www/
	Stuff to do with the front end website, inc templates and handlers.go
	
/xstate/
	Contains the "StateMachine" and the main "App object"
	with instances of FlightsStore and MpServersStore
	 
	
/websocket/
	Stuff to do with the websocket server
	
	 	 
	
	 