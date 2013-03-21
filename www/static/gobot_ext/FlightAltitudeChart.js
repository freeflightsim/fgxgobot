Ext.define("GB.FlightAltitudeChart", {
	
extend: "Ext.chart.Chart", 

initComponent: function(){
	
	Ext.apply(this, {
		style: 'background:#fff',
		animate: true,
		//store: Ext.getStore(this.x_store_id),
		shadow: true,
		theme: 'Category1',
		legend: {
			position: 'right'
		},
		axes: [
			{type: 'Numeric',
				minimum: 0,
				position: 'left',
				fields: ['alt_ft'],
				title: 'Altitude',
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
				fields: ['idx'],
				title: 'Positions',
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
				xField: 'idx',
				yField: 'alt_ft',
				markerConfig: {
					type: 'cross',
					size: 4,
					radius: 4,
					'stroke-width': 0
				}
			}
		]
	});
	this.callParent();
	
}

});