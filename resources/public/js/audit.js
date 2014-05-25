
$(document).ready(function () {
    showLoader();
    $.get("/api/audit/"+currentPage(), function (auditEntries) {
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
        if (currentPage() <= 0) {
        	$("#audit .pager .previous").addClass("disabled");
        }
        if (auditEntries.length == 0) {
        	$("#audit .pager .next").addClass("disabled");        	
        }
        $("#audit .pager .previous").not(".disabled").find("a").click(function() {
            window.location.href = "/web/audit/"+(currentPage() - 1);
        });
        $("#audit .pager .next").not(".disabled").find("a").click(function() {
            window.location.href = "/web/audit/"+(currentPage() + 1);
        });
        $("#audit .pager .disabled a").click(function() {
            return false;
        });
    }
});	
