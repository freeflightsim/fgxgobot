Ext.define("GB.FlightPanel", {
	
extend: "Ext.panel.Panel",

//= Refresh flight stuff
refresh_rate: 5,
runner:  Ext.create("Ext.util.TaskRunner", {}),
		   
initComponent: function(){
	
	Ext.create('Ext.data.Store', {
		storeId: 'stoFlight-' + this.Flight.callsign,
		model: "GB.model.Position",
		proxy: {
			type: 'ajax',
			url: "/flight/" + this.Flight.callsign,
			reader: {
				type: 'json',
				root: 'positions'
			}
		},
		pageSize: 1000,
		autoLoad: false
	});
	Ext.getStore('stoFlight-' + this.Flight.callsign).sort("elapsed", "ASC");
	
	Ext.apply(this, {
		border: false, frame: false, padding: 0, margin: 0,
		layout: "border", 
		items: [
			{xtype: "panel", region: "north", layout: "hbox", padding: 0, margin: 0,
				title: "=", c_name: "callsign_words",
				padding: 5, defaults: {margin: 2, padding: 0},
				
				items: [
					{xtype: "fieldset", title: "Heading", flex: 1,
						items: [
							{xtype: "displayfield", hideLabel: true, c_name: "hdg_t", value: 0}
						]
					},
					{xtype: "fieldset", title: "Altitude Ft", flex: 1,
						items: [
							{xtype: "displayfield", hideLabel: true, c_name: "alt_ft", value: 0}
						]
					},
					{xtype: "fieldset", title: "Speed Kt", flex: 1,
						items: [
							{xtype: "displayfield", hideLabel: true, c_name: "spd_kt", value: 0}
						]
					}
				]
			},
			{xtype: "panel", region: "center", flex: 3, 
				dockedItems: [
					{xtype: 'pagingtoolbar',
						store: Ext.getStore('stoFlight-' + this.Flight.callsign),
						dock: 'bottom', displayInfo: true
					}
				],
				layout: {type: "border"}, 
				items: [
					Ext.create("GB.FlightSpeedChart", {
						flex: 1,  region: "center",
						store: Ext.getStore('stoFlight-' + this.Flight.callsign)
					}),
					Ext.create("GB.FlightAltitudeChart", {
						flex: 2,  region: "north",
						store: Ext.getStore('stoFlight-' + this.Flight.callsign)
					})
				]
			},
		
			//= Positions Grid
			/*
			Ext.create('Ext.grid.Panel', {
				title: 'Positions', flex: 1,
				region: "east", width: 300,
				store: Ext.getStore('stoFlight-' + this.Flight.callsign),
				columns: [
					{text: 'Elap', dataIndex: 'elapsed', menuDisabled: true, width: 50},
					//{text: 'Ts', dataIndex: 'ts', xtype: 'datecolumn', format:'H:i:s', menuDisabled: true},
					{text: 'Alt Ft',  dataIndex: 'alt_ft', flex: 1, menuDisabled: true, align:"right"},
					{text: 'Spd Kt', dataIndex: 'spd_kt', flex: 1, menuDisabled: true, align:"right"},
					{text: 'Lat', dataIndex: 'lat', flex: 1, menuDisabled: true},
					{text: 'Lon', dataIndex: 'lon', flex: 1, menuDisabled: true},
					
				],
				dockedItems: [
					{xtype: 'pagingtoolbar',
						store: Ext.getStore('stoFlight-' + this.Flight.callsign),
						dock: 'bottom', displayInfo: true
					}
				]
			})*/
		] 
	});
	this.callParent();
	
	this.runner.start({
		interval: this.refresh_rate * 1000,
		run: function(){
			this.fetch_data();
		},
		scope: this		
	});
	
},

fetch_data: function(){
	//console.log("fetch data");
	Ext.Ajax.request({
		scope: this, 
		url : "/flight/" + this.Flight.callsign,
		params: {
			
		},
		method: 'GET',
		scope: this,
		success: function(resp){	
			var data = Ext.decode(resp.responseText);

			//console.log(data);
			var poss = data.flight.positions;
			Ext.getStore('stoFlight-' + this.Flight.callsign).loadData(poss);
			var last = poss[ poss.length - 1];
			//console.log("last", last);
			this.down("[c_name=callsign_words]").setTitle(data.flight.callsign_words);
			
			this.down("[c_name=hdg_t]").setValue(last.hdg_t);
			this.down("[c_name=alt_ft]").setValue(last.alt_ft);
			this.down("[c_name=spd_kt]").setValue(last.spd_kt);
			//this.fireEvent("FLIGHT", data);
		},
		failure: function(resp){
			//console.log("OOP", "Failed");
		}
	});
}

})//= End define