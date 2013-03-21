
var WS_URL = "ws://localhost:9999/ws";

//== Flight Model
Ext.define('GB.model.Alphabet', {
    extend: 'Ext.data.Model',
    fields: [
        {name: 'char',  type: 'string', sortType: "asText"},
        {name: 'word',  type: 'string', sortType: "asUCText"},
		{name: 'phonetic',  type: 'string', sortType: "asUCText"},
		{name: 'morse',  type: 'string'}
    ]
});


//== Flight Model
Ext.define('GB.model.Flight', {
    extend: 'Ext.data.Model',
    fields: [
        {name: 'fid',  type: 'string', sortType: "asText"},
        {name: 'callsign',  type: 'string', sortType: "asUCText"},
		{name: 'callsign_words',  type: 'string', sortType: "asUCText"},
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
		{name: 'elapsed', type: 'int', sortType: "asInt"},
        {name: 'lat',  type: 'string',},
        {name: 'lon',   type: 'string'},
        {name: 'alt_ft', type: 'int', sortType: "asInt"},
        {name: 'spd_kt', type: 'int', sortType: "asInt"},
		{name: 'ts', type: 'int', sortType: "asInt"}
    ]
});