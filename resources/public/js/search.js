
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
    		$("#searchResults").append('<div xmlns="http://www.w3.org/1999/xhtml" class="popover instance right" data-nodeid="'+instance.id+'"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>');

    		var popoverElement = $("#searchResults [data-nodeid='" + instance.id + "'].popover");
    		renderInstanceElement(popoverElement, instance, "search");
    	});        	
        
        $("div.popover").popover();
        $("div.popover").show();
	
        if (instances.length == 0) {
        	addAlert("No search results found for "+currentSearchString());
        }
    }
});	
