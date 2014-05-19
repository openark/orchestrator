
$(document).ready(function () {
    showLoader();
    $.get("/api/problems", function (instances) {
			normalizeInstances(instances, []);
	        displayProblemInstances(instances);
	    }, "json");
    function displayProblemInstances(instances) {
        hideLoader();
        
        function SortByProblemOrder(instance0, instance1){
        	return instance0.problemOrder - instance1.problemOrder; 
        }
        instances.sort(SortByProblemOrder);

        var problemInstancesFound = false;
    	instances.forEach(function (instance) {
    		if (instance.hasProblem) {
	    		$("#instance_problems_content .panel-body").append('<div xmlns="http://www.w3.org/1999/xhtml" class="popover instance" data-nodeid="'+instance.id+'"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>');
	
	    		var popoverElement = $("#instance_problems [data-nodeid='" + instance.id + "'].popover");
	    		renderInstanceElement(popoverElement, instance, "problems");
	    		problemInstancesFound = true;
    		}
    	});        	
    	if (problemInstancesFound) {
    		$("#instance_problems .collapse").addClass("in");
    	}
        
        $("div.popover").popover();
        $("div.popover").show();
    }
});	
