Ext.define("GB.radio.RadioPanel", {
	
extend: "Ext.panel.Panel",

		   
initComponent: function(){
	
	
	Ext.apply(this, {
		title: "Radio",
		border: false, frame: false, padding: 0, margin: 0,
		layout: "fit",
		items: [
		
			Ext.create("GB.radio.AlphabetGrid", {})

		]
	});
	this.callParent();

	
}

});