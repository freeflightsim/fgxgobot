/*
Package flights - Flight Data, Store and definitions

	Flights Data and Store
	==============================================
	
	* flight.go contains the main struct
		* XFlight - representaion of a flight
		* XPos    - represents a position in time
		* AjaxFlight - object to spool out ajax record
		*  AjaxFlightsPayload - object to spool out ajax Payload
	
	* flights_store.go
		* The main FlightsStore is the database
		* The store is autoloaded every few moments via crossfeed	
	

*/
package flights