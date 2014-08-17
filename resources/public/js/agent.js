
$(document).ready(function () {
    showLoader();
    
    $.get("/api/agent/"+currentAgentHost(), function (agent) {
    	displayAgent(agent);
    }, "json");
    function displayAgent(agent) {
    	if (!agent.Hostname) {
    		$("[data-agent=hostname]").html('<span class="alert-danger">Not found</span>');
    		return;
    	}
    	$("[data-agent=hostname]").html(agent.Hostname)
    	$("[data-agent=hostname_search]").html(
    			'<a href="/web/search?s='+agent.Hostname+'">'+agent.Hostname+'</a>'
    	)
    	$("[data-agent=port]").html(agent.Port)
    	$("[data-agent=last_submitted]").html(agent.LastSubmitted)
    	
    	var mySQLStatus = "" + agent.MySQLRunning + '<div class="pull-right">' +
    		(agent.MySQLRunning ? '<button class="btn btn-xs btn-danger" data-command="mysql-stop">Stop</button>' : 
    			'<button class="btn btn-xs btn-success" data-command="mysql-start">Start</button>') + 
    		'</div>';
    	$("[data-agent=mysql_running]").html(mySQLStatus)
    	
    	
    	function beautifyAvailableSnapshots(hostnames) {
        	var result = hostnames.map(function(hostname) {
        		if (hostname == agent.Hostname) {
        			return '<td><span class="">'+hostname+'</span></td><td></td>';
        		}
        		var isLocal = $.inArray(hostname, agent.AvailableLocalSnapshots) >= 0;
        		var btnType = (isLocal ? "btn-success" : "btn-warning"); 
        		return '<td><a href="/web/agent/'+hostname+'">'+hostname+'</a><div class="pull-right"><button class="btn btn-xs '+btnType+'">Reseed</button></div></td>';
        	});
        	result = result.map(function(entry) {
        		return '<tr>'+entry+'</tr>';
        	});
        	return result;
    	}
    	beautifyAvailableSnapshots(agent.AvailableLocalSnapshots).forEach(function (entry) {
        	$("[data-agent=available_local_snapshots]").append(entry)
    	});
    	availableRemoteSnapshots = agent.AvailableSnapshots.filter(function(snapshot) {
    		return agent.AvailableLocalSnapshots.indexOf(snapshot) < 0;
    	});
    	beautifyAvailableSnapshots(availableRemoteSnapshots).forEach(function (entry) {
        	$("[data-agent=available_remote_snapshots]").append(entry)
    	});
//    	$("[data-agent=local_snapshots]").append(beautifyAvailableSnapshots(agent.AvailableLocalSnapshots).join("<br/>"));
    	
//    	beautifyAvailableSnapshots(agent.AvailableSnapshots).forEach(function (entry) {
//        	$("[data-agent=available_snapshots]").append(entry)
//    	});
//    	$("[data-agent=snapshots]").append(beautifyAvailableSnapshots(agent.AvailableSnapshots).join("<br/>"));
    	
    	var mountMessage = "No snapshot mounted"
    	var mountedVolume = ""
    	if (agent.MountPoint) {
	    	if (agent.MountPoint.IsMounted) { 
	    		mountedVolume = agent.MountPoint.LVPath;
	    		mountMessage = '<code>'+mountedVolume+'</code> mounted on '+
	    			'<code>'+agent.MountPoint.Path+'</code>, size '+agent.MountPoint.DiskUsage+'b';
	    	}
    	}
		$("[data-agent=mount_point]").append(mountMessage);

    	if (agent.LogicalVolumes) {
	    	var lvSnapshots = agent.LogicalVolumes.filter(function (logicalVolume) {
	    		return logicalVolume.IsSnapshot;
	    	}).map(function(logicalVolume) {
	    		return logicalVolume.Path;
	    	});
        	var result = lvSnapshots.map(function(volume) {
        		var volumeText = '';
        		var volumeTextType = 'text-info';
        		if (volume == mountedVolume) {
        			volumeText = '<button class="btn btn-xs btn-danger" data-command="unmount">Unmount</button>';
        			volumeTextType = 'text-success';
        		} else if (mountedVolume == "") {
        			volumeText += '<button class="btn btn-xs btn-success" data-command="mountlv" data-lv="'+volume+'">Mount</button>'
        		} else {
        		}
        		volumeText = '<td><code class="'+volumeTextType+'"><strong>'+volume+'</strong></code><div class="pull-right">' + volumeText + '</div></td>';
        		return volumeText;
        	});
        	result = result.map(function(entry) {
        		return '<tr>'+entry+'</tr>';
        	});

        	result.forEach(function (entry) {
	        	$("[data-agent=lv_snapshots]").append(entry)
	    	});
			//$("[data-agent=lv_snapshots]").append(lvSnapshots.join("<br/>"));
    	}
		
        hideLoader();
    }
    $("body").on("click", "button[data-command=unmount]", function(event) {
    	showLoader();
        $.get("/api/agent-umount/"+currentAgentHost(), function (operationResult) {
			hideLoader();
			if (operationResult.Code == "ERROR") {
				addAlert(operationResult.Message)
			} else {
				location.reload();
			}	
        }, "json");	
    });
    $("body").on("click", "button[data-command=mountlv]", function(event) {
    	var lv = $(event.target).attr("data-lv")
    	showLoader();
        $.get("/api/agent-mount/"+currentAgentHost()+"?lv="+encodeURIComponent(lv), function (operationResult) {
			hideLoader();
			if (operationResult.Code == "ERROR") {
				addAlert(operationResult.Message)
			} else {
				location.reload();
			}	
        }, "json");	
    });
    $("body").on("click", "button[data-command=mysql-stop]", function(event) {
    	var message = "Are you sure you wish to shut down MySQL service on <code><strong>" + 
    		currentAgentHost() + "</strong></code>?";
		bootbox.confirm(message, function(confirm) {
			if (confirm) {
		    	showLoader();
		        $.get("/api/agent-mysql-stop/"+currentAgentHost(), function (operationResult) {
					hideLoader();
					if (operationResult.Code == "ERROR") {
						addAlert(operationResult.Message)
					} else {
						location.reload();
					}	
		        }, "json");
			}
		});
    });
    $("body").on("click", "button[data-command=mysql-start]", function(event) {
    	showLoader();
        $.get("/api/agent-mysql-start/"+currentAgentHost(), function (operationResult) {
			hideLoader();
			if (operationResult.Code == "ERROR") {
				addAlert(operationResult.Message)
			} else {
				location.reload();
			}	
        }, "json");	
    });
});	
