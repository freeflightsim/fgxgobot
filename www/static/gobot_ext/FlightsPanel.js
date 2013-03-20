Ext.define("GB.FlightsPanel", {
	
extend: "Ext.panel.Panel", 

initComponent: function(){
	
	Ext.apply(this, {
		title: "Flights",
		border: false, frame: false,
		layout: "border",
		items: [
			Ext.create("GB.FlightsGrid", {
				region: "center", flex: 3, x_name: "flights_grid"
			}),
			Ext.create("GB.FlightsAltitudeChart", {
				region: "east", flex: 2, DEADheight: window.innerHeight - 50
			})
		]
	});
	this.callParent();
	
	this.down("[x_name=flights_grid]").on("OPEN_FLIGHT", function(fly){
		//console.log("flightspanel.OPEN_FLIGHT", fly);
		this.fireEvent("OPEN_FLIGHT", fly);
	}, this);
},


})//= End define
