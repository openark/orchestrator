
$(document).ready(function () {
    showLoader();
    $.get(appUrl("/api/long-queries/"+currentFilter()), function (processes) {
            displayProcesses(processes);
            if (isAuthorizedForAction()) {
            	// Read-only users don't get auto-refresh. Sorry!
            	activateRefreshTimer();
            }
    	}, "json");
    function displayProcesses(processes) {
        hideLoader();
        processes.forEach(function (process) {
        	$("#long_queries .panel-body").append('<table class="table table-bordered table-condensed"><tbody></tbody></table>');    

        	process.Host = process.Host.split(':')[0];
        	
        	function addInfo(title, value) {
        		$("#long_queries tbody:last").append('<tr><td>'+title+'</td><td><code class="text-primary"><strong>'+value+'</strong></code></td></tr>')
        	}
        	
        	addInfo('Instance', '<a href="' + appUrl('/web/search?s='+process.InstanceHostname+':'+process.InstancePort) + '"><code class="text-primary">'+process.InstanceHostname+":"+process.InstancePort+'</code></a>');
        	addInfo('Process ID <siv class="pull-right"><button class="btn btn-xs btn-danger" data-command="kill_query" data-host="'+process.InstanceHostname+'" data-port="'+process.InstancePort+'" data-process="'+process.Id+'">Kill query</button></div>', process.Id);
        	addInfo('Started at', process.StartedAt);
        	addInfo('Runtime seconds', process.Time);
        	addInfo('User', process.User + ":" + process.Host);
        	$("#long_queries .panel-body").append('<div style="overflow:auto"><pre class="text-primary">'+process.Info.replace(/\n/g, '\n')+'</pre></div>')
        	/*
        	var rowDiv = jQuery('<div class="row col-xs-12"/>');
        	var infoDiv = jQuery('<div/>');
        	infoDiv.html(infoDiv.html() + '<div><a href="' + appUrl('/web/search?s='+process.InstanceHostname+":"+process.InstancePort) + '"><code class="text-primary">'+process.InstanceHostname+":"+process.InstancePort+'</code></a></div>');
        	infoDiv.html(infoDiv.html() + '<div>Id: '+process.Id+'</div>');
        	infoDiv.html(infoDiv.html() + '<div>Started: '+process.StartedAt+'</div>');
        	infoDiv.html(infoDiv.html() + '<div>Time: '+process.Time+'</div>');
        	infoDiv.html(infoDiv.html() + '<div>User: '+process.User + ":" + process.Host+'</div>');
        	infoDiv.appendTo(rowDiv)
        	rowDiv.append('<div class="text-primary"><pre>'+process.Info.replace(/\n/g, '<br/>')+'</pre></div>')
        	rowDiv.appendTo("#long_queries .panel-body")
        	*/
    	});
    }
    $("body").on("click", "button[data-command=kill_query]", function(event) {
    	var host = $(event.target).attr("data-host");
    	var port = $(event.target).attr("data-port");
    	var processId = $(event.target).attr("data-process");
    	var message = "Are you sure you wish to kill query on <code><strong>" + host +':'+ port + "</strong></code>?";
    	bootbox.confirm(message, function(confirm) {
			if (confirm) {
		    	showLoader();
		        $.get(appUrl("/api/kill-query/"+host + "/" + port + "/" + processId), function (operationResult) {
					hideLoader();
					if (operationResult.Code == "ERROR") {
						addAlert(operationResult.Message)
					} else {
						location.reload();
					}	
		        }, "json");
			}
        });
    });
});	
