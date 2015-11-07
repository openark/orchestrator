
$(document).ready(function () {
    showLoader();
    $.get("/api/audit-failure-detection/"+currentPage(), function (auditEntries) {
        $.get("/api/replication-analysis-changelog", function (analysisChangelog) {
        	displayAudit(auditEntries, analysisChangelog);
        }, "json");
    }, "json");
    function displayAudit(auditEntries, analysisChangelog) {
    	var changelogMap = {}
    	analysisChangelog.forEach(function (changelogEntry) {
    		changelogMap[getInstanceId(changelogEntry.AnalyzedInstanceKey.Hostname, changelogEntry.AnalyzedInstanceKey.Port)] = changelogEntry.Changelog;
    	});
    	
        hideLoader();
        auditEntries.forEach(function (audit) {
        	var analyzedInstanceDisplay = audit.AnalysisEntry.AnalyzedInstanceKey.Hostname+":"+audit.AnalysisEntry.AnalyzedInstanceKey.Port;
    		var row = jQuery('<tr/>');
    		$('<td/>', { text: audit.AnalysisEntry.Analysis }).prepend('<span class="more-detection-info pull-right glyphicon glyphicon-info-sign text-primary" data-toggle="popover" data-html="true" title="" data-content=""></span>').appendTo(row);
    		$('<a/>',  { text: analyzedInstanceDisplay, href: "/web/search/" + analyzedInstanceDisplay }).wrap($("<td/>")).parent().appendTo(row);
    		$('<td/>', { text: audit.AnalysisEntry.CountSlaves }).appendTo(row);
    		$('<a/>',  { text: audit.AnalysisEntry.ClusterDetails.ClusterName, href: "/web/cluster/"+audit.AnalysisEntry.ClusterDetails.ClusterName}).wrap($("<td/>")).parent().appendTo(row);
    		$('<a/>',  { text: audit.AnalysisEntry.ClusterDetails.ClusterAlias, href: "/web/cluster/alias/"+audit.AnalysisEntry.ClusterDetails.ClusterAlias}).wrap($("<td/>")).parent().appendTo(row);
    		$('<td/>', { text: audit.RecoveryStartTimestamp }).appendTo(row);

    		var moreInfo = "";
    		moreInfo += '<div>Detected: '+audit.RecoveryStartTimestamp+'</div>';
    		if (audit.AnalysisEntry.SlaveHosts.length > 0) {
    			moreInfo += '<div>'+audit.AnalysisEntry.CountSlaves+' slave hosts :<ul>';
        		audit.AnalysisEntry.SlaveHosts.forEach(function(instanceKey) {
        			moreInfo += "<li><code>"+getInstanceTitle(instanceKey.Hostname, instanceKey.Port)+"</code></li>";    			
        		});
        		moreInfo += "</ul></div>";
    		}
    		var changelog = changelogMap[getInstanceId(audit.AnalysisEntry.AnalyzedInstanceKey.Hostname, audit.AnalysisEntry.AnalyzedInstanceKey.Port)];
    		if (changelog) {
    			moreInfo += '<div>Changelog :<ul>';
    			changelog.split(",").reverse().forEach(function(changelogEntry) {
    				var changelogEntryTokens = changelogEntry.split(';');
    				var changelogEntryTimestamp = changelogEntryTokens[0];
    				var changelogEntryAnalysis = changelogEntryTokens[1];

    				if (changelogEntryTimestamp > audit.RecoveryStartTimestamp) {
    					// This entry is newer than the detection time; irrelevant
    					return;
    				}
        			moreInfo += "<li><code>"+changelogEntryTimestamp + " <strong>" + changelogEntryAnalysis + "</strong></code></li>";
        		});
        		moreInfo += "</ul></div>";
    		}
    		moreInfo += "<div>Proccessed by <code>"+audit.ProcessingNodeHostname+"</code></div>";
    		row.find(".more-detection-info").attr("title", audit.AnalysisEntry.Analysis+'<ul><li><code>'+analyzedInstanceDisplay+'</code></li></ul>');
    		row.find(".more-detection-info").attr("data-content", moreInfo);
    		row.find('[data-toggle="popover"]').popover();    		
    		
    		row.appendTo('#audit tbody');
    	});
        if (currentPage() <= 0) {
        	$("#audit .pager .previous").addClass("disabled");
        }
        if (auditEntries.length == 0) {
        	$("#audit .pager .next").addClass("disabled");        	
        }
        $("#audit .pager .previous").not(".disabled").find("a").click(function() {
            window.location.href = "/web/audit-failure-detection/"+(currentPage() - 1);
        });
        $("#audit .pager .next").not(".disabled").find("a").click(function() {
            window.location.href = "/web/audit-failure-detection/"+(currentPage() + 1);
        });
        $("#audit .pager .disabled a").click(function() {
            return false;
        });
    }
});	
