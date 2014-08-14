
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
    	$("[data-agent=port]").html(agent.Port)
    	$("[data-agent=last_submitted]").html(agent.LastSubmitted)
    	
    	function beautifyAvailableSnapshots(hostnames) {
        	var result = hostnames.map(function(hostname) {
        		if (hostname == agent.Hostname) {
        			return '<td><span class="">'+hostname+'</span></td><td></td>';
        		}
        		var isLocal = $.inArray(hostname, agent.AvailableLocalSnapshots) >= 0;
        		var btnType = (isLocal ? "btn-success" : "btn-warning"); 
        		return '<td><a href="/web/agent/'+hostname+'">'+hostname+'</a></td><td><button class="btn btn-xs '+btnType+'">Reseed</button></td>';
        	});
        	result = result.map(function(entry) {
        		return '<tr>'+entry+'</tr>';
        	});
        	return result;
    	}
//    	beautifyAvailableSnapshots(agent.AvailableLocalSnapshots).forEach(function (entry) {
//        	$("[data-agent=local_snapshots]").append(entry)
//    	});
//    	$("[data-agent=local_snapshots]").append(beautifyAvailableSnapshots(agent.AvailableLocalSnapshots).join("<br/>"));
    	
    	beautifyAvailableSnapshots(agent.AvailableSnapshots).forEach(function (entry) {
        	$("[data-agent=available_snapshots]").append(entry)
    	});
//    	$("[data-agent=snapshots]").append(beautifyAvailableSnapshots(agent.AvailableSnapshots).join("<br/>"));
    	
    	var mountMessage = "No snapshot mounted"
    	var mountedVolume = ""
    	if (agent.MountPoint) {
	    	if (agent.MountPoint.IsMounted) { 
	    		mountedVolume = agent.MountPoint.LVPath;
	    		mountMessage = '<code>'+mountedVolume+'</code>';
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
        			volumeText += '<td><button class="btn btn-xs btn-danger">Unmount</button></td>';
        			volumeTextType = 'text-success';
        		} else if (mountedVolume == "") {
        			volumeText += '<td><button class="btn btn-xs btn-success">Mount</button></td>'
        		} else {
        			volumeText += '<td></td>'
        		}
        		volumeText = '<td><code class="'+volumeTextType+'"><strong>'+volume+'</strong></code></td>' + volumeText;
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
});	
