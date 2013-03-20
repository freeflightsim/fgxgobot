

//== Flight Model
Ext.define('GB.model.Flight', {
    extend: 'Ext.data.Model',
    fields: [
        {name: 'fid',  type: 'string', sortType: "asText"},
        {name: 'callsign',  type: 'string', sortType: "asUCText"},
		
		{name: 'model',  type: 'string', sortType: "asUCText"},
		{name: 'aero',  type: 'string', sortType: "asUCText"},
		
        {name: 'alt_ft',   type: 'int',sortType: "asInt"},
        {name: 'spd_kt', type: 'int', sortType: "asInt"},
        {name: 'hdg_t', type: 'int', sortType: "asInt"},
        
		{name: 'lat',  type: 'string'},
        {name: 'lon',  type: 'string'},
		{name: 'positions_count',  type: 'int',sortType: "asInt"},
    ]
});




//== MpServer Model 
Ext.define('GB.model.MpServer', {
    extend: 'Ext.data.Model',
    fields: [
        {name: 'no',  type: 'int', sortType: "asInt"},
        {name: 'ip',   type: 'string', convert: null, sortType: "asUCText"},
        {name: 'subdomain', type: 'string'},
        {name: 'status', type: 'string'}
    ]
});

//== MpServer Model 
Ext.define('GB.model.Position', {
    extend: 'Ext.data.Model',
    fields: [
        {name: 'lat',  type: 'string',},
        {name: 'lon',   type: 'string'},
        {name: 'alt_ft', type: 'int', sortType: "asInt"},
        {name: 'spd_kt', type: 'int', sortType: "asInt"},
		{name: 'ts', type: 'date'}
    ]
});