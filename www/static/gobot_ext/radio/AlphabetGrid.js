Ext.define("GB.radio.AlphabetGrid", {
	
extend: "Ext.grid.Panel", 

initComponent: function(){
	
		// Flights Store
	Ext.create("Ext.data.JsonStore", {
		model: "GB.model.Alphabet",
		storeId: "stoAlphabet",
		autoLoad: true,
		pageSize: 100,
		proxy: {
			type: "ajax",
			url: "/radio/alphabet",
			reader: {
				type: 'json',
				root: 'alphabet',
				idProperty: 'char'
			}
		}
	});
	Ext.getStore("stoAlphabet").sort("code", "ASC");
	
	Ext.apply(this, {
		store: Ext.getStore("stoAlphabet"),
		viewConfig: {
			emptyText: "No items in this view",
			deferEmpty: false
		},
		columns: [
			{text: 'Char',  dataIndex: 'char', width: 50, menuDisabled: true,
				renderer: function(v, meta, rec){
					meta.style = "font-weight: bold;";
					return v;
				}
			},
			{text: 'Word', dataIndex: 'word', flex: 2, align: "left", menuDisabled: true},
			{text: 'Phonetic', dataIndex: 'phonetic', flex: 1, align: "left", menuDisabled: true},
			
			{text: 'Morse', dataIndex: 'morse', width: 60, align: "right", menuDisabled: true}
			
		]
	});
	this.callParent();
}


})//= End define
