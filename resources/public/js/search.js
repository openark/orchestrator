
$(document).ready(function () {
	$("#searchInput").val(currentSearchString());
	
	
    showLoader();
    $.get("/api/search/"+currentSearchString(), function (instances) {
        $.get("/api/maintenance",
            function (maintenanceList) {
        		normalizeInstances(instances, maintenanceList);
                displaySearchInstances(instances);
            }, "json");
    }, "json");
    function displaySearchInstances(instances) {
        hideLoader();
    	instances.forEach(function (instance) {
            var instanceEl = Instance.createElement(instance).addClass("instance-search").appendTo("#searchResults");
    		renderInstanceElement(instanceEl, instance, "search");
    	    instanceEl.find("h3").click(function () {
    	    	openNodeModal(instance);
    	    	return false;
    	    });	
    	});        	
	
        if (instances.length == 0) {
        	addAlert("No search results found for "+currentSearchString());
        }
    }
});	
