
$(document).ready(function () {
    showLoader();
    
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
	    	$("#clusters [data-cluster-name='" + clusterName + "'].popover").find(".popover-content .pull-right").append('<span class="badge '+badgeClass+'" title="' + title + '">' + count + '</span> ');
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
    		$("#clusters").append('<div xmlns="http://www.w3.org/1999/xhtml" class="popover instance right" data-cluster-name="'+cluster.ClusterName+'"><div class="arrow"></div><h3 class="popover-title"><div class="pull-left"><a href="/web/cluster/'+cluster.ClusterName+'"><span>'+cluster.ClusterName+'</span></a></div><div class="pull-right"></div>&nbsp;<br/>&nbsp;</h3><div class="popover-content"></div></div>');
    		var popoverElement = $("#clusters [data-cluster-name='" + cluster.ClusterName + "'].popover");

            if (typeof removeTextFromHostnameDisplay != "undefined" && removeTextFromHostnameDisplay()) {
              	var title = cluster.ClusterName.replace(removeTextFromHostnameDisplay(), '');
                popoverElement.find("h3 .pull-left a span").html(title);
            } 
    		if (cluster.ClusterAlias != "") {
                popoverElement.find("h3 .pull-left a span").addClass("small");
                popoverElement.find("h3 .pull-left").prepend('<a href="/web/cluster/alias/'+encodeURIComponent(cluster.ClusterAlias)+'"><strong>'+cluster.ClusterAlias+'</strong></a><br/>');
                popoverElement.find("h3 .pull-right").append('<a href="/web/cluster/alias/'+encodeURIComponent(cluster.ClusterAlias)+'?compact=true"><span class="glyphicon glyphicon-compressed" title="Compact display"></span></a>');
    		}
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

    if (isAuthorizedForAction()) {
    	// Read-only users don't get auto-refresh. Sorry!
    	activateRefreshTimer();
    }
});	
