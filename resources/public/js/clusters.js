
$(document).ready(function () {
    showLoader();
    activateRefreshTimer();
    
    $.get("/api/clusters-info", function (clusters) {
        $.get("/api/problems", function (problemInstances) {
	    		normalizeInstances(problemInstances, []);
	    		displayClusters(clusters, problemInstances);
        }, "json");
    }, "json");
    function displayClusters(clusters, problemInstances) {
        hideLoader();
        var clustersProblems = {};
        clusters.forEach(function (cluster) {
        	clustersProblems[cluster.ClusterName] = {};
        });
        
	    function addInstancesBadge(clusterName, count, badgeClass) {
	    	$("#clusters [data-cluster-name='" + clusterName + "'].popover").find(".popover-content .pull-right").append('<span class="badge '+badgeClass+'">' + count + '</span> ');
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
	        	incrementClusterProblems(instance.ClusterName, "label-info")
	        }
	        //
	        if (instance.lastCheckInvalidProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "label-fatal")
	        } else if (instance.notRecentlyCheckedProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "label-stale")
	        } else if (instance.notReplicatingProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "label-danger")
	        } else if (instance.replicationLagProblem()) {
	        	incrementClusterProblems(instance.ClusterName, "label-warning")
	        }
	    });

        clusters.forEach(function (cluster) {
    		$("#clusters").append('<div xmlns="http://www.w3.org/1999/xhtml" class="popover instance right" data-cluster-name="'+cluster.ClusterName+'"><div class="arrow"></div><h3 class="popover-title"><a href="/web/cluster/'+cluster.ClusterName+'">'+cluster.ClusterName+'</a></h3><div class="popover-content"></div></div>');
    		var popoverElement = $("#clusters [data-cluster-name='" + cluster.ClusterName + "'].popover");
    		popoverElement.find("h3 a").html(cluster.ClusterName);
    	    var contentHtml = ''
    				+ '<div>Instances: <div class="pull-right"></div></div>'
    			;
    	    popoverElement.find(".popover-content").html(contentHtml);
    	    addInstancesBadge(cluster.ClusterName, cluster.CountInstances, "alert-default");
    	    for (var problemType in clustersProblems[cluster.ClusterName]) {
    	    	addInstancesBadge(cluster.ClusterName, clustersProblems[cluster.ClusterName][problemType], problemType);
    	    }
        });     
        
        $("div.popover").popover();
        $("div.popover").show();
	
        if (clusters.length == 0) {
        	addAlert("No clusters found");
        }
    }
});	
