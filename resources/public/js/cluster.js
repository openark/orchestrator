
function generateInstanceDivs(nodesList) {
    var nodesMap = nodesList.reduce(function (map, node) {
        map[node.id] = node;
        return map;
    }, {});

    $("[data-fo-id]").each(function () {
        var id = $(this).attr("data-fo-id");
        $(this).html(
        	'<div xmlns="http://www.w3.org/1999/xhtml" class="popover right instance"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>'
        );
    });
    nodesList.forEach(function (node) {
    	var popoverElement = $("[data-fo-id='" + node.id + "'] .popover");
    	renderInstanceElement(popoverElement, node, "cluster");
    });
    $("[data-fo-id]").each(
        function () {
            var id = $(this).attr("data-fo-id");
            var popoverDiv = $("[data-fo-id='" + id + "'] div.popover");

            popoverDiv.attr("x", $(this).attr("x"));
            $(this).attr("y",
                0 - popoverDiv.height() / 2 - 2);
            popoverDiv.attr("y", $(this).attr("y"));
            $(this).attr("width",
                popoverDiv.width() + 30);
            $(this).attr("height",
                popoverDiv.height() +16);
        });
    $("div.popover").popover();
    $("div.popover").show();
    
    $("[data-fo-id]").on("mouseenter", ".popover[data-nodeid]", function() {
    	if ($(".popover.instance[data-duplicate-node]").hasClass("ui-draggable-dragging")) {
    		// Do not remove & recreate while dragging. Ignore any mouseenter
    		return false;
    	}
    	var draggedNodeId = $(this).attr("data-nodeid"); 
    	if (draggedNodeId == $(".popover.instance[data-duplicate-node]").attr("data-nodeid")) {
    		return false;
    	}
    	$(".popover.instance[data-duplicate-node]").remove();
    	var duplicate = $(this).clone().appendTo("#cluster_container");
    	$(duplicate).attr("data-duplicate-node", "true");
    	$(".popover.instance[data-duplicate-node] h3").addClass("label-primary");
    	$(duplicate).css({"margin-left": "0"});
    	$(duplicate).css($(this).offset());
    	$(duplicate).width($(this).width());
    	$(duplicate).height($(this).height());
    	$(duplicate).popover();
        $(duplicate).show();
        $(".popover.instance[data-duplicate-node] h3 a").click(function () {
        	openNodeModal(nodesMap[draggedNodeId]);
        	return false;
        });
        $(duplicate).draggable({ 
        	addClasses: true, 
        	opacity: 0.67,
        	start: function(event, ui) {
        		//$(this).data("startingScrollTop",$(this).parent().scrollTop());
        		$("#cluster_container .popover.instance").droppable({
        			accept: function(draggable) {
        				var draggedNode = nodesMap[draggedNodeId];
        				var targetNode = nodesMap[$(this).attr("data-nodeid")];
        				return moveInstance(draggedNode, targetNode, false);
        			},
        			hoverClass: "draggable-hovers",
					drop: function( event, ui ) {
				        $(".popover.instance[data-duplicate-node]").remove();
				        moveInstance(nodesMap[draggedNodeId], nodesMap[$(this).attr("data-nodeid")], true);
					}
        		});
        	},
	    	drag: function(event, ui) {
	    		//var st = parseInt($(this).data("startingScrollTop"));
	    		//ui.position.top -= $(this).parent().scrollTop() - st;
	    	},
	    	stop: function(event, ui) {
        		//$(".popover.instance").droppable("disable");
	    	}
        });
    	$(duplicate).on("mouseleave", function() {
    		if (!$(this).hasClass("ui-draggable-dragging")) {
	        	//$(".popover.instance[data-duplicate-node]").remove();
	    		$(this).remove();
    		}
    	});
    });
}

function moveInstance(node, droppableNode, shouldApply) {
	if (instancesAreSiblings(node, droppableNode)) {
		if (node.hasProblem || droppableNode.hasProblem) {
			return false;
		}
		if (shouldApply) {
			moveBelow(node, droppableNode);
		}
		return true;
	}
	if (instanceIsGrandchild(node, droppableNode)) {
		if (node.hasProblem) {
			return false;
		}
		if (shouldApply) {
			moveUp(node, droppableNode);
		}
		return true;
	}
	
	if (shouldApply) {
		addAlert(
				"Cannot move <code><strong>" + 
					node.Key.Hostname + ":" + node.Key.Port +
					"</strong></code> under <code><strong>" +
					droppableNode.Key.Hostname + ":" + droppableNode.Key.Port +
					"</strong></code>. " +
				"You may only move a node down below its sibling or up below its grandparent."
			);
	}
	return false;
}

function moveBelow(node, siblingNode) {
	var message = "Are you sure you wish to turn <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> into a slave of <code><strong>" +
		siblingNode.Key.Hostname + ":" + siblingNode.Key.Port +
		"</strong></code>?";
	bootbox.confirm(message, function(confirm) {
		if (confirm) {
			showLoader();
			var apiUrl = "/api/move-below/" + node.Key.Hostname + "/" + node.Key.Port + "/" + siblingNode.Key.Hostname + "/" + siblingNode.Key.Port;
		    $.get(apiUrl, function (operationResult) {
	    			hideLoader();
	    			if (operationResult.Code == "ERROR") {
	    				addAlert(operationResult.Message)
	    			} else {
	    				location.reload();
	    			}	
	            }, "json");					
		}
	}); 
	return false;
}


function moveUp(node, grandparentNode) {
	var message = "Are you sure you wish to turn <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> into a slave of <code><strong>" +
		grandparentNode.Key.Hostname + ":" + grandparentNode.Key.Port +
		"</strong></code>?"
	bootbox.confirm(message, function(confirm) {
		if (confirm) {
			showLoader();
			var apiUrl = "/api/move-up/" + node.Key.Hostname + "/" + node.Key.Port;
		    $.get(apiUrl, function (operationResult) {
	    			hideLoader();
	    			if (operationResult.Code == "ERROR") {
	    				addAlert(operationResult.Message)
	    			} else {
	    				location.reload();
	    			}	
	            }, "json");					
		}
	}); 
	return false;
}


function instancesAreSiblings(node1, node2) {
	if (node1.id == node2.id) return false;
	if (node1.parent == null ) return false;
	if (node2.parent == null ) return false;
	if (node1.parent.id != node2.parent.id) return false;
	return true;
}


function instanceIsGrandchild(node, grandparentNode, nodesMap) {
	if (!node.hasMaster) {
		return false;
	}
	var parentNode = node.parent;
	if (!parentNode.hasMaster) {
		return false;
	}
	if (parentNode.parent.id != grandparentNode.id) {
		return false;
	}
	return true;
}


$(document)
    .ready(
        function () {

            $.get("/api/cluster/"+currentClusterName(), function (instances) {
                $.get("/api/maintenance",
                    function (maintenanceList) {
                		var instancesMap = normalizeInstances(instances, maintenanceList);
                        visualizeInstances(instances, instancesMap);
                        generateInstanceDivs(instances);
                    }, "json");
            }, "json");
        });
