
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
    		$("#searchResults").append('<div xmlns="http://www.w3.org/1999/xhtml" class="popover instance right" data-nodeid="'+instance.id+'"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>');

            $("[data-nodeid='" + instance.id + "'].popover h3").html(
            		instance.canonicalTitle + '<div class="pull-right"><a href="#"><span class="glyphicon glyphicon-cog"></span></a></div>');
            if (instance.inMaintenance) {
                $("[data-nodeid='" + instance.id + "'].popover h3").addClass("label-info");
	        } else if (!instance.SecondsSinceLastSeen.Valid || instance.SecondsSinceLastSeen.Int64 > 3600) {
	            $("[data-nodeid='" + instance.id + "'].popover h3").addClass("label-fatal");
            } else if (!instance.SecondsBehindMaster.Valid) {
                $("[data-nodeid='" + instance.id + "'].popover h3").addClass("label-danger");
	        } else if (instance.SecondsBehindMaster.Valid && instance.SecondsBehindMaster.Int64 > 10) {
	            $("[data-nodeid='" + instance.id + "'].popover h3").addClass("label-warning");
	        }
            $("[data-nodeid='" + instance.id + "'].popover .popover-content").html(''
                	+ '<div class="pull-right">' + instance.SecondsBehindMaster.Int64 + ' seconds lag</div>'
           		+ '<p>' 
        			+ instance.Version + " " + instance.Binlog_format 
                	+ '<br/>Cluster: <a href="/web/cluster/'+instance.ClusterName+'">'+instance.ClusterName+'</a>'
                + '</p>'
            );
            $("[data-nodeid='" + instance.id + "'].popover h3 a").click(function () {
            	openNodeModal(instance);
            	return false;
            });
    	});        	
        
        $("div.popover").popover();
        $("div.popover").show();
	
        if (instances.length == 0) {
        	addAlert("No search results found for "+currentSearchString());
        }
    }
});	
