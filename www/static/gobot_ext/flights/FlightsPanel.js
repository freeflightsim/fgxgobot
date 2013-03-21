Ext.define("GB.flights.FlightsPanel", {
	
extend: "Ext.panel.Panel", 

initComponent: function(){
	
	Ext.apply(this, {
		title: "Flights",
		border: false, frame: false,
		layout: "border",
		items: [
			Ext.create("GB.flights.FlightsGrid", {
				region: "center", width: 500, x_name: "flights_grid"
			}),
			/*Ext.create("GB.FlightsAltitudeChart", {
				region: "east", flex: 2, DEADheight: window.innerHeight - 50
			})*/
			Ext.create("GB.map.MapPanel", {
				region: "east", flex: 3, x_name: "map_panel"
			})
		]
	});
	this.callParent();
	
	this.down("[x_name=flights_grid]").on("OPEN_FLIGHT", function(fly){
		//console.log("flightspanel.OPEN_FLIGHT", fly);
		this.fireEvent("OPEN_FLIGHT", fly);
	}, this);
	Ext.getStore("stoFlights").on("load", function(sto, recs){
			//console.log("xFlightsStore.load");
			this.down("[x_name=map_panel]").show_flights(recs);
			
	}, this);
},


})//= End define
