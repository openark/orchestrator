
$(document).ready(function () {
    showLoader();

    var defaultAgentsUrl = "/api/agents"
    var currentPage = 0

    var autocompleteSources = {
        hostname: "/api/agents-hosts/",
        clusteralias: "/api/clusters-aliases/",
        status: "/api/agents-statuses/"
    }
    
    var autocompleteTerms = {
        hostname: "hostname",
        clusteralias: "clusteralias",
        status: "status"
    }

    var autocompleteMinLength = {
        hostname: 4,
        clusteralias: 4,
        status: 0
    }
    
    var $agents_filter_autocomplete = $("#agents_filter").autocomplete({
        source: function (request, response) {
            $.getJSON(appUrl("/api/agents-hosts/"), {
                hostname: request.term
            }, response);
        },
        minLength: 4,
		appendTo: "#agents_filter_input",
		change: function (_, ui) {
			if(!ui.item){
				$("#agents_filter").val("");
			} else {
				$("#agents_filter").val(ui.item.label)
			}
		},
		select: function(event, ui) {
			event.stopPropagation();
		}
    });

    $('#agents_filter_type').change(function() {
        $("#agents_filter").val('');
        var src = $(this).find("option:selected").val();
        $agents_filter_autocomplete.autocomplete('option', 'source', function (request, response) {
            var data = {}
            data[autocompleteTerms[src]] = request.term
            $.ajax({
                url: appUrl(autocompleteSources[src]),
                dataType: "json",
                cache: false,
                type: "get",
                data: data,
            }).done(function(data) {
                response(data);
            });
        });
        $agents_filter_autocomplete.autocomplete("option", "minLength",  autocompleteMinLength[src]);
    });

    $.get(appUrl(getUrl(currentPage)), function (agents) {
    	displayAgents(agents);
    }, "json");

    function getUrl(page) {
        var filtervalue=$.trim($("#agents_filter").val());
        apiUrl = defaultAgentsUrl + "?page="+page;
        if(filtervalue.length>0)
        {
         var filtertype = $('#agents_filter_type').val();
         apiUrl = defaultAgentsUrl+"?"+filtertype+"="+filtervalue + "&page="+page;
        };
        return apiUrl;
    }

    function displayAgents(agents) {
        hideLoader();
        $("#agents .pager .next").removeClass("disabled");
        $("#agents .pager .previous").removeClass("disabled");
        $("#agents_info").empty();
        
        agents.forEach(function (agent) {
            var row = '<tr>';
            row += '<td><a href="' + appUrl('/web/cluster/' + agent.ClusterAlias) + '">' + agent.ClusterAlias + '</a></td>';
            row += '<td><a href="' + appUrl('/web/agent/' + agent.Info.Hostname) + '">' + agent.Info.Hostname + '</a></td>';
            row += '<td>' + agent.Status + '</td>';
            row += '<td>' + moment.utc(agent.LastSeen).format("YYYY-MM-DD HH:mm:ss") + '</td>';
            row += '<td>' + agent.MySQLVersion + '</td>';
            row += '<td>' + Object.keys(agent.Data.MySQLDatabases).length + '</td>';
            row += '<td>' + agent.Data.OsName + '</td>';
            row += '<td>' + agent.Data.NumCPU + '</td>';
            row += '<td>' + toHumanFormat(agent.Data.MemTotal) + '</td>';
            row += '<td>' + toHumanFormat(agent.Data.BackupDirDiskFree) + '</td>';
            row += '<td>' + toHumanFormat(agent.Data.MySQLDatadirDiskUsed) + '</td>';
            row += '<td>' + toHumanFormat(agent.Data.MySQLDatadirDiskFree) + '</td>';
            row += '<td>' + Object.keys(agent.Data.AvailiableSeedMethods) + '</td>';
            row += '</tr>';
            $("#agents_info").append(row);
            hideLoader();
        });
        if (agents.length == 0) {
            addAlert("No agents found");
            $("#agents .pager .next").addClass("disabled");
        }
        if (agents.length < 20) {
            $("#agents .pager .next").addClass("disabled");
        }
        if (currentPage <= 0) {
            $("#agents .pager .previous").addClass("disabled");
        }
    }

    $("#agents .pager .previous").click(function() {
        if ($("#agents .pager .previous").hasClass("disabled") == true) {
            return false;
        } else {
            currentPage -= 1;
            $.get(appUrl(getUrl(currentPage)), function (agents) {
                $("#alerts_container").empty();
                displayAgents(agents);
            }, "json");
        }
    });
    $("#agents .pager .next").click(function() {
        if ($("#agents .pager .next").hasClass("disabled") == true) {
            return false;
        } else {
            currentPage += 1;
            $.get(getUrl(currentPage), function (agents) {
                $("#alerts_container").empty();
                displayAgents(agents);
            }, "json");
        }
    });
    
    $("#agents_filter").keypress(function(event){
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if(keycode == '13'){
            $("#agent_filter_button").click();
        }
    });

    $("#agent_filter_button").click(function(){
        currentPage = 0
        $.get(appUrl(getUrl(currentPage)), function (agents) {
            $("#alerts_container").empty()
            displayAgents(agents);
        }, "json");
    });

});	

