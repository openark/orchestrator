
$(document).ready(function () {
    showLoader();
    $.get("/api/audit", function (auditEntries) {
            displayAudit(auditEntries);
    	}, "json");
    function displayAudit(auditEntries) {
        hideLoader();
        auditEntries.forEach(function (audit) {
    		var row = jQuery('<tr/>');
    		jQuery('<td/>', { text: audit.AuditTimestamp }).appendTo(row);
    		jQuery('<td/>', { text: audit.AuditType }).appendTo(row);
    		jQuery('<td/>', { text: audit.AuditInstanceKey.Hostname+":"+audit.AuditInstanceKey.Port }).appendTo(row);
    		jQuery('<td/>', { text: audit.Message }).appendTo(row);
    		row.appendTo('#audit tbody');    		
    	});        	
    }
});	
