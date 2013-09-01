/*global Ext: false, console: false, FGx: false */


Ext.define("GB.MapPanel", {

extend: "GeoExt.panel.Map", 

	
//===========================================================
initComponent: function() {
	
	var mapX = new OpenLayers.Map({});
	
	var wms = new OpenLayers.Layer.WMS(
		"OpenLayers WMS",
		"http://vmap0.tiles.osgeo.org/wms/vmap0?",
		{layers: 'basic'}
	);
	
	var context = {
          getColor: function(feature) {
                if (feature.attributes.elevation < 2000) {
                    return 'green';
                }
                if (feature.attributes.elevation < 2300) {
                    return 'orange';
                }
                return 'red';
            }
        };
        var template = {
            cursor: "pointer",
            fillOpacity: 0.5,
            fillColor: "${getColor}",
            pointRadius: 5,
            strokeWidth: 1,
            strokeOpacity: 1,
            strokeColor: "${getColor}",
            graphicName: "triangle"
        };
        var style = new OpenLayers.Style(template, {context: context});
        var vecLayer = new OpenLayers.Layer.Vector("vector", {
            styleMap: new OpenLayers.StyleMap({
                'default': style
            }),
            protocol: new OpenLayers.Protocol.HTTP({
                url: "../data/summits.json",
                format: new OpenLayers.Format.GeoJSON()
            }),
            strategies: [new OpenLayers.Strategy.Fixed()]
        });
		
	mapX.addLayers([wms, vecLayer]);
	
	//var ll;
	//if( config.lat && config.lon){
	//	ll =  new OpenLayers.Geometry.Point(config.lon, config.lat
	//		).transform(this.get_display_projection(), this.get_map().getProjectionObject() ); 
	//	ll.xFlag = "  SET:";
	//}else{
		ll = new OpenLayers.LonLat(939262.20344, 5938898.34882);
	//	ll.xFlag = "DEAFAUT: "
	//}

	Ext.apply(this, {
		title: 'The GeoExt.panel.Map-class',
            map: mapX,
            center: '12.3046875,51.48193359375',
            zoom: 6,
            stateful: true,
            stateId: 'mappanel',
//            extent: '12.87,52.35,13.96,52.66',
            dockedItems: [{
                xtype: 'toolbar',
                dock: 'top',
                items: [{
                    text: 'Current center of the map',
                    handler: function(){
                        var c = GeoExt.panel.Map.guess().map.getCenter();
                        Ext.Msg.alert(this.getText(), c.toString());
                    }
                }]
            }]
     });
	this.callParent();
	
}, // << initComponent	


set_base_layer: function(layer_name){
	var layer = this.map.getLayersByName(layer_name)[0];
	this.map.setBaseLayer( layer );
},


pan_to: function(obj){
	var lonLat = new OpenLayers.LonLat(obj.lon, obj.lat
			).transform(this.get_display_projection(),  this.get_map().getProjectionObject() );
	this.map.setCenter(lonLat, zoom);
	
},

show_blip: function(obj){
	
	this.L.blip.removeAllFeatures();
	if(!obj){
		return;
	}
	var lonLat = new OpenLayers.LonLat(obj.lon, obj.lat
		).transform(new OpenLayers.Projection("EPSG:4326"), new OpenLayers.Projection("EPSG:3857"));
	this.map.panTo( lonLat );
	
	var pt =  new OpenLayers.Geometry.Point(obj.lon, obj.lat
				).transform(new OpenLayers.Projection("EPSG:4326"), new OpenLayers.Projection("EPSG:3857") );	
	var circle = OpenLayers.Geometry.Polygon.createRegularPolygon(
		pt,
			0, // wtf. .I want a larger cicle
			20
	);
	var style = {
		strokeColor: "red",
		strokeOpacity: 1,
		strokeWidth: 4,
		fillColor: "yellow",
		fillOpacity: 0.8 
	};
	var feature = new OpenLayers.Feature.Vector(circle, null, style);
	this.L.blip.addFeatures( [feature] );
},




});
