
$(document).ready(function () {
    showLoader();
    
    $.get("/api/clusters-info", function (clusters) {
        $.get("/api/replication-analysis", function (replicationAnalysis) {
	    	displayClustersAnalysis(clusters, replicationAnalysis);
        }, "json");
    }, "json");
    function sortByCountInstances(cluster1, cluster2) {
    	var diff = cluster2.CountInstances - cluster1.CountInstances;
    	if (diff != 0) {
    		return diff;
    	}
    	return cluster1.ClusterName.localeCompare(cluster2.ClusterName);
    }
    
    var interestingAnalysis = {
        	"DeadMaster": true,	
        	"DeadMasterAndSlaves": true,	
        	"DeadMasterAndSomeSlaves": true,	
        	"DeadIntermediateMaster": true
    };
    
    function displayClustersAnalysis(clusters, replicationAnalysis) {
        hideLoader();
        
        clusters.sort(sortByCountInstances);

	    function addInstancesBadge(clusterName, text, count, badgeClass, title) {
	    	$("#clusters_analysis [data-cluster-name='" + clusterName + "'].popover").find(".popover-content>div").append('<div>'+text+':<div class="pull-right"><span class="badge '+badgeClass+'" title="' + title + '">' + count + '</span></div></div>');
	    }
	    function addAnalysisEntry(analysisEntry, countInstances) {
	    	if (!(analysisEntry.Analysis in interestingAnalysis)) {
	    		return;
	    	}
	    	var displayText = "<hr/><span><strong>"+analysisEntry.Analysis+"</strong></span>" 
	    		+ "<br/>" + "<span>" + analysisEntry.AnalyzedInstanceKey.Hostname+":"+analysisEntry.AnalyzedInstanceKey.Port+ "</span>" 
	    		;
	    	$("#clusters_analysis [data-cluster-name='" + analysisEntry.ClusterName + "'].popover").find(".popover-content>div").append('<div class="divider"></div><div>' 
	    			+ displayText +	'</div> ');
    		addInstancesBadge(analysisEntry.ClusterName, "Affected slaves", analysisEntry.CountSlaves, "label-danger", "Slaves of failed instance");
	    }
	    
        var clustersWithProblems = {};

        replicationAnalysis.Details.forEach(function (analysisEntry) {
        	if (!(analysisEntry.Analysis in interestingAnalysis)) {
	    		return;
	    	}
        	clustersWithProblems[analysisEntry.ClusterName] = true;
        });


        clusters.forEach(function (cluster) {
        	if (cluster.ClusterName in clustersWithProblems) {
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
	    	    addInstancesBadge(cluster.ClusterName, "Instances", cluster.CountInstances, "label-primary", "Total instances in cluster");
        	}
        });     
        
        replicationAnalysis.Details.forEach(function (analysisEntry) {
        	//addInstancesBadge(analysisEntry.ClusterName, analysisEntry.Analysis, "label-error", analysisEntry.Analysis);
        	addAnalysisEntry(analysisEntry)
        	console.log(analysisEntry);
        });
        
        if (Object.keys(clustersWithProblems).length == 0) {
        	// No problems
        	addInfo("No incidents which require a failvoer to report. Orchestrator reports dead-master and dead-intermediate master issues.");
        }

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
