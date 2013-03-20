Ext.define("GB.MpServersPanel", {
	
extend: "Ext.grid.Panel", 

initComponent: function(){
	
	Ext.apply(this, {

		title: "Mp Servers",
		store: Ext.getStore("stoMpServers"),
		viewConfig: {
			emptyText: "No flights in this view",
			deferEmpty: false
		},
		columns: [
			{text: 'No',  dataIndex: 'no', menuDisabled: true, width: 50},
			{text: 'MpServer',  dataIndex: 'subdomain', menuDisabled: true, flex: 1},
			{text: 'Ip', dataIndex: 'ip', flex: 1, menuDisabled: true, flex: 1},
			{text: 'Status', dataIndex: 'status', menuDisabled: true, flex: 1}
		],
		dockedItems: [
			{xtype: 'pagingtoolbar',
				store: Ext.getStore("stoMpServers"), 
				dock: 'bottom', displayInfo: true
			}
		]
	});
	this.callParent();
},


})//= End define
	