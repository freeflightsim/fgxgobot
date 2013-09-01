


Ext.define("FGx.mpnet.MpStatusGrid", {

extend: "Ext.grid.Panel",
tbw: 35,

//= Triggered when a refresh toolbar button is clicked
DEADon_refresh_toggled: function(butt, checked){
	butt.setIconCls( checked ? "icoOn" : "icoOff");
	if(checked){
		this.runner.stopAll(); // stop if already ruinning
		this.refresh_rate = parseInt(butt.ref_rate, 10);
		if(this.refresh_rate === 0){
			//this.runner.stop()
		}else{
			this.runner.start( { run: this.update_flights, interval: this.refresh_rate * 1000 });
		}
	}
},

//= Riggered for reshresh now
DEADon_refresh_now: function(){
	this.store.load();
},





	
//===========================================================
//== Renderers 
// @todo: pete to put in css
DEADrender_callsign: function(v, meta, rec){
	return "<b>" + v + "</b>";
},

DEADrender_altitude: function(v, meta, rec){
	return Ext.util.Format.number(v, '00,000');	
},


//== Grid
initComponent: function() {
	
	Ext.apply(this, {
		iconCls: 'icoMpServers',
		fgxType: "MpStatusGrid",
		autoScroll: true,
		autoWidth: true,
		enableHdMenu: false,
		viewConfig: {
			emptyText: 'No servers in view', 
			deferEmptyText: false,
			forceFit: true
		}, 
		stripeRows: true,
        store: Ext.getStore("mpservers_store"),
		loadMask: false,
		
		columns: [ 
            {header: 'No',  dataIndex:'no', sortable: true, width: 40,  menuDisabled: true},
			{header: 'Domain',  dataIndex:'domain', sortable: true, menuDisabled: true, flex: 2},
            {header: 'IP Address',  dataIndex:'ip', sortable: true,  menuDisabled: true},

			{header: 'Last Checked', dataIndex:'last_checked', sortable: true, 
                align: 'right',  menuDisabled: true
			},
			{header: 'Last Seen', dataIndex:'last_seen', sortable: true, 
                align: 'right',  menuDisabled: true
			},
           
           {header: 'Last Telnet', dataIndex:'last_telnet', sortable: true, 
                align: 'right',  menuDisabled: true},
           {header: 'Telnet Lag', dataIndex:'telnet_lag', sortable: true, 
               align: 'right',  menuDisabled: true,
				renderer: function(v){
					if(v > 0){
						return v;
					}
					return "-";
				}
			},
           {header: 'UDP Lag', dataIndex:'udp_lag', sortable: true, 
               align: 'right',  menuDisabled: true,
                renderer: function(v){
                    if(v > 0){
                        return v;
                    }
                    return "-";
                }
           },
			{header: 'Country',  dataIndex:'country', sortable: true, hidden: false,
				width: 100,
			},
			{header: 'Time Zone',  dataIndex:'time_zone', sortable: true, hidden: false,
				width: 100,
			},
           {header: 'Last Error',  dataIndex:'last_err_msg', sortable: true, hidden: false,
				width: 100,
			},
		],
		
		bbar: [
			new Ext.PagingToolbar({
				//frame: false, plain: true, 
                store: Ext.getStore("mpservers_store"),
				displayInfo: true,
				pageSize: 500,
				prependButtons: true	
			})
		]
	});
	this.callParent();
}, // initComponent

});

