initClusterTopology = function(connectDivs) {
	
	jsPlumb.ready(function() {			

		var color = "gray";

		var jsPlumbInstance = jsPlumb.getInstance({
			// notice the 'curviness' argument to this Bezier curve.  the curves on this page are far smoother
			// than the curves on the first demo, which use the default curviness value.			
			Connector : [ "Bezier", { curviness:50 } ],
			DragOptions : { cursor: "pointer", zIndex:2000 },
			PaintStyle : { strokeStyle:color, lineWidth:2 },
			EndpointStyle : { radius:9, fillStyle:color },
			HoverPaintStyle : {strokeStyle:"#ec9f2e" },
			EndpointHoverStyle : {fillStyle:"#ec9f2e" },
			Container:"instances"
		});
			
		// suspend drawing and initialise.
		jsPlumbInstance.doWhileSuspended(function() {		
			// declare some common values:
			var arrowCommon = { foldback:0.7, fillStyle:color, width:14 },
				// use three-arg spec to create two different arrows with the common values:
				overlays = [
					[ "Arrow", { location:0.7 }, arrowCommon ],
					[ "Arrow", { location:0.3, direction:-1 }, arrowCommon ]
				];

			// add endpoints, giving them a UUID.
			// you DO NOT NEED to use this method. You can use your library's selector method.
			// the jsPlumb demos use it so that the code can be shared between all three libraries.
			
			var windows = jsPlumb.getSelector(".chart-demo .window");
			for (var i = 0; i < windows.length; i++) {
				jsPlumbInstance.addEndpoint(windows[i], {
					uuid:windows[i].getAttribute("id") + "-bottom",
					anchor:"Bottom",
					maxConnections:-1
				});
				jsPlumbInstance.addEndpoint(windows[i], {
					uuid:windows[i].getAttribute("id") + "-top",
					anchor:"Top",
					maxConnections:-1
				});
			}
			connectDivs(jsPlumbInstance, overlays);
					
			//jsPlumbInstance.draggable(windows);		
		});
	});
	
};
