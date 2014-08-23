
$(document).ready(function () {
    showLoader();
    
    $.get("/api/agent/"+currentAgentHost(), function (agent) {
    		showLoader();
	    	displayAgent(agent);
	    }, "json");

    $.get("/api/agent-active-seeds/"+currentAgentHost(), function (activeSeeds) {
        showLoader();
    	activeSeeds.forEach(function (activeSeed) {
    		appendActiveSeed(activeSeed);
    	});
    	if (activeSeeds.length == 0) {
    		$("div.active_seeds").parent().hide();
    		$("div.seed_states").parent().hide();
    	}
    	if (activeSeeds.length > 0) {
    	    activateRefreshTimer();

    	    $.get("/api/agent-seed-states/"+activeSeeds[0].SeedId, function (seedStates) {
    	        showLoader();
    	        seedStates.forEach(function (seedState) {
    	        	appendSeedState(seedState);
    	    	});
    	    }, "json");    		
    	}
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
    	$("[data-agent=mysql_port]").html(agent.MySQLPort)
    	$("[data-agent=mysql_disk_usage]").html(toHumanFormat(agent.MySQLDiskUsage))    	
    	
    	function beautifyAvailableSnapshots(hostnames) {
        	var result = hostnames.filter(function(hostname) {
        		return hostname.trim() != "";
        	});
        	result = result.map(function(hostname) {
        		if (hostname == agent.Hostname) {
        			return '<td><span class="">'+hostname+'</span></td><td></td>';
        		}
        		var isLocal = $.inArray(hostname, agent.AvailableLocalSnapshots) >= 0;
        		var btnType = (isLocal ? "btn-success" : "btn-warning"); 
        		return '<td><a href="/web/agent/'+hostname+'">'+hostname+'</a><div class="pull-right"><button class="btn btn-xs '+btnType+'" data-command="seed" data-seed-source-host="'+hostname+'" data-seed-local="'+isLocal+'" data-mysql-running="'+agent.MySQLRunning+'">Seed</button></div></td>';
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
    	
    	var mountedVolume = ""
    	if (agent.MountPoint) {
	    	if (agent.MountPoint.IsMounted) { 
	    		mountedVolume = agent.MountPoint.LVPath;
	    		var mountMessage = '<td>';
	    		mountMessage += '<code>'+mountedVolume+'</code> mounted on '+
	    			'<code>'+agent.MountPoint.Path+'</code>, size '+toHumanFormat(agent.MountPoint.DiskUsage);
	    		mountMessage += '<br/>MySQL data path: <code>'+agent.MountPoint.MySQLDataPath+'</code>';
	    		mountMessage += '</td><td><div class="pull-right"><button class="btn btn-xs btn-danger" data-command="unmount">Unmount</button></div></td>';
	    		$("[data-agent=mount_point]").append(mountMessage);
	    	}
    	}

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
        		} else if (!(agent.MountPoint && agent.MountPoint.IsMounted)) {
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
    	}
		
        hideLoader();
    }
    
    function appendActiveSeed(seed) {    	
    	var row = '<tr>';
    	row += '<td>' + seed.SeedId + '</td>';
    	row += '<td>' + (seed.TargetHostname == currentAgentHost() ? seed.TargetHostname : '<a href="/web/agent/'+seed.TargetHostname+'">'+seed.TargetHostname+'</a>') + '</td>';
    	row += '<td>' + (seed.SourceHostname == currentAgentHost() ? seed.SourceHostname : '<a href="/web/agent/'+seed.SourceHostname+'">'+seed.SourceHostname+'</a>') + '</td>';
    	row += '<td>' + seed.StartTimestamp + '</td>';
    	row += '</tr>';
    	$("[data-agent=active_seeds]").append(row);
        hideLoader();
    }
    
    function appendSeedState(seedState) {    	
    	var row = '<tr>';
    	row += '<td>' + seedState.StateTimestamp + '</td>';
    	row += '<td>' + seedState.Action + '</td>';
    	row += '<td>' + seedState.ErrorMessage + '</td>';
    	row += '</tr>';
    	$("[data-agent=seed_states]").append(row);
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
    $("body").on("click", "button[data-command=seed]", function(event) {
    	if ($(event.target).attr("data-mysql-running") == "true") {
			addAlert("MySQL is running on this host. Please first stop the MySQL service");
			return;
    	}
    	var sourceHost = $(event.target).attr("data-seed-source-host");
    	var isLocalSeed = ($(event.target).attr("data-seed-local") == "true");
    	
    	var message = "Are you sure you wish to destroy data on <code><strong>" + 
    		currentAgentHost() + "</strong></code> and seed from <code><strong>" + 
    		sourceHost + "</strong></code>?";
    	if (isLocalSeed) {
    		message += '<p/><span class="text-success">This seed is dc-local</span>';
    	} else {
    		message += '<p/><span class="text-danger">This seed is non-local and will require cross-DC data transfer!</span>';
    	}
    	
		bootbox.confirm(message, function(confirm) {
			if (confirm) {
		    	showLoader();
		        $.get("/api/agent-seed/"+currentAgentHost()+"/"+sourceHost, function (operationResult) {
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

});	
