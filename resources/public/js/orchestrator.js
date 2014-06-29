
var refreshIntervalSeconds = 60 ; // seconds
var secondsTillRefresh = refreshIntervalSeconds; 
var nodeModalVisible = false;

function startRefreshTimer() {
    setInterval(function() {
    	if (nodeModalVisible) {
    		return;
    	}
    	secondsTillRefresh = Math.max(secondsTillRefresh - 1, 0);
    	if (secondsTillRefresh <= 0) {
    		location.reload(true);
    	}
    	$("#refreshCountdown").html('<span class="glyphicon glyphicon-repeat"></span> ' + secondsTillRefresh + 's');
    }, 1*1000);
}

function resetRefreshTimer() {
	secondsTillRefresh = refreshIntervalSeconds;
}
		
function showLoader() {
    $("#ajaxLoader").css('visibility', 'visible');
}
function hideLoader() {
    $("#ajaxLoader").css('visibility', 'hidden');
}


function booleanString(b) {
	return (b ? "true" : "false");
}

function getInstanceId(host, port) {
    return "instance" + host.replace(/[.]/g, "_") + "__" + port
}

function commonSuffixLength(strings) {
	if (strings.length == 0) {
		return 0;
	}
	if (strings.length == 1) {
		return 0;
	}
	var longestSuffixLength = 0;
	var maxLength = 0;
	strings.forEach(function(s) {
		maxLength = ((maxLength == 0) ? s.length : Math
				.min(maxLength, s.length));
	});
	var suffixLength = 0;
	while (suffixLength < maxLength) {
		suffixLength++
		var suffixes = strings.map(function(s) {
			return s.substring(s.length - suffixLength)
		});
		var uniqueSuffixes = suffixes.filter(function(elem, pos) {
			return suffixes.indexOf(elem) == pos;
		})
		if (uniqueSuffixes.length > 1) {
			// lost it. keep last longestSuffixLength value
			break;
		}
		// we're still good
		longestSuffixLength = suffixLength;
	}
	return longestSuffixLength;
}


function addAlert(alertText, alertClass) {
	if(typeof(alertClass)==='undefined') alertClass = "danger";
	$("#alerts_container").append(
		'<div class="alert alert-'+alertClass+' alert-dismissable">'
				+ '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>'
				+ alertText + '</div>');
	$(".alert").alert();
	return false;
}


function addInfo(alertText) {
	return addAlert(alertText, "info");
}


// Modal

function addNodeModalDataAttribute(name, value) {
    $('#modalDataAttributesTable').append(
        '<tr><td>' + name + '</td><td><code class="text-primary"><strong>' + value + '</strong></code></td></tr>');
}

function addModalAlert(alertText) {
	$("#node_modal .modal-body").append(
		'<div class="alert alert-danger alert-dismissable">'
				+ '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>'
				+ alertText + '</div>');
	$(".alert").alert();
	return false;
}

function openNodeModal(node) {
	nodeModalVisible = true;
    $('#node_modal .modal-title').html(node.title);
    $('#modalDataAttributesTable').html("");

    if (node.MasterKey.Hostname) {
        addNodeModalDataAttribute("Master", node.masterTitle);
        addNodeModalDataAttribute("Replication running", booleanString(node.replicationRunning));
        addNodeModalDataAttribute("Seconds behind master", node.SecondsBehindMaster.Valid ? node.SecondsBehindMaster.Int64 : "null");
        addNodeModalDataAttribute("Replication lag", node.SlaveLagSeconds.Valid ? node.SlaveLagSeconds.Int64 : "null");
    }
    addNodeModalDataAttribute("Num slaves", node.SlaveHosts.length);
    addNodeModalDataAttribute("Server ID", node.ServerID);
    addNodeModalDataAttribute("Version", node.Version);
    addNodeModalDataAttribute("Binlog format", node.Binlog_format);
    addNodeModalDataAttribute("Has binary logs", booleanString(node.LogBinEnabled));
    addNodeModalDataAttribute("Logs slave updates", booleanString(node.LogSlaveUpdatesEnabled));
    addNodeModalDataAttribute("Cluster",
            '<a href="/web/cluster/'+node.ClusterName+'">'+node.ClusterName+'</a>');
    
    $('#node_modal button[data-btn=begin-maintenance]').unbind("click");
    $('#node_modal button[data-btn=end-maintenance]').unbind("click");
    $('#node_modal button[data-btn=refresh-instance]').unbind("click");
    $('#node_modal button[data-btn=forget-instance]').unbind("click");
    $('#node_modal button[data-btn=start-slave]').unbind("click");
    $('#node_modal button[data-btn=stop-slave]').unbind("click");
    $('#node_modal button[data-btn=begin-maintenance]').click(function() {
    	if (!$("#beginMaintenanceOwner").val()) {
    		return addModalAlert("You must fill the owner field");
    	}
    	if (!$("#beginMaintenanceReason").val()) {
    		return addModalAlert("You must fill the reason field");
    	}
    	showLoader();
    	var uri = "/api/begin-maintenance/"+node.Key.Hostname+"/"+node.Key.Port + "/" + $("#beginMaintenanceOwner").val() + "/" + $("#beginMaintenanceReason").val();
        $.get(uri, function (operationResult) {
			hideLoader();
			if (operationResult.Code == "ERROR") {
				addAlert(operationResult.Message)
			} else {
				location.reload();
			}	
        }, "json");	
    });
    $('#node_modal button[data-btn=end-maintenance]').click(function(){
    	showLoader();
        $.get("/api/end-maintenance/"+node.Key.Hostname+"/"+node.Key.Port, function (operationResult) {
			hideLoader();
			if (operationResult.Code == "ERROR") {
				addAlert(operationResult.Message)
			} else {
				location.reload();
			}	
        }, "json");	
    });
    $('#node_modal button[data-btn=refresh-instance]').click(function(){
    	showLoader();
        $.get("/api/refresh/"+node.Key.Hostname+"/"+node.Key.Port, function (operationResult) {
			hideLoader();
			if (operationResult.Code == "ERROR") {
				addAlert(operationResult.Message)
			} else {
				location.reload();
			}	
        }, "json");	
    });
    $('#node_modal button[data-btn=start-slave]').click(function(){
    	showLoader();
        $.get("/api/start-slave/"+node.Key.Hostname+"/"+node.Key.Port, function (operationResult) {
			hideLoader();
			if (operationResult.Code == "ERROR") {
				addAlert(operationResult.Message)
			} else {
				location.reload();
			}	
        }, "json");	
    });
    $('#node_modal button[data-btn=stop-slave]').click(function(){
    	showLoader();
        $.get("/api/stop-slave/"+node.Key.Hostname+"/"+node.Key.Port, function (operationResult) {
			hideLoader();
			if (operationResult.Code == "ERROR") {
				addAlert(operationResult.Message)
			} else {
				location.reload();
			}	
        }, "json");	
    });
    $('#node_modal button[data-btn=forget-instance]').click(function(){
    	var message = "<p>Are you sure you wish to forget <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
			"</strong></code>?" +
			"<p>It may be re-discovered if accessible from an existing instance through replication topology."
			;
    	bootbox.confirm(message, function(confirm) {
				if (confirm) {
					showLoader();
					var apiUrl = "/api/forget/" + node.Key.Hostname + "/" + node.Key.Port;
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
    });

    if (node.inMaintenance) {
    	$('#node_modal [data-panel-type=maintenance]').html("In maintenance");
    	$('#node_modal [data-description=maintenance-status]').html(
    			"Started " + node.maintenanceEntry.BeginTimestamp + " by "+node.maintenanceEntry.Owner + ".<br/>Reason: "+node.maintenanceEntry.Reason
    	);    	
    	$('#node_modal [data-panel-type=begin-maintenance]').hide();
    	$('#node_modal [data-panel-type=end-maintenance]').show();
    } else {
    	$('#node_modal [data-panel-type=maintenance]').html("Maintenance");
    	$('#node_modal [data-panel-type=begin-maintenance]').show();
    	$('#node_modal [data-panel-type=end-maintenance]').hide();
    }
	$('#node_modal button[data-btn=start-slave]').hide();
	$('#node_modal button[data-btn=stop-slave]').hide();
    if (node.MasterKey.Hostname) {
        if (node.replicationRunning) {
        	$('#node_modal button[data-btn=stop-slave]').show();
        } else {
        	$('#node_modal button[data-btn=start-slave]').show();
        }
    }
    
    $('#node_modal').modal({})
    $('#node_modal').unbind('hidden.bs.modal');
    $('#node_modal').on('hidden.bs.modal', function () {
    	nodeModalVisible = false;
    })
}


function normalizeInstance(instance) {
    instance.id = getInstanceId(instance.Key.Hostname, instance.Key.Port);
    instance.title= instance.Key.Hostname+':'+instance.Key.Port;
    instance.canonicalTitle = instance.title;
    instance.masterTitle = instance.MasterKey.Hostname + ":" + instance.MasterKey.Port;
    instance.masterId = getInstanceId(instance.MasterKey.Hostname,
            instance.MasterKey.Port);

    instance.replicationRunning = instance.Slave_SQL_Running && instance.Slave_IO_Running;
    instance.replicationLagReasonable = instance.SlaveLagSeconds.Int64 <= 10;
    instance.isSeenRecently = instance.SecondsSinceLastSeen.Valid && instance.SecondsSinceLastSeen.Int64 <= 3600;

    // used by cluster-tree
    instance.children = [];
    instance.parent = null;
    instance.hasMaster = true;
    instance.inMaintenance = false;
    instance.maintenanceEntry = null;

    instance.isMaster = (instance.title == instance.ClusterName);
    instance.isCoMaster = false;
    instance.isVirtualCoMastersRoot = false;
}

function normalizeInstanceProblem(instance) {
    instance.problem = null;
    instance.problemOrder = 0;
    if (instance.inMaintenance) {
    	instance.problem = "in_maintenance";
    	instance.problemOrder = 1;
    } else if (!instance.IsLastCheckValid) {
    	instance.problem = "last_check_invalid";
    	instance.problemOrder = 2;
    } else if (!instance.IsRecentlyChecked) {
    	instance.problem = "not_recently_checked";
    	instance.problemOrder = 3;
    } else if (!instance.isMaster && !instance.replicationRunning) {
    	// check slaves only; where not replicating
    	instance.problem = "not_replicating";
    	instance.problemOrder = 4;
    } else if (!instance.replicationLagReasonable) {
    	instance.problem = "replication_lag";
    	instance.problemOrder = 5;
    }
    instance.hasProblem = (instance.problem != null) ;
}

function normalizeInstances(instances, maintenanceList) {
    var virtualCoMastersRoot = {
    	id: "virtualCoMastersRoot",
        children : [],
        parent: null,
        hasMaster: false,
        inMaintenance: false,
        maintenanceEntry: null,
        isMaster: false,
        isCoMaster: false,
        isVirtualCoMastersRoot: true,
        SlaveLagSeconds: 0,
        SecondsSinceLastSeen: 0
    }
    instances.forEach(function(instance) {
    	normalizeInstance(instance);
    });
    // Take canonical host name: strip down longest common suffix of all hosts
    // (experimental; you may not like it)
    var hostNames = instances.map(function (instance) {
        return instance.title
    });
    var suffixLength = commonSuffixLength(hostNames);
    instances.forEach(function (instance) {
    	instance.canonicalTitle = instance.title.substring(0, instance.title.length - suffixLength);
    });
    var instancesMap = instances.reduce(function (map, instance) {
        map[instance.id] = instance;
        return map;
    }, {});
    // mark maintenance instances
    maintenanceList.forEach(function (maintenanceEntry) {
        var instanceId = getInstanceId(maintenanceEntry.Key.Hostname, maintenanceEntry.Key.Port)
        if (instanceId in instancesMap) {
        	instancesMap[instanceId].inMaintenance = true;
        	instancesMap[instanceId].maintenanceEntry = maintenanceEntry;
        }
    });
    instances.forEach(function(instance) {
    	// Now that we also know about maintenance
    	normalizeInstanceProblem(instance);
    });
     // create the tree array
    instances.forEach(function (instance) {
        // add to parent
        var parent = instancesMap[instance.masterId];
        if (parent) {
        	instance.parent = parent;
            // create child array if it doesn't exist
            parent.children.push(instance);
            //(parent.contents || (parent.contents = [])).push(instance);
        } else {
            // parent is null or missing
        	instance.hasMaster = false;
            instance.parent = null;
        }
    });

    instances.forEach(function (instance) {
    	if (instance.isMaster && instance.parent != null && instance.parent.parent != null && instance.parent.parent.id == instance.id) {
    	    // In case there's a master-master setup, break the circle
    		// This is for visualization purposes...
    		instance.isCoMaster = true
    		instance.parent.isCoMaster = true
    		var index = instance.parent.children.indexOf(instance);
    		if (index >= 0)
    			instance.parent.children.splice(index, 1);
    		instance.hasMaster = false;
    		instance.masterId = "";
    		instance.parent = null;
    	} 
    });

    return instancesMap;
}

function renderInstanceElement(popoverElement, instance, renderType) {
	popoverElement.attr("data-nodeid", instance.id);
	popoverElement.find("h3").html(
    		instance.canonicalTitle + '<div class="pull-right"><a href="#"><span class="glyphicon glyphicon-cog"></span></a></div>');
	var indicateLastSeenInStatus = false;
    if (instance.inMaintenance) {
    	popoverElement.find("h3").addClass("label-info");
    } else if (!instance.IsLastCheckValid) {
    	popoverElement.find(" h3").addClass("label-fatal");
    	indicateLastSeenInStatus = true;
    } else if (!instance.IsRecentlyChecked) {
    	popoverElement.find(" h3").addClass("label-stale");
    	indicateLastSeenInStatus = true;
    } else if (!instance.isMaster && !instance.replicationRunning) {
    	// check slaves only; where not replicating
    	popoverElement.find("h3").addClass("label-danger");
    } else if (!instance.replicationLagReasonable) {
    	popoverElement.find("h3").addClass("label-warning");
    }
	var statusMessage = instance.SlaveLagSeconds.Int64 + ' seconds lag';
	if (indicateLastSeenInStatus) {
		statusMessage = 'seen ' + instance.SecondsSinceLastSeen.Int64 + ' seconds ago';
	}
    var contentHtml = ''
        	+ '<div class="pull-right">' + statusMessage + ' </div>'
   		+ '<p>' 
			+ instance.Version + " " + instance.Binlog_format 
        + '</p>';
    if (instance.isCoMaster) {
    	contentHtml += '<p><strong>Co master</strong></p>';
    }
    else if (instance.isMaster) {
    	contentHtml += '<p><strong>Master</strong></p>';
    }
    if (renderType == "search") {
    	contentHtml += '<p>' 
        	+ 'Cluster: <a href="/web/cluster/'+instance.ClusterName+'">'+instance.ClusterName+'</a>'
        + '</p>';
    }  
    if (renderType == "problems") {
    	contentHtml += '<p>' 
        	+ 'Problem: <strong>'+instance.problem.replace(/_/g, ' ') + '</strong>'
        + '</p>';
    }  
    
    popoverElement.find(".popover-content").html(contentHtml);
    
    popoverElement.find("h3 a").click(function () {
    	openNodeModal(instance);
    	return false;
    });	
}

var onClustersListeners = [];

function onClusters(func) {
	onClustersListeners.push(func);
} 

$(document).ready(function() {
	$.get("/api/clusters", function(clusters) {
		clusters.forEach(function(cluster) {                      
	        $("#dropdown-clusters").append('<li><a href="/web/cluster/'+cluster+'">'+cluster+'</a></li>');
	    });                 
		onClustersListeners.forEach(function(func) {
			func(clusters);
		});
	}, "json");
	$("#ajaxLoader").click(function() {
        return false;
    });
	$("#refreshCountdown").click(function() {
		location.reload(true);
    });
});



