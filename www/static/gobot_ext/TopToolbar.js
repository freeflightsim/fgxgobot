Ext.define("GB.TopToolbar", {
	
extend: "Ext.panel.Panel", 

initComponent: function(){
		
	Ext.apply(this, {

		frame: false, plain: true, border: false, hideBorders: true,
		margins: {top:0, right:0, bottom:5, left:0},
		hideHeader: true,
		tbar: [
						
			//== Refresh MP
			{xtype: 'tbspacer', width: 20},
			{text: "Refresh Flights Interval >&nbsp;", tooltip: "Refresh now now",
				handler: function(){
					
				}
			},
			{text:  "Off", iconCls: "icoOff", enableToggle: true,   
				width: this.tbw,   allowDepress: false,
				toggleGroup: "refresh_rate",  refresh_rate: 0,  
				toggleHandler: this.on_refresh_toggled,	scope: this
			},
			{text:  "1" ,  iconCls: "icoOff", enableToggle: true,    
				width: this.tbw, allowDepress: false,
				toggleGroup: "refresh_rate",  refresh_rate: 1, 
				toggleHandler: this.on_refresh_toggled,	scope: this
			},
			{text:  "2", iconCls:  "icoOn", enableToggle: true,    
				width: this.tbw, allowDepress: false, 
				toggleGroup: "refresh_rate",  refresh_rate: 2, 
				toggleHandler: this.on_refresh_toggled,	scope: this
			},
			{text:  "3", iconCls:  "icoOff", enableToggle: true,   
				width: this.tbw, allowDepress: false,  pressed: true,
				toggleGroup: "refresh_rate",  refresh_rate: 3, 
				toggleHandler: this.on_refresh_toggled,	scope: this
			},
			{text:  "6", iconCls:  "icoOff", enableToggle: true,   
				width: this.tbw, allowDepress: false, 
				toggleGroup: "refresh_rate",  refresh_rate: 6,
				toggleHandler: this.on_refresh_toggled,	scope: this
			},
			{text:  "10", iconCls: "icoOff", enableToggle: true,   
				width: this.tbw, allowDepress: false, 
				toggleGroup: "refresh_rate",  refresh_rate: 10, 
				toggleHandler: this.on_refresh_toggled,	scope: this
			},
			"-",
			"->",
		
			//== FlightGear Menu
			"-",
			{text: "FlightGear", iconCls: "icoFlightGear", disabled: true
			
			},
			"-",	
				
			//== FGx Menu
			{text: "FGx", iconCls: "icoFgx", 
				menu: [
					{text: "Issues", url: "https://github.com/fgx/fgx-go/issues",
						handler: this.on_open_url, scope: this,
					},
					{text: "Git View", url: "https://github.com/fgx/fgx-go",
						handler: this.on_open_url, scope: this 
					}
				]
			},
			{xtype: 'tbspacer', width: 10}
						
		]
			
	});
	this.callParent();
},

on_refresh_toggled: function(butt, checked){
	butt.setIconCls( checked ? "icoOn" : "icoOff" );
	//console.log("checked", checked, butt.refresh_rate)
	if(checked){	
		this.fireEvent("REFRESH_RATE", butt.refresh_rate);
	}
}

})//= End define