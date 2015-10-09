
$(document).ready(function () {
    showLoader();
    var apiUri = "/api/audit-recovery/"+currentPage();
    if (auditCluster()) {
    	apiUri = "/api/audit-recovery/cluster/"+auditCluster()+"/"+currentPage();
    }
    $.get(apiUri, function (auditEntries) {
            displayAudit(auditEntries);
    	}, "json");
    function displayAudit(auditEntries) {
    	var baseWebUri = "/web/audit-recovery/";
    	if (auditCluster()) {
    		baseWebUri += "cluster/"+auditCluster()+"/";
        }

    	hideLoader();
        auditEntries.forEach(function (audit) {
        	var analyzedInstanceDisplay = audit.AnalysisEntry.AnalyzedInstanceKey.Hostname+":"+audit.AnalysisEntry.AnalyzedInstanceKey.Port;
        	var sucessorInstanceDisplay = audit.SuccessorKey.Hostname+":"+audit.SuccessorKey.Port;
    		var row = jQuery('<tr/>');
    		$('<td/>', { text: audit.AnalysisEntry.Analysis }).prepend('<span class="more-recovery-info pull-right glyphicon glyphicon-info-sign text-primary" data-toggle="tooltip" data-placement="right" data-html="true" title=""></span>').appendTo(row);
    		$('<a/>',  { text: analyzedInstanceDisplay, href: "/web/search/" + analyzedInstanceDisplay }).wrap($("<td/>")).parent().appendTo(row);
    		$('<td/>', { text: audit.AnalysisEntry.CountSlaves }).appendTo(row);
    		$('<a/>',  { text: audit.AnalysisEntry.ClusterDetails.ClusterName, href: "/web/cluster/"+audit.AnalysisEntry.ClusterDetails.ClusterName}).wrap($("<td/>")).parent().appendTo(row);
    		$('<a/>',  { text: audit.AnalysisEntry.ClusterDetails.ClusterAlias, href: "/web/cluster/alias/"+audit.AnalysisEntry.ClusterDetails.ClusterAlias}).wrap($("<td/>")).parent().appendTo(row);
    		$('<td/>', { text: audit.RecoveryStartTimestamp }).appendTo(row);
    		$('<td/>', { text: audit.RecoveryEndTimestamp }).appendTo(row);
    		if (audit.RecoveryEndTimestamp && !audit.IsSuccessful && !audit.SuccessorKey.Hostname) {
        		$('<td/>', { text: "FAIL" }).appendTo(row);
    		} else if (audit.SuccessorKey.Hostname) {
    			$('<a/>',  { text: sucessorInstanceDisplay, href: "/web/search/" + sucessorInstanceDisplay }).wrap($("<td/>")).parent().appendTo(row);
    		} else {
    			$('<td/>', { text: "pending" }).appendTo(row);
    		}
    		var moreInfo = "<h6>"+audit.AnalysisEntry.Analysis+"</h6>";
    		if (audit.LostSlaves.length > 0) {
    			moreInfo += "Lost slaves:<ul>";
        		audit.LostSlaves.forEach(function(instanceKey) {
        			moreInfo += "<li>"+getInstanceTitle(instanceKey.Hostname, instanceKey.Port);    			
        		});
        		moreInfo += "</ul>";
    		}
    		if (audit.ParticipatingInstanceKeys.length > 0) {
    			moreInfo += "Participating instances:<ul>";
        		audit.ParticipatingInstanceKeys.forEach(function(instanceKey) {
        			moreInfo += "<li>"+getInstanceTitle(instanceKey.Hostname, instanceKey.Port);    			
        		});
        		moreInfo += "</ul>";
    		}
    		if (audit.AllErrors.length > 0 && audit.AllErrors[0]) {
    			moreInfo += "All errors:<ul>";
        		audit.AllErrors.forEach(function(err) {
        			moreInfo += "<li>"+err;    			
        		});
        		moreInfo += "</ul>";
    		}
    		row.find(".more-recovery-info").attr("title", moreInfo);
    		row.find('[data-toggle="tooltip"]').tooltip();
    		row.appendTo('#audit tbody');
    	});
        if (currentPage() <= 0) {
        	$("#audit .pager .previous").addClass("disabled");
        }
        if (auditEntries.length == 0) {
        	$("#audit .pager .next").addClass("disabled");        	
        }
        $("#audit .pager .previous").not(".disabled").find("a").click(function() {
            window.location.href = baseWebUri+(currentPage() - 1);
        });
        $("#audit .pager .next").not(".disabled").find("a").click(function() {
            window.location.href = baseWebUri+(currentPage() + 1);
        });
        $("#audit .pager .disabled a").click(function() {
            return false;
        });
    }
});	
