
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
        			return ' <button class="btn alert-info disabled">'+hostname+' [this]</button>';
        		}
        		return ' <button class="btn alert-success">'+hostname+'</button>'
        	});
        	return result;
    	}
    	beautifyAvailableSnapshots(agent.AvailableLocalSnapshots).forEach(function (entry) {
        	$("[data-agent=local_snapshots]").append(entry)
    	});
    	
    	beautifyAvailableSnapshots(agent.AvailableSnapshots).forEach(function (entry) {
        	$("[data-agent=snapshots]").append(entry)
    	});

        hideLoader();
    }
});	
