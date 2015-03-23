clusterOperationPseudoGTIDMode = false;

var dcColors = ["#ff8c00", "#4682b4", "#9acd32", "#dc143c", "#9932cc", "#ffd700", "#191970", "#7fffd4", "#808080", "#dda0dd"];
var dcColorsMap = {};


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
        $(".popover.instance[data-duplicate-node] a[data-command=recover-auto]").click(function () {
        	return apiCommand("/api/recover/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port);

        });
        $(".popover.instance[data-duplicate-node] a[data-command=match-up-slaves]").click(function () {
        	return apiCommand("/api/match-up-slaves/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port);
        });
        $(".popover.instance[data-duplicate-node] a[data-command=regroup-slaves]").click(function () {
        	return apiCommand("/api/regroup-slaves/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port);
        });
        $(".popover.instance[data-duplicate-node] a[data-command=multi-match-slaves]").click(function (event) {
            var belowHost = $(event.target).attr("data-below-host");
            var belowPort = $(event.target).attr("data-below-port");
        	return apiCommand("/api/multi-match-slaves/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port+"/"+belowHost+"/"+belowPort);
        });
        $(".popover.instance[data-duplicate-node] button[data-command=make-master]").click(function () {
        	makeMaster(nodesMap[draggedNodeId]);
        	return false;
        });
        $(".popover.instance[data-duplicate-node] button[data-command=make-local-master]").click(function () {
        	makeLocalMaster(nodesMap[draggedNodeId]);
        	return false;
        });
        /*
        $(".popover.instance[data-duplicate-node] h3").click(function () {
        	var popoverElement = $("[data-fo-id='" + draggedNodeId + "'] .popover");
        	if (popoverElement.hasClass("selected")) {
        		popoverElement.removeClass("selected");
        	} else {
        		popoverElement.addClass("selected");
        	}
       		return false;
        });
        */
        if (!isAuthorizedForAction() || nodesMap[draggedNodeId].lastCheckInvalidProblem() || nodesMap[draggedNodeId].notRecentlyCheckedProblem()) {
            $(".popover.instance[data-duplicate-node] h3").click(function () {
               	openNodeModal(nodesMap[draggedNodeId]);
            	return false;
            });
        } else {
            $(duplicate).draggable({
            	addClasses: true, 
            	opacity: 0.67,
            	cancel: "#cluster_container .popover.instance h3 a",
            	start: function(event, ui) {
            		resetRefreshTimer();
            		$("#cluster_container .accept_drop").removeClass("accept_drop");
             		$("#cluster_container .accept_drop_warning").removeClass("accept_drop_warning");
             		$("#cluster_container .popover.instance").droppable({
            			accept: function(draggable) {
            				var draggedNode = nodesMap[draggedNodeId];
            				var targetNode = nodesMap[$(this).attr("data-nodeid")];
            				var acceptDrop =  moveInstance(draggedNode, targetNode, false);
            				if (acceptDrop == "ok") {
            					$(this).addClass("accept_drop");
            				}
            				if (acceptDrop == "warning") {
            					$(this).addClass("accept_drop_warning");
            				}
            				return acceptDrop != null;
            			},
            			hoverClass: "draggable-hovers",
					    over: function( event, ui ) {
					    },
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
            		$("#cluster_container .accept_drop_warning").removeClass("accept_drop_warning");
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
            	$("#cluster_container .accept_drop").removeClass("accept_drop_warning");
            	return false;
            });	
        }
    });
}

function moveInstance(node, droppableNode, shouldApply) {
    if (!isAuthorizedForAction()) {
    	// Obviously this is also checked on server side, no need to try stupid hacks
		return null;
    }
	if (clusterOperationPseudoGTIDMode) {
		if (node.hasConnectivityProblem || droppableNode.hasConnectivityProblem) {
			// Obviously can't handle.
			return null;
		}
		
		if (node.id == droppableNode.id) {
			return null;
		}
		if (instanceIsDescendant(droppableNode, node)) {
			// Wrong direction!
			return null;
		}
		if (isReplicationStrictlyBehindSibling(droppableNode, node)) {
			// Sibling known to be less advanced. Wrong direction!
			return null;
		}
		if (instanceIsDescendant(node, droppableNode)) {
			// clearly node cannot be more up to date than droppableNode
			if (shouldApply) {
				matchBelow(node, droppableNode);
			}
			return "ok";
		}
		if (isReplicationBehindSibling(node, droppableNode)) {
			// verified that node isn't more up to date than droppableNode
			if (shouldApply) {
				matchBelow(node, droppableNode);
			}
			return "ok";
		}
		// TODO: the general case, where there's no clear family connection, meaning we cannot infer
		// which instance is more up to date. It's under the user's responsibility!
		if (shouldApply) {
			matchBelow(node, droppableNode);
		}
		return "warning";
	}
	// Not pseudo-GTID mode
	if (node.isCoMaster) {
		// Cannot move. RESET SLAVE on one of the co-masters.
		return null;
	}
	if (instancesAreSiblings(node, droppableNode)) {
		if (node.hasProblem || droppableNode.hasProblem) {
			return null;
		}
		if (shouldApply) {
			moveBelow(node, droppableNode);
		}
		return "ok";
	}
	if (instanceIsGrandchild(node, droppableNode)) {
		if (node.hasProblem) {
			// Typically, when a node has a problem we do not allow moving it up.
			// But there's a special situation when allowing is desired: when the parent has personal issues,
			// (say disk issue or otherwise something heavyweight running which slows down replication)
			// and you want to move up the slave which is only delayed by its master.
			// So to help out, if the instance is identically at its master's trail, it is allowed to move up.
			if (!node.isSQLThreadCaughtUpWithIOThread) { 
				return null;
			}
		}
		if (shouldApply) {
			moveUp(node, droppableNode);
		}
		return "ok";
	}
	if (instanceIsChild(droppableNode, node) && node.isMaster) {
		if (node.hasProblem) {
			return null;
		}
		if (shouldApply) {
			makeCoMaster(node, droppableNode);
		}
		return "ok";
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
	return null;
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
	    				reloadWithOperationResult(operationResult);
	    			}	
	            }, "json");					
		}
		$("#cluster_container .accept_drop").removeClass("accept_drop");
    	$("#cluster_container .accept_drop").removeClass("accept_drop_warning");
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
	    				reloadWithOperationResult(operationResult);
	    			}	
	            }, "json");					
		}
		$("#cluster_container .accept_drop").removeClass("accept_drop");
	}); 
	return false;
}


function makeCoMaster(node, childNode) {
	var message = "Are you sure you wish to make <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> and <code><strong>" +
		childNode.Key.Hostname + ":" + childNode.Key.Port +
		"</strong></code> co-masters?"
	bootbox.confirm(message, function(confirm) {
		if (confirm) {
			showLoader();
			var apiUrl = "/api/make-co-master/" + childNode.Key.Hostname + "/" + childNode.Key.Port;
		    $.get(apiUrl, function (operationResult) {
	    			hideLoader();
	    			if (operationResult.Code == "ERROR") {
	    				addAlert(operationResult.Message)
	    			} else {
	    				reloadWithOperationResult(operationResult);
	    			}	
	            }, "json");					
		}
		$("#cluster_container .accept_drop").removeClass("accept_drop");
	}); 
	return false;
}



function matchBelow(node, otherNode) {
	var message = "<h4>PSEUDO-GTID MODE</h4>Are you sure you wish to turn <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> into a slave of <code><strong>" +
		otherNode.Key.Hostname + ":" + otherNode.Key.Port +
		"</strong></code>?";
	bootbox.confirm(message, function(confirm) {
		if (confirm) {
			showLoader();
			var apiUrl = "/api/match-below/" + node.Key.Hostname + "/" + node.Key.Port + "/" + otherNode.Key.Hostname + "/" + otherNode.Key.Port;
		    $.get(apiUrl, function (operationResult) {
	    			hideLoader();
	    			if (operationResult.Code == "ERROR") {
	    				addAlert(operationResult.Message)
	    			} else {
	    				reloadWithOperationResult(operationResult);
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


function instanceIsChild(node, parentNode) {
	if (!node.hasMaster) {
		return false;
	}
	if (node.masterNode.id != parentNode.id) {
		return false;
	}
	if (node.id == parentNode.id) {
		return false;
	}
	return true;
}


function instanceIsGrandchild(node, grandparentNode) {
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


function instanceIsDescendant(node, nodeAtQuestion) {
	if (nodeAtQuestion == null) {
		return false;
	}
	if (node.id == nodeAtQuestion.id) {
		return false;
	}
	if (!node.hasMaster) {
		return false;
	}
	if (node.masterNode.id == nodeAtQuestion.id) {
		return true;
	}
	return instanceIsDescendant(node.masterNode, nodeAtQuestion)
}

// Returns true when the two instances are siblings, and 'node' is behind or at same position
// (in reltation to shared master) as its 'sibling'.
// i.e. 'sibling' is same as, or more up to date by master than 'node'.
function isReplicationBehindSibling(node, sibling) { 
	if (!instancesAreSiblings(node, sibling)) {
		return false;
	} 
	return compareInstancesExecBinlogCoordinates(node, sibling) <= 0;
}

function isReplicationStrictlyBehindSibling(node, sibling) { 
	if (!instancesAreSiblings(node, sibling)) {
		return false;
	} 
	return compareInstancesExecBinlogCoordinates(node, sibling) < 0;
}

function compareInstancesExecBinlogCoordinates(i0, i1) {
	if (i0.ExecBinlogCoordinates.LogFile == i1.ExecBinlogCoordinates.LogFile) {
		// executing from same master log file
		return i0.ExecBinlogCoordinates.LogPos - i1.ExecBinlogCoordinates.LogPos;
	}
	return (getLogFileNumber(i0.ExecBinlogCoordinates.LogFile) - getLogFileNumber(i1.ExecBinlogCoordinates.LogFile));
	
}

function getLogFileNumber(logFileName) {
	logFileTokens = logFileName.split(".")
	return parseInt(logFileTokens[logFileTokens.length-1])
}


function analyzeClusterInstances(nodesMap) {
	instances = []
    for (var nodeId in nodesMap) {
    	instances.push(nodesMap[nodeId]);
    } 

    instances.forEach(function (instance) {
	    if (instance.hasConnectivityProblem) {
	    	// The instance has a connectivity problem! Do a client-size recommendation of most advanced slave:
	    	// a direct child of the master, with largest exec binlog coordinates.
	    	var sortedChildren = instance.children.slice(); 
	    	sortedChildren.sort(compareInstancesExecBinlogCoordinates)
	    	
	    	instance.children.forEach(function(child) {
	    		if (!child.hasConnectivityProblem) {
		    		if (compareInstancesExecBinlogCoordinates(child, sortedChildren[sortedChildren.length - 1]) == 0) {
		    			child.isMostAdvancedOfSiblings = true;
		    	    	if (instance.isMaster && !instance.isCoMaster) {
		    	    		// Moreover, the instance is the (only) master!
		    	    		// Therefore its most advanced slaves are candidate masters 
		    	    		child.isCandidateMaster = true;
		    		    }
		    		}
	    		}
	    	})
		    	
    	}
    });
    instances.forEach(function (instance) {
    	if (instance.children && instance.children.length > 0) {
            instance.children[0].isFirstChildInDisplay = true
        }
    });
}

function postVisualizeInstances(nodesMap) {
    // DC colors
    var knownDCs = [];
    instances.forEach(function (instance) {
        knownDCs.push(instance.DataCenter)
    });
    function uniq(a) {
        return a.sort().filter(function(item, pos) {
            return !pos || item != a[pos - 1];
        })
    }
    knownDCs = uniq(knownDCs);
    for (i = 0 ; i < knownDCs.length ; ++i) {
        dcColorsMap[knownDCs[i]] = dcColors[i % dcColors.length];
    }
    instances.forEach(function (instance) {
    	var draggedNodeId = $(this).attr("data-nodeid"); 
    	$(".popover.instance[data-nodeid="+instance.id+"]").attr("data-dc-color", dcColorsMap[instance.DataCenter]);
    });
}


function refreshClusterOperationModeButton() {
	if (clusterOperationPseudoGTIDMode) {
		$("#cluster_operation_mode_button").html("Pseudo-GTID mode");
		$("#cluster_operation_mode_button").removeClass("btn-success").addClass("btn-warning");
	} else {
		$("#cluster_operation_mode_button").html("Classic mode");
		$("#cluster_operation_mode_button").removeClass("btn-warning").addClass("btn-success");
	}
}

function makeMaster(instance) {
	var message = "Are you sure you wish to make <code><strong>" + instance.Key.Hostname+":"+instance.Key.Port + "</strong></code> the new master?"
	+ "<p>Siblings of <code><strong>" + instance.Key.Hostname+":"+instance.Key.Port + "</strong></code> will turn to be its children, "
	+ "via Pseudo-GTID."
	+ "<p>The instance will be set to be writeable (<code><strong>read_only = 0</strong></code>)."
	+ "<p>Replication on this instance will be stopped, but not reset. You should run <code><strong>RESET SLAVE</strong></code> yourself "
	+ "if this instance will indeed become the master."
	+ "<p>Pointing your application servers to the new master is on you."
	;
	bootbox.confirm(message, function(confirm) {
		if (confirm) {
	    	showLoader();
	        $.get("/api/make-master/"+instance.Key.Hostname+"/"+instance.Key.Port, function (operationResult) {
				hideLoader();
				if (operationResult.Code == "ERROR") {
					addAlert(operationResult.Message)
				} else {
    				reloadWithOperationResult(operationResult);
				}	
	        }, "json");
		}
	});
}

function makeLocalMaster(instance) {
	var message = "Are you sure you wish to make <code><strong>" + instance.Key.Hostname+":"+instance.Key.Port + "</strong></code> a local master?"
	+ "<p>Siblings of <code><strong>" + instance.Key.Hostname+":"+instance.Key.Port + "</strong></code> will turn to be its children, "
	+ "via Pseudo-GTID."
	+ "<p>The instance will replicate from its grandparent."
	;
	bootbox.confirm(message, function(confirm) {
		if (confirm) {
	    	showLoader();
	        $.get("/api/make-local-master/"+instance.Key.Hostname+"/"+instance.Key.Port, function (operationResult) {
				hideLoader();
				if (operationResult.Code == "ERROR") {
					addAlert(operationResult.Message)
				} else {
                    reloadWithOperationResult(operationResult);
				}	
	        }, "json");
		}
	});
}


function promptForAlias(oldAlias) {
	bootbox.prompt({
		title: "Enter alias for this cluster",
		value: oldAlias,
		callback: function(result) {
			if (result !== null) {
		    	showLoader();
		        $.get("/api/set-cluster-alias/"+currentClusterName()+"?alias="+encodeURIComponent(result), function (operationResult) {
					hideLoader();
					if (operationResult.Code == "ERROR") {
						addAlert(operationResult.Message)
					} else {
						location.reload();
					}	
		        }, "json");	    				
			} 
		}
	}); 
}

$(document).ready(function () {
    $.get("/api/cluster/"+currentClusterName(), function (instances) {
        $.get("/api/maintenance",
            function (maintenanceList) {
        		var instancesMap = normalizeInstances(instances, maintenanceList);
                analyzeClusterInstances(instancesMap);
                visualizeInstances(instancesMap);
                generateInstanceDivs(instancesMap);
                postVisualizeInstances(instancesMap);
            }, "json");
    }, "json");
    $.get("/api/cluster-info/"+currentClusterName(), function (clusterInfo) {    
    	var alias = clusterInfo.ClusterAlias
    	var visualAlias = (alias ? alias : currentClusterName())
    	document.title = document.title.split(" - ")[0] + " - " + visualAlias;
    	$("#cluster_container").append('<div class="floating_background">'+visualAlias+'</div>');
        $("#dropdown-context").append('<li><a data-command="change-cluster-alias" data-alias="'+clusterInfo.ClusterAlias+'">Alias: '+alias+'</a></li>');           
        $("#dropdown-context").append('<li><a data-command="anonymize">Anonymize</a></li>');           
        $("#dropdown-context").append('<li><a data-command="colorize-dc">Colorize DC</a></li>');           
    }, "json");
    
    if (isPseudoGTIDModeEnabled()) {
        $("ul.navbar-nav").append('<li><a class="cluster_operation_mode"><button type="button" class="btn btn-xs" id="cluster_operation_mode_button"></button></a></li>');
        refreshClusterOperationModeButton();
        
	    $("body").on("click", "#cluster_operation_mode_button", function() {
	    	clusterOperationPseudoGTIDMode = !clusterOperationPseudoGTIDMode;
	    	refreshClusterOperationModeButton(); 
	    });
    }
    $("#instance_problems button").html("Cluster " + $("#instance_problems button").html())
    
    
    $("body").on("click", "a[data-command=change-cluster-alias]", function(event) {    	
    	promptForAlias($(event.target).attr("data-alias"));
    });    
    $("body").on("click", "a[data-command=anonymize]", function(event) {    	
        var _ = function() {
            var counter = 0;  
            var port = 3306;
            $("h3.popover-title .pull-left").each(function() {
               jQuery(this).html("instance-"+(counter++)+":"+port)
            });
        }();
    	$("#cluster_container div.floating_background").html("");
    });    
    $("body").on("click", "a[data-command=colorize-dc]", function(event) {
        console.log("here");
        $(".popover.instance[data-dc-color]").each(function () {
            console.log($(this).attr("data-dc-color"))
            $(this).css("border-color", $(this).attr("data-dc-color"));
            $(this).css("border-width", 2);
        });
    });    

    activateRefreshTimer();
});
