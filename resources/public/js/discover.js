
$(document).ready(function () {
    $('button[data-btn=discover-instance]').unbind("click");
    $('button[data-btn=discover-instance]').click(function() {

        if (!$("#discoverHostName").val()) {
            return addAlert("You must enter a host name");
        }
        if (!$("#discoverPort").val()) {
            return addAlert("You must enter a port");
        }
        showLoader();
        var uri = "/api/discover/"+$("#discoverHostName").val()+"/"+$("#discoverPort").val();
        $.get(uri, function (operationResult) {
            hideLoader();
            if (operationResult.Code == "ERROR") {
                addAlert(operationResult.Message)
            } else {
                //location.reload();
                addInfo(""+$("#discoverHostName").val()+":"+$("#discoverPort").val() 
                		+ " sent for discovery. This should reflect in the topology listing or in topology instances."
                		+ ' <a href="/web/search?s='+$("#discoverHostName").val()+":"+$("#discoverPort").val()+'" class="alert-link">Search</a> for this instance');
            }   
        }, "json"); 
        return false;
    });
});

