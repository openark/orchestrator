
$(document).ready(function () {
    showLoader();
    activateRefreshTimer();
    
    var errorMapping = {
   		"inMaintenanceProblem": {"badge": "label-info", "description": "In maintenance"}, 
   		"lastCheckInvalidProblem": {"badge": "label-fatal", "description": "Last check invalid"}, 
   		"notRecentlyCheckedProblem": {"badge": "label-stale", "description": "Not recently checked (stale)"}, 
   		"notReplicatingProblem": {"badge": "label-danger", "description": "Not replicating"}, 
   		"replicationLagProblem": {"badge": "label-warning", "description": "Replication lag"}
	};
    
    $.get("/api/clusters-info", function (clusters) {
        $.get("/api/problems", function (problemInstances) {
	    		normalizeInstances(problemInstances, []);
	    		displayClusters(clusters, problemInstances);
        }, "json");
    }, "json");
    function sortByCountInstances(cluster1, cluster2) {
    	var diff = cluster2.CountInstances - cluster1.CountInstances;
    	if (diff != 0) {
    		return diff;
    	}
    	return cluster1.ClusterName.localeCompare(cluster2.ClusterName);
    }
    
    function displayClusters(clusters, problemInstances) {
        hideLoader();
        
        clusters.sort(sortByCountInstances);
        var clustersProblems = {};
        clusters.forEach(function (cluster) {
        	clustersProblems[cluster.ClusterName] = {};
        });
        
	    function addInstancesBadge(clusterName, count, badgeClass, title) {
	    	$("#clusters [data-cluster-name='" + clusterName + "'].popover").find(".popover-content .pull-right").append('<span class="badge '+badgeClass+'" title="' + title + '"">' + count + '</span> ');
	    }
        
        function incrementClusterProblems(clusterName, problemType) {
        	if (clustersProblems[clusterName][problemType] > 0) {
        		clustersProblems[clusterName][problemType] = clustersProblems[clusterName][problemType] + 1;
        	} else {
        		clustersProblems[clusterName][problemType] = 1;
        	}
        }
        problemInstances.forEach(function(instance) {
	        if (instance.inMaintenanceProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "inMaintenanceProblem")
	        }
	        //
	        if (instance.lastCheckInvalidProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "lastCheckInvalidProblem")
	        } else if (instance.notRecentlyCheckedProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "notRecentlyCheckedProblem")
	        } else if (instance.notReplicatingProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "notReplicatingProblem")
	        } else if (instance.replicationLagProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "replicationLagProblem")
	        }
	    });

        clusters.forEach(function (cluster) {
    		$("#clusters").append('<div xmlns="http://www.w3.org/1999/xhtml" class="popover instance right" data-cluster-name="'+cluster.ClusterName+'"><div class="arrow"></div><h3 class="popover-title"><a href="/web/cluster/'+cluster.ClusterName+'">'+cluster.ClusterName+'</a></h3><div class="popover-content"></div></div>');
    		var popoverElement = $("#clusters [data-cluster-name='" + cluster.ClusterName + "'].popover");
    		var title = cluster.ClusterName;
            if (typeof removeTextFromHostnameDisplay != "undefined" && removeTextFromHostnameDisplay()) {
              	title = title.replace(removeTextFromHostnameDisplay(), '');
            } 
    		if (cluster.ClusterAlias != "") {
    			title = '<strong>' + cluster.ClusterAlias + '</strong>, <span class="small">' + title + '</span>';
    		}
    		popoverElement.find("h3 a").html(title);
    	    var contentHtml = ''
    				+ '<div>Instances: <div class="pull-right"></div></div>'
    			;
    	    popoverElement.find(".popover-content").html(contentHtml);
    	    addInstancesBadge(cluster.ClusterName, cluster.CountInstances, "label-primary", "Total instances in cluster");
    	    for (var problemType in clustersProblems[cluster.ClusterName]) {
    	    	addInstancesBadge(cluster.ClusterName, clustersProblems[cluster.ClusterName][problemType], errorMapping[problemType]["badge"], errorMapping[problemType]["description"]);
    	    }
        });     
        
        $("div.popover").popover();
        $("div.popover").show();
	
        if (clusters.length == 0) {
        	addAlert("No clusters found");
        }
    }
});	
