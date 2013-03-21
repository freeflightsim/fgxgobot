Ext.define("GB.dev.DevPanel", {
	
extend: "Ext.panel.Panel",

		   
initComponent: function(){
	
	
	Ext.apply(this, {
		title: "Developer",
		border: false, frame: false, padding: 0, margin: 0,
		layout: "fit",
		tbar: [
			{xtype: 'buttongroup',
				title: 'Websocket',
				columns: 3,
				defaults: {
					scale: 'small'
				},
				items: [
					{text: "Connect", handler: this.on_connect, scope: this, iconCls: "icoConnect"},
					{text: "Disconnect", handler: this.on_disconnect, scope: this, iconCls: "icoDisconnect"},
					{xtype: "displayfield", value: "Status: Unknown", width: "100", x_name: "ws_status"}
				]
			}
		],
		items: [
		
			//Ext.create("GB.radio.AlphabetGrid", {})

		]
	});
	this.callParent();

	
},
get_ws: function(){
	var ws = Ext.getCmp("main_viewport").ws();
	console.log("ws=", ws);
	return ws;
},

on_connect: function(butt){
	var ws = this.get_ws();
	console.log("conn", butt, ws);
},
on_disconnect: function(butt){
	console.log("dicconn", butt);
},


});