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
			{text: 'Callsign',  dataIndex: 'callsign', width: 100, menuDisabled: true,
				renderer: function(v, meta, rec){
					meta.style = "font-weight: bold;";
					return v;
				}
			},
			{text: 'Alt Ft', dataIndex: 'alt_ft', flex: 1, align: "right", menuDisabled: true},
			{text: 'Spd Kt', dataIndex: 'spd_kt', flex: 1, align: "right", menuDisabled: true},
			{text: 'Hdg True', dataIndex: 'hdg_t', flex: 1, align: "right", menuDisabled: true},
			{text: 'Lat', dataIndex: 'lat', flex: 1, align: "right", menuDisabled: true},
			{text: 'Lon', dataIndex: 'lon', flex: 1, align: "right", menuDisabled: true},
			{text: 'Positions', dataIndex: 'positions_count', width: 100, align: "right", menuDisabled: true},
			{text: 'Model', dataIndex: 'aero', flex: 1, align: "left", menuDisabled: true}
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
