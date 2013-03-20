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
	
	The Servers are discovered from DNS by looking up the address
*/
package mpserver
