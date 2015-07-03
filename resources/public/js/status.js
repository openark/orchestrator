
function addStatusTableData(name, content) {
	$("#orchestratorStatusTable").append(
	        '<tr><td>' + name + '</td><td><code class="text-info"><strong>' + content + '</strong></code></td></tr>'
	);
}
function addStatusActionButton(name, uri) {
	$("#orchestratorStatus .panel-footer").append(
		'<button type="button" class="btn btn-sm btn-info">'+name+'</button> '
	);
	var button = $('#orchestratorStatus .panel-footer button:last');
	button.click(function(){
    	apiCommand("/api/"+uri);
    });

	console.log(button)
}

$(document).ready(function () {
	var statusObject = $("#orchestratorStatus .panel-body");
    $.get("/api/health/", function (health) {
    	statusObject.prepend('<h4>'+health.Message+'</h4>')
    	health.Details.AvailableNodes.forEach(function(node) {
    		var hostname = node.split(";")[0];
    		var message = hostname;
    		if (node == health.Details.ActiveNode) {
    			message += ' <span class="text-success">[Active]</span>';
    		}
    		if (hostname == health.Details.Hostname) {
    			message += ' <span class="text-primary">[This node]</span>';
    		}
    		addStatusTableData("Available node", message);
    	})
    	
    	var userId = getUserId();
    	if (userId == "") {
    		userId = "[unknown]"
    	}
    	var userStatus = (isAuthorizedForAction() ? "admin" : "read only");
    	addStatusTableData("You", userId + ", " + userStatus);

    	if (isAuthorizedForAction()) {
    		addStatusActionButton("Reload configuration", "reload-configuration");
    		addStatusActionButton("Reset hostname resolve cache", "reset-hostname-resolve-cache");
    		addStatusActionButton("Reelect", "reelect");
    	}
    
    }, "json");   
});
