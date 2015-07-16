
$(document).ready(function () {
    showLoader();
    
    $.get("/api/clusters-info", function (clusters) {
        $.get("/api/replication-analysis", function (replicationAnalysis) {
	    	displayClustersAnalysis(clusters, replicationAnalysis);
        }, "json");
    }, "json");
    
    function sortByCountInstances(cluster1, cluster2) {
    	if (cluster2.allAnalysisDowntimed && !cluster1.allAnalysisDowntimed) {
    		return 1
    	}
    	if (cluster1.allAnalysisDowntimed && !cluster2.allAnalysisDowntimed) {
    		return 1
    	}
    	var diff = cluster2.CountInstances - cluster1.CountInstances;
    	if (diff != 0) {
    		return diff;
    	}
    	return cluster1.ClusterName.localeCompare(cluster2.ClusterName);
    }
    
    function displayClustersAnalysis(clusters, replicationAnalysis) {
        hideLoader();
        
        var clustersMap = {};
        clusters.forEach(function (cluster) {
        	cluster.analysisEntries = Array();
        	cluster.allAnalysisDowntimed = true;
        	clustersMap[cluster.ClusterName] = cluster;
        });
	    
        // Apply/associate analysis to clusters
        replicationAnalysis.Details.forEach(function (analysisEntry) {
        	if (!(analysisEntry.Analysis in interestingAnalysis)) {
	    		return;
	    	}
        	clustersMap[analysisEntry.ClusterDetails.ClusterName].analysisEntries.push(analysisEntry);
        	if (!analysisEntry.IsDowntimed) {
        		clustersMap[analysisEntry.ClusterDetails.ClusterName].allAnalysisDowntimed = false;
        	}
        });
        // Only keep clusters with some analysis (the rest are fine, no need to include them)
        clusters = clusters.filter(function(cluster) {
        	return (cluster.analysisEntries.length > 0);
        });

	    function displayInstancesBadge(popoverElement, text, count, badgeClass, title) {
	    	popoverElement.find(".popover-content>div").append('<div>'+text+':<div class="pull-right"><span class="badge '+badgeClass+'" title="' + title + '">' + count + '</span></div></div>');
	    }
	    function displayAnalysisEntry(analysisEntry, popoverElement) {
	    	if (!(analysisEntry.Analysis in interestingAnalysis)) {
	    		return;
	    	}
	    	var displayText = '<hr/><span><strong>'+analysisEntry.Analysis 
	    		+ (analysisEntry.IsDowntimed ? '<br/>[<i>downtime till '+analysisEntry.DowntimeEndTimestamp+'</i>]': '')
	    		+ "</strong></span>" 
	    		+ "<br/>" + "<span>" + analysisEntry.AnalyzedInstanceKey.Hostname+":"+analysisEntry.AnalyzedInstanceKey.Port+ "</span>" 
	    		;
	    	if (analysisEntry.IsDowntimed) {
	    		displayText = '<div class="downtimed">'+displayText+'</div>';
	    	}
	    	popoverElement.find(".popover-content>div").append('<div class="divider"></div><div>' 
	    			+ displayText +	'</div> ');
	    	displayInstancesBadge(popoverElement, "Affected slaves", analysisEntry.CountSlaves, "label-danger", "Slaves of failed instance");
	    }
        function displayCluster(cluster) {
    		$("#clusters_analysis").append('<div xmlns="http://www.w3.org/1999/xhtml" class="popover instance right" data-cluster-name="'+cluster.ClusterName+'"><div class="arrow"></div><h3 class="popover-title"><div class="pull-left"><a href="/web/cluster/'+cluster.ClusterName+'"><span>'+cluster.ClusterName+'</span></a></div><div class="pull-right"></div>&nbsp;<br/>&nbsp;</h3><div class="popover-content"><div></div></div></div>');
    		var popoverElement = $("#clusters_analysis [data-cluster-name='" + cluster.ClusterName + "'].popover");

            if (typeof removeTextFromHostnameDisplay != "undefined" && removeTextFromHostnameDisplay()) {
              	var title = cluster.ClusterName.replace(removeTextFromHostnameDisplay(), '');
                popoverElement.find("h3 .pull-left a span").html(title);
            } 
    		if (cluster.ClusterAlias != "") {
                popoverElement.find("h3 .pull-left a span").addClass("small");
                popoverElement.find("h3 .pull-left").prepend('<a href="/web/cluster/alias/'+encodeURIComponent(cluster.ClusterAlias)+'"><strong>'+cluster.ClusterAlias+'</strong></a><br/>');
                popoverElement.find("h3 .pull-right").append('<a href="/web/cluster/alias/'+encodeURIComponent(cluster.ClusterAlias)+'?compact=true"><span class="glyphicon glyphicon-compressed" title="Compact display"></span></a>');
    		}
    	    displayInstancesBadge(popoverElement, "Instances", cluster.CountInstances, "label-primary", "Total instances in cluster");
    	    
    	    cluster.analysisEntries.forEach(function (analysisEntry) {
    	    	displayAnalysisEntry(analysisEntry, popoverElement)
            });
        }

        clusters.sort(sortByCountInstances);
        clusters.forEach(function (cluster) {
        	displayCluster(cluster);
        });
                
        if (clusters.length == 0) {
        	// No problems
        	var info = "No incidents which require a failover to report. Orchestrator reports the following incidents:<ul>";
        	for ( var analysis in interestingAnalysis) {
				if (interestingAnalysis[analysis]) {
					info += "<li>" + analysis + "</li>";
				}
			}
        	info += "</ul>";
        	addInfo(info);
        }

        $("div.popover").popover();
        $("div.popover").show();
    }

    if (isAuthorizedForAction()) {
    	// Read-only users don't get auto-refresh. Sorry!
    	activateRefreshTimer();
    }
});	
