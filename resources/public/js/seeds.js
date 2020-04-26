
$(document).ready(function () {
	showLoader();
	var defaultSeedsUrl = "/api/seeds";
	var currentPage = 0;
	var sourceAgentSeedMethods = [];
	var targetAgentSeedMethods = [];
	var sourceAgentCluster = "";
	var targetAgentCluster = "";

	var autocompleteSources = {
        targethostname: "/api/agents-hosts/",
        sourcehostname: "/api/agents-hosts/",
        status: "/api/seeds-statuses"
    }
    
    var autocompleteTerms = {
        targethostname: "hostname",
        sourcehostname: "hostname",
        status: "status"
    }

    var autocompleteMinLength = {
        targethostname: 4,
        sourcehostname: 4,
        status: 0
	}
	
    var $seeds_filter_autocomplete = $("#seeds_filter").autocomplete({
        source: function (request, response) {
            $.getJSON(appUrl("/api/agents-hosts/"), {
                hostname: request.term
            }, response);
		},
		minLength: 4,
		appendTo: "#seeds_filter_input",
		change: function (_, ui) {
			if(!ui.item){
				$("#seeds_filter").val("");
			} else {
				$("#seeds_filter").val(ui.item.label)
			}
		},
		select: function(event, ui) {
			event.stopPropagation();
		}
	});
	
    $('#seeds_filter_type').change(function() {
        $("#seeds_filter").val('');
        var src = $(this).find("option:selected").val();
        $seeds_filter_autocomplete.autocomplete('option', 'source', function (request, response) {
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
        $seeds_filter_autocomplete.autocomplete('option', 'minLength', autocompleteMinLength[src]);
    });

	$.get(appUrl(getUrl(currentPage)), function (seeds) {
    	displaySeeds(seeds);
    }, "json");

    function getUrl(page) {
        var filtervalue=$.trim($("#seeds_filter").val());
        apiUrl = defaultSeedsUrl + "?page="+page;
        if(filtervalue.length>0) {
         var filtertype = $("#seeds_filter_type").val();
         apiUrl = defaultSeedsUrl+"?"+filtertype+"="+filtervalue+ "&page="+page;
        };
        return apiUrl;
	}
	
	function displaySeeds(seeds) {
		hideLoader();
		$("#seeds .pager .next").removeClass("disabled");
		$("#seeds .pager .previous").removeClass("disabled");
		$("[data-agent=seeds_seed_details]").empty();
		var hasActive = false;
	    seeds.forEach(function (seed) {
	    	appendSeedDetails(seed, "[data-agent=seeds_seed_details]");
	    	if (!seed.IsComplete) {
	    		hasActive = true;
	    	}
		});
		if (hasActive) {
			activateRefreshTimer();
		}
		if (seeds.length == 0) {
			addAlert("No seeds found");
			$("#seeds .pager .next").addClass("disabled");
		}
		if (seeds.length < 20) {
			$("#seeds .pager .next").addClass("disabled");
		}
		if (currentPage <= 0) {
			$("#seeds .pager .previous").addClass("disabled");
		}
	}

	$("#seeds .pager .previous").click(function() {
		if ($("#seeds .pager .previous").hasClass("disabled") == true) {
			return false;
		} else {
			currentPage -= 1
			$.get(appUrl(getUrl(currentPage)), function (seeds) {
				$("#alerts_container").empty();
				displaySeeds(seeds);
			}, "json");
		}
	});

	$("#seeds .pager .next").click(function() {
		if ($("#seeds .pager .next").hasClass("disabled") == true) {
			return false;
		} else {
			currentPage += 1
			$.get(appUrl(getUrl(currentPage)), function (seeds) {
				$("#alerts_container").empty();
				displaySeeds(seeds);
			}, "json");
		}
	});
	
	$("#seeds_filter").keypress(function(event){
		var keycode = (event.keyCode ? event.keyCode : event.which);
		if(keycode == '13'){
			$("#seeds_filter_button").click();
		}
	});
	
	$("#seeds_filter_button").click(function(){
		currentPage = 0
		$.get(appUrl(getUrl(currentPage)), function (seeds) {
			$("#alerts_container").empty()
			displaySeeds(seeds);
		}, "json");
	});

	$("#seeds_new_seed").click(function(){
		if ( $("#target_agent_hostname").val().length == 0 && $("#source_agent_hostname").val().length == 0) {
			$("#seeds_seed_method").attr('disabled','disabled');
			$("#source_agent_hostname").val("");
			$("#target_agent_hostname").val("");
			$("#seeds_seed_method").empty();
		}
		
	});

	function enableSeedMethods(agent, side) {
		$("#seeds_seed_method").removeAttr('disabled');
		$.get(appUrl("/api/agent/" + agent), function(agent) {
			if (side == "source") {
				sourceAgentSeedMethods = Object.keys(agent.Data.AvailiableSeedMethods);
				sourceAgentCluster = agent.ClusterAlias;
				seedMethods = sourceAgentSeedMethods;
			} else {
				targetAgentSeedMethods = Object.keys(agent.Data.AvailiableSeedMethods);
				targetAgentCluster = agent.ClusterAlias;
				seedMethods = targetAgentSeedMethods;
			}
			if (sourceAgentSeedMethods.length ==0 || targetAgentSeedMethods.length == 0) {
				seedMethods.forEach(function(element) {
					entry += '<option value="' + element +'" selected="">'+element+'</option>';
				  });
				  $("#seeds_seed_method").append(entry);
			} else {
				intersection = sourceAgentSeedMethods.filter(element => targetAgentSeedMethods.includes(element));
				$("#seeds_seed_method").empty();
				var entry = "";
				intersection.forEach(function(element) {
					entry += '<option value="' + element +'" selected="">'+element+'</option>';
				  });
				  $("#seeds_seed_method").append(entry);
			}
		}, "json");
	}

	$("#source_agent_hostname").autocomplete({
		source: function (request, response) {
			$.getJSON(appUrl("/api/agents-hosts/"), {
				hostname: request.term
			}, response);
		},
		appendTo: "#source_agent_input",
		minLength: 4,
		change: function (_, ui) {
			if(!ui.item){
				$("#source_agent_hostname").val("");
			} else {
				$("#source_agent_hostname").val(ui.item.label)
			}
		},
		select: function(event, ui) {
			enableSeedMethods(ui.item.label, "source");
			event.stopPropagation();
		}
	});

	$("#target_agent_hostname").autocomplete({
		source: function (request, response) {
			$.getJSON(appUrl("/api/agents-hosts/"), {
				hostname: request.term
			}, response);
		},
		appendTo: "#target_agent_input",
		minLength: 4,
		change: function (_, ui) {
			if(!ui.item){
				$("#target_agent_hostname").val("");
			} else {
				$("#target_agent_hostname").val(ui.item.label)
			}
		},
		select: function(event, ui) {
			enableSeedMethods(ui.item.label, "target");
			event.stopPropagation();
		}
	});

	$("#seeds_start_seed").click(function() {
		var seed_method = $('#seeds_seed_method').val();
		var source_hostname = $('#source_agent_hostname').val();
		if (source_hostname.length == 0) {
			addAlert("Please specify source agent hostname for seed");
			return;
		  }
		var target_hostname = $('#target_agent_hostname').val();
		if (target_hostname.length == 0) {
			addAlert("Please specify target agent hostname for seed");
			return;
		  }
		var url = "/api/agent-seed/"+seed_method+"/"+target_hostname+"/"+source_hostname;
		var message = "Are you sure you wish to seed data from <code><strong>" + source_hostname + "</strong></code> to <code><strong>"+ target_hostname + "</strong></code> using <code><strong>"+ seed_method + "</strong></code> seed method? <br><code><strong>This will destroy all data on host "+target_hostname+"</strong></code>";
		if (targetAgentCluster != sourceAgentCluster) {
			message += "<br><br><code><strong>You are going to seed data to a host in a different cluster</strong></code><br>Source agent cluster:<code><strong>" +sourceAgentCluster+"</strong></code><br>Target agent cluster:<code><strong>"+targetAgentCluster+"</strong></code>"
		}
		bootbox.confirm(message, function(confirm) {
		  if (confirm) {
			showLoader();
			$.get(appUrl(url), function() {
			  hideLoader();
			  location.reload(true);
			}, "json").fail(function(operationResult) {
			  hideLoader();
			  if (operationResult.responseJSON.Code == "ERROR") {
				addAlert(operationResult.responseJSON.Message);
			  }
			});
		  }
		});
	  });

});	

