
$(document).ready(function () {
    showLoader();
    $.get("/api/audit-recovery/"+currentPage(), function (auditEntries) {
            displayAudit(auditEntries);
    	}, "json");
    function displayAudit(auditEntries) {
        hideLoader();
        auditEntries.forEach(function (audit) {
    		var row = jQuery('<tr/>');
    		jQuery('<td/>', { text: audit.AnalysisEntry.Analysis }).appendTo(row);
    		jQuery('<td/>', { text: audit.AnalysisEntry.AnalyzedInstanceKey.Hostname+":"+audit.AnalysisEntry.AnalyzedInstanceKey.Port }).appendTo(row);
    		jQuery('<td/>', { text: audit.AnalysisEntry.CountSlaves }).appendTo(row);
    		jQuery('<td/>', { text: audit.AnalysisEntry.ClusterDetails.ClusterName }).appendTo(row);
    		jQuery('<td/>', { text: audit.AnalysisEntry.ClusterDetails.ClusterAlias }).appendTo(row);
    		jQuery('<td/>', { text: audit.RecoveryStartTimestamp }).appendTo(row);
    		jQuery('<td/>', { text: audit.RecoveryEndTimestamp }).appendTo(row);
    		jQuery('<td/>', { text: audit.SuccessorKey.Hostname+":"+audit.SuccessorKey.Port }).appendTo(row);
    		row.appendTo('#audit tbody');
    	});
        if (currentPage() <= 0) {
        	$("#audit .pager .previous").addClass("disabled");
        }
        if (auditEntries.length == 0) {
        	$("#audit .pager .next").addClass("disabled");        	
        }
        $("#audit .pager .previous").not(".disabled").find("a").click(function() {
            window.location.href = "/web/audit-recovery/"+(currentPage() - 1);
        });
        $("#audit .pager .next").not(".disabled").find("a").click(function() {
            window.location.href = "/web/audit-recovery/"+(currentPage() + 1);
        });
        $("#audit .pager .disabled a").click(function() {
            return false;
        });
    }
});	
