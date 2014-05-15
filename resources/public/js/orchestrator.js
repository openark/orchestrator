function booleanString(b) {
	return (b ? "true" : "false");
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


function addAlert(alertText) {
	$("#alerts_container").append(
		'<div class="alert alert-danger alert-dismissable">'
				+ '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>'
				+ alertText + '</div>');
	$(".alert").alert();
	return false;
}


// Modal

function addNodeModalDataAttribute(name, value) {
    $('#modalDataAttributesTable').append(
        '<tr><td>' + name + '</td><td><code>' + value + '</code></td></tr>');
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
    
    $('#node_modal button[data-btn=begin-maintenance]').unbind("click");
    $('#node_modal button[data-btn=end-maintenance]').unbind("click");
    $('#node_modal button[data-btn=forget-instance]').unbind("click");
    $('#node_modal button[data-btn=begin-maintenance]').click(function() {
    	console.log($("#beginMaintenanceReason").val());
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
				addAlert("<strong>Error</strong>: " + operationResult.Message)
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
				addAlert("<strong>Error</strong>: " + operationResult.Message)
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
			    				addAlert("<strong>Error</strong>: " + operationResult.Message)
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
    
    $('#node_modal').modal({})
}

