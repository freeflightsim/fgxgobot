Ext.define("GB.FlightsPanel", {
	
extend: "Ext.panel.Panel", 

initComponent: function(){
	
	Ext.apply(this, {
		title: "Flights",
		border: false, frame: false,
		layout: "hbox",
		items: [
			Ext.create("GB.FlightsGrid", {flex: 3}),
			Ext.create("GB.FlightsAltitudeChart", {layout: "fit", flex: 2, height: window.innerHeight - 50})
		]
	});
	this.callParent();
},


})//= End define
