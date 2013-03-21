Ext.define("GB.FlightsGrid", {
	
extend: "Ext.grid.Panel", 

initComponent: function(){
	
	Ext.apply(this, {
		store: Ext.getStore("stoFlights"),
		viewConfig: {
			emptyText: "No flights in this view",
			deferEmpty: false
		},
		columns: [
			{text: 'Callsign',  dataIndex: 'callsign', width: 70, menuDisabled: true,
				renderer: function(v, meta, rec){
					meta.style = "font-weight: bold;";
					return v;
				}
			},
			{text: 'Radio', dataIndex: 'callsign_words', flex: 2, align: "left", menuDisabled: true},
			{text: 'Aero', dataIndex: 'aero', width: 80, align: "left", menuDisabled: true},
			
			{text: 'Alt Ft', dataIndex: 'alt_ft', width: 60, align: "right", menuDisabled: true},
			{text: 'Spd Kt', dataIndex: 'spd_kt', width: 50, align: "right", menuDisabled: true},
			{text: 'Hdg T', dataIndex: 'hdg_t', width: 50, align: "right", menuDisabled: true},
		   
			{text: 'Lat', dataIndex: 'lat', width: 80, align: "right", menuDisabled: true},
			{text: 'Lon', dataIndex: 'lon', width: 80, align: "right", menuDisabled: true},
		   
			{text: 'Pos', dataIndex: 'positions_count', width: 30, align: "right", menuDisabled: true},
			
		],
		dockedItems: [
			{xtype: 'pagingtoolbar',
				store: Ext.getStore("stoFlights"), 
				dock: 'bottom', displayInfo: true
			}
		],
		listeners: {
			scope: this,
			itemdblclick: function(grid, rec, item, idx, e, eOpts){
				this.fireEvent("OPEN_FLIGHT", rec.getData());
			}			
		}
	});
	this.callParent();
}


})//= End define
