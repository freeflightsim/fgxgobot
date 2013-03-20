/*
Package www - Deals with web front end

	Web Front End
	======================================================
	
	The Front End requests for 
		* ajax and json 
		* html pages
		* static stuff eg js
		* This is all dodgy and under testing
		
	The handlers.go file spools out the main stuff atmo
	
	As an initial naming convertion then ajax is:
	- www.Ajax_foo where foo == http://server.com/foo
	- www.Ajax_foo_bar would be == http://server.com/foo_bar
	- eg http.HandleFunc("/radio/callsign2words", www.Ajax_radio_callsign2words)
	
	The www directory contains:
	* ajax_handlers.go 
		- for ajax related stuff, Urls such as 
			/mpservers
			/flights
			/radio/alphabet
			/radio/callsign2words?callsign=FG129
		- ajax is returned with Same-Origin-Policy disabled = everyone
		- Every ajax request returns {success: true} as extjs required this
			
	* handlers.go 
		- for the others stuff eg web pages
	          
	* /static/ 
		- for static stuff such at css and js
		- NOTE: This stuff does not all atmo
*/
package www