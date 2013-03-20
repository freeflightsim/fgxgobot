Ext.define("GB.FlightsPanel", {
	
extend: "Ext.panel.Panel", 

initComponent: function(){
	
	Ext.apply(this, {
		title: "Flights",
		border: false, frame: false,
		layout: "border",
		items: [
			Ext.create("GB.FlightsGrid", {
				region: "center", flex: 3
			}),
			Ext.create("GB.FlightsAltitudeChart", {
				region: "east", flex: 2, DEADheight: window.innerHeight - 50
			})
		]
	});
	this.callParent();
},


})//= End define
