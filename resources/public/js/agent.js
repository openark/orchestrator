$(document).ready(function() {
  showLoader();

  var hasActiveSeeds = false;
  var currentAgentSeedMethods = [];
  var targetAgentCluster = "";
  var sourceAgentCluster = "";

  $.get(appUrl("/api/agent/" + currentAgentHost()), function(agent) {
    showLoader();
    agent.AvailableLocalSnapshots || (agent.AvailableLocalSnapshots = []);
    agent.AvailableSnapshots || (agent.AvailableSnapshots = []);
    displayAgent(agent);
  }, "json").fail(function(operationResult) {
    hideLoader();
      if (operationResult.responseJSON.Code == "ERROR") {
        addAlert(operationResult.responseJSON.Message)
        return
      }
  });

  $.get(appUrl("/api/agent-active-seeds/" + currentAgentHost()), function(activeSeeds) {
    showLoader();
    activeSeeds.forEach(function(activeSeed) {
      appendSeedDetails(activeSeed, "#agent_active_seeds");
    });
    if (activeSeeds.length == 0) {
      $("#agent_active_seeds_container").parent().hide();
      $("#agent_active_seeds_states_container").parent().hide();
    }
    if (activeSeeds.length > 0) {
      hasActiveSeeds = true;
      activateRefreshTimer();

      $.get(appUrl("/api/agent-seed-states/" + activeSeeds[0].SeedID), function(seedStates) {
        showLoader();
        seedStates.forEach(function(seedState) {
          appendSeedState(seedState, "#seed_states");
        });
      }, "json");
    }
  }, "json");

  $.get(appUrl("/api/agent-recent-seeds/" + currentAgentHost()), function(recentSeeds) {
    showLoader();
    recentSeeds.forEach(function(recentSeed) {
      appendSeedDetails(recentSeed, "#agent_recent_seeds");
    });
    if (recentSeeds.length == 0) {
      $("#agent_recent_seeds_container").parent().hide();
    }
  }, "json");

  function displayAgent(agent) {
    currentAgentSeedMethods = Object.keys(agent.Data.AvailiableSeedMethods);
    if (!(Object.keys(agent.Data.AvailiableSeedMethods).includes("LVM"))) {
      $("#agent_snapshots_info_container").hide();
    }
    if (Object.keys(agent.Data.MySQLDatabases).length == 0) {
      $("#agent_databases_container").hide();
    }
    $("#agent_hostname").html(agent.Info.Hostname);
    $("#agent_hostname").html(agent.Info.Hostname + '<div class="pull-right btn-group"><button class="btn btn-xs btn-success" id="agent_discover_button" data-hostname="' + agent.Info.Hostname + '" data-mysql-port="' + agent.Info.MySQLPort + '">Discover</button><button class="btn btn-xs btn-info" id="agent_refresh_button" data-hostname="' + agent.Info.Hostname + '">Refresh</button></div>'
    );
    $("#agent_port").html(agent.Info.Port);
    $("#agent_cluster_name").html('<a href="' + appUrl('/web/cluster/' + agent.ClusterAlias + '">' + agent.ClusterAlias + '</a>'))
    sourceAgentCluster = agent.ClusterAlias
    $("#agent_status").html(agent.Status)
    $("#agent_last_seen").html(moment.utc(agent.LastSeen).format("YYYY-MM-DD HH:mm:ss"))
    $("#agent_os").html(agent.Data.OsName);
    $("#agent_cpu").html(agent.Data.NumCPU);
    $("#agent_ram").html(toHumanFormat(agent.Data.MemTotal));

    $("#agent_avaliable_seed_methods").html(Object.keys(agent.Data.AvailiableSeedMethods).join())
    if (agent.Data.AgentCommands !== null) {
      $("#agent_commands").html('<div class="input-group" id="agent_commands_input"><select class="form-control" id="agent_agent_commands" style="height: 25px"></select><span class="input-group-addon agent-filter-empty-space"></span><span class="input-group-btn"><button class="btn btn-xs btn-info" id="agent_execute_command_button" data-hostname="' + agent.Info.Hostname + '">Execute</button></span></div>');
      agent.Data.AgentCommands.forEach(function(command) {
        $('<option/>', { value : command }).text(command).appendTo('#agent_agent_commands');
      });
    } else {
      $("#agent_commands_container").remove();
    }
    $("#agent_backup_dir").html(agent.Data.BackupDir)
    $("#agent_backup_dir_used").html(toHumanFormat(agent.Data.BackupDirDiskUsed))
    $("#agent_backup_dir_free").html(toHumanFormat(agent.Data.BackupDirDiskFree))
    $("#agent_mysql_version").html(agent.MySQLVersion)
    var mySQLStatus = "" + agent.Data.MySQLRunning + '<div class="pull-right">' +
      (agent.Data.MySQLRunning ? '<button class="btn btn-xs btn-danger" id="agent_mysql_stop_button">Stop</button>' :
        '<button class="btn btn-xs btn-success" id="agent_mysql_start_button">Start</button>') +
      '</div>';
    $("#agent_mysql_running").html(mySQLStatus)
    $("#agent_mysql_port").html(agent.Info.MySQLPort)
    $("#agent_mysql_datadir").html(agent.Data.MySQLDatadir)
    $("#agent_mysql_disk_usage").html(toHumanFormat(agent.Data.MySQLDatadirDiskUsed))
    $("#agent_mysql_disk_free").html(toHumanFormat(agent.Data.MySQLDatadirDiskFree))
    $("#agent_error_log_tail").html(
      '<div class="pull-right btn-group"><button class="btn btn-xs btn-success" id="agent_error_log_tail_button" data-hostname="' + agent.Info.Hostname + '">View</button></div>'
    );

    $("body").on("click", "#agent_execute_command_button", function(event) {
      var hostname = $(event.target).attr("data-hostname")
      var command = $("#agent_agent_commands").val()
      var message = "Are you sure you wish to run <code><strong>" +
      command + "</strong></code> command on <code><strong>" + hostname + "</strong></code>?";
      bootbox.confirm(message, function(confirm) {
        if (confirm) {
          showLoader();
          $.get(appUrl("/api/agent-custom-command/"+hostname+"/"+command), function() {
            hideLoader();
            $("#agent_refresh_button").click();
          }, "json").fail(function(operationResult) {
            hideLoader();
            if (operationResult.responseJSON.Code == "ERROR") {
              addAlert(operationResult.responseJSON.Message);
            }
          });
        }
      });
    });

    $("body").on("click", "#agent_discover_button", function(event) {
      var hostname = $(event.target).attr("data-hostname")
      var mySQLPort = $(event.target).attr("data-mysql-port")
      discover(hostname, mySQLPort)
    });

    $("body").on("click", "#agent_refresh_button", function(event) {
      var hostname = $(event.target).attr("data-hostname");
      showLoader();
      $.get(appUrl("/api/agent-update/"+hostname), function() {
        hideLoader();
        location.reload(true);
      }, "json").fail(function(operationResult) {
        hideLoader();
        if (operationResult.responseJSON.Code == "ERROR") {
          addAlert(operationResult.responseJSON.Message);
        }
      });
    });

    $("body").on("click", "#agent_error_log_tail_button", function(event) {
      var hostname = $(event.target).attr("data-hostname");
      showLoader();
      $.get(appUrl("/api/agent-mysql-error-log/"+hostname), function(rows) {
        hideLoader();
        rows = rows.slice(1,-1).split("\\n").slice(0,-1);
        rows = rows.map(function(row) {
          return '<strong>' + row + '</strong>';
        });
        rows = rows.map(function(row) {
          if (row.indexOf("[ERROR]") >= 0) {
            return '<code class="text-danger">' + row + '</code>';
          } else if (row.indexOf("[Warning]") >= 0) {
            return '<code class="text-warning">' + row + '</code>';
          } else if (row.indexOf("[Note]") >= 0) {
            return '<code class="text-info">' + row + '</code>';
          } else {
            return '<code class="text-primary">' + row + '</code>';
          }
        });
        bootbox.alert('<div style="overflow: auto">' + rows.join("<br/>") + '</div>').find("div.modal-dialog").addClass("error-log-large-width");
        return false;
      }, "json").fail(function(operationResult) {
        hideLoader();
        if (operationResult.responseJSON.Code == "ERROR") {
          addAlert(operationResult.responseJSON.Message);
        }
      });
    });

    Object.keys(agent.Data.AvailiableSeedMethods).forEach(function(key) {
      $('<option/>', { value : key }).text(key).appendTo('#agent_seed_seed_method');
    });

    $("body").on("click", "#agent_start_seed_button", function() {
      var seed_method = $('#agent_seed_seed_method').val();
      var seed_hostname = $('#agent_seed_hostname').val();
      if (seed_hostname.length == 0) {
        addAlert("Please specify hostname for seed");
        return;
      }
      var seedSide = $('#agent_seed_side').val();
      if (seedSide == "Seed from") {
        url = "/api/agent-seed/"+seed_method+"/"+currentAgentHost()+"/"+seed_hostname;
        var message = "Are you sure you wish to seed data from <code><strong>" + seed_hostname + "</strong></code> to <code><strong>"+ currentAgentHost() + "</strong></code> using <code><strong>"+ seed_method + "</strong></code> seed method? <code><strong>This will destroy all data on host "+currentAgentHost()+"</strong></code>";
      } else {
        url = "/api/agent-seed/"+seed_method+"/"+seed_hostname+"/"+currentAgentHost()
        var message = "Are you sure you wish to seed data from <code><strong>" + currentAgentHost() + "</strong></code> to <code><strong>"+ seed_hostname + "</strong></code> using <code><strong>"+ seed_method + "</strong></code> seed method? <code><strong>This will destroy all data on host "+seed_hostname+"</strong></code>";
      }
      if (targetAgentCluster != sourceAgentCluster) {
        message += "<br><br><code><strong>You are going to seed data to a host in a different cluster</strong></code><br>Source agent cluster:<code><strong>" +sourceAgentCluster+"</strong></code><br>Target agent cluster:<code><strong>"+targetAgentCluster+"</strong></code>"
      }
      bootbox.confirm(message, function(confirm) {
        if (confirm) {
          showLoader();
          $.get(appUrl(url), function() {
            hideLoader();
            $("#agent_refresh_button").click();
          }, "json").fail(function(operationResult) {
            hideLoader();
            if (operationResult.responseJSON.Code == "ERROR") {
              addAlert(operationResult.responseJSON.Message);
            }
          });
        }
      });
    });

    if (agent.Data.LocalSnapshotsHosts.length > 0) {
      local_snaphost_entry = '<tr><td>Available locally</td><td>';
      local_snaphost_entry += beautifyAvailableSnapshots(agent.Data.LocalSnapshotsHosts);
      local_snaphost_entry += '</td></tr>';
      $("#agent_snaphots_data").append(local_snaphost_entry);
    } 
    if (agent.Data.SnaphostHosts.length > 0) {
      remote_snaphost_entry = '<tr><td>Available remote</td><td>';
      availableRemoteSnapshots = agent.Data.SnaphostHosts.filter(function(snapshot) {
        return agent.Data.LocalSnapshotsHosts.indexOf(snapshot) < 0;
      });
      remote_snaphost_entry += beautifyAvailableSnapshots(availableRemoteSnapshots);
      remote_snaphost_entry += '</td></tr>';
      $('#agent_snaphots_data').append(remote_snaphost_entry);
    }
    snaphost_entry = '<tr><td>Snapshots taken on this host</td><td><button id="agent_create_snapshot_button" class="btn btn-xs btn-info">Create snapshot</button></td></tr>';
    $('#agent_snaphots_data').append(snaphost_entry);

    if (agent.Data.LogicalVolumes) {
      var lvSnapshots = agent.Data.LogicalVolumes.filter(function(logicalVolume) {
        return logicalVolume.IsSnapshot;
      }).map(function(logicalVolume) {
        return logicalVolume.Path;
      });
      var mountedVolume = ""
      if (agent.Data.MountPoint) {
        if (agent.Data.MountPoint.IsMounted) {
          mountedVolume = agent.Data.MountPoint.LVPath;
        }
      }

      var result = lvSnapshots.map(function(volume) {
        var volumeText = '';
        var volumeTextType = 'text-info';
        if (volume == mountedVolume) {
          volumeText = '<button class="btn btn-xs btn-danger" id="agent_unmount_button">Unmount</button>';
          volumeTextType = 'text-success';
        } else if (!(agent.Data.MountPoint && agent.Data.MountPoint.IsMounted)) {
          volumeText += '<button class="btn btn-xs btn-success" id="agent_mountlv_button" data-lv="' + volume + '">Mount</button>';
          volumeText += '<button class="btn btn-xs btn-danger" id="agent_removelv_button" data-lv="' + volume + '">Remove</button>';
        } else if (agent.Data.MountPoint.IsMounted) {
          // Do nothing
          volumeText += '<button class="btn btn-xs btn-danger" id="agent_removelv_button" data-lv="' + volume + '">Remove</button>';
        }
        volumeText = '<tr><td style="border-top:0px"><code class="' + volumeTextType + '"><strong>' + volume + '</strong></code></td><td style="border-top:0px"><div class="pull-left btn-group">'+ volumeText + '</div></td></tr>';
        return volumeText;
      });

      result.forEach(function(entry) {
        $("#agent_snaphots_data").append(entry);
      });
    }

    if (Object.keys(agent.Data.MySQLDatabases).length > 0) {
      entry = "";
      Object.keys(agent.Data.MySQLDatabases).forEach(function(key) {
        entry += '<tr><td>'+key+'</td><td>'+toHumanFormat(agent.Data.MySQLDatabases[key].Size)+'</td></tr>';
      });
      $("#agent_databases").append(entry);
    }

    hideLoader();
  }

  $("body").on("click", "#agent_unmount_button", function() {
    if (hasActiveSeeds) {
      addAlert("This agent participates in an active seed; please await or abort active seed before unmounting");
      return;
    }
    $.get(appUrl("/api/agent-umount/" + currentAgentHost()), function() {
      hideLoader();
      $("#agent_refresh_button").click();
    }, "json").fail(function(operationResult) {
      hideLoader();
      if (operationResult.responseJSON.Code == "ERROR") {
        addAlert(operationResult.responseJSON.Message)
      }
    });
  });

  $("body").on("click", "#agent_mountlv_button", function(event) {
    var lv = $(event.target).attr("data-lv")
    showLoader();
    $.get(appUrl("/api/agent-mount/" + currentAgentHost() + "?lv=" + encodeURIComponent(lv)), function() {
      hideLoader();
      $("#agent_refresh_button").click();
    }, "json").fail(function(operationResult) {
      hideLoader();
      if (operationResult.responseJSON.Code == "ERROR") {
        addAlert(operationResult.responseJSON.Message)
      }
    });
  });

  $("body").on("click", "#agent_removelv_button", function(event) {
    var lv = $(event.target).attr("data-lv")
    var message = "Are you sure you wish to remove logical volume <code><strong>" + lv + "</strong></code>?";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        showLoader();
        $.get(appUrl("/api/agent-removelv/" + currentAgentHost() + "?lv=" + encodeURIComponent(lv)), function() {
          hideLoader();
          $("#agent_refresh_button").click();
        }, "json").fail(function(operationResult) {
          hideLoader();
          if (operationResult.responseJSON.Code == "ERROR") {
            addAlert(operationResult.responseJSON.Message)
          }
        });
      }
    });
  });

  $("body").on("click", "#agent_create_snapshot_button", function() {
    var message = "Are you sure you wish to create a new snapshot on <code><strong>" +
      currentAgentHost() + "</strong></code>?";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        showLoader();
        $.get(appUrl("/api/agent-create-snapshot/" + currentAgentHost()), function() {
          hideLoader();
          $("#agent_refresh_button").click();
        }, "json").fail(function(operationResult) {
          hideLoader();
          if (operationResult.responseJSON.Code == "ERROR") {
            addAlert(operationResult.responseJSON.Message);
          }
        });
      }
    });
  });


  $("body").on("click", "#agent_mysql_stop_button", function() {
    var message = "Are you sure you wish to shut down MySQL service on <code><strong>" +
      currentAgentHost() + "</strong></code>?";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        showLoader();
        $.get(appUrl("/api/agent-mysql-stop/" + currentAgentHost()), function() {
          hideLoader();
          $("#agent_refresh_button").click();
        }, "json").fail(function(operationResult) {
          hideLoader();
          if (operationResult.responseJSON.Code == "ERROR") {
            addAlert(operationResult.responseJSON.Message);
          }
        });
      }
    });
  });

  $("body").on("click", "#agent_mysql_start_button", function() {
    showLoader();
    $.get(appUrl("/api/agent-mysql-start/" + currentAgentHost()), function() {
      hideLoader();
      $("#agent_refresh_button").click();
    }, "json").fail(function(operationResult) {
      hideLoader();
      if (operationResult.responseJSON.Code == "ERROR") {
        addAlert(operationResult.responseJSON.Message);
      }
    });
  });

	function updateSeedMethods(agent) {
		$.get(appUrl("/api/agent/" + agent), function(agent) {
      seedMethods = Object.keys(agent.Data.AvailiableSeedMethods);
      targetAgentCluster = agent.ClusterAlias;
			intersection = currentAgentSeedMethods.filter(element => seedMethods.includes(element));
			$("#agent_seed_seed_method").empty();
			var entry = "";
			intersection.forEach(function(element) {
				entry += '<option value="' + element +'" selected="">'+element+'</option>';
			  });
			  $("#agent_seed_seed_method").append(entry);
			
		}, "json");
	}

	$("#agent_seed_hostname").autocomplete({
    source: function (request, response) {
			$.getJSON(appUrl("/api/agents-hosts/"), {
				hostname: request.term
			}, response);
		},
		appendTo: "#agent_seed_hostname_input",
		minLength: 4,
		change: function (_,ui) {
			if(!ui.item){
				$("#agent_seed_hostname").val("");
			} else {
				$("#agent_seed_hostname").val(ui.item.label)
			}
		},
		select: function(event, ui) {
			updateSeedMethods(ui.item.label);
			event.stopPropagation();
		}
  });
  
});

function fillNewSeed(hostname) {
  $('#agent_seed_hostname').val(hostname);
  $('#agent_seed_seed_method').val("LVM").change();
  $('#agent_seed_side').val("Seed from").change();
  window.scrollTo(0, 0);
}

function beautifyAvailableSnapshots(hostnames) {
  var result = hostnames.filter(function(hostname) {
    return hostname.trim() != "";
  });
  result = result.map(function(hostname) {
    if (hostname == agent.Hostname) {
      return '<p>' + hostname + '</p>';
    }
    return '<p><a href="#" onclick="fillNewSeed(\''+hostname+'\');" data-hostname="'+ hostname + '">' + hostname + '</a></p>';
  });
  return result.join("");
}