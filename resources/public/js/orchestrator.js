
var refreshIntervalSeconds = 60 ; // seconds
var secondsTillRefresh = refreshIntervalSeconds; 
var nodeModalVisible = false;

var errorMapping = {
   		"inMaintenanceProblem": {"badge": "label-info", "description": "In maintenance"}, 
   		"lastCheckInvalidProblem": {"badge": "label-fatal", "description": "Last check invalid"}, 
   		"notRecentlyCheckedProblem": {"badge": "label-stale", "description": "Not recently checked (stale)"}, 
   		"notReplicatingProblem": {"badge": "label-danger", "description": "Not replicating"}, 
   		"replicationLagProblem": {"badge": "label-warning", "description": "Replication lag"}
	};

function startRefreshTimer() {
    setInterval(function() {
    	if (nodeModalVisible) {
    		return;
    	}
    	secondsTillRefresh = Math.max(secondsTillRefresh - 1, 0);
    	if (secondsTillRefresh <= 0) {
    		$(".navbar-nav li[data-nav-page=refreshCountdown]").addClass("active");
    		showLoader();
    		location.reload(true);
    	}
    	$("#refreshCountdown").html('<span class="glyphicon glyphicon-repeat"></span> ' + secondsTillRefresh + 's');
    }, 1*1000);
}

function resetRefreshTimer() {
	secondsTillRefresh = refreshIntervalSeconds;
}

function activateRefreshTimer() {
    startRefreshTimer();
    $(document).click(function() {
    	resetRefreshTimer();
    });
    $(document).mousemove(function() {
    	resetRefreshTimer();
    });
}
		
function showLoader() {
    $(".ajaxLoader").css('visibility', 'visible');
}
function hideLoader() {
    $(".ajaxLoader").css('visibility', 'hidden');
}

function showContextMenu() {
    $("[data-nav-page=context]").css('visibility', 'visible');
}

function booleanString(b) {
	return (b ? "true" : "false");
}

function toHumanFormat(bytes) {
    var s = ['bytes', 'kB', 'MB', 'GB', 'TB', 'PB'];
    var e = Math.floor(Math.log(bytes) / Math.log(1024));
    return (bytes / Math.pow(1024, e)).toFixed(2) + " " + s[e];
}

function getInstanceId(host, port) {
    return "instance__" + host.replace(/[.]/g, "_") + "__" + port
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
    if ($.cookie("anonymize") == "true") {
        return false;
    }
	if(typeof(alertClass)==='undefined') {
        alertClass = "danger";
    }
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

function apiCommand(uri) {
	showLoader();
    $.get(uri, function (operationResult) {
		hideLoader();
		if (operationResult.Code == "ERROR") {
			addAlert(operationResult.Message)
		} else {
			reloadWithOperationResult(operationResult);
		}	
    }, "json");	
    return false;
}


function reloadWithMessage(msg) {
    window.location.href = window.location.href.split("#")[0].split("?")[0] + "?orchestrator-msg="+ encodeURIComponent(msg);
}

function reloadWithOperationResult(operationResult) {
    var msg = operationResult.Message;
    reloadWithMessage(msg);
}


// Modal

function addNodeModalDataAttribute(name, value) {
	var codeClass = "text-primary";
	if (value == "true" || value == true) {
		codeClass = "text-success";
	}
	if (value == "false" || value === false) {
		codeClass = "text-danger";
	}
    $('#modalDataAttributesTable').append(
        '<tr><td>' + name + '</td><td><code class="'+codeClass+'"><strong>' + value + '</strong></code><div class="pull-right attributes-buttons"></div></td></tr>');
    return $('#modalDataAttributesTable tr:last td:last');
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
	if (node.isAggregate) {
		return false;
	}
	nodeModalVisible = true;
    $('#node_modal #modalDataAttributesTable button[data-btn][data-grouped!=true]').appendTo("#node_modal .modal-footer");
    $('#node_modal #modalDataAttributesTable [data-btn-group]').appendTo("#node_modal .modal-footer");
    $('#node_modal .modal-title').html(node.title);
    $('#modalDataAttributesTable').html("");

    if (node.UnresolvedHostname) {
    	addNodeModalDataAttribute("Unresolved hostname", node.UnresolvedHostname);
    }
    if (node.MasterKey.Hostname) {
        var td = addNodeModalDataAttribute("Master", node.masterTitle);
        $('#node_modal button[data-btn=reset-slave]').appendTo(td.find("div"))
        
        td = addNodeModalDataAttribute("Replication running", booleanString(node.replicationRunning));
        $('#node_modal button[data-btn=start-slave]').appendTo(td.find("div"))
        $('#node_modal [data-btn-group=stop-slave]').appendTo(td.find("div"))
        
        if (!node.replicationRunning) {
            td = addNodeModalDataAttribute("Last SQL error", node.LastSQLError);
            $('#node_modal button[data-btn=skip-query]').appendTo(td.find("div"))
            addNodeModalDataAttribute("Last IO error", node.LastIOError);
        }
        addNodeModalDataAttribute("Seconds behind master", node.SecondsBehindMaster.Valid ? node.SecondsBehindMaster.Int64 : "null");
        addNodeModalDataAttribute("Replication lag", node.SlaveLagSeconds.Valid ? node.SlaveLagSeconds.Int64 : "null");
        addNodeModalDataAttribute("SQL delay", node.SQLDelay);
    }
    var td = addNodeModalDataAttribute("Num slaves", node.SlaveHosts.length);
    $('#node_modal button[data-btn=move-up-slaves]').appendTo(td.find("div"))
    $('#node_modal button[data-btn=match-up-slaves]').appendTo(td.find("div"))
    $('#node_modal button[data-btn=regroup-slaves]').appendTo(td.find("div"))
    addNodeModalDataAttribute("Server ID", node.ServerID);
    addNodeModalDataAttribute("Version", node.Version);
    var td = addNodeModalDataAttribute("Read only", booleanString(node.ReadOnly));
    $('#node_modal button[data-btn=set-read-only]').appendTo(td.find("div"))
    $('#node_modal button[data-btn=set-writeable]').appendTo(td.find("div"))

    addNodeModalDataAttribute("Binlog format", node.Binlog_format);
    addNodeModalDataAttribute("Has binary logs", booleanString(node.LogBinEnabled));
    var td = addNodeModalDataAttribute("Logs slave updates", booleanString(node.LogSlaveUpdatesEnabled));
    $('#node_modal button[data-btn=enslave-siblings]').appendTo(td.find("div"))
        
    addNodeModalDataAttribute("Uptime", node.Uptime);
    addNodeModalDataAttribute("Cluster",
            '<a href="/web/cluster/'+node.ClusterName+'">'+node.ClusterName+'</a>');
    addNodeModalDataAttribute("Agent",
            '<a href="/web/agent/'+node.Key.Hostname+'">'+node.Key.Hostname+'</a>');
    addNodeModalDataAttribute("Long queries",
            '<a href="/web/long-queries?filter='+node.Key.Hostname+'">on '+node.Key.Hostname+'</a>');
    
    $('#node_modal [data-btn]').unbind("click");
    
    $('#node_modal button[data-btn=begin-maintenance]').click(function() {
    	if (!$("#beginMaintenanceOwner").val()) {
    		return addModalAlert("You must fill the owner field");
    	}
    	if (!$("#beginMaintenanceReason").val()) {
    		return addModalAlert("You must fill the reason field");
    	}
    	var uri = "/api/begin-maintenance/"+node.Key.Hostname+"/"+node.Key.Port + "/" + $("#beginMaintenanceOwner").val() + "/" + $("#beginMaintenanceReason").val();
    	apiCommand(uri);
    });
    $('#node_modal button[data-btn=end-maintenance]').click(function(){
    	apiCommand("/api/end-maintenance/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal button[data-btn=refresh-instance]').click(function(){
    	apiCommand("/api/refresh/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal button[data-btn=skip-query]').click(function(){
    	apiCommand("/api/skip-query/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal button[data-btn=start-slave]').click(function(){
    	apiCommand("/api/start-slave/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal [data-btn=stop-slave]').click(function(){
    	apiCommand("/api/stop-slave/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal [data-btn=stop-slave-nice]').click(function(){
    	apiCommand("/api/stop-slave-nice/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal button[data-btn=reset-slave]').click(function(){
    	var message = "<p>Are you sure you wish to reset <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
			"</strong></code>?" +
			"<p>This will stop and break the replication." +
			"<p>FYI, this is a destructive operation that cannot be easily reverted"
			;
    	bootbox.confirm(message, function(confirm) {
			if (confirm) {
		    	apiCommand("/api/reset-slave/"+node.Key.Hostname+"/"+node.Key.Port);
			}
		}); 
		return false;
    });
    $('#node_modal button[data-btn=set-read-only]').click(function(){
    	apiCommand("/api/set-read-only/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal button[data-btn=set-writeable]').click(function(){
    	apiCommand("/api/set-writeable/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal button[data-btn=forget-instance]').click(function(){
    	var message = "<p>Are you sure you wish to forget <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
			"</strong></code>?" +
			"<p>It may be re-discovered if accessible from an existing instance through replication topology."
			;
    	bootbox.confirm(message, function(confirm) {
			if (confirm) {
		    	apiCommand("/api/forget/"+node.Key.Hostname+"/"+node.Key.Port);
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
	$('#node_modal button[data-btn=skip-query]').hide();
	$('#node_modal button[data-btn=start-slave]').hide();
	$('#node_modal [data-btn-group=stop-slave]').hide();
	
    if (node.MasterKey.Hostname) {
        if (node.replicationRunning || node.replicationAttemptingToRun) {
        	$('#node_modal [data-btn-group=stop-slave]').show();
        } 
        if (!node.replicationRunning) {
        	$('#node_modal button[data-btn=start-slave]').show();
        }
        if (!node.Slave_SQL_Running && node.LastSQLError) {
        	$('#node_modal button[data-btn=skip-query]').show();
        }
    }

	$('#node_modal button[data-btn=set-read-only]').hide();
	$('#node_modal button[data-btn=set-writeable]').hide();
    if (node.ReadOnly) {
    	$('#node_modal button[data-btn=set-writeable]').show();
    } else {
    	$('#node_modal button[data-btn=set-read-only]').show();
    }

    $('#node_modal button[data-btn=move-up-slaves]').hide();
    if (!node.lastCheckInvalidProblem()) {
        if (node.SlaveHosts.length > 0) {
            $('#node_modal button[data-btn=move-up-slaves]').show();
        }
    }
    $('#node_modal button[data-btn=move-up-slaves]').click(function(){
    	apiCommand("/api/move-up-slaves/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal button[data-btn=match-up-slaves]').hide();
    $('#node_modal button[data-btn=regroup-slaves]').hide();
    if (node.UsingPseudoGTID) {
        $('#node_modal button[data-btn=match-up-slaves]').show();
        if (node.SlaveHosts.length > 1) {
            $('#node_modal button[data-btn=regroup-slaves]').show();
        }
    }
    $('#node_modal button[data-btn=match-up-slaves]').click(function(){
    	apiCommand("/api/match-up-slaves/"+node.Key.Hostname+"/"+node.Key.Port);
    });
    $('#node_modal button[data-btn=regroup-slaves]').click(function(){
    	apiCommand("/api/regroup-slaves/"+node.Key.Hostname+"/"+node.Key.Port);
    });

   	$('#node_modal button[data-btn=enslave-siblings]').hide();
    if (node.LogBinEnabled && node.LogSlaveUpdatesEnabled) {
    	$('#node_modal button[data-btn=enslave-siblings]').show();
    }
    $('#node_modal button[data-btn=enslave-siblings]').click(function(){
    	apiCommand("/api/enslave-siblings/"+node.Key.Hostname+"/"+node.Key.Port);
    });

    $('#node_modal button[data-btn=recover]').hide();
    if (node.lastCheckInvalidProblem() && node.children && node.children.length > 0) {
        $('#node_modal button[data-btn=recover]').show();
    }
    $('#node_modal button[data-btn=recover]').click(function(){
    	apiCommand("/api/recover/"+node.Key.Hostname+"/"+node.Key.Port);
    });

    $('#node_modal').modal({})
    $('#node_modal').unbind('hidden.bs.modal');
    $('#node_modal').on('hidden.bs.modal', function () {
    	nodeModalVisible = false;
    })
}


function normalizeInstance(instance) {
    instance.id = getInstanceId(instance.Key.Hostname, instance.Key.Port);
    instance.title = instance.Key.Hostname+':'+instance.Key.Port;
    instance.canonicalTitle = instance.title;
    instance.masterTitle = instance.MasterKey.Hostname + ":" + instance.MasterKey.Port;
    instance.masterId = getInstanceId(instance.MasterKey.Hostname,
            instance.MasterKey.Port);

    instance.replicationRunning = instance.Slave_SQL_Running && instance.Slave_IO_Running;
    instance.replicationAttemptingToRun = instance.Slave_SQL_Running || instance.Slave_IO_Running;
    instance.replicationLagReasonable = Math.abs(instance.SlaveLagSeconds.Int64 - instance.SQLDelay) <= 10;
    instance.isSeenRecently = instance.SecondsSinceLastSeen.Valid && instance.SecondsSinceLastSeen.Int64 <= 3600;
    instance.usingGTID = instance.UsingOracleGTID || instance.UsingMariaDBGTID;
    instance.isMaxScale = (instance.Version.indexOf("maxscale") >= 0); 

    // used by cluster-tree
    instance.children = [];
    instance.parent = null;
    instance.hasMaster = true;
    instance.masterNode = null;
    instance.inMaintenance = false;
    instance.maintenanceEntry = null;
    instance.isFirstChildInDisplay = false

    instance.isMaster = (instance.title == instance.ClusterName);
    instance.isCoMaster = false;
    instance.isCandidateMaster = false;
    instance.isMostAdvancedOfSiblings = false;
    instance.isVirtual = false;
    instance.isAggregate = false;
    
    instance.renderHint = "";
}

function normalizeInstanceProblem(instance) {
    instance.inMaintenanceProblem = function() { return instance.inMaintenance; }
    instance.lastCheckInvalidProblem = function() { return !instance.IsLastCheckValid; }
    instance.notRecentlyCheckedProblem = function() { return !instance.IsRecentlyChecked; }
    instance.notReplicatingProblem = function() { return !instance.replicationRunning && !(instance.isMaster && !instance.isCoMaster); }
    instance.replicationLagProblem = function() { return !instance.replicationLagReasonable; }

    instance.problem = null;
    instance.problemOrder = 0;
    if (instance.inMaintenanceProblem()) {
    	instance.problem = "in_maintenance";
    	instance.problemOrder = 1;
    } else if (instance.lastCheckInvalidProblem()) {
    	instance.problem = "last_check_invalid";
    	instance.problemOrder = 2;
    } else if (instance.notRecentlyCheckedProblem()) {
    	instance.problem = "not_recently_checked";
    	instance.problemOrder = 3;
    } else if (instance.notReplicatingProblem()) {
    	// check slaves only; where not replicating
    	instance.problem = "not_replicating";
    	instance.problemOrder = 4;
    } else if (instance.replicationLagProblem()) {
    	instance.problem = "replication_lag";
    	instance.problemOrder = 5;
    }
    instance.hasProblem = (instance.problem != null) ;
    instance.hasConnectivityProblem = (!instance.IsLastCheckValid || !instance.IsRecentlyChecked);
}

var virtualInstanceCounter = 0;
function createVirtualInstance() {
    var virtualInstance = {
        	id: "orchestrator-virtual-instance-" + (virtualInstanceCounter++),
            children : [],
            parent: null,
            hasMaster: false,
            inMaintenance: false,
            maintenanceEntry: null,
            isMaster: false,
            isCoMaster: false,
            isVirtual: true,
            SlaveLagSeconds: 0,
            SecondsSinceLastSeen: 0
        }
    normalizeInstanceProblem(virtualInstance);
    return virtualInstance;
}

function normalizeInstances(instances, maintenanceList) {
    instances.forEach(function(instance) {
    	normalizeInstance(instance);
    });
    // Take canonical host name: strip down longest common suffix of all hosts
    // (experimental; you may not like it)
    var hostNames = instances.map(function (instance) {
        return instance.title
    });
    if (typeof removeTextFromHostnameDisplay != "undefined" && removeTextFromHostnameDisplay()) {
        instances.forEach(function (instance) {
        	instance.canonicalTitle = instance.title.replace(removeTextFromHostnameDisplay(), '');
        });
    } else {
        var suffixLength = commonSuffixLength(hostNames);
        instances.forEach(function (instance) {
        	instance.canonicalTitle = instance.title.substring(0, instance.title.length - suffixLength);
        });
    }
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
        	instance.masterNode = parent;
            // create child array if it doesn't exist
            parent.children.push(instance);
            // (parent.contents || (parent.contents = [])).push(instance);
        } else {
            // parent is null or missing
        	instance.hasMaster = false;
            instance.parent = null;
            instance.masterNode = null;
        }
    });

    instances.forEach(function (instance) {
    	if (instance.masterNode != null) {
		    instance.isSQLThreadCaughtUpWithIOThread = (instance.ExecBinlogCoordinates.LogFile == instance.ReadBinlogCoordinates.LogFile &&
		    		instance.ExecBinlogCoordinates.LogPos == instance.ReadBinlogCoordinates.LogPos);
    	} else {
    		instance.isSQLThreadCaughtUpWithIOThread = false;
    	}
    });
    
    instances.forEach(function (instance) {
    	if (instance.isMaster && instance.parent != null && instance.parent.parent != null && instance.parent.parent.id == instance.id) {
    	    // In case there's a master-master setup, introduce a virtual node
			// that is parent of both.
    		// This is for visualization purposes...
    	    var virtualCoMastersRoot = createVirtualInstance();
    		coMaster = instance.parent;
    		
    		function setAsCoMaster(instance, coMaster) {
        		instance.isCoMaster = true;
        		instance.hasMaster = true;
        		instance.masterId = coMaster.id;
        		instance.masterNode = coMaster;

        		var index = coMaster.children.indexOf(instance);
        		if (index >= 0)
        			coMaster.children.splice(index, 1);
        			
        		instance.parent = virtualCoMastersRoot;
        		virtualCoMastersRoot.children.push(instance);
    		}
    		setAsCoMaster(instance, coMaster);
    		setAsCoMaster(coMaster, instance);

    		instancesMap[virtualCoMastersRoot.id] = virtualCoMastersRoot;
    	} 
    });
    return instancesMap;
}


function renderInstanceElement(popoverElement, instance, renderType) {
	popoverElement.attr("data-nodeid", instance.id);
	popoverElement.find("h3").attr('title', instance.title);
	popoverElement.find("h3").html('&nbsp;<div class="pull-left">'+
			instance.canonicalTitle + '</div><div class="pull-right"><a href="#"><span class="glyphicon glyphicon-cog" title="Open config dialog"></span></a></div>');
	var indicateLastSeenInStatus = false;

	if (instance.isAggregate) {
		popoverElement.find("h3 div.pull-right span").remove();
	    popoverElement.find(".popover-content").append('<div>Instances: <div class="pull-right"></div></div>');
	    
	    function addInstancesBadge(count, badgeClass, title) {
	    	popoverElement.find(".popover-content .pull-right").append('<span class="badge '+badgeClass+'" title="' + title + '"">' + count + '</span> ');
	    }
	    addInstancesBadge(instance.aggregatedInstances.length, "label-primary", "Aggregated instances");
	    for (var problemType in instance.aggregatedProblems) {
	    	addInstancesBadge(instance.aggregatedProblems[problemType], errorMapping[problemType]["badge"], errorMapping[problemType]["description"]);
	    }
	}
	if (!instance.isAggregate) {
	    if (instance.isFirstChildInDisplay) {
	    	popoverElement.addClass("first-child-in-display");
	        popoverElement.attr("data-first-child-in-display", "true");
	    } 
	    if (instance.usingGTID) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-globe" title="Using GTID"></span> ');
	    } 
	    if (instance.UsingPseudoGTID) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-globe" title="Using Pseudo GTID"></span> ');
	    } 
	    if (!instance.ReadOnly) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-pencil" title="Writeable"></span> ');
	    } 
	    if (instance.isMostAdvancedOfSiblings) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-star" title="Most advanced slave"></span> ');
	    } 
	    if (instance.CountMySQLSnapshots > 0) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-camera" title="'+instance.CountMySQLSnapshots +' snapshots"></span> ');
	    } 
	    if (instance.HasReplicationFilters) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-filter" title="Using replication filters"></span> ');
	    } 
	    if (instance.LogBinEnabled && instance.LogSlaveUpdatesEnabled && !(instance.isMaster && !instance.isCoMaster)) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-forward" title="Logs slave updates"></span> ');
	    } 
	    if (instance.IsCandidate) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-heart" title="Candidate"></span> ');
	    } 
	    if (instance.inMaintenanceProblem()) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-wrench"></span> ');
	    } 
	    if (instance.IsDowntimed) {
	    	popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-volume-off"></span> ');
	    } 
	
	    if (instance.lastCheckInvalidProblem()) {
	    	instance.renderHint = "fatal";
	    	indicateLastSeenInStatus = true;
	    } else if (instance.notRecentlyCheckedProblem()) {
	    	instance.renderHint = "stale";
	    	indicateLastSeenInStatus = true;
	    } else if (instance.notReplicatingProblem()) {
	    	// check slaves only; check master only if it's co-master where not
			// replicating
	    	instance.renderHint = "danger";
	    } else if (instance.replicationLagProblem()) {
	    	instance.renderHint = "warning";
	    }
	    if (instance.renderHint != "") {
	    	popoverElement.find("h3").addClass("label-" + instance.renderHint);
	    }
		var statusMessage = instance.SlaveLagSeconds.Int64 + ' seconds lag';
		if (indicateLastSeenInStatus) {
			statusMessage = 'seen ' + instance.SecondsSinceLastSeen.Int64 + ' seconds ago';
		}
	    var contentHtml = ''
				+ instance.Version + " " + instance.Binlog_format
				;
	    
	    contentHtml = ''
	    	+ '<div class="pull-right">' + statusMessage + ' </div>'
			+ '<p>' 
			+ contentHtml
			+ '</p>'
			;
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
	}
    // if (instance.isCandidateMaster) {
    // popoverElement.append('<h4 class="popover-footer"><strong>Master
	// candidate</strong><div class="pull-right"><button class="btn btn-xs
	// btn-default" data-command="make-master"><span class="glyphicon
	// glyphicon-play"></span> Make master</button></div></h4>');
    // } else if (instance.isMostAdvancedOfSiblings) {
    // popoverElement.append('<h4
	// class="popover-footer"><strong>Candidate</strong><div
	// class="pull-right"><button class="btn btn-xs btn-default"
	// data-command="make-local-master"><span class="glyphicon
	// glyphicon-play"></span> Make local master</button></div></h4>');
    // }
    
    popoverElement.find("h3 a").click(function () {
    	openNodeModal(instance);
    	return false;
    });	
}

var onClustersListeners = [];

function onClusters(func) {
	onClustersListeners.push(func);
} 


function getParameterByName(name) {
    name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
    var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
        results = regex.exec(location.search);
    return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
}


$(document).ready(function() {
	$(".navbar-nav li").removeClass("active");
	$(".navbar-nav li[data-nav-page='" + activePage() + "']").addClass("active");
	
	$.get("/api/clusters-info", function(clusters) {
		clusters.forEach(function(cluster) {             
    		var title = '<span class="small">' + cluster.ClusterName + '</span>';
    		title = ((cluster.ClusterAlias != "") ? '<strong>' + cluster.ClusterAlias + '</strong>, ' + title : title);
	        $("#dropdown-clusters").append('<li><a href="/web/cluster/'+cluster.ClusterName+'">'+title+'</a></li>');
	    });                 
		onClustersListeners.forEach(function(func) {
			func(clusters);
		});
	}, "json");
	$(".ajaxLoader").click(function() {
        return false;
    });
	$("#refreshCountdown").click(function() {
		location.reload(true);
    });
	if (agentsHttpActive() == "true") {
		$("#nav_agents").show();
	}
	if (contextMenuVisible() == "true") {
		showContextMenu();
	}
	if (!isAuthorizedForAction()) {
	    $("[data-nav-page=read-only]").css('display', 'inline-block');
	}
	if (getUserId() != "") {
		$("[data-nav-page=user-id]").css('display', 'inline-block');
		$("[data-nav-page=user-id] a").html(" "+getUserId());
	}
    var orchestratorMsg = getParameterByName("orchestrator-msg")
    if (orchestratorMsg) {
        addInfo(orchestratorMsg)
        history.pushState(null, document.title, location.href.split("?orchestrator-msg=")[0])
    }
    $("#searchInput").focus();
});



