/*
Package flights - Flight Data, Store and definitions

	Flights Data and Store
	==============================================
	
	flight.go - contains the main struct
		* Flight - representaion of a flight
		* Pos    - represents a position in time
		* AjaxFlight - object to spool out ajax record
		* AjaxFlightsPayload - object to spool out ajax Payload
	
	flights_store.go - the data store
		* The FlightsStore is the database
		* This is initiates in xtate.GStateMachine
		* The store is autoloaded every few moments via crossfeed	
	

*/
package flights