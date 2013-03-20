Ext.define("GB.Viewport", {
	
extend: "Ext.container.Viewport", 

initComponent: function(){
	
	//=== Initialise the stores
	/*
	Ext.create("Ext.data.JsonStore", {
		model: "Flight",
		storeId: "stoFlights",
		autoLoad: true,
		proxy: {
			type: "ajax",
			url: "/flights",
			reader: {
				type: 'json',
				root: 'flights',
				idProperty: 'callsign' // later FID ?
			}
		},
		
	});
	Ext.getStore("stoFlights").sort("callsign", "ASC");



	Ext.create("Ext.data.JsonStore", {
		model: "MpServer",
		storeId: "stoMpServers",
		autoLoad: true,
		proxy: {
			type: "ajax",
			url: "/mpservers",
			reader: {
				type: 'json',
				root: 'mpservers',
				idProperty: 'no' 
			}
		}
	});
	Ext.getStore("stoMpServers").sort("subdomain", "ASC");
	*/
	
	Ext.apply(this, {
		layout: 'border',
		items: [
			{xtype: "tabpanel", region: "center", 
				items: [
				
					//== Flights Tab
					Ext.create("GB.FlightsGrid", {
						title: "Flights",
					}),
					
					
				]
			}//= center tabs
		]
		
	});
	this.callParent();
}
	
	
});