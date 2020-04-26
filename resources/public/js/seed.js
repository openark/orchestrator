
$(document).ready(function () {
    showLoader();
    
    $.get(appUrl("/api/agent-seed-details/"+currentSeedId()), function (seed) {
	        showLoader();
	    	appendSeedDetails(seed, "#seed_details");
	    	if (!seed.IsComplete) {
	    		activateRefreshTimer();
	    	}
	    }, "json");

    $.get(appUrl("/api/agent-seed-states/"+currentSeedId()), function (seedStates) {
	        showLoader();
	        seedStates.forEach(function (seedState) {
	        	appendSeedState(seedState, "#seed_states");
	    	});
	    }, "json");
});	
