
function generateInstanceDivs(nodesMap) {
    nodesList = []
    for (var nodeId in nodesMap) {
        nodesList.push(nodesMap[nodeId]);
    } 

    $("[data-fo-id]").each(function () {
        var isVirtual = $(this).attr("data-fo-is-virtual") == "true";
        if (!isVirtual) {
	        $(this).html(
	        	'<div xmlns="http://www.w3.org/1999/xhtml" class="popover right instance"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>'
	        );
        }
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
    	//$(".popover.instance[data-duplicate-node] h3").addClass("label-primary");
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
        	cancel: "#cluster_container .popover.instance h3 a",
        	start: function(event, ui) {
        		resetRefreshTimer();
        		$("#cluster_container .accept_drop").removeClass("accept_drop");
        		$("#cluster_container .popover.instance").droppable({
        			accept: function(draggable) {
        				var draggedNode = nodesMap[draggedNodeId];
        				var targetNode = nodesMap[$(this).attr("data-nodeid")];
        				var acceptDrop =  moveInstance(draggedNode, targetNode, false);
        				if (acceptDrop) {
        					$(this).addClass("accept_drop");
        				}
        				return acceptDrop;
        			},
        			hoverClass: "draggable-hovers",
					drop: function( event, ui ) {
				        $(".popover.instance[data-duplicate-node]").remove();
				        moveInstance(nodesMap[draggedNodeId], nodesMap[$(this).attr("data-nodeid")], true);
					}
        		});
        	},
	    	drag: function(event, ui) {
	    		resetRefreshTimer();
	    	},
	    	stop: function(event, ui) {
	    		resetRefreshTimer();
        		$("#cluster_container .accept_drop").removeClass("accept_drop");
	    	}
        });
    	$(duplicate).on("mouseleave", function() {
    		if (!$(this).hasClass("ui-draggable-dragging")) {
	    		$(this).remove();
    		}
    	});
    	// Don't ask why the following... jqueryUI recognizes the click as start drag, but fails to stop...
    	$(duplicate).on("click", function() {
        	$("#cluster_container .accept_drop").removeClass("accept_drop");
        	return false;
        });	
    });
}

function moveInstance(node, droppableNode, shouldApply) {
	if (node.isCoMaster) {
		return false;
	}
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
		$("#cluster_container .accept_drop").removeClass("accept_drop");
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
		$("#cluster_container .accept_drop").removeClass("accept_drop");
	}); 
	return false;
}


function instancesAreSiblings(node1, node2) {
	if (node1.id == node2.id) return false;
	if (node1.masterNode == null ) return false;
	if (node2.masterNode == null ) return false;
	if (node1.masterNode.id != node2.masterNode.id) return false;
	return true;
}


function instanceIsGrandchild(node, grandparentNode, nodesMap) {
	if (!node.hasMaster) {
		return false;
	}
	var masterNode = node.masterNode;
	if (!masterNode.hasMaster) {
		return false;
	}
	if (masterNode.masterNode.id != grandparentNode.id) {
		return false;
	}
	if (node.id == grandparentNode.id) {
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
                        visualizeInstances(instancesMap);
                        generateInstanceDivs(instancesMap);
                    }, "json");
            }, "json");
            
            startRefreshTimer();
            $(document).click(function() {
            	resetRefreshTimer();
            });
            $(document).mousemove(function() {
            	resetRefreshTimer();
            });
        });
