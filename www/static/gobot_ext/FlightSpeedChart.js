Ext.define("GB.FlightSpeedChart", {
	
extend: "Ext.chart.Chart", 

initComponent: function(){
	
	Ext.apply(this, {
		style: 'background:#fff',
		animate: true,
		//store: Ext.getStore(this.x_store_id),
		shadow: true,
		theme: 'Category1',
		//legend: {
		//	position: 'right'
		//},
		axes: [
			{type: 'Numeric',
				minimum: 0,
				position: 'left',
				fields: ['spd_kt'],
				title: 'Speed Kt',
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
				//minorTickSteps: 60,
				//majorTickSteps: 100	
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
				yField: 'spd_kt',
				markerConfig: {
					type: 'circle',
					size: 2,
					radius: 2,
					'stroke-width': 0,
					fill: "#0F0"
				}
			}
		]
	});
	this.callParent();
	
}

});