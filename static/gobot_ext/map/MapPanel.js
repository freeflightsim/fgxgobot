/*global Ext: false, console: false, FGx: false */


Ext.define("GB.map.MapPanel", {

extend: "GeoExt.panel.Map", 

L: {},	
last_pos: null,

//===========================================================
initComponent: function() {
	
	var mapX = new OpenLayers.Map({});
	
	var wms = new OpenLayers.Layer.WMS(
		"OpenLayers WMS",
		"http://vmap0.tiles.osgeo.org/wms/vmap0?",
		{layers: 'basic'}
	);
	
	
	this.L.blip = new OpenLayers.Layer.Vector("HighLight Markers");
	this.L.track = new OpenLayers.Layer.Vector("Track Lines Layer");	

	this.L.radarBlip = new OpenLayers.Layer.Vector(
		"Radar Markers", 
		{styleMap: new OpenLayers.StyleMap({
				"default": {
					strokeColor: "lime",
					strokeWidth: 1,
					fillColor: "lime",

					externalGraphic: "/static/images/radar_blip2.png",
					graphicWidth: 8,
					graphicHeight: 24,
					graphicOpacity: 1,
					graphicXOffset: 0,
					graphicYOffset: -20,
					
					fontColor: "black",
					fontSize: "12px",
					fontFamily: "Helvetica, Arial, sans-serif",
					fontWeight: "bold",
					rotation : "${hdg}"
				},
				"select": {
					fillColor: "black",
					strokeColor: "yellow",
					pointRadius: 12,
					fillOpacity: 1
				}
			})
		}, {  visibility: true}
	);

	this.L.radarLbl =  new OpenLayers.Layer.Vector(
		"Radar Label", 
		{
			styleMap:  new OpenLayers.StyleMap({
				"default": {
					fill: true,
					fillOpacity: 1,
					fillColor: "black",
					strokeColor: "green",
					strokeWidth: 1,

					//graphic: false,
					externalGraphic: "/static/images/fgx-background-black.png",
					graphicWidth: 50,
					graphicHeight: 12,
					graphicOpacity: 0.8,
					graphicXOffset: "${gxOff}",
					graphicYOffset: "${gyOff}",
					
					
					fontColor: "white",
					fontSize: "10px",
					fontFamily: "sans-serif",
					fontWeight: "bold",
					labelAlign: "left",
					labelXOffset: "${lxOff}", 
					labelYOffset: "${lyOff}", 
					label : "${callsign}"
					//rotation : "${planerotation}",

				},
				"select": {
					fillColor: "black",
					strokeColor: "yellow",
					pointRadius: 12,
					fillOpacity: 1
				}

			})
		}
	);
	mapX.addLayers([wms, this.L.blip, this.L.track, this.L.radarBlip, this.L.radarLbl]);
	
	//var ll;
	//if( config.lat && config.lon){
	//	ll =  new OpenLayers.Geometry.Point(config.lon, config.lat
	//		).transform(this.get_display_projection(), this.get_map().getProjectionObject() ); 
	//	ll.xFlag = "  SET:";
	//}else{
		//ll = new OpenLayers.LonLat(939262.20344, 5938898.34882);
	//	ll.xFlag = "DEAFAUT: "
	//}

	Ext.apply(this, {
		map: mapX,
		center: '12.3046875,51.48193359375',
		zoom: 4,
		stateful: true,
		stateId: 'mappanel',
		tbar: this.get_toolbar()
		//            extent: '12.87,52.35,13.96,52.66',
		/* dockedItems: [{
			xtype: 'toolbar',
			dock: 'top',
			items: [{
				text: 'Current center of the map',
				handler: function(){
					var c = GeoExt.panel.Map.guess().map.getCenter();
					Ext.Msg.alert(this.getText(), c.toString());
				}
			}]
		}]*/
     });
	this.callParent();
	

	
	
}, // << initComponent	

get_toolbar: function(){
	var a = [];
	a.push({
            xtype: 'buttongroup',
            title: 'Tracking',
            columns: 4,
            defaults: {
                scale: 'small'
            },
            items: [
				{text: "None", toggleHandler: this.on_zoom_track, scope: this, zoo: "none", enableToggle: true, toggleGroup: "zoom", pressed: true},
				{text: "Zoom Track", toggleHandler: this.on_zoom_track, scope: this, zoo: "track", enableToggle: true, toggleGroup: "zoom"},
				{text: "Zoom Curr", toggleHandler: this.on_zoom_track, scope: this, zoo: "curr", enableToggle: true, toggleGroup: "zoom"},
				{text: "Zoom Close", toggleHandler: this.on_zoom_track, scope: this, zoo: "close", enableToggle: true, toggleGroup: "zoom"}
			]
	});
	return a;
},

set_base_layer: function(layer_name){
	var layer = this.map.getLayersByName(layer_name)[0];
	this.map.setBaseLayer( layer );
},

on_zoom_track: function(butt, checked){
	console.log(butt.text, butt.zoo, checked);
	//this.last_pos
	
},
on_zoom_curr: function(butt){
	console.log(butt.text);
	this.pan_to(this.last_pos, 10)
},
on_zoom_close: function(butt){
	console.log(butt.text);
	this.pan_to(this.last_pos, 15)
},
pan_to: function(obj, zoom){
	var lonLat = new OpenLayers.LonLat(obj.lon, obj.lat
			).transform(new OpenLayers.Projection("EPSG:4326"),  this.map.getProjectionObject() );
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

show_flights: function(recs){
	this.L.radarBlip.removeAllFeatures();
	this.L.radarLbl.removeAllFeatures();
	var i, r;
	var rec_len = recs.length;
	for(i=0; i < rec_len; i++){
		r = recs[i].data;
		this.show_radar(r.callsign, r.lat, r.lon, r.hdg_t, r.alt_ft);
	}
			
},
show_flight: function(fly){
	//console.log("show_flight", r);
	this.L.radarBlip.removeAllFeatures();
	this.L.radarLbl.removeAllFeatures();

	var r = fly.positions[ fly.positions.length - 1];
	this.last_pos = r;
	console.log("show_flight", r, fly);
	//var rec_len = recs.length;
	//for(i=0; i < rec_len; i++){
	//	r = recs[i].data;
	this.show_radar(fly.callsign, r.lat, r.lon, r.hdg_t, r.alt_ft);
	//this.map.setCenter
	//this.pan_to(r);
	this.show_positions(fly.positions);
	//}
			
},
show_positions: function(tracks){
	
	this.L.track.removeAllFeatures();
	var trk_length = tracks.length;
	var points = [];
	//var points;
	var p;
	for(var i =0; i < trk_length; i++){
		p = tracks[i];
		points.push(
				new OpenLayers.Geometry.Point(p.lon, p.lat
					).transform(new OpenLayers.Projection("EPSG:4326"),  this.map.getProjectionObject() )
			  );
	}
	var line = new OpenLayers.Geometry.LineString(points);

	var style = { 
		strokeColor: '#FF2285', 
		strokeOpacity: 0.9,
		strokeWidth: 1
	};

	//var lineFeature = new OpenLayers.Feature.Vector(line, null, style);
	this.L.track.addFeatures([ new OpenLayers.Feature.Vector(line, null, style) ]);
	this.map.zoomToExtent(this.L.track.getDataExtent()); 
	this.map.zoomOut();
	
},
//==========================================================
// Shows aircraft on the RADAR map, with callsign (two features, poor openlayer)
show_radar: function show_radar(mcallsign, mlat, mlon, mheading, maltitude){

	// remove xisting iamge/label if exist
	/*
	var existing_img = radarImageMarkers.getFeatureBy("_callsign", mcallsign);
	if(existing_img){
		radarImageMarkers.removeFeatures(existing_img);
	}
	var existing_lbl  = radarLabelMarkers.getFeatureBy("_callsign", mcallsign);
	if(existing_lbl){
		radarLabelMarkers.removeFeatures(existing_lbl);
	}
	*/
	//c//onsole.log(mcallsign, mlat, mlon, mheading, maltitude)
	var pointImg = new OpenLayers.Geometry.Point(mlon, mlat
						).transform(new OpenLayers.Projection("EPSG:4326"), this.map.getProjectionObject() );	
	//if(!this.get_map().getExtent().containsPixel(pointImg, false)){
		//return; //alert(map.getExtent().containsLonLat(pointImg, false));
	//}

	// Add Image
	var imgFeat = new OpenLayers.Feature.Vector(pointImg, {
				hdg: mheading
				}); 
	imgFeat._callsign = mcallsign;
	this.L.radarBlip.addFeatures([imgFeat]);	
	//console.log(mcallsign, mlat, mlon, mheading, maltitude);
	
	var gxOff = 4;
	var gyOff = -8;

	var lxOff = 6;
	var lyOff = 2;
	
	// move the label offset
	if(mheading > 0  && mheading < 90){
		lyOff = lyOff - 15;
		gyOff = gyOff  + 15 ;
	}else if( mheading > 90 && mheading < 150){
		lyOff = lyOff + 5;
		gyOff = gyOff - 5;
	}else if( mheading > 270 && mheading < 360){
		lyOff = lyOff - 10;
		gyOff = gyOff  + 10;
		
	}

	// Add callsign label as separate feature, to have a background color (graphic) with offset
	var pointLabel = new OpenLayers.Geometry.Point(mlon, mlat
					).transform(new OpenLayers.Projection("EPSG:4326"),  this.map.getProjectionObject() );
	var lblFeat = new OpenLayers.Feature.Vector(pointLabel, {
                callsign: mcallsign,
				lxOff: lxOff, lyOff: lyOff,
				gxOff: gxOff, gyOff: gyOff
				});
	lblFeat._callsign = mcallsign;
	this.L.radarLbl.addFeatures([lblFeat]);	
	
},


});
