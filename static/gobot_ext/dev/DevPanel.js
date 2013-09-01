Ext.define("GB.dev.DevPanel", {
	
extend: "Ext.panel.Panel",

		   
initComponent: function(){
	
	
	Ext.apply(this, {
		title: "Developer",
		border: false, frame: false, padding: 0, margin: 0,
		layout: "fit",
		tbar: [
			{xtype: 'buttongroup',
				title: 'WebSock Actions',
				columns: 2,
				defaults: {
					scale: 'small'
				},
				items: [
					{text: "Connect", handler: this.on_connect, scope: this, iconCls: "icoConnect"},
					{text: "Disconnect", handler: this.on_disconnect, scope: this, iconCls: "icoDisconnect"}
				]
			},
		   	{xtype: 'buttongroup',
				title: 'WebSock Status',
				columns: 1,
				defaults: {
					scale: 'small'
				},
				items: [
					{xtype: "displayfield", value: "Status: Unknown", width: "100", x_name: "ws_status"}
				]
			},
		   	{xtype: 'buttongroup',
				title: 'WebSock Test',
				columns: 2,
				defaults: {
					scale: 'small'
				},
				items: [
					{xtype: "textfield", value: "Hello"},
					{xtype: "button", text: "Send", iconCls: "icoSend", scope: this, 
						handler: function(){
							this.get_ws().send("Hellp");
						}
					}
				]
			}
		],
		items: [
		
			Ext.create("GB.dev.WebSockMessGrid", {})

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