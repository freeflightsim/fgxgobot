Ext.define("GB.FlightPanel", {
	
extend: "Ext.panel.Panel",


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
		autoLoad: true
	});
	Ext.getStore('stoFlight-' + this.Flight.callsign).sort("elapsed", "DESC");
	
	//var Dt = new Date()
	//var rStart = Ext.Date.add(Dt, Ext.Date.MINUTE, -10);
	//var rEnd = Ext.Date.add(Dt, Ext.Date.MINUTE, 10);
	//console.log(Dt, rStart, rEnd);
	Ext.apply(this, {
		border: false, frame: false,
		layout: "border",
		items: [
			{xtype: "fieldset", title: "Info", region: "north"},
			{xtype: "container", region: "center", flex: 3, 
				layout: {type: "border"}, ssheight: 600, title: "Charts",
				items: [
					Ext.create("GB.FlightSpeedChart", {
						flex: 1, ssheight: 200, region: "center",
						//x_store_id: 'stoFlight-' + this.Flight.callsign
						store: Ext.getStore('stoFlight-' + this.Flight.callsign)
					}),
					Ext.create("GB.FlightAltitudeChart", {
						flex: 2, ssheight: 200, region: "north",
						//x_store_id: 'stoFlight-' + this.Flight.callsign
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
				region: "east",
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
	
	
},

fetch_data: function(callsign){
	this.callsign = "29Delta"
	Ext.Ajax.request({
		scope: this, 
		url : "/flight/" + this.callsign,
		params: {
			
		},
		method: 'GET',
		success: function(resp){	
			var data = Ext.decode(resp.responseText);

			console.log(data);
			this.fireEvent("FLIGHT", data);
		},
		failure: function(resp){
			G2.msg("OOP", "Failed");
		}
	});
}

})//= End define