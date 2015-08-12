
$(document).ready(function () {
    showLoader();
    $.get("/api/audit-failure-detection/"+currentPage(), function (auditEntries) {
            displayAudit(auditEntries);
    	}, "json");
    function displayAudit(auditEntries) {
        hideLoader();
        auditEntries.forEach(function (audit) {
        	var analyzedInstanceDisplay = audit.AnalysisEntry.AnalyzedInstanceKey.Hostname+":"+audit.AnalysisEntry.AnalyzedInstanceKey.Port;
    		var row = jQuery('<tr/>');
    		$('<td/>', { text: audit.AnalysisEntry.Analysis }).appendTo(row);
    		$('<a/>',  { text: analyzedInstanceDisplay, href: "/web/search/" + analyzedInstanceDisplay }).wrap($("<td/>")).parent().appendTo(row);
    		$('<td/>', { text: audit.AnalysisEntry.CountSlaves }).appendTo(row);
    		$('<a/>',  { text: audit.AnalysisEntry.ClusterDetails.ClusterName, href: "/web/cluster/"+audit.AnalysisEntry.ClusterDetails.ClusterName}).wrap($("<td/>")).parent().appendTo(row);
    		$('<a/>',  { text: audit.AnalysisEntry.ClusterDetails.ClusterAlias, href: "/web/cluster/alias/"+audit.AnalysisEntry.ClusterDetails.ClusterAlias}).wrap($("<td/>")).parent().appendTo(row);
    		$('<td/>', { text: audit.RecoveryStartTimestamp }).appendTo(row);
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
