Ext.define("GB.flight.FlightAltitudeChart", {
	
extend: "Ext.chart.Chart", 

initComponent: function(){
	
	Ext.apply(this, {
		style: 'background:#fff',
		animate: false,
		//store: Ext.getStore(this.x_store_id),
		shadow: false,
		theme: 'Category1',
		//legend: {
		//	position: 'right'
		//},
		axes: [
			{type: 'Numeric',
				minimum: 0,
				position: 'left',
				fields: ['alt_ft'],
				title: 'Altitude ',
				minorTickSteps: 1,
				grid: {
					odd: {
						opacity: 1,
						fill: '#ddd',
						stroke: '#bbb',
						'stroke-width': 0.5
					}
				}
			}, 
		   {type: 'Category',
				position: 'bottom',
				fields: ['elapsed'],
				title: 'Elapsed',
				step: 1
			}
		],
		series: [
			{type: 'line',
				highlight: {
					size: 7,
					radius: 7
				},
				axis: 'left',
				xField: 'elapsed',
				yField: 'alt_ft',
				//showMarkers: false,
				markerConfig: {
					type: 'circle',
					size: 2,
					radius: 2,
					'stroke-width': 0
				}
			}
		]
	});
	this.callParent();
	
}

});