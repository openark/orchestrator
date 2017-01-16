
function addStatusTableData(name, column1, column2) {
	$("#orchestratorStatusTable").append(
	        '<tr><td>' + name + '</td>' +
                '<td><code class="text-info"><strong>' + column1 + '</strong></code></td>' +
                '<td><code class="text-info">' + column2 + '</code></td></tr>'
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
    $.get(appUrl("/api/health/"), function (health) {
    	statusObject.prepend('<h4>'+health.Message+'</h4>')
    	health.Details.AvailableNodes.forEach(function(node) {
				var app_version = node.AppVersion;
				if (app_version == "") {
					app_version = "unknown version";
				}
				var message = node.Hostname;
				if (node.Hostname == health.Details.ActiveNode.Hostname && node.Token == health.Details.ActiveNode.Token) {
					message += ' <span class="text-success">[Active since '+health.Details.ActiveNode.FirstSeenActive+']</span>';
				}
				if (node.Hostname == health.Details.Hostname) {
    			message += ' <span class="text-primary">[This node]</span>';
    		}
				message += ' <span class="text-info">[Running since '+node.FirstSeenActive+']</span>';

    		addStatusTableData("Available node", message, app_version);
    	})

    	var userId = getUserId();
    	if (userId == "") {
    		userId = "[unknown]"
    	}
    	var userStatus = (isAuthorizedForAction() ? "admin" : "read only");
    	addStatusTableData("You", userId + ", " + userStatus, "");

    	if (isAuthorizedForAction()) {
    		addStatusActionButton("Reload configuration", "reload-configuration");
    		addStatusActionButton("Reset hostname resolve cache", "reset-hostname-resolve-cache");
    		addStatusActionButton("Reelect", "reelect");
    	}

    }, "json");
});
