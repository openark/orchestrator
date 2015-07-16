
$(document).ready(function () {
    $('button[data-btn=discover-instance]').unbind("click");
    $('button[data-btn=discover-instance]').click(function() {

        if (!$("#discoverHostName").val()) {
            return addAlert("You must enter a host name");
        }
        if (!$("#discoverPort").val()) {
            return addAlert("You must enter a port");
        }
        discover($("#discoverHostName").val(), $("#discoverPort").val())
        return false;
    });
    $("#discoverHostName").focus();
});

function discover(hostname, port) {
    showLoader();
    var uri = "/api/discover/"+hostname+"/"+port;
    $.get(uri, function (operationResult) {
        hideLoader();
        if (operationResult.Code == "ERROR") {
            addAlert(operationResult.Message)
        } else {
            //location.reload();
            addInfo(""+hostname+":"+port
            		+ " sent for discovery. This should reflect in the topology listing or in topology instances."
            		+ ' <a href="/web/search?s='+hostname+":"+port+'" class="alert-link">Search</a> for this instance');
        }   
    }, "json"); 
	
}