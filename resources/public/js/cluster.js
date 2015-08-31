var moveInstanceMethod = $.cookie("move-instance-method") || "smart";
var droppableIsActive = false;

var renderColors = ["#ff8c00", "#4682b4", "#9acd32", "#dc143c", "#9932cc", "#ffd700", "#191970", "#7fffd4", "#808080", "#dda0dd"];
var dcColorsMap = {};


function getInstanceDiv(instanceId) {
    var popoverDiv = $("#cluster_container .popover.instance[data-nodeid='" + instanceId + "']");
	return popoverDiv
}

function repositionIntanceDiv(id) {
    if (!id) {
    	return false;
    }
 
    var popoverDiv = getInstanceDiv(id);
    {
	    var pos = getSvgPos($(popoverDiv).data("svg-instance-wrapper"));
	    pos.left += 20;
	    pos.top -= popoverDiv.height()/2 + 2;
	    popoverDiv.css({left:pos.left+"px", top:pos.top+"px"});
	
	    popoverDiv.popover();
	    popoverDiv.show();
    }
    $(".instance-trailer[data-nodeid='" + id + "']").each(function() {
    	var pos = popoverDiv.position();
    	$(this).css({top: pos.top, left: pos.left + popoverDiv.outerWidth(true)});
		$(this).css("height", popoverDiv.outerHeight(true));
	});
}

function repositionIntanceDivs() {
    $("[data-fo-id]").each(function () {
    	repositionIntanceDiv($(this).attr("data-fo-id"));
    });
}

// All instance dragging/dropping code goes here
function activateInstanceDraggableDroppable(nodesMap, draggedNodeId, originalNode, duplicate) {
	function clearDroppable() {
		$(".original-dragged").removeClass("original-dragged");
		resetRefreshTimer();
		$("#cluster_container .accept_drop").removeClass("accept_drop");
		$("#cluster_container .accept_drop_warning").removeClass("accept_drop_warning");
		$(".being-dragged").removeClass("being-dragged");
		$(".instance-trailer").show();
		droppableIsActive = false;
	}
    $(duplicate).draggable({
    	addClasses: true, 
    	opacity: 1,
    	cancel: "#cluster_container .popover.instance h3 a,.instance-trailer",
    	snap: "#cluster_container .popover.instance",
    	snapMode: "inner",
    	snapTolerance: 10,
    	containment: $("#cluster_container"),
    	start: function(event, ui) {
    		// dragging begins
    		clearDroppable();
    		droppableIsActive = true;
    		originalNode.addClass("original-dragged");
    		if (originalNode.data("instance-trailer")) {
    			originalNode.data("instance-trailer").addClass("original-dragged");
    		}
    		duplicate.addClass("being-dragged");
     		$("#cluster_container .popover.instance").droppable({
    			accept: function(draggable) {
    				// Find the objects that accept a draggable (i.e. valid droppables)
    				if (!droppableIsActive) {
    					return false
    				}
    				var draggedNode = nodesMap[draggedNodeId];
    				var targetNode = nodesMap[$(this).attr("data-nodeid")];
    				var acceptDrop = moveInstance(draggedNode, targetNode, false);
    				if (acceptDrop.accept == "ok") {
    					$(this).addClass("accept_drop");
    				}
    				if (acceptDrop.accept == "warning") {
    					$(this).addClass("accept_drop_warning");
    				}
   					$(this).attr("data-drop-comment", acceptDrop.accept ? acceptDrop.type : "");
    				return acceptDrop.accept != null;
    			},
    			hoverClass: "draggable-hovers",
			    over: function( event, ui ) {
			    	// Called once when dragged object is over another object
			    	if ($(this).attr("data-drop-comment")) {
			    		$(duplicate).addClass("draggable-msg");
			    		$(duplicate).find(".popover-content").html($(this).attr("data-drop-comment"))
			    	} else {
			    		$(duplicate).find(".popover-content").html("Cannot drop here")
			    	}
			    },
			    out: function( event, ui ) {
			    	// Called once when dragged object leaves other object
		    		$(duplicate).removeClass("draggable-msg");
			    	$(duplicate).find(".popover-content").html("")
			    },
			    drop: function( event, ui ) {
		            $(".popover.instance[data-duplicate-node]").remove();
		            moveInstance(nodesMap[draggedNodeId], nodesMap[$(this).attr("data-nodeid")], true);
		            clearDroppable();
			    }
    		});
    	},
    	drag: function(event, ui) {
    		resetRefreshTimer();
    	},
    	stop: function(event, ui) {
    		clearDroppable();
    	}
    });
	$(duplicate).on("mouseleave", function() {
		if (!$(this).hasClass("ui-draggable-dragging")) {
    		$(this).remove();
		}
	});
	// Don't ask why the following... jqueryUI recognizes the click as start drag, but fails to stop...
	$(duplicate).on("click", function() {
		clearDroppable();
    	return false;
    });		
}


//All instance dragging/dropping code goes here
function activateInstanceChildrenDraggableDroppable(nodesMap, draggedNodeId, originalNode, duplicate) {
	function clearDroppable() {
		$(".original-dragged").removeClass("original-dragged");
		$(".instance-trailer").show();
		resetRefreshTimer();
		$("#cluster_container .accept_drop").removeClass("accept_drop");
		$("#cluster_container .accept_drop_warning").removeClass("accept_drop_warning");
		$(".being-dragged").removeClass("being-dragged");
		droppableIsActive = false;
	}
	var draggedNode = nodesMap[draggedNodeId];
	$(duplicate).draggable({
	 	addClasses: true, 
	 	opacity: 1,
	 	snap: "#cluster_container .popover.instance",
	 	snapMode: "inner",
	 	snapTolerance: 10,
	 	start: function(event, ui) {
            //TODO: $(event.target).closest(".popup").is(".dfsdfsd")
	 		// dragging begins
	 		clearDroppable();
	 		droppableIsActive = true;

	 		draggedNode.children.forEach(function f(instance) {
	 			var popoverElement = getInstanceDiv(instance.id);
	 			popoverElement.addClass("original-dragged");
	 		});
            duplicate.addClass("being-dragged");
	 		$(".instance-trailer[data-nodeid="+draggedNodeId+"]").each(function() {
	 			if ($(this).is("[data-duplicate-node]")) {
	 				return;
	 			}
	 			$(this).addClass("original-dragged");
	 		});
	 		
	  		$("#cluster_container .popover.instance").droppable({
	 			accept: function(draggable) {
	 				// Find the objects that accept a draggable (i.e. valid droppables)
	 				if (!droppableIsActive) {
	 					return false
	 				}
	 				if ($(this).is("[data-duplicate-node]")) {
	 					// The duplicate of the node whose children we are dragging.
	 					return false;
	 				}
	 				var targetNode = nodesMap[$(this).attr("data-nodeid")];
	 				
	 				var acceptDrop = moveChildren(draggedNode, targetNode, false);
	 				if (acceptDrop.accept == "ok") {
	 					$(this).addClass("accept_drop");
	 				}
	 				if (acceptDrop.accept == "warning") {
	 					$(this).addClass("accept_drop_warning");
	 				}
	 				$(this).attr("data-drop-comment", acceptDrop.accept ? acceptDrop.type : "");
	 				return acceptDrop.accept != null;
	 			},
	 			hoverClass: "draggable-hovers",
				    over: function( event, ui ) {
				    	// Called once when dragged object is over another object
				    	if ($(this).attr("data-drop-comment")) {
				    		$(duplicate).addClass("draggable-msg");
				    		$(duplicate).find(".popover-content").html($(this).attr("data-drop-comment"))
				    	} else {
				    		$(duplicate).find(".popover-content").html("Cannot drop here")
				    	}
				    },
				    out: function( event, ui ) {
				    	// Called once when dragged object leaves other object
			    		$(duplicate).removeClass("draggable-msg");
				    	$(duplicate).find(".popover-content").html("")
				    },
				    drop: function( event, ui ) {
			            $(".instance-trailer[data-duplicate-node]").remove();
			            moveChildren(nodesMap[draggedNodeId], nodesMap[$(this).attr("data-nodeid")], true);
			            clearDroppable();
				    }
	 		});
	 	},
	 	drag: function(event, ui) {
	 		resetRefreshTimer();
	 	},
	 	stop: function(event, ui) {
	 		clearDroppable();
	 	}
	});
	$(duplicate).on("mouseleave", function() {
		if (!$(this).hasClass("ui-draggable-dragging")) {
			$(this).remove();
		}
	});
	// Don't ask why the following... jqueryUI recognizes the click as start drag, but fails to stop...
	$(duplicate).on("click", function() {
		clearDroppable();
 	return false;
 });		
}

function generateInstanceDiv(svgInstanceWrapper, nodesMap) {
	var isVirtual = $(svgInstanceWrapper).attr("data-fo-is-virtual") == "true";
	var isAnchor = $(svgInstanceWrapper).attr("data-fo-is-anchor") == "true";
    if (!isVirtual && !isAnchor) {
    	var popoverElement = $('<div class="popover right instance"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>').appendTo("#cluster_container");
    	$(popoverElement).hide();
    	$(svgInstanceWrapper).data("instance-popover", popoverElement);
    	$(popoverElement).data("svg-instance-wrapper", svgInstanceWrapper);
    	
    	var id = $(svgInstanceWrapper).attr("data-fo-id");
    	var node = nodesMap[id];
		renderInstanceElement(popoverElement, node, "cluster");
		if (node.children) {
            var trailerElement = $('<div class="popover left instance-trailer" data-nodeid="'+node.id+'"><div><span class="glyphicon glyphicon-chevron-left" title="Drag and drop slaves of this instance"></span></div></div>').appendTo("#cluster_container");
            popoverElement.data("instance-trailer", trailerElement);
		}
    }	
}


function prepareDraggable(nodesMap) {
    $("body").on("mouseenter", "#cluster_container .popover.instance[data-nodeid]", function() {
    	if ($("[data-duplicate-node]").hasClass("ui-draggable-dragging")) {
    		// Do not remove & recreate while dragging. Ignore any mouseenter
    		return false;
    	}
    	var draggedNodeId = $(this).attr("data-nodeid"); 
    	if (draggedNodeId == $(".popover.instance[data-duplicate-node]").attr("data-nodeid")) {
    		return false;
    	}
    	$(".popover.instance[data-duplicate-node]").remove();
    	var originalNode = $(this);
    	var duplicate = originalNode.clone().appendTo("#cluster_container");
    	$(duplicate).attr("data-duplicate-node", "true");
    	$(duplicate).css({top:originalNode[0].style.top, left:originalNode[0].style.left});
    	$(duplicate).width(originalNode.width());
    	$(duplicate).height(originalNode.height());
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
        $(".popover.instance[data-duplicate-node] a[data-command=recover-auto-lite]").click(function () {
        	apiCommand("/api/recover-lite/"+nodesMap[draggedNodeId].Key.Hostname+"/"+nodesMap[draggedNodeId].Key.Port);
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
        	activateInstanceDraggableDroppable(nodesMap, draggedNodeId, originalNode, duplicate);
        }
    });
    $("body").on("mouseenter", ".instance-trailer[data-nodeid]", function() {
    	if ($("[data-duplicate-node]").hasClass("ui-draggable-dragging")) {
    		// Do not remove & recreate while dragging. Ignore any mouseenter
    		return false;
    	}
    	var draggedNodeId = $(this).attr("data-nodeid");
    	if (draggedNodeId == $(".instance-trailer[data-duplicate-node]").attr("data-nodeid")) {
    		return false;
    	}

    	$(".instance-trailer[data-duplicate-node]").remove();
    	var trailerElement = $(this); // The trailer, a small div
    	var instanceElement = getInstanceDiv(draggedNodeId); // However we copy the instance node element
    	var duplicate = trailerElement.clone().appendTo("#cluster_container");
    	$(duplicate).attr("data-duplicate-node", "true");
    	//$(duplicate).css({"margin-left": "0"});
    	//$(duplicate).css({top:instanceElement[0].style.top, left:instanceElement[0].offsetLeft + trailerElement[0].offsetLeft + 50});
    	//$(duplicate).css({right:"auto"});
    	//$(duplicate).html($(instanceElement).html());
    	$(duplicate).width(instanceElement.width());
    	$(duplicate).height(instanceElement.height());
    	var numSlaves = nodesMap[draggedNodeId].children.length;
    	var numSlavesMessage = ((numSlaves == 1)? "1 slave" : ""+numSlaves+" slaves");
    	$(duplicate).find("div").append(" "+numSlavesMessage);
    	$(duplicate).find("h3 .pull-left").html(numSlavesMessage);
    	$(duplicate).find("h3 .pull-right").html('<span class="glyphicon glyphicon-chevron-left"></span>');
    	$(duplicate).append('<div class="popover-content"></div>');
    	$(duplicate).find(".popover-content").html("Drag to move slaves");
    	$(duplicate).find(".popover-footer").remove();
    	$(duplicate).popover();
        $(duplicate).show();

        if (!isAuthorizedForAction()) {
        	return false;
        }
        activateInstanceChildrenDraggableDroppable(nodesMap, draggedNodeId, trailerElement, duplicate);
    });
}

// moveInstance checks whether an instance (node) can be dropped on another (droppableNode).
// The function consults with the current moveInstanceMethod; the type of action taken is based on that.
// For example, actions can be repoint, match-below, relocate, move-up, enslave-master etc.
// When shouldApply is false nothing gets executed, and the function merely serves as a predictive
// to the possibility of the drop.
function moveInstance(node, droppableNode, shouldApply) {
    if (!isAuthorizedForAction()) {
    	// Obviously this is also checked on server side, no need to try stupid hacks
		return {accept: false} ;
    }
    var droppableTitle = getInstanceDiv(droppableNode.id).find("h3 .pull-left").html();
    var isUsingGTID = (node.usingGTID && droppableNode.usingGTID);
	if (moveInstanceMethod == "smart") {
		// Moving via GTID or Pseudo GTID
		if (node.hasConnectivityProblem || droppableNode.hasConnectivityProblem || droppableNode.isAggregate) {
			// Obviously can't handle.
			return {accept: false};
		}
		if (!droppableNode.LogSlaveUpdatesEnabled) {
			// Obviously can't handle.
			return {accept: false};
		}
		
		if (node.id == droppableNode.id) {
			return {accept: false};
		}
		if (instanceIsDescendant(droppableNode, node)) {
			// Wrong direction!
			return {accept: false};
		}
		// the general case
		if (shouldApply) {
			relocate(node, droppableNode);
		}
		return {accept: "warning", type: "relocate < " + droppableTitle};
	}
	if (moveInstanceMethod == "pseudo-gtid") {
		var gtidBelowFunc = matchBelow
		//~~~TODO: when GTID is fully supported: gtidBelowFunc = moveBelow
		// Moving via GTID or Pseudo GTID
		if (node.hasConnectivityProblem || droppableNode.hasConnectivityProblem || droppableNode.isAggregate) {
			// Obviously can't handle.
			return {accept: false};
		}
		if (!droppableNode.LogSlaveUpdatesEnabled) {
			// Obviously can't handle.
			return {accept: false};
		}
		
		if (node.id == droppableNode.id) {
			return {accept: false};
		}
		if (instanceIsDescendant(droppableNode, node)) {
			// Wrong direction!
			return {accept: false};
		}
		if (instanceIsDescendant(node, droppableNode)) {
			// clearly node cannot be more up to date than droppableNode
			if (shouldApply) {
				gtidBelowFunc(node, droppableNode);
			}
			return {accept: "ok", type: gtidBelowFunc.name + " " + droppableTitle};
		}
		if (isReplicationBehindSibling(node, droppableNode)) {
			// verified that node isn't more up to date than droppableNode
			if (shouldApply) {
				gtidBelowFunc(node, droppableNode);
			}
			return {accept: "ok", type: gtidBelowFunc.name + " " + droppableTitle};
		}
		// TODO: the general case, where there's no clear family connection, meaning we cannot infer
		// which instance is more up to date. It's under the user's responsibility!
		if (shouldApply) {
			gtidBelowFunc(node, droppableNode);
		}
		return {accept: "warning", type: gtidBelowFunc.name + " " + droppableTitle};
	}
	if (moveInstanceMethod == "classic") {
		// Not pseudo-GTID mode, non GTID mode
		if (node.id == droppableNode.id) {
			return {accept: false};
		}
		if (node.isCoMaster) {
			// Cannot move. RESET SLAVE on one of the co-masters.
			return {accept: false};
		}
		if (instancesAreSiblings(node, droppableNode)) {
			if (node.hasProblem || droppableNode.hasProblem || droppableNode.isAggregate || !droppableNode.LogSlaveUpdatesEnabled) {
				return {accept: false};
			}
			if (shouldApply) {
				moveBelow(node, droppableNode);
			}
			return {accept: "ok", type: "moveBelow " + droppableTitle};
		}
		if (instanceIsGrandchild(node, droppableNode)) {
			if (node.hasProblem) {
				// Typically, when a node has a problem we do not allow moving it up.
				// But there's a special situation when allowing is desired: when the parent has personal issues,
				// (say disk issue or otherwise something heavyweight running which slows down replication)
				// and you want to move up the slave which is only delayed by its master.
				// So to help out, if the instance is identically at its master's trail, it is allowed to move up.
				if (!node.isSQLThreadCaughtUpWithIOThread) { 
					return {accept: false};
				}
			}
			if (shouldApply) {
				moveUp(node, droppableNode);
			}
			return {accept: "ok", type: "moveUp under " + droppableTitle};
		}
		if (instanceIsChild(node, droppableNode) && !droppableNode.isMaster) {
			if (node.hasProblem) {
				// Typically, when a node has a problem we do not allow moving it up.
				// But there's a special situation when allowing is desired: when
				// this slave is completely caught up;
				if (!node.isSQLThreadCaughtUpWithIOThread) { 
					return {accept: false};
				}
			}
			if (shouldApply) {
				enslaveMaster(node, droppableNode);
			}
			return {accept: "ok", type: "enslaveMaster " + droppableTitle};
		}
		if (instanceIsChild(droppableNode, node) && node.isMaster) {
			if (node.hasProblem) {
				return {accept: false};
			}
			if (shouldApply) {
				makeCoMaster(node, droppableNode);
			}
			return {accept: "ok", type: "makeCoMaster with " + droppableTitle};
		}
		return {accept: false};
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
	return {accept: false};
}


// moveChildren checks whether an children of an instance (node) can be dropped on another (droppableNode).
// The function consults with the current moveInstanceMethod; the type of action taken is based on that.
// For example, actions can be repoint-slaves, multi-match-slaves, relocate-slaves, move-up-slaves etc.
// When shouldApply is false nothing gets executed, and the function merely serves as a predictive
// to the possibility of the drop.
function moveChildren(node, droppableNode, shouldApply) {
    if (!isAuthorizedForAction()) {
    	// Obviously this is also checked on server side, no need to try stupid hacks
		return {accept: false} ;
    }
    var droppableTitle = getInstanceDiv(droppableNode.id).find("h3 .pull-left").html();
    var isUsingGTID = (node.usingGTID && droppableNode.usingGTID);
	if (moveInstanceMethod == "smart") {
		// Moving via GTID or Pseudo GTID
		if (droppableNode.hasConnectivityProblem || droppableNode.isAggregate) {
			// Obviously can't handle.
			return {accept: false};
		}
		if (!droppableNode.LogSlaveUpdatesEnabled) {
			// Obviously can't handle.
			return {accept: false};
		}
		
		if (node.id == droppableNode.id) {
			if (shouldApply) {
				relocateSlaves(node, droppableNode);
			}
			return {accept: "ok", type: "relocate < " + droppableTitle};
		}
		if (instanceIsDescendant(droppableNode, node) && node.children.length <= 1) {
			// Can generally move slaves onto one of them, but there needs to be at least two slaves...
			// Otherwise we;re trying to mvoe a slave under itself which is clearly an error.
			return {accept: false};
		}
		// the general case
		if (shouldApply) {
			relocateSlaves(node, droppableNode);
		}
		return {accept: "warning", type: "relocate < " + droppableTitle};
	}

	if (moveInstanceMethod == "pseudo-gtid") {
		var gtidBelowFunc = matchSlaves
		//~~~TODO: when GTID is fully supported: gtidBelowFunc = moveBelow
		// Moving via GTID or Pseudo GTID
		if (droppableNode.hasConnectivityProblem || droppableNode.isAggregate) {
			// Obviously can't handle.
			return {accept: false};
		}
		if (!droppableNode.LogSlaveUpdatesEnabled) {
			// Obviously can't handle.
			return {accept: false};
		}		
		if (node.id == droppableNode.id) {
			if (shouldApply) {
				gtidBelowFunc(node, droppableNode);
			}
			return {accept: "ok", type: gtidBelowFunc.name + " < " + droppableTitle};
		}
		if (instanceIsDescendant(droppableNode, node) && node.children.length <= 1) {
			// Can generally move slaves onto one of them, but there needs to be at least two slaves...
			// Otherwise we;re trying to mvoe a slave under itself which is clearly an error.
			// Wrong direction!
			return {accept: false};
		}
		if (instanceIsDescendant(node, droppableNode)) {
			// clearly node cannot be more up to date than droppableNode
			if (shouldApply) {
				gtidBelowFunc(node, droppableNode);
			}
			return {accept: "ok", type: gtidBelowFunc.name + " < " + droppableTitle};
		}
		// TODO: the general case, where there's no clear family connection, meaning we cannot infer
		// which instance is more up to date. It's under the user's responsibility!
		if (shouldApply) {
			gtidBelowFunc(node, droppableNode);
		}
		return {accept: "warning", type: gtidBelowFunc.name + " < " + droppableTitle};
	}
	if (moveInstanceMethod == "classic") {
		// Not pseudo-GTID mode, non GTID mode
		if (node.id == droppableNode.id) {
			if (shouldApply) {
				repointSlaves(node);
			}
			return {accept: "ok", type: "repointSlaves < " + droppableTitle};
		}
		if (instanceIsChild(node, droppableNode)) {
			if (shouldApply) {
				moveUpSlaves(node, droppableNode);
			}
			return {accept: "ok", type: "moveUpSlaves < " + droppableTitle};
		}
		return {accept: false};
	}
	if (shouldApply) {
		addAlert(
				"Cannot move slaves of <code><strong>" + 
					node.Key.Hostname + ":" + node.Key.Port +
					"</strong></code> under <code><strong>" +
					droppableNode.Key.Hostname + ":" + droppableNode.Key.Port +
					"</strong></code>. " +
				"You may only repoint or move up the slaves of an instance. Otherwise try Smart Mode."
			);
	}
	return {accept: false};
}


function executeMoveOperation(message, apiUrl) {
	bootbox.confirm(anonymizeIfNeedBe(message), function(confirm) {
		if (confirm) {
			showLoader();
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

function relocate(node, siblingNode) {
	var message = "<h4>relocate</h4>Are you sure you wish to turn <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> into a slave of <code><strong>" +
		siblingNode.Key.Hostname + ":" + siblingNode.Key.Port +
		"</strong></code>?"+
		"<h4>Note</h4><p>Orchestrator will try and figure out the best relocation path. This may involve multiple steps. " +
		"<p>In case multiple steps are involved, failure of one would leave your instance hanging in a different location than you expected, " +
		"but it would still be in a <i>valid</i> state.";
	var apiUrl = "/api/relocate/" + node.Key.Hostname + "/" + node.Key.Port + "/" + siblingNode.Key.Hostname + "/" + siblingNode.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

function relocateSlaves(node, siblingNode) {
	var message = "<h4>relocate-slaves</h4>Are you sure you wish to relocate slaves of <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> below <code><strong>" +
		siblingNode.Key.Hostname + ":" + siblingNode.Key.Port +
		"</strong></code>?"+
		"<h4>Note</h4><p>Orchestrator will try and figure out the best relocation path. This may involve multiple steps. " +
		"<p>In case multiple steps are involved, failure of one may leave some instances hanging in a different location than you expected, " +
		"but they would still be in a <i>valid</i> state.";
	var apiUrl = "/api/relocate-slaves/" + node.Key.Hostname + "/" + node.Key.Port + "/" + siblingNode.Key.Hostname + "/" + siblingNode.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

function repointSlaves(node, siblingNode) {
	var message = "<h4>repoint-slaves</h4>Are you sure you wish to repoint slaves of <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code>?";
	var apiUrl = "/api/repoint-slaves/" + node.Key.Hostname + "/" + node.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

function moveUpSlaves(node, masterNode) {
	var message = "<h4>move-up-slaves</h4>Are you sure you wish to move up slaves of <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> below <code><strong>" +
		masterNode.Key.Hostname + ":" + masterNode.Key.Port +
		"</strong></code>?";
	var apiUrl = "/api/move-up-slaves/" + node.Key.Hostname + "/" + node.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

function matchSlaves(node, otherNode) {
	var message = "<h4>multi-match-slaves</h4>Are you sure you wish to match slaves of <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> below <code><strong>" +
		otherNode.Key.Hostname + ":" + otherNode.Key.Port +
		"</strong></code>?";
	var apiUrl = "/api/multi-match-slaves/" + node.Key.Hostname + "/" + node.Key.Port + "/" + otherNode.Key.Hostname + "/" + otherNode.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

function moveBelow(node, siblingNode) {
	var message = "<h4>move-below</h4>Are you sure you wish to turn <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> into a slave of <code><strong>" +
		siblingNode.Key.Hostname + ":" + siblingNode.Key.Port +
		"</strong></code>?";
	var apiUrl = "/api/move-below/" + node.Key.Hostname + "/" + node.Key.Port + "/" + siblingNode.Key.Hostname + "/" + siblingNode.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

function moveUp(node, grandparentNode) {
	var message = "<h4>move-up</h4>Are you sure you wish to turn <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> into a slave of <code><strong>" +
		grandparentNode.Key.Hostname + ":" + grandparentNode.Key.Port +
		"</strong></code>?";
	var apiUrl = "/api/move-up/" + node.Key.Hostname + "/" + node.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

function enslaveMaster(node, masterNode) {
	var message = "<h4>enslave-master</h4>Are you sure you wish to make <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> master of <code><strong>" +
		masterNode.Key.Hostname + ":" + masterNode.Key.Port +
		"</strong></code>?";
	var apiUrl = "/api/enslave-master/" + node.Key.Hostname + "/" + node.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

function makeCoMaster(node, childNode) {
	var message = "<h4>make-co-master</h4>Are you sure you wish to make <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> and <code><strong>" +
		childNode.Key.Hostname + ":" + childNode.Key.Port +
		"</strong></code> co-masters?";
	var apiUrl = "/api/make-co-master/" + childNode.Key.Hostname + "/" + childNode.Key.Port;
	return executeMoveOperation(message, apiUrl);
}


function matchBelow(node, otherNode) {
	var message = "<h4>PSEUDO-GTID MODE, match-below</h4>Are you sure you wish to turn <code><strong>" + 
		node.Key.Hostname + ":" + node.Key.Port +
		"</strong></code> into a slave of <code><strong>" +
		otherNode.Key.Hostname + ":" + otherNode.Key.Port +
		"</strong></code>?";
	var apiUrl = "/api/match-below/" + node.Key.Hostname + "/" + node.Key.Port + "/" + otherNode.Key.Hostname + "/" + otherNode.Key.Port;
	return executeMoveOperation(message, apiUrl);
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
    	$(".popover.instance[data-nodeid="+instance.id+"]").attr("data-dc-color", dcColorsMap[instance.DataCenter]);
    	$(".instance-trailer[data-nodeid="+instance.id+"]").attr("data-dc-color", dcColorsMap[instance.DataCenter]);
    });
}


function refreshClusterOperationModeButton() {
	if (moveInstanceMethod == "smart") {
		$("#move-instance-method-button").removeClass("btn-success").removeClass("btn-warning").addClass("btn-info");
	} else if (moveInstanceMethod == "classic") {
		$("#move-instance-method-button").removeClass("btn-info").removeClass("btn-warning").addClass("btn-success");
	} else if (moveInstanceMethod == "pseudo-gtid") {
		$("#move-instance-method-button").removeClass("btn-success").removeClass("btn-info").addClass("btn-warning");
	} 
	$("#move-instance-method-button").html(moveInstanceMethod + ' mode <span class="caret"></span>')
}

// This is legacy and will be removed
function makeMaster(instance) {
	var message = "Are you sure you wish to make <code><strong>" + instance.Key.Hostname+":"+instance.Key.Port + "</strong></code> the new master?"
	+ "<p>Siblings of <code><strong>" + instance.Key.Hostname+":"+instance.Key.Port + "</strong></code> will turn to be its children, "
	+ "via Pseudo-GTID."
	+ "<p>The instance will be set to be writeable (<code><strong>read_only = 0</strong></code>)."
	+ "<p>Replication on this instance will be stopped, but not reset. You should run <code><strong>RESET SLAVE</strong></code> yourself "
	+ "if this instance will indeed become the master."
	+ "<p>Pointing your application servers to the new master is on you."
	;
	var apiUrl = "/api/make-master/"+instance.Key.Hostname+"/"+instance.Key.Port;
	return executeMoveOperation(message, apiUrl);
}

//This is legacy and will be removed
function makeLocalMaster(instance) {
	var message = "Are you sure you wish to make <code><strong>" + instance.Key.Hostname+":"+instance.Key.Port + "</strong></code> a local master?"
	+ "<p>Siblings of <code><strong>" + instance.Key.Hostname+":"+instance.Key.Port + "</strong></code> will turn to be its children, "
	+ "via Pseudo-GTID."
	+ "<p>The instance will replicate from its grandparent."
	;
	var apiUrl = "/api/make-local-master/"+instance.Key.Hostname+"/"+instance.Key.Port;
	return executeMoveOperation(message, apiUrl);
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


function anonymizeInstanceId(instanceId) {
	var tokens = instanceId.split("__");
    return "instance-"+md5(tokens[1]).substring(0, 4)+":"+tokens[2];
}

function anonymizeIfNeedBe(message) {
	if ($.cookie("anonymize") == "true") {
    	message = message.replace(/<strong>.*?<\/strong>/g, "############");
    }
    return message;
}

function anonymize() {
    var _ = function() {
        $("#cluster_container .instance[data-nodeid]").each(function() {
        	var instanceId = $(this).attr("data-nodeid");
        	$(this).find("h3 .pull-left").html(anonymizeInstanceId(instanceId));
        	$(this).find("h3").attr("title", anonymizeInstanceId(instanceId));
        });
        $(".popover-content p").each(function() {
        	tokens = jQuery(this).html().split(" ", 2)
            jQuery(this).html(tokens[0].match(/[^.]+[.][^.]+/) + " " + tokens[1])
         });
    }();
	$("#cluster_container div.floating_background").html("");	
}

function colorize_dc() {
    $(".popover[data-dc-color]").each(function () {
        $(this).css("border-color", $(this).attr("data-dc-color"));
        $(this).css("border-width", 2);
    });	
    $(".popover.instance-trailer[data-dc-color]").each(function () {
   		$(this).css("height", getInstanceDiv($(this).attr("data-nodeid")).outerHeight(true));
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
		var content = 'Domain: '+clusterInfo.ClusterDomain+'';
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
    recoveryListing.append('<li><a href="#" data-btn="auto-lite" data-command="recover-auto-lite">Auto (do not execute hooks/processes)</a></li>');
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
                visualizeInstances(instancesMap, generateInstanceDiv);
                prepareDraggable(instancesMap);
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
                if (reloadPageHint.hint == "refresh") {
                	var instanceId = getInstanceId(reloadPageHint.hostname, reloadPageHint.port);
            		var instance = instancesMap[instanceId]        	
            		if (instance) { 
            			openNodeModal(instance);
            		}
                }
        	}, "json");
    	}, "json");
    }, "json");
    $.get("/api/cluster-info/"+currentClusterName(), function (clusterInfo) {    
    	var alias = clusterInfo.ClusterAlias
    	var visualAlias = (alias ? alias : currentClusterName())
    	document.title = document.title.split(" - ")[0] + " - " + visualAlias;
    	
    	if (!($.cookie("anonymize") == "true")) {
        	$("#cluster_container").append('<div class="floating_background">'+visualAlias+'</div>');
            $("#dropdown-context").append('<li><a data-command="change-cluster-alias" data-alias="'+clusterInfo.ClusterAlias+'">Alias: '+alias+'</a></li>');
        }
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

    $("#li-move-instance-method").appendTo("ul.navbar-nav").show();
    $("#move-instance-method a").click(function() {
    	moveInstanceMethod = $(this).attr("data-method");
    	refreshClusterOperationModeButton();
    	$.cookie("move-instance-method", moveInstanceMethod, { path: '/', expires: 1 });
    });
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
    		return;
        }
    	$.cookie("pool-indicator", "true", { path: '/', expires: 1 });
		location.reload();
    });    
    $("body").on("click", "a[data-command=anonymize]", function(event) {
    	if ($.cookie("anonymize") == "true") {
    		$.cookie("anonymize", "false", { path: '/', expires: 1 });
    		location.reload();
    		return;
        }
    	anonymize();
    	$("#dropdown-context a[data-command=anonymize]").prepend('<span class="glyphicon glyphicon-ok small"></span> ');
    	$.cookie("anonymize", "true", { path: '/', expires: 1 });
    });    
    $("body").on("click", "a[data-command=colorize-dc]", function(event) {
    	if ($.cookie("colorize-dc") == "true") {
    		$.cookie("colorize-dc", "false", { path: '/', expires: 1 });
    		location.reload();
    		return;
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
    refreshClusterOperationModeButton();
});

function getHtmlPos(el) {
    return {left:el.offsetLeft, top:el.offsetTop};
}

function getSvgPos(el) {
    var svg = $(el).closest("svg")[0];
    if (!svg) {
    	return false;
    }
    var pt = svg.createSVGPoint();
    var matrix = el.getCTM();
    var box = el.getBBox();
    pt.x = box.x;
    pt.y = box.y;
    var pt2 = pt.matrixTransform(matrix);
    return {left:pt2.x, top:pt2.y};
}