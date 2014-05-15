
$(document).ready(function () {
	$("#searchInput").val(currentSearchString());
	
	
    showLoader();
    $.get("/api/search/"+currentSearchString(), function (instances) {
        $.get("/api/maintenance",
            function (maintenanceList) {
                displaySearchInstances(instances,
                		maintenanceList);
            }, "json");
    }, "json");
    function displaySearchInstances(instances, maintenanceList) {
        hideLoader();
        instances.forEach(function (instance) {
            instance.id = getInstanceId(instance.Key.Hostname, instance.Key.Port);
            instance.title= instance.Key.Hostname+':'+instance.Key.Port;
            instance.canonicalTitle = instance.title;
        });
        var nodesMap = instances.reduce(function (map, node) {
            map[node.id] = node;
            return map;
        }, {});
        // mark maintenance instances
        maintenanceList.forEach(function (maintenanceEntry) {
            var instanceId = getInstanceId(maintenanceEntry.Key.Hostname, maintenanceEntry.Key.Port)
            if (instanceId in nodesMap) {
                nodesMap[instanceId].inMaintenance = true;
                nodesMap[instanceId].maintenanceEntry = maintenanceEntry;
            }
        });
    	instances.forEach(function (instance) {
    		var node=instance;
    		$("#searchResults").append('<div xmlns="http://www.w3.org/1999/xhtml" class="popover instance right" data-nodeid="'+instance.id+'"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>');

            $("[data-nodeid='" + instance.id + "'].popover h3").html(
            		instance.canonicalTitle + '<div class="pull-right"><a href="#"><span class="glyphicon glyphicon-cog"></span></a></div>');
            if (instance.inMaintenance) {
                $("[data-nodeid='" + instance.id + "'].popover h3").addClass("label-info");
            } else if (instance.SecondsBehindMaster > 10) {
                $("[data-nodeid='" + instance.id + "'].popover h3").addClass("label-danger");
            }
            $("[data-nodeid='" + instance.id + "'].popover .popover-content").html(
           		'<p>' 
                	+ '<div class="pull-right">' + instance.SecondsBehindMaster + ' seconds lag</div>'
        			+ instance.Version + " " + instance.Binlog_format 
                	+ '<br/><br/>Cluster: <a href="/web/cluster/'+instance.ClusterName+'">'+instance.ClusterName+'</a>'
                + '</p>'
            );
            $("[data-nodeid='" + instance.id + "'].popover h3 a").click(function () {
            	openNodeModal(node);
            	return false;
            });
    	});        	
        
        $("div.popover").popover();
        $("div.popover").show();
	
    }
});	
