
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
    		var row = $('<tr/>');
    		var ack = $('<span class="pull-left glyphicon acknowledge-indicator" title=""></span>');
    		if (audit.Acknowledged) {
    			ack.addClass("glyphicon-ok-sign").addClass("text-primary");
    			var ackTitle = "Acknowledged by "+audit.AcknowledgedBy+" at "+audit.AcknowledgedAt+": "+audit.AcknowledgedComment;
    			ack.attr("title", ackTitle);
    		} else {
    			ack.addClass("glyphicon-question-sign").addClass("text-danger").addClass("unacknowledged");
    			ack.attr("data-recovery-id", audit.Id);
    			ack.attr("title", "Unacknowledged. Click to acknowledge");
    		}
    		$('<td/>', { text: audit.AnalysisEntry.Analysis }).prepend(ack).prepend('<span class="more-recovery-info pull-right glyphicon glyphicon-info-sign text-primary" data-toggle="popover" data-html="true" title="" data-content=""></span>').appendTo(row);
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
    		var moreInfo = "";
    		if (audit.Acknowledged) {
    			moreInfo += '<div>Acknowledged by '+audit.AcknowledgedBy+', '+audit.AcknowledgedAt+'<ul>';
    			moreInfo += "<li>"+audit.AcknowledgedComment+"</li>";    		
    			moreInfo += '</ul></div>';
    		} else {
    			moreInfo += '<div><strong>Unacknowledged</strong></div>';
    		}
    		if (audit.LostSlaves.length > 0) {
    			moreInfo += "<div>Lost slaves:<ul>";
        		audit.LostSlaves.forEach(function(instanceKey) {
        			moreInfo += "<li><code>"+getInstanceTitle(instanceKey.Hostname, instanceKey.Port)+"</code></li>";    			
        		});
        		moreInfo += "</ul></div>";
    		}
    		if (audit.ParticipatingInstanceKeys.length > 0) {
    			moreInfo += "<div>Participating instances:<ul>";
        		audit.ParticipatingInstanceKeys.forEach(function(instanceKey) {
        			moreInfo += "<li><code>"+getInstanceTitle(instanceKey.Hostname, instanceKey.Port)+"</code></li>";    			
        		});
        		moreInfo += "</ul></div>";
    		}
    		if (audit.AnalysisEntry.SlaveHosts.length > 0) {
    			moreInfo += '<div>'+audit.AnalysisEntry.CountSlaves+' slave hosts :<ul>';
        		audit.AnalysisEntry.SlaveHosts.forEach(function(instanceKey) {
        			moreInfo += "<li><code>"+getInstanceTitle(instanceKey.Hostname, instanceKey.Port)+"</code></li>";    			
        		});
        		moreInfo += "</ul></div>";
    		}
    		if (audit.AllErrors.length > 0 && audit.AllErrors[0]) {
    			moreInfo += "All errors:<ul>";
        		audit.AllErrors.forEach(function(err) {
        			moreInfo += "<li>"+err;    			
        		});
        		moreInfo += "</ul>";
    		}
    		moreInfo += "<div>Proccessed by <code>"+audit.ProcessingNodeHostname+"</code></div>";
    		row.find(".more-recovery-info").attr("title", audit.AnalysisEntry.Analysis);
    		row.find(".more-recovery-info").attr("data-content", moreInfo);
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
            window.location.href = baseWebUri+(currentPage() - 1);
        });
        $("#audit .pager .next").not(".disabled").find("a").click(function() {
            window.location.href = baseWebUri+(currentPage() + 1);
        });
        $("#audit .pager .disabled a").click(function() {
            return false;
        });
        
        $("body").on("click", ".acknowledge-indicator.unacknowledged", function (event) {
        	var recoveryId = $(event.target).attr("data-recovery-id");
            bootbox.prompt({
                title: "Acknowledge recovery",
                value: "comment",
                callback: function (result) {
                    if (result !== null) {
                        showLoader();
                        $.get("/api/ack-recovery/"+recoveryId+"?comment="+encodeURIComponent(result), function (operationResult) {
                            hideLoader();
                            if (operationResult.Code == "ERROR") {
                                addAlert(operationResult.Message)
                            } else {
                                location.reload();
                            }
                        }, "json");
                    }
                }
            });
        });
    }
});	
