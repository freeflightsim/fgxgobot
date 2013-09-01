Ext.define("GB.dev.WebSockMessGrid", {
	
extend: "Ext.grid.Panel", 

initComponent: function(){
	

	Ext.apply(this, {
		store: Ext.getStore("stoWebSockMess"),
		viewConfig: {
			emptyText: "No items in this view",
			deferEmpty: false
		},
		columns: [
			{text: 'Status', dataIndex: 'status', flex: 1, align: "left", menuDisabled: true},
			{text: 'Message', dataIndex: 'msg', flex: 3, align: "left", menuDisabled: true},
			
			{text: 'Ts', dataIndex: 'ts', width: 1, align: "right", menuDisabled: true}
			
		]
	});
	this.callParent();
}


})//= End define
