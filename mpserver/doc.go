/* 
Package mpserver - The FlightGear Multiplayer Servers

	=====================================================
	About FlightGear Multiplayer Server (fgms) 	
	=====================================================
	
	The Mp Servers are running instances of fgms
	* Newbie: http://wiki.flightgear.org/Howto:Multiplayer
	* Info:   http://wiki.flightgear.org/FlightGear_Multiplayer_Server
	* Code:   http://gitorious.org/fgms/fgms-0-x/trees/master
	
	Pilots flying in the FlightGear sim connect to a running instance and speak
	to each other with udp packets.
	
	The server addresses are in the flightgear.org TLD are are numbered
	- mpserver01.flightgear.org to mpserverNN.flightgear.org
	- currently its up to around mpserver20.flightgear.org
	- The Servers are discovered from DNS by looking up the address
	
	A running fgms has the following "net" interfaces
	* port 5000 - this is the UDP socket for sim and relay
	* port 5001 - is the admin port this returns info on state
	              eg telnet mpserver14.flightgear.org 5001
	* crossfeed - these are udp packets "crossfeed" to another channel eg 5555
	            - Note: 
					The http://godoc.org/github.com/fgx/fgxgobot/crossfeed uses this  
*/
package mpserver
