
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
    	addStatusTableData("Active node", health.Details.ActiveNode.split(";")[0]);
    	addStatusTableData("This node", health.Details.Hostname);
    	
    	var userId = getUserId();
    	if (userId == "") {
    		userId = "[unknown]"
    	}
    	var userStatus = (isAuthorizedForAction() ? "admin" : "read only");
    	addStatusTableData("You", userId + ", " + userStatus);

    	if (isAuthorizedForAction()) {
    		addStatusActionButton("Reload configuration", "reload-configuration");
    		addStatusActionButton("Reset hostname resolve cache", "reset-hostname-resolve-cache");
    	}
    
    }, "json");   
});
