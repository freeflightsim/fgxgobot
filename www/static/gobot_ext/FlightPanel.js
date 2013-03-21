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
	
	//var Dt = new Date()
	//var rStart = Ext.Date.add(Dt, Ext.Date.MINUTE, -10);
	//var rEnd = Ext.Date.add(Dt, Ext.Date.MINUTE, 10);
	//console.log(Dt, rStart, rEnd);
	Ext.apply(this, {
		border: false, frame: false,
		layout: "border",
		items: [
			{xtype: "container", region: "north", layout: "hbox", 
				padding: 5, defaults: {margin: 2, padding: 0},
				items: [
					{xtype: "fieldset", title: "Heading", flex: 1,
						items: [
							{xtype: "displayfield", hideLabel: true, name: "hdg_t", value: 0}
						]
					},
					{xtype: "fieldset", title: "Altitude Ft", flex: 1,
						items: [
							{xtype: "displayfield", hideLabel: true, name: "alt_ft", value: 0}
						]
					},
					{xtype: "fieldset", title: "Speed Kt", flex: 1,
						items: [
							{xtype: "displayfield", hideLabel: true, name: "spd_kt", value: 0}
						]
					}
				]
			},
			{xtype: "container", region: "center", flex: 3, 
				layout: {type: "border"}, ssheight: 600, title: "Charts",
				items: [
					Ext.create("GB.FlightSpeedChart", {
						flex: 1,  region: "center",
						store: Ext.getStore('stoFlight-' + this.Flight.callsign)
					}),
					Ext.create("GB.FlightAltitudeChart", {
						flex: 2,  region: "north",
						store: Ext.getStore('stoFlight-' + this.Flight.callsign)
					}),
		   			
					/*Ext.create('Ext.chart.Chart', {
						flex: 1,
						xtype: 'chart',
						style: 'background:#fff',
						animate: true,
						store: Ext.getStore('stoFlight-' + this.Flight.callsign),
						shadow: true,
						theme: 'Category1',
						legend: {
							position: 'right'
						},
						axes: [{
							type: 'Numeric',
							minimum: 0,
							position: 'left',
							fields: ['alt_ft', 'speed_kt'],
							title: 'Number of Hits',
							minorTickSteps: 1,
							grid: {
								odd: {
									opacity: 1,
									fill: '#ddd',
									stroke: '#bbb',
									'stroke-width': 0.5
								}
							}
						}, {
							type: 'Category',
							position: 'bottom',
							fields: ['idx'],
							title: 'Positions',
							step: 1
						}],
						series: [{
							type: 'line',
							highlight: {
								size: 7,
								radius: 7
							},
							axis: 'left',
							xField: 'idx',
							yField: 'alt_ft',
							markerConfig: {
								type: 'cross',
								size: 4,
								radius: 4,
								'stroke-width': 0
							}
						}, {
							type: 'line',
							highlight: {
								size: 7,
								radius: 7
							},
							axis: 'left',
							smooth: true,
							xField: 'idx',
							yField: 'spd_kt',
							markerConfig: {
								type: 'circle',
								size: 4,
								radius: 4,
								'stroke-width': 0
							}
						} 
						]
					}), */  //= End chart
				]
			},
		
			//= Positions Grid
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
			})
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
	console.log("fetch data");
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
			Ext.getStore('stoFlight-' + this.Flight.callsign).loadData(data.positions);
			var last = data.positions[ data.positions.length - 1];
			//console.log("last", last);
			this.down("[name=hdg_t]").setValue(last.hdg_t);
			this.down("[name=alt_ft]").setValue(last.alt_ft);
			this.down("[name=spd_kt]").setValue(last.spd_kt);
			//this.fireEvent("FLIGHT", data);
		},
		failure: function(resp){
			//console.log("OOP", "Failed");
		}
	});
}

})//= End define