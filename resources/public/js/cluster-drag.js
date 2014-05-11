{
	var selectedNode = null;
	var draggingNode = null;
	// panning variables
	var panSpeed = 200;
	var panBoundary = 20; // Within 20px from edges will pan when dragging.

	function pan(domNode, direction) {
	    var speed = panSpeed;
	    if (panTimer) {
	        clearTimeout(panTimer);
	        translateCoords = d3.transform(svgGroup.attr("transform"));
	        if (direction == 'left' || direction == 'right') {
	            translateX = direction == 'left' ? translateCoords.translate[0] + speed : translateCoords.translate[0] - speed;
	            translateY = translateCoords.translate[1];
	        } else if (direction == 'up' || direction == 'down') {
	            translateX = translateCoords.translate[0];
	            translateY = direction == 'up' ? translateCoords.translate[1] + speed : translateCoords.translate[1] - speed;
	        }
	        scaleX = translateCoords.scale[0];
	        scaleY = translateCoords.scale[1];
	        scale = zoomListener.scale();
	        svgGroup.transition().attr("transform", "translate(" + translateX + "," + translateY + ")scale(" + scale + ")");
	        d3.select(domNode).select('g.node').attr("transform", "translate(" + translateX + "," + translateY + ")");
	        zoomListener.scale(zoomListener.scale());
	        zoomListener.translate([translateX, translateY]);
	        panTimer = setTimeout(function () {
	            pan(domNode, speed, direction);
	        }, 50);
	    }
	}
	
	function initiateDrag(d, domNode) {
		draggingNode = d;
		d3.select(domNode).select('.ghostCircle')
				.attr('pointer-events', 'none');
		d3.selectAll('.ghostCircle').attr('class', 'ghostCircle show');
		d3.select(domNode).attr('class', 'node activeDrag');

		svgGroup.selectAll("g.node").sort(function(a, b) { // select the parent
															// and
			// sort the path's
			if (a.id != draggingNode.id)
				return 1; // a is not the hovered element, send "a" to the
							// back
			else
				return -1; // a is the hovered element, bring "a" to the front
		});
		// if nodes has children, remove the links and nodes
		if (nodes.length > 1) {
			// remove link paths
			links = tree.links(nodes);
			nodePaths = svgGroup.selectAll("path.link").data(links,
					function(d) {
						return d.target.id;
					}).remove();
			// remove child nodes
			nodesExit = svgGroup.selectAll("g.node").data(nodes, function(d) {
				return d.id;
			}).filter(function(d, i) {
				if (d.id == draggingNode.id) {
					return false;
				}
				return true;
			}).remove();
		}

		// remove parent link
		parentLink = tree.links(tree.nodes(draggingNode.parent));
		svgGroup.selectAll('path.link').filter(function(d, i) {
			if (d.target.id == draggingNode.id) {
				return true;
			}
			return false;
		}).remove();

		dragStarted = null;
	}

	// Define the drag listeners for drag/drop behaviour of nodes.
	dragListener = d3.behavior.drag().on("dragstart", function(d) {
		if (d == root) {
			return;
		}
		dragStarted = true;
		nodes = tree.nodes(d);
		d3.event.sourceEvent.stopPropagation();
		// it's important that we suppress the mouseover event on the node being
		// dragged. Otherwise it will absorb the mouseover event and the
		// underlying
		// node will not detect it d3.select(this).attr('pointer-events',
		// 'none');
	}).on("drag", function(d) {
		if (d == root) {
			return;
		}
		if (dragStarted) {
			domNode = this;
			initiateDrag(d, domNode);
		}

		// get coords of mouseEvent relative to svg container to allow for
		// panning
		relCoords = d3.mouse($('svg').get(0));
		if (relCoords[0] < panBoundary) {
			panTimer = true;
			pan(this, 'left');
		} else if (relCoords[0] > ($('svg').width() - panBoundary)) {

			panTimer = true;
			pan(this, 'right');
		} else if (relCoords[1] < panBoundary) {
			panTimer = true;
			pan(this, 'up');
		} else if (relCoords[1] > ($('svg').height() - panBoundary)) {
			panTimer = true;
			pan(this, 'down');
		} else {
			try {
				clearTimeout(panTimer);
			} catch (e) {

			}
		}

		d.x0 += d3.event.dy;
		d.y0 += d3.event.dx;
		var node = d3.select(this);
		node.attr("transform", "translate(" + d.y0 + "," + d.x0 + ")");
		updateTempConnector();
	}).on(
			"dragend",
			function(d) {
				if (d == root) {
					return;
				}
				domNode = this;
				if (selectedNode) {
					// now remove the element from the parent, and insert it
					// into
					// the new elements children
					var index = draggingNode.parent.children
							.indexOf(draggingNode);
					if (index > -1) {
						draggingNode.parent.children.splice(index, 1);
					}
					if (typeof selectedNode.children !== 'undefined'
							|| typeof selectedNode._children !== 'undefined') {
						if (typeof selectedNode.children !== 'undefined') {
							selectedNode.children.push(draggingNode);
						} else {
							selectedNode._children.push(draggingNode);
						}
					} else {
						selectedNode.children = [];
						selectedNode.children.push(draggingNode);
					}
					// Make sure that the node being added to is expanded so
					// user
					// can see added node is correctly moved
					//expand(selectedNode);
					//sortTree();
					endDrag();
				} else {
					endDrag();
				}
			});

	function endDrag() {
		selectedNode = null;
		d3.selectAll('.ghostCircle').attr('class', 'ghostCircle');
		d3.select(domNode).attr('class', 'node');
		// now restore the mouseover event or we won't be able to drag a 2nd
		// time
		d3.select(domNode).select('.ghostCircle').attr('pointer-events', '');
		updateTempConnector();
		if (draggingNode !== null) {
			update(root);
			centerNode(draggingNode);
			draggingNode = null;
		}
	}
	// Function to update the temporary connector indicating dragging
	// affiliation
	var updateTempConnector = function() {
		var data = [];
		if (draggingNode !== null && selectedNode !== null) {
			// have to flip the source coordinates since we did this for the
			// existing connectors on the original tree
			data = [ {
				source : {
					x : selectedNode.y0,
					y : selectedNode.x0
				},
				target : {
					x : draggingNode.y0,
					y : draggingNode.x0
				}
			} ];
		}
		var link = svgGroup.selectAll(".templink").data(data);

		link.enter().append("path").attr("class", "templink").attr("d",
				d3.svg.diagonal()).attr('pointer-events', 'none');

		link.attr("d", d3.svg.diagonal());

		link.exit().remove();
	};

	// Function to center node when clicked/dropped so node doesn't get lost
	// when
	// collapsing/moving with large amount of children.

	function centerNode(source) {
		scale = zoomListener.scale();
		x = -source.y0;
		y = -source.x0;
		x = x * scale + viewerWidth / 2;
		y = y * scale + viewerHeight / 2;
		d3.select('g').transition().duration(duration).attr("transform",
				"translate(" + x + "," + y + ")scale(" + scale + ")");
		zoomListener.scale(scale);
		zoomListener.translate([ x, y ]);
	}

	var svgGroup = $("svg").append("g");
}
