 Ext.define("GB.flights.FlightsAltitudeChart", {
	
extend: "Ext.chart.Chart", 

initComponent: function(){
	
	Ext.apply(this, {
		style: 'background:#fff',
		animate: true,
		shadow: true,
		store: Ext.getStore("stoFlights"),
		axes: [
			{type: 'Numeric',
				position: 'bottom',
				fields: ['alt_ft'],
				label: {
					renderer: Ext.util.Format.numberRenderer('0,0')
				},
				title: 'Altitude',
				grid: true,
				minimum: 0
			}, 
			{type: 'Category',
				position: 'left',
				fields: ['callsign'],
				title: 'Flight'
			}
		],
		series: [
			{type: 'bar',
				axis: 'bottom',
				highlight: true,
				tips: {
					trackMouse: true,
					width: 140,
					height: 28,
					renderer: function(storeItem, item) {
						this.setTitle(storeItem.get('callsign') + ': ' + storeItem.get('alt_ft') + ' ft');
					}
				},
				label: {
					display: 'insideEnd',
					'text-anchor': 'middle',
					field: 'data1',
					renderer: Ext.util.Format.numberRenderer('0'),
					orientation: 'vertical',
					color: '#333'
				},
				xField: 'callsign',
				yField: 'alt_ft'
			}
		]
	});
	this.callParent();
}

}); // end define
	