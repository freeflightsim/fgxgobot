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
	Ext.getStore('stoFlight-' + this.Flight.callsign).sort("ts", "DESC");
	
	
	Ext.apply(this, {
		border: false, frame: false,
		layout: "border",
		items: [
			Ext.create('Ext.grid.Panel', {
				title: 'Positions',
				region: "center",
				store: Ext.getStore('stoFlight-' + this.Flight.callsign),
				columns: [
					{text: 'Ts', dataIndex: 'ts', xtype: 'datecolumn', format:'H:i:s', menuDisabled: true},
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
				],
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