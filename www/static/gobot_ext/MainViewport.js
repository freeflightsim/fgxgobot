Ext.define("GB.MainViewport", {
	
extend: "Ext.container.Viewport", 

//= Refresh flight stuff
refresh_rate: 0,
runner:  Ext.create("Ext.util.TaskRunner", {}),

id: "main_viewport",

webSock: null,


initComponent: function(){
	
	
	//=== Initialise the stores  and sort ====
	
	// Flights Store
	Ext.create("Ext.data.JsonStore", {
		model: "GB.model.Flight",
		storeId: "stoFlights",
		autoLoad: true,
		pageSize: 100,
		proxy: {
			type: "ajax",
			url: "/flights",
			reader: {
				type: 'json',
				root: 'flights',
				idProperty: 'callsign' // later FID ?
			}
		}
	});
	Ext.getStore("stoFlights").sort("callsign", "ASC");

	// MpServers store
	Ext.create("Ext.data.JsonStore", {
		model: "GB.model.MpServer",
		storeId: "stoMpServers",
		autoLoad: true,
		proxy: {
			type: "ajax",
			url: "/mpservers",
			reader: {
				type: 'json',
				root: 'mpservers',
				idProperty: 'no' 
			}
		}
	});
	Ext.getStore("stoMpServers").sort("subdomain", "ASC");
	
	
	
	// Ws Store
	Ext.create("Ext.data.Store", {
		model: "GB.model.WebSockMess",
		storeId: "stoWebSockMess",
		//autoLoad: true,
		pageSize: 100,
		//proxy: {
			//type: "ajax",
			//url: "/flights",
			//reader: {
			//	type: 'json',
			//	root: 'flights',
				//idProperty: 'callsign' // later FID ?
			//}
		//}
	});
	//Ext.getStore("stoFlights").sort("callsign", "ASC");
	
	
	//==== MainViewport for app ===
	Ext.apply(this, {
		layout: 'border',
		items: [
			Ext.create("GB.TopToolbar", {region: "north", x_name: "top_toolbar"}),
			{xtype: "tabpanel", region: "center", x_name: "tab_panel", 
				header: false, 
				items: [
					//== Flights Tab
					Ext.create("GB.dev.DevPanel", {
						x_name: "dev_panel"
					}),
					//== Flights Tab
					Ext.create("GB.flights.FlightsPanel", {
						x_name: "flights_panel"
					}),
					//== MpServers Tab
					//Ext.create("GB.mpservers.MpServersPanel", {
					//}),
					//== Radio Tab
					//Ext.create("GB.radio.RadioPanel", {
					//})
					
				]
			}//= center tabs
		]
		
	});
	this.callParent();
	
	// Setup trigger of toolbar button to the runner
	this.down("[x_name=top_toolbar]").on("REFRESH_RATE", function(rate){
		
		this.refresh_rate = rate;
		
		//= stop any runners.. but this will/might callback though.. does not cancel sent request..
		this.runner.stopAll();
		
		//= start again with new rate..
		if(this.refresh_rate > 0){
			this.runner.start({
				interval: this.refresh_rate * 1000,
				run: function(){
					Ext.getStore("stoFlights").load()
					
				},
				scope: this		
			});
		}
	}, this);
	
	// Setup trigger of toolbar button to the runner
	this.down("[x_name=flights_panel]").on("OPEN_FLIGHT", function(fly){
		console.log("MainVireport.OPEN_FLIGHT", fly);
		var d = Ext.create("GB.flight.FlightPanel", {
			Flight: fly, title: fly.callsign, closable: true
		});
		var tabPanel = this.down("[x_name=tab_panel]");
		tabPanel.add(d);
		tabPanel.setActiveTab(d);
	}, this);
	
},

//=========================================================
ws: function(){
	console.log("viewport.ws()", this.webSock);
	if( !this.webSock && "WebSocket" in window){
		console.log("MAKE websock");
		
		var Vp = Ext.getCmp("main_viewport");
		this.webSock = new WebSocket(WS_URL);
		this.ws_add_log("Created", "New Ws Creted")
		this.webSock.onclose = function(evt) {
            //appendLog($("<div><b>Connection closed.</b></div>"))
            this.ws_add_log("Close", "Close event");
        }
        this.webSock.onmessage = function(evt) {
            console.log(">> ws >>", evt.data);
			Vp.ws_add_log("Mess", evt.data);
        }
        
	}
	return this.webSock;
},

ws_add_log: function(status, mess){
	console.log("add_log", status, mess);
	Ext.getStore("stoWebSockMess").add( {status: status, msg: mess, ts: new Date()} );
}
	
	
});