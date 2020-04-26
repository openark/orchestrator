
function appendSeedDetails(seed, selector) {    	
	var row = '<tr>';
	var statusMessage;
	if ((seed.Status == "Failed") || (seed.Status == "Error")) {
		statusMessage = '<span class="text-danger">'+seed.Status+'</span>';
	} else if (seed.Status == "Completed") {
		statusMessage = '<span class="text-success">'+seed.Status+'</span>';
	} else {
		statusMessage = '<span class="text-info">'+seed.Status+'</span>';
	}
	row += '<td><a href="' + appUrl('/web/seed-details/' + seed.SeedID) + '">' + seed.SeedID + '</a></td>';
	row += '<td>' + statusMessage + '</td>';
	row += '<td><a href="' + appUrl('/web/agent/'+seed.TargetHostname) + '">'+seed.TargetHostname+'</a></td>';
	row += '<td><a href="' + appUrl('/web/agent/'+seed.SourceHostname) + '">'+seed.SourceHostname+'</a></td>';
	row += '<td>' + seed.SeedMethod + '</td>';
	row += '<td>' + seed.Stage + '</td>';
	row += '<td>' + seed.Retries + '</td>';
	row += '<td>' + moment.utc(seed.UpdatedAt).format("YYYY-MM-DD HH:mm:ss") + '</td>';
	row += '<td>' + moment.utc(seed.StartTimestamp).format("YYYY-MM-DD HH:mm:ss") + '</td>';
	row += '<td>' + ((seed.Status == "Completed" || seed.Status == "Failed") ? moment.utc(seed.EndTimestamp).format("YYYY-MM-DD HH:mm:ss"): '<button class="btn btn-xs btn-danger" id="abort_seed_button" data-seed-source-host="'+seed.SourceHostname+'" data-seed-target-host="'+seed.TargetHostname+'" data-seed-id="' + seed.SeedID + '">Abort</button>') + '</td>';
	row += '</tr>';
	$(selector).append(row);
    hideLoader();
}

function appendSeedState(seedState,selector) {
	var row = '<tr>';
	row += '<td>' + moment.utc(seedState.Timestamp).format("YYYY-MM-DD HH:mm:ss") + '</td>';
	row += '<td>' + seedState.Stage + '</td>';
	row += '<td>' + seedState.Hostname + '</td>';
	row += '<td>' + seedState.Status + '</td>';
	row += '<td>' + seedState.Details + '</td>';
	row += '</tr>';
	$(selector).append(row);
    hideLoader();
}

$("body").on("click", "#abort_seed_button", function(event) {
	var seedID = $(event.target).attr("data-seed-id");
	var sourceHost = $(event.target).attr("data-seed-source-host");
	var targetHost = $(event.target).attr("data-seed-target-host");

	var message = "Are you sure you wish to abort seed " + seedID + " from <code><strong>" + 
		sourceHost + "</strong></code> to <code><strong>" + 
		targetHost + "</strong></code> ?";
	bootbox.confirm(message, function(confirm) {
		if (confirm) {
	    	showLoader();
	        $.get(appUrl("/api/agent-abort-seed/"+seedID), function (operationResult) {
				hideLoader();
				location.reload();
			}, "json").fail(function(operationResult) {
				hideLoader();
				if (operationResult.responseJSON.Code == "ERROR") {
				  addAlert(operationResult.responseJSON.Message);
				}
			});
		}
	});
});