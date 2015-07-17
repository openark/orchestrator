clusterOperationPseudoGTIDMode = ($.cookie("operation-pgtid-mode") == "true");

var renderColors = ["#ff8c00", "#4682b4", "#9acd32", "#dc143c", "#9932cc", "#ffd700", "#191970", "#7fffd4", "#808080", "#dda0dd"];
var dcColorsMap = {};


function getInstanceDiv(instanceId) {
    var popoverDiv = $("[data-fo-id='" + instanceId + "'] div.popover");
	return popoverDiv
}

function repositionIntanceDivs() {
    $("[data-fo-id]").each(
            function () {
                var id = $(this).attr("data-fo-id");
                var popoverDiv = getInstanceDiv(id);

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
}

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
    	var popoverElement = getInstanceDiv(node.id);
   		renderInstanceElement(popoverElement, node, "cluster");
    });

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
    	$(duplicate).find(".dropdown-toggle").dropdown();
    	$(duplicate).popover();
        $(duplicate).show();
        $(".popover.instance[data-duplicate-node] h3 a").click(function () {
        	openNodeModal(nodesMap[draggedNodeId]);
        	return false;
        });
        $(".popover.instance[data-duplicate-node] a[data-command=recover-auto]").click(function () {
        	apiCommand("/api/recover/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port);
        	return true;
        });
        $(".popover.instance[data-duplicate-node] a[data-command=match-up-slaves]").click(function () {
        	apiCommand("/api/match-up-slaves/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port);
        	return true;
        });
        $(".popover.instance[data-duplicate-node] a[data-command=regroup-slaves]").click(function () {
        	apiCommand("/api/regroup-slaves/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port);
        	return true;
        });
        $(".popover.instance[data-duplicate-node] a[data-command=recover-suggested-successor]").click(function (event) {
            var suggestedSuccessorHost = $(event.target).attr("data-suggested-successor-host");
            var suggestedSuccessorPort = $(event.target).attr("data-suggested-successor-port");
        	apiCommand("/api/recover/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port+"/"+suggestedSuccessorHost+"/"+suggestedSuccessorPort);
        	return true;
        });
        $(".popover.instance[data-duplicate-node] a[data-command=multi-match-slaves]").click(function (event) {
            var belowHost = $(event.target).attr("data-below-host");
            var belowPort = $(event.target).attr("data-below-port");
        	apiCommand("/api/multi-match-slaves/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port+"/"+belowHost+"/"+belowPort);
        	return true;
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
        if (!isAuthorizedForAction() || nodesMap[draggedNodeId].lastCheckInvalidProblem() || nodesMap[draggedNodeId].notRecentlyCheckedProblem() || nodesMap[draggedNodeId].isAggregate) {
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
    var isUsingGTID = (node.usingGTID && droppableNode.usingGTID);
    var gtidBelowFunc = null
	if (clusterOperationPseudoGTIDMode) {
		// definitely peudo-gtid match
		gtidBelowFunc = matchBelow
	} else if (isUsingGTID) {
		//gtidBelowFunc = moveBelow
	}
	if (gtidBelowFunc != null) {
		// Moving via GTID or Pseudo GTID
		if (node.hasConnectivityProblem || droppableNode.hasConnectivityProblem || droppableNode.isAggregate) {
			// Obviously can't handle.
			return null;
		}
		if (!droppableNode.LogSlaveUpdatesEnabled) {
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
		if (instanceIsDescendant(node, droppableNode)) {
			// clearly node cannot be more up to date than droppableNode
			if (shouldApply) {
				gtidBelowFunc(node, droppableNode);
			}
			return "ok";
		}
		if (isReplicationBehindSibling(node, droppableNode)) {
			// verified that node isn't more up to date than droppableNode
			if (shouldApply) {
				gtidBelowFunc(node, droppableNode);
			}
			return "ok";
		}
		// TODO: the general case, where there's no clear family connection, meaning we cannot infer
		// which instance is more up to date. It's under the user's responsibility!
		if (shouldApply) {
			gtidBelowFunc(node, droppableNode);
		}
		return "warning";
	}
	// Not pseudo-GTID mode, non GTID mode
	if (node.isCoMaster) {
		// Cannot move. RESET SLAVE on one of the co-masters.
		return null;
	}
	if (instancesAreSiblings(node, droppableNode)) {
		if (node.hasProblem || droppableNode.hasProblem || droppableNode.isAggregate || !droppableNode.LogSlaveUpdatesEnabled) {
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
	if (instanceIsChild(node, droppableNode) && !droppableNode.isMaster) {
		if (node.hasProblem) {
			// Typically, when a node has a problem we do not allow moving it up.
			// But there's a special situation when allowing is desired: when
			// this slave is completely caught up;
			if (!node.isSQLThreadCaughtUpWithIOThread) { 
				return null;
			}
		}
		if (shouldApply) {
			enslaveMaster(node, droppableNode);
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
	var message = "<h4>move-below</h4>Are you sure you wish to turn <code><strong>" + 
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
	var message = "<h4>move-up</h4>Are you sure you wish to turn <code><strong>" + 
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



function enslaveMaster(node, masterNode) {
	var message = "<h4>enslave-master</h4>Are you sure you wish to make <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> master of <code><strong>" +
		masterNode.Key.Hostname + ":" + masterNode.Key.Port +
		"</strong></code>?"
	bootbox.confirm(message, function(confirm) {
		if (confirm) {
			showLoader();
			var apiUrl = "/api/enslave-master/" + node.Key.Hostname + "/" + node.Key.Port;
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
	var message = "<h4>make-co-master</h4>Are you sure you wish to make <code><strong>" + 
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
	var message = "<h4>PSEUDO-GTID MODE, match-below</h4>Are you sure you wish to turn <code><strong>" + 
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

function compactInstances(instances, instancesMap) {
    instances.forEach(function (instance) {
    	if (instance.children) {
    		// Aggregating children who are childless
        	childlessChildren = instance.children.filter(function(child) {
        		if (child.children && child.children.length > 0) {
        			return false
        		}
    			return true;
    		});
        	if (childlessChildren.length > 1) {
        		// OK, more than one childless child. Aggregate!
        		var aggregatedChild = childlessChildren[0]
        		aggregatedChild.isAggregate = true;
        		aggregatedChild.title = "[aggregation]";
        		aggregatedChild.canonicalTitle = aggregatedChild.title;
        		var aggregatedProblems = {}
        		aggregatedChild.aggregatedInstances = childlessChildren; // includes itself

                function incrementProblems(problemType) {
                	if (aggregatedProblems[problemType] > 0) {
                		aggregatedProblems[problemType] = aggregatedProblems[problemType] + 1;
                	} else {
                		aggregatedProblems[problemType] = 1;
                	}
                }
        		aggregatedChild.aggregatedProblems = aggregatedProblems;
        		
				childlessChildren.forEach(function (instance) {
			        if (instance.inMaintenanceProblem()) {
			        	incrementProblems(instance.ClusterName, "inMaintenanceProblem")
			        }
			        if (instance.lastCheckInvalidProblem()) {
			        	incrementProblems("lastCheckInvalidProblem")
			        } else if (instance.notRecentlyCheckedProblem()) {
			        	incrementProblems("notRecentlyCheckedProblem")
			        } else if (instance.notReplicatingProblem()) {
			        	incrementProblems("notReplicatingProblem")
			        } else if (instance.replicationLagProblem()) {
			        	incrementProblems("replicationLagProblem")
			        }
        		});

				childlessChildren.forEach(function (child) {
        			if (!child.isAggregate) {
        				instance.children.splice( $.inArray(child, instance.children), 1 );
        				delete instancesMap[child.id];
        			}
        		});
    		}

    	}
    });	
    return instancesMap;
}

function analyzeClusterInstances(nodesMap) {
	instances = []
    for (var nodeId in nodesMap) {
    	instances.push(nodesMap[nodeId]);
    } 

    instances.forEach(function (instance) {
	    if (instance.hasConnectivityProblem) {
	    	// The instance has a connectivity problem! Do a client-side recommendation of most advanced slave:
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
        dcColorsMap[knownDCs[i]] = renderColors[i % renderColors.length];
    }
    instances.forEach(function (instance) {
    	var draggedNodeId = $(this).attr("data-nodeid"); 
    	$(".popover.instance[data-nodeid="+instance.id+"]").attr("data-dc-color", dcColorsMap[instance.DataCenter]);
    });
    repositionIntanceDivs()
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

function showOSCSlaves() {
    $.get("/api/cluster-osc-slaves/"+currentClusterName(), function (instances) {
		var instancesMap = normalizeInstances(instances, Array());
		var instancesTitles = Array();
	    instances.forEach(function (instance) {
	    	instancesTitles.push(instance.title);
	    });
	    var instancesTitlesConcatenates = instancesTitles.join(" ");
	    bootbox.alert("Heuristic list of OSC controller slaves: <pre>"+instancesTitlesConcatenates+"</pre>");
    }, "json");
}

function anonymize() {
    var _ = function() {
        var counter = 0;  
        var port = 3306;
        $("h3.popover-title .pull-left").each(function() {
            jQuery(this).html("instance-"+(counter++)+":"+port)
         });
        $(".popover-content p").each(function() {
        	tokens = jQuery(this).html().split(" ", 2)
            jQuery(this).html(tokens[0].match(/[^.]+[.][^.]+/) + " " + tokens[1])
         });
    }();
	$("#cluster_container div.floating_background").html("");	
}

function colorize_dc() {
    $(".popover.instance[data-dc-color]").each(function () {
        $(this).css("border-color", $(this).attr("data-dc-color"));
        $(this).css("border-width", 2);
    });	
}

function addSidebarInfoPopoverContent(content, prepend) {
	if (prepend === true) {
		var wrappedContent = '<div>'+content+'</div>';
		$("#cluster_sidebar [data-bullet=info] [data-toggle=popover]").attr("data-content",
			wrappedContent + $("#cluster_sidebar [data-bullet=info] [data-toggle=popover]").attr("data-content"));
		
	} else {
		var wrappedContent = '<div><hr/>'+content+'</div>';
		$("#cluster_sidebar [data-bullet=info] [data-toggle=popover]").attr("data-content",
			$("#cluster_sidebar [data-bullet=info] [data-toggle=popover]").attr("data-content") + wrappedContent);
	}
}

function populateSidebar(clusterInfo) {
	var content = '';

	{
		var content = 'Alias: '+clusterInfo.ClusterAlias+'';
		addSidebarInfoPopoverContent(content, false);
	}
	{
		var content = 'Heuristic lag: '+clusterInfo.HeuristicLag+'s';
		addSidebarInfoPopoverContent(content, false);
	}
	{
		var content = '';
		if (clusterInfo.HasAutomatedMasterRecovery === true) {
			content += '<span class="glyphicon glyphicon-heart text-info" title="Automated master recovery for this cluster ENABLED"></span>';
		} else {
			content += '<span class="glyphicon glyphicon-heart text-muted" title="Automated master recovery for this cluster DISABLED"></span>';
		}
		if (clusterInfo.HasAutomatedIntermediateMasterRecovery === true) {
			content += '<span class="glyphicon glyphicon-heart-empty text-info" title="Automated intermediate master recovery for this cluster ENABLED"></span>';
		} else {
			content += '<span class="glyphicon glyphicon-heart-empty text-muted" title="Automated intermediate master recovery for this cluster DISABLED"></span>';
		}
		addSidebarInfoPopoverContent(content, true);
	}
}

function onAnalysisEntry(analysisEntry, instance) {
	var content = '<span><strong>'+analysisEntry.Analysis 
		+ (analysisEntry.IsDowntimed ? '<br/>[<i>downtime till '+analysisEntry.DowntimeEndTimestamp+'</i>]': '')
		+ "</strong></span>" 
		+ "<br/>" + "<span>" + analysisEntry.AnalyzedInstanceKey.Hostname+":"+analysisEntry.AnalyzedInstanceKey.Port+ "</span>" 
	;
	addSidebarInfoPopoverContent(content);        
	
	var popoverElement = getInstanceDiv(instance.id);

	popoverElement.append('<h4 class="popover-footer"><div class="dropdown"></div></h4>');
	popoverElement.find(".popover-footer .dropdown").append('<button type="button" class="btn btn-xs btn-default dropdown-toggle" id="recover_dropdown_'+instance.id+'" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true"><span class="glyphicon glyphicon-heart text-danger"></span> Recover <span class="caret"></span></button><ul class="dropdown-menu" aria-labelledby="recover_dropdown_'+instance.id+'"></ul>');
	popoverElement.find(".popover-footer .dropdown").append('<ul class="dropdown-menu" aria-labelledby="recover_dropdown_'+instance.id+'"></ul>');
	var recoveryListing = popoverElement.find(".dropdown ul");
    recoveryListing.append('<li><a href="#" data-btn="auto" data-command="recover-auto">Auto (implies running external hooks/processes)</a></li>');
    recoveryListing.append('<li role="separator" class="divider"></li>');
    
    if (!instance.isMaster) {
        recoveryListing.append('<li><a href="#" data-btn="match-up-slaves" data-command="match-up-slaves">Match up slaves to <code>'+instance.masterTitle+'</code></a></li>');
    }
    if (instance.children.length > 1) {
    	recoveryListing.append('<li><a href="#" data-btn="regroup-slaves" data-command="regroup-slaves">Regroup slaves (auto pick best slave, only heals topology, no external processes)</a></li>');
    }
    if (instance.isMaster) {
    	// Suggest successor
	    instance.children.forEach(function(slave) {
            if (!slave.LogBinEnabled) {
                return
            }
            if (slave.SQLDelay > 0) {
                return
            }
            if (!slave.LogSlaveUpdatesEnabled) {
                return
            }
            if (slave.lastCheckInvalidProblem()) {
                return
            }
            if (slave.notRecentlyCheckedProblem()) {
                return
            }
            recoveryListing.append(
                '<li><a href="#" data-btn="recover-suggested-successor" data-command="recover-suggested-successor" data-suggested-successor-host="'+slave.Key.Hostname
                +'" data-suggested-successor-port="'+slave.Key.Port+'">Regroup slaves, try to promote <code>'+slave.title+'</code></a></li>');
        });                 
    }
    if (instance.masterNode) {
    	// Intermediate master; suggest successor
	    instance.masterNode.children.forEach(function(sibling) {
            if (sibling.id == instance.id) {
                return
            }
            if (!sibling.LogBinEnabled) {
                return
            }
            if (!sibling.LogSlaveUpdatesEnabled) {
                return
            }
            if (sibling.lastCheckInvalidProblem()) {
                return
            }
            if (sibling.notRecentlyCheckedProblem()) {
                return
            }
            recoveryListing.append(
                '<li><a href="#" data-btn="multi-match-slaves" data-command="multi-match-slaves" data-below-host="'+sibling.Key.Hostname
                +'" data-below-port="'+sibling.Key.Port+'">Match all slaves below <code>'+sibling.title+'</code></a></li>');
        });                 
    }
}


function reviewReplicationAnalysis(replicationAnalysis, instancesMap) {
	var clusterHasReplicationAnalysisIssue = false;
    replicationAnalysis.Details.forEach(function (analysisEntry) {
    	if (!(analysisEntry.Analysis in interestingAnalysis)) {
    		return;
    	}
    	if (analysisEntry.ClusterDetails.ClusterName == currentClusterName()) {
    		clusterHasReplicationAnalysisIssue = true;
    		
    		var instanceId = getInstanceId(analysisEntry.AnalyzedInstanceKey.Hostname, analysisEntry.AnalyzedInstanceKey.Port);
    		var instance = instancesMap[instanceId]
    		onAnalysisEntry(analysisEntry, instance);
    	}
    });
    if (clusterHasReplicationAnalysisIssue) {
		$("#cluster_sidebar [data-bullet=info] div span").addClass("text-danger");
    } else {
    	$("#cluster_sidebar [data-bullet=info] div span").addClass("text-info");
    }	
}


function indicateClusterPoolInstances(clusterPoolInstances, instancesMap) {
	for (var pool in clusterPoolInstances.Details) {
		if (clusterPoolInstances.Details.hasOwnProperty(pool)) {
			clusterPoolInstances.Details[pool].forEach(function(instanceKey) {
				var instanceId = getInstanceId(instanceKey.Hostname, instanceKey.Port)
				var instance  = instancesMap[instanceId];
				if (!instance.IsInPool) {
					instance.IsInPool = true;
					getInstanceDiv(instance.id).find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-tint" title="pools:"></span> ');
				}
				var indicatorElement = getInstanceDiv(instance.id).find("h3 div.pull-right span.glyphicon-tint");
				indicatorElement.attr("title", indicatorElement.attr("title") + " " + pool);
			});
		}
	}
}

$(document).ready(function () {
    $.get("/api/cluster/"+currentClusterName(), function (instances) {
        $.get("/api/replication-analysis", function (replicationAnalysis) {
        	$.get("/api/maintenance", function (maintenanceList) {
        		var instancesMap = normalizeInstances(instances, maintenanceList);
        	    if (isCompactDisplay()) {
        	    	instancesMap = compactInstances(instances, instancesMap);
        	    }
                analyzeClusterInstances(instancesMap);
                visualizeInstances(instancesMap);
                generateInstanceDivs(instancesMap);
                reviewReplicationAnalysis(replicationAnalysis, instancesMap);
                postVisualizeInstances(instancesMap);
                if ($.cookie("anonymize") == "true") {
                	anonymize();
                }
                if ($.cookie("colorize-dc") == "true") {
                	colorize_dc();
                }
                
                instances.forEach(function (instance) {
                	if (instance.isMaster) {
                	    $.get("/api/recently-active-instance-recovery/"+ instance.Key.Hostname + "/" + instance.Key.Port, function (recoveries) {
                	        // Result is an array: either empty (no active recovery) or with multiple entries
                	    	recoveries.forEach(function (recoveryEntry) {
                	    		addInfo("<strong>" + instance.title + "</strong> has just recently been promoted as result of <strong>" + recoveryEntry.AnalysisEntry.Analysis + "</strong>. It may still take some time to rebuild topology graph.");
                	        });
                	    }, "json");                		
                    }
                });            
                if ($.cookie("pool-indicator") == "true") {
                    $.get("/api/cluster-pool-instances/" + currentClusterName(), function (clusterPoolInstances) {
               	    	indicateClusterPoolInstances(clusterPoolInstances, instancesMap);
                    }, "json");                	
                }
        	}, "json");
    	}, "json");
    }, "json");
    $.get("/api/cluster-info/"+currentClusterName(), function (clusterInfo) {    
    	var alias = clusterInfo.ClusterAlias
    	var visualAlias = (alias ? alias : currentClusterName())
    	document.title = document.title.split(" - ")[0] + " - " + visualAlias;
    	$("#cluster_container").append('<div class="floating_background">'+visualAlias+'</div>');
        $("#dropdown-context").append('<li><a data-command="change-cluster-alias" data-alias="'+clusterInfo.ClusterAlias+'">Alias: '+alias+'</a></li>');
        $("#dropdown-context").append('<li><a href="/web/cluster-pools/'+currentClusterName()+'">Pools</a></li>');    
        if (isCompactDisplay()) {
            $("#dropdown-context").append('<li><a data-command="expand-display" href="'+location.href.split("?")[0]+'?compact=false">Expand display</a></li>');    
        } else {
            $("#dropdown-context").append('<li><a data-command="compact-display" href="'+location.href.split("?")[0]+'?compact=true">Compact display</a></li>');    
        }
        $("#dropdown-context").append('<li><a data-command="pool-indicator">Pool indicator</a></li>');    
        $("#dropdown-context").append('<li><a data-command="colorize-dc">Colorize DC</a></li>');    
        $("#dropdown-context").append('<li><a data-command="anonymize">Anonymize</a></li>');    
        if ($.cookie("pool-indicator") == "true") {
        	$("#dropdown-context a[data-command=pool-indicator]").prepend('<span class="glyphicon glyphicon-ok small"></span> ');
        } 
        if ($.cookie("anonymize") == "true") {
        	$("#dropdown-context a[data-command=anonymize]").prepend('<span class="glyphicon glyphicon-ok small"></span> ');
        } 
        if ($.cookie("colorize-dc") == "true") {
        	$("#dropdown-context a[data-command=colorize-dc]").prepend('<span class="glyphicon glyphicon-ok small"></span> ');
        } 
        populateSidebar(clusterInfo);

    }, "json");
    
    $.get("/api/active-cluster-recovery/"+currentClusterName(), function (recoveries) {
        // Result is an array: either empty (no active recovery) or with multiple entries
    	recoveries.forEach(function (recoveryEntry) {
    		addInfo("<strong>" + recoveryEntry.AnalysisEntry.Analysis + " active recovery in progress.</strong> Topology is subject to change in the next moments.");
        });
    }, "json");
    $.get("/api/recently-active-cluster-recovery/"+currentClusterName(), function (recoveries) {
        // Result is an array: either empty (no active recovery) or with multiple entries
    	recoveries.forEach(function (recoveryEntry) {
    		addInfo("This cluster just recently recovered from <strong>" + recoveryEntry.AnalysisEntry.Analysis + "</strong>. It may still take some time to rebuild topology graph.");
        });
    }, "json");
    $.get("/api/cluster-osc-slaves/"+currentClusterName(), function (instances) {
		var instancesMap = normalizeInstances(instances, Array());
		var instancesTitles = Array();
	    instances.forEach(function (instance) {
	    	instancesTitles.push(instance.title);
	    });
	    var instancesTitlesConcatenates = instancesTitles.join(" ");
    	var content = "Heuristic list of OSC controller slaves: <pre>"+instancesTitlesConcatenates+"</pre>"; 
		;
    	addSidebarInfoPopoverContent(content);
    }, "json");
    
    if (isPseudoGTIDModeEnabled()) {
        $("ul.navbar-nav").append('<li><a class="cluster_operation_mode"><button type="button" class="btn btn-xs" id="cluster_operation_mode_button"></button></a></li>');
        refreshClusterOperationModeButton();
        
	    $("body").on("click", "#cluster_operation_mode_button", function() {
	    	clusterOperationPseudoGTIDMode = !clusterOperationPseudoGTIDMode;
	    	$.cookie("operation-pgtid-mode", ""+clusterOperationPseudoGTIDMode, { path: '/', expires: 1 });
	    	refreshClusterOperationModeButton(); 
	    });
    }
    $("#instance_problems_button").attr("title", "Cluster Problems");
    
    $("body").on("click", "a[data-command=change-cluster-alias]", function(event) {    	
    	promptForAlias($(event.target).attr("data-alias"));
    });    
    $("body").on("click", "a[data-command=cluster-osc-slaves]", function(event) {    	
    	showOSCSlaves();
    });    
    $("body").on("click", "a[data-command=pool-indicator]", function(event) {
    	if ($.cookie("pool-indicator") == "true") {
    		$.cookie("pool-indicator", "false", { path: '/', expires: 1 });
    		location.reload();
    		return
        }
    	$.cookie("pool-indicator", "true", { path: '/', expires: 1 });
		location.reload();
    });    
    $("body").on("click", "a[data-command=anonymize]", function(event) {
    	if ($.cookie("anonymize") == "true") {
    		$.cookie("anonymize", "false", { path: '/', expires: 1 });
    		location.reload();
    		return
        }
    	anonymize();
    	$("#dropdown-context a[data-command=anonymize]").prepend('<span class="glyphicon glyphicon-ok small"></span> ');
    	$.cookie("anonymize", "true", { path: '/', expires: 1 });
    });    
    $("body").on("click", "a[data-command=colorize-dc]", function(event) {
    	if ($.cookie("colorize-dc") == "true") {
    		$.cookie("colorize-dc", "false", { path: '/', expires: 1 });
    		location.reload();
    		return
        }
    	colorize_dc();
    	$("#dropdown-context a[data-command=colorize-dc]").prepend('<span class="glyphicon glyphicon-ok small"></span> ');
    	$.cookie("colorize-dc", "true", { path: '/', expires: 1 });
    });    

    $("[data-toggle=popover]").popover();
    $("[data-toggle=popover]").show();
    
    if (isAuthorizedForAction()) {
    	// Read-only users don't get auto-refresh. Sorry!
    	activateRefreshTimer();
    }
});
