function getInstanceId(host, port) {
    return "instance" + host.replace(/[.]/g, "_") + "__" + port
}

function addNodeModalDataAttribute(name, value) {
    $('#modalDataAttributesTable').append(
        '<tr><td>' + name + '</td><td><code>' + value + '</code></td></tr>');
}

function openNodeModal(node) {
    $('#node_modal .modal-title').html(node.title);
    $('#modalDataAttributesTable').html("");

    if (node.MasterKey.Hostname) {
        addNodeModalDataAttribute("Master",
            node.MasterKey.Hostname + ":" + node.MasterKey.Port);
        addNodeModalDataAttribute("Replication running",
            booleanString(node.Slave_SQL_Running && node.Slave_IO_Running));
        addNodeModalDataAttribute("Replication lag",
            node.SecondsBehindMaster);
    }
    addNodeModalDataAttribute("Num slaves",
        node.SlaveHosts.length);
    addNodeModalDataAttribute("Server ID", node.ServerID);
    addNodeModalDataAttribute("Version", node.Version);
    addNodeModalDataAttribute("Binlog format",
        node.Binlog_format);
    addNodeModalDataAttribute("Has binary logs",
        booleanString(node.LogBinEnabled));
    addNodeModalDataAttribute("Logs slave updates",
        booleanString(node.LogSlaveUpdatesEnabled));
    $('#node_modal').modal({})
}

function generateInstanceDivs(nodesList) {
    var nodesMap = nodesList.reduce(function (map, node) {
        map[node.id] = node;
        return map;
    }, {});

    $("[data-fo-id]")
        .each(
            function () {
                var id = $(this).attr("data-fo-id");
                $(this)
                    .html(
                        '<div xmlns="http://www.w3.org/1999/xhtml" class="popover right instance"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"><p></p></div></div>');

            });
    nodesList
        .forEach(function (node) {
            $("[data-fo-id='" + node.id + "'] .popover").attr("data-nodeid", node.id);
            $(
                    "[data-fo-id='" + node.id + "'] .popover h3")
                    .html(
                        node.canonicalTitle + '<div class="pull-right"><a href="#"><span class="glyphicon glyphicon-cog"></span></a></div>');
            if (node.inMaintenance) {
                $(
                    "[data-fo-id='" + node.id + "'] .popover h3")
                    .addClass("label-info");
            } else if (node.secondsBehindMaster > 10) {
                $(
                    "[data-fo-id='" + node.id + "'] .popover h3")
                    .addClass("label-danger");
            }
            $(
                "[data-fo-id='" + node.id + "'] .popover .popover-content p")
                .html(
                    node.Version + " " + node.Binlog_format + '<div class="pull-right">' + node.SecondsBehindMaster + ' seconds lag</div>');
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
        		$(".popover.instance").droppable({
        			accept: function(draggable) {
        				//".popover.instance[data-duplicate-node]",
        				return moveInstance(nodesMap[draggedNodeId], nodesMap[$(this).attr("data-nodeid")], false);
        			},
        			hoverClass: "draggable-hovers",
					drop: function( event, ui ) {
				        $(".popover.instance[data-duplicate-node]").remove();
				        moveInstance(nodesMap[draggedNodeId], nodesMap[$(this).attr("data-nodeid")], true);
					}
        		});
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
		if (shouldApply) {
			moveBelow(node, droppableNode);
		}
		return true;
	}
	if (instanceIsGrandchild(node, droppableNode)) {
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
	bootbox.confirm("Are you sure you wish to turn <code><strong>" + 
				node.Key.Hostname + ":" + node.Key.Port +
				"</strong></code> into a slave of <code><strong>" +
				siblingNode.Key.Hostname + ":" + siblingNode.Key.Port +
				"</strong></code>?"
			, function(confirm) {
				if (confirm) {
					showLoader();
					var apiUrl = "/api/move-below/" + node.Key.Hostname + "/" + node.Key.Port + "/" + siblingNode.Key.Hostname + "/" + siblingNode.Key.Port;
				    $.get(apiUrl, function (operationResult) {
			    			hideLoader();
			    			if (operationResult.Code == "ERROR") {
			    				addAlert("<strong>Error</strong>: " + operationResult.Message)
			    			} else {
			    				location.reload();
			    			}	
			            }, "json");					
				}
			}); 
	return false;
}


function moveUp(node, grandparentNode) {
	bootbox.confirm("Are you sure you wish to turn <code><strong>" + 
				node.Key.Hostname + ":" + node.Key.Port +
				"</strong></code> into a slave of <code><strong>" +
				grandparentNode.Key.Hostname + ":" + grandparentNode.Key.Port +
				"</strong></code>?"
			, function(confirm) {
				if (confirm) {
					showLoader();
					var apiUrl = "/api/move-up/" + node.Key.Hostname + "/" + node.Key.Port;
				    $.get(apiUrl, function (operationResult) {
			    			hideLoader();
			    			if (operationResult.Code == "ERROR") {
			    				addAlert("<strong>Error</strong>: " + operationResult.Message)
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
                    function (maintenanceInstances) {
                        visualizeInstances(instances,
                            maintenanceInstances);
                        generateInstanceDivs(instances);
                    }, "json");
            }, "json");
        });
