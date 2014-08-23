
function appendSeedDetails(seed, selector) {    	
	var row = '<tr>';
	var statusMessage;
	if (seed.IsComplete) {
		statusMessage = (seed.IsSuccessful ? '<span class="text-success">Success</span>' : '<span class="text-danger">Fail</span>');
	} else {
		statusMessage = '<span class="text-info">Active</span>';
	}
	row += '<td>' + statusMessage + '</td>';
	row += '<td><a href="/web/agent-seed-details/' + seed.SeedId + '">' + seed.SeedId + '</a></td>';
	row += '<td><a href="/web/agent/'+seed.TargetHostname+'">'+seed.TargetHostname+'</a></td>';
	row += '<td><a href="/web/agent/'+seed.SourceHostname+'">'+seed.SourceHostname+'</a></td>';
	row += '<td>' + seed.StartTimestamp + '</td>';
	row += '<td>' + (seed.IsComplete ? seed.EndTimestamp : '-') + '</td>';
	row += '</tr>';
	$(selector).append(row);
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
