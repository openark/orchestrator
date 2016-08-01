var refreshIntervalSeconds = 60; // seconds
var secondsTillRefresh = refreshIntervalSeconds;
var nodeModalVisible = false;

reloadPageHint = {
  hint: "",
  hostname: "",
  port: ""
}

var errorMapping = {
  "inMaintenanceProblem": {
    "badge": "label-info",
    "description": "In maintenance"
  },
  "lastCheckInvalidProblem": {
    "badge": "label-fatal",
    "description": "Last check invalid"
  },
  "notRecentlyCheckedProblem": {
    "badge": "label-stale",
    "description": "Not recently checked (stale)"
  },
  "notReplicatingProblem": {
    "badge": "label-danger",
    "description": "Not replicating"
  },
  "replicationLagProblem": {
    "badge": "label-warning",
    "description": "Replication lag"
  }
};

function updateCountdownDisplay() {
  if ($.cookie("auto-refresh") == "true") {
    $("#refreshCountdown").html('<span class="glyphicon glyphicon-repeat" title="Click to pause"></span> ' + secondsTillRefresh + 's');
  } else {
    secondsTillRefresh = refreshIntervalSeconds;
    $("#refreshCountdown").html('<span class="glyphicon glyphicon-pause" title="Click to countdown"></span> ' + secondsTillRefresh + 's');
  }
}

function startRefreshTimer() {
  var refreshFunction = function() {
    if (nodeModalVisible) {
      return;
    }
    secondsTillRefresh = Math.max(secondsTillRefresh - 1, 0);
    if (secondsTillRefresh <= 0) {
      $(".navbar-nav li[data-nav-page=refreshCountdown]").addClass("active");
      showLoader();
      location.reload(true);
    }
    updateCountdownDisplay();
  }
  setInterval(refreshFunction, 1 * 1000);
}

function resetRefreshTimer() {
  secondsTillRefresh = refreshIntervalSeconds;
}

function activateRefreshTimer() {
  startRefreshTimer();
  $(document).click(function() {
    resetRefreshTimer();
  });
  $(document).mousemove(function() {
    resetRefreshTimer();
  });
}

function showLoader() {
  $(".ajaxLoader").css('visibility', 'visible');
}

function hideLoader() {
  $(".ajaxLoader").css('visibility', 'hidden');
}

function appUrl(url) {
  // Create an absolute URL that respects URL-rewriting proxies,
  // such as the Kubernetes apiserver proxy.
  var here = window.location.pathname;
  var pos = here.lastIndexOf('/web/');
  if (pos < 0) {
    return url;
  }
  // This assumes the Orchestrator UI is accessed as ".../web/...".
  // The part before the "/web/" should be prefixed onto any URLs we write.
  return here.substr(0, pos) + url;
}

function visualizeBrand() {
  var img = $("<img>");

  img.attr("src", appUrl("/images/octocat-logo-32.png")).attr("alt", "GitHub");

  if (document.domain && document.domain.indexOf("outbrain.com") >= 0) {
    img.attr("src", appUrl("/images/outbrain-logo-32.png")).attr("alt", "Outbrain");
  }
  if (document.domain && document.domain.indexOf("booking.com") >= 0) {
    img.attr("src", appUrl("/images/booking-logo-32.png")).attr("alt", "Booking.com");
  }
  $(".orchestrator-brand").prepend(img)
}

function showContextMenu() {
  $("[data-nav-page=context]").css('visibility', 'visible');
}

function booleanString(b) {
  return (b ? "true" : "false");
}

function toHumanFormat(bytes) {
  var s = ['bytes', 'kB', 'MB', 'GB', 'TB', 'PB'];
  var e = Math.floor(Math.log(bytes) / Math.log(1024));
  return (bytes / Math.pow(1024, e)).toFixed(2) + " " + s[e];
}

function getInstanceId(host, port) {
  return "instance__" + host.replace(/[.]/g, "_") + "__" + port
}


function canonizeInstanceTitle(title) {
  if (typeof removeTextFromHostnameDisplay != "undefined" && removeTextFromHostnameDisplay()) {
    return title.replace(removeTextFromHostnameDisplay(), '');
  }
  return title;
}

function getInstanceTitle(host, port) {
  return canonizeInstanceTitle(host + ":" + port);
}



function commonSuffixLength(strings) {
  if (strings.length == 0) {
    return 0;
  }
  if (strings.length == 1) {
    return 0;
  }
  var longestSuffixLength = 0;
  var maxLength = 0;
  strings.forEach(function(s) {
    maxLength = ((maxLength == 0) ? s.length : Math
      .min(maxLength, s.length));
  });
  var suffixLength = 0;
  while (suffixLength < maxLength) {
    suffixLength++
    var suffixes = strings.map(function(s) {
      return s.substring(s.length - suffixLength)
    });
    var uniqueSuffixes = suffixes.filter(function(elem, pos) {
      return suffixes.indexOf(elem) == pos;
    })
    if (uniqueSuffixes.length > 1) {
      // lost it. keep last longestSuffixLength value
      break;
    }
    // we're still good
    longestSuffixLength = suffixLength;
  }
  return longestSuffixLength;
}


function addAlert(alertText, alertClass) {
  if ($.cookie("anonymize") == "true") {
    return false;
  }
  if (typeof(alertClass) === 'undefined') {
    alertClass = "danger";
  }
  $("#alerts_container").append(
    '<div class="alert alert-' + alertClass + ' alert-dismissable">' + '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>' + alertText + '</div>');
  $(".alert").alert();
  return false;
}


function addInfo(alertText) {
  return addAlert(alertText, "info");
}

function apiCommand(uri, hint) {
  showLoader();
  $.get(appUrl(uri), function(operationResult) {
    hideLoader();
    if (operationResult.Code == "ERROR") {
      addAlert(operationResult.Message)
    } else {
      reloadWithOperationResult(operationResult, hint);
    }
  }, "json");
  return false;
}


function reloadWithMessage(msg, details, hint) {
  var hostname = "";
  var port = "";
  if (details) {
    hostname = details.Hostname || hostname
    port = details.Port || port
  }
  hint = hint || "";
  var newUri = window.location.href.split("#")[0].split("?")[0] + "?orchestrator-msg=" + encodeURIComponent(msg) + "&hostname=" + hostname + "&port=" + port + "&hint=" + hint;
  if (isCompactDisplay && isCompactDisplay()) {
    newUri += "&compact=true";
  }
  window.location.href = newUri;
}

function reloadWithOperationResult(operationResult, hint) {
  var msg = operationResult.Message;
  reloadWithMessage(msg, operationResult.Details, hint);
}


// Modal

function addNodeModalDataAttribute(name, value) {
  var codeClass = "text-primary";
  if (value == "true" || value == true) {
    codeClass = "text-success";
  }
  if (value == "false" || value === false) {
    codeClass = "text-danger";
  }
  $('#modalDataAttributesTable').append(
    '<tr><td>' + name + '</td><td><code class="' + codeClass + '"><strong>' + value + '</strong></code><div class="pull-right attributes-buttons"></div></td></tr>');
  return $('#modalDataAttributesTable tr:last td:last');
}

function addModalAlert(alertText) {
  $("#node_modal .modal-body").append(
    '<div class="alert alert-danger alert-dismissable">' + '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>' + alertText + '</div>');
  $(".alert").alert();
  return false;
}

function openNodeModal(node) {
  if (!node) {
    return false;
  }
  if (node.isAggregate) {
    return false;
  }
  nodeModalVisible = true;
  var hiddenZone = $('#node_modal .hidden-zone');
  $('#node_modal #modalDataAttributesTable button[data-btn][data-grouped!=true]').appendTo("#node_modal .modal-footer");
  $('#node_modal #modalDataAttributesTable [data-btn-group]').appendTo("#node_modal .modal-footer");

  $('#node_modal .modal-title').html('<code class="text-primary">' + node.title + "</code>");

  $('#modalDataAttributesTable').html("");

  if (node.InstanceAlias) {
    addNodeModalDataAttribute("Instance Alias", node.InstanceAlias);
  }
  addNodeModalDataAttribute("Last seen", node.LastSeenTimestamp + " (" + node.SecondsSinceLastSeen.Int64 + "s ago)");
  if (node.UnresolvedHostname) {
    addNodeModalDataAttribute("Unresolved hostname", node.UnresolvedHostname);
  }
  $('#node_modal [data-btn-group=move-equivalent]').appendTo(hiddenZone);
  if (node.MasterKey.Hostname) {
    var td = addNodeModalDataAttribute("Master", node.masterTitle);
    if (node.IsDetachedMaster) {
      $('#node_modal button[data-btn=reattach-slave-master-host]').appendTo(td.find("div"))
    }
    $('#node_modal button[data-btn=reset-slave]').appendTo(td.find("div"))

    td = addNodeModalDataAttribute("Replication running", booleanString(node.replicationRunning));
    $('#node_modal button[data-btn=start-slave]').appendTo(td.find("div"))
    $('#node_modal button[data-btn=restart-slave]').appendTo(td.find("div"))
    $('#node_modal [data-btn-group=stop-slave]').appendTo(td.find("div"))

    if (!node.replicationRunning) {
      td = addNodeModalDataAttribute("Last SQL error", node.LastSQLError);
      $('#node_modal button[data-btn=skip-query]').appendTo(td.find("div"))
      addNodeModalDataAttribute("Last IO error", node.LastIOError);
    }
    addNodeModalDataAttribute("Seconds behind master", node.SecondsBehindMaster.Valid ? node.SecondsBehindMaster.Int64 : "null");
    addNodeModalDataAttribute("Replication lag", node.SlaveLagSeconds.Valid ? node.SlaveLagSeconds.Int64 : "null");
    addNodeModalDataAttribute("SQL delay", node.SQLDelay);

    var masterCoordinatesEl = addNodeModalDataAttribute("Master coordinates", node.ExecBinlogCoordinates.LogFile + ":" + node.ExecBinlogCoordinates.LogPos);
    $('#node_modal [data-btn-group=move-equivalent] ul').empty();
    $.get(appUrl("/api/master-equivalent/") + node.MasterKey.Hostname + "/" + node.MasterKey.Port + "/" + node.ExecBinlogCoordinates.LogFile + "/" + node.ExecBinlogCoordinates.LogPos, function(equivalenceResult) {
      if (!equivalenceResult.Details) {
        return false;
      }
      equivalenceResult.Details.forEach(function(equivalence) {
        if (equivalence.Key.Hostname == node.Key.Hostname && equivalence.Key.Port == node.Key.Port) {
          // This very instance; will not move below itself
          return;
        }
        var title = canonizeInstanceTitle(equivalence.Key.Hostname + ':' + equivalence.Key.Port);
        $('#node_modal [data-btn-group=move-equivalent] ul').append('<li><a href="#" data-btn="move-equivalent" data-hostname="' + equivalence.Key.Hostname + '" data-port="' + equivalence.Key.Port + '">' + title + '</a></li>');
      });

      if ($('#node_modal [data-btn-group=move-equivalent] ul li').length) {
        $('#node_modal [data-btn-group=move-equivalent]').appendTo(masterCoordinatesEl.find("div"));
      }
    }, "json");
    if (node.IsDetached) {
      $('#node_modal button[data-btn=detach-slave]').appendTo(hiddenZone)
      $('#node_modal button[data-btn=reattach-slave]').appendTo(masterCoordinatesEl.find("div"))
    } else {
      $('#node_modal button[data-btn=detach-slave]').appendTo(masterCoordinatesEl.find("div"))
      $('#node_modal button[data-btn=reattach-slave]').appendTo(hiddenZone)
    }
  } else {
    $('#node_modal button[data-btn=reset-slave]').appendTo(hiddenZone);
    $('#node_modal button[data-btn=reattach-slave-master-host]').appendTo(hiddenZone);
    $('#node_modal button[data-btn=skip-query]').appendTo(hiddenZone);
    $('#node_modal button[data-btn=detach-slave]').appendTo(hiddenZone)
    $('#node_modal button[data-btn=reattach-slave]').appendTo(hiddenZone)
  }
  if (node.LogBinEnabled) {
    addNodeModalDataAttribute("Self coordinates", node.SelfBinlogCoordinates.LogFile + ":" + node.SelfBinlogCoordinates.LogPos);
  }
  var td = addNodeModalDataAttribute("Num slaves", node.SlaveHosts.length);
  $('#node_modal button[data-btn=regroup-slaves]').appendTo(td.find("div"))
  addNodeModalDataAttribute("Server ID", node.ServerID);
  if (node.ServerUUID) {
    addNodeModalDataAttribute("Server UUID", node.ServerUUID);
  }
  addNodeModalDataAttribute("Version", node.Version);
  var td = addNodeModalDataAttribute("Read only", booleanString(node.ReadOnly));
  $('#node_modal button[data-btn=set-read-only]').appendTo(td.find("div"))
  $('#node_modal button[data-btn=set-writeable]').appendTo(td.find("div"))

  addNodeModalDataAttribute("Has binary logs", booleanString(node.LogBinEnabled));
  if (node.LogBinEnabled) {
    addNodeModalDataAttribute("Binlog format", node.Binlog_format);
    var td = addNodeModalDataAttribute("Logs slave updates", booleanString(node.LogSlaveUpdatesEnabled));
    $('#node_modal button[data-btn=enslave-siblings]').appendTo(td.find("div"))
  }

  var td = addNodeModalDataAttribute("GTID based replication", booleanString(node.usingGTID));
  $('#node_modal button[data-btn=enable-gtid]').appendTo(td.find("div"))
  $('#node_modal button[data-btn=disable-gtid]').appendTo(td.find("div"))
  if (node.usingGTID) {
    addNodeModalDataAttribute("Executed GTID set", node.ExecutedGtidSet);
    addNodeModalDataAttribute("GTID purged", node.GtidPurged);
  }

  addNodeModalDataAttribute("Semi-sync enforced", booleanString(node.SemiSyncEnforced));

  addNodeModalDataAttribute("Uptime", node.Uptime);
  addNodeModalDataAttribute("Allow TLS", node.AllowTLS);
  addNodeModalDataAttribute("Cluster",
    '<a href="' + appUrl('/web/cluster/' + node.ClusterName) + '">' + node.ClusterName + '</a>');
  addNodeModalDataAttribute("Audit",
    '<a href="' + appUrl('/web/audit/instance/' + node.Key.Hostname + '/' + node.Key.Port) + '">' + node.title + '</a>');
  addNodeModalDataAttribute("Agent",
    '<a href="' + appUrl('/web/agent/' + node.Key.Hostname) + '">' + node.Key.Hostname + '</a>');
  addNodeModalDataAttribute("Long queries",
    '<a href="' + appUrl('/web/long-queries?filter=' + node.Key.Hostname) + '">on ' + node.Key.Hostname + '</a>');

  $('#node_modal [data-btn]').unbind("click");

  $("#beginDowntimeOwner").val(getUserId());
  $('#node_modal button[data-btn=begin-downtime]').click(function() {
    if (!$("#beginDowntimeOwner").val()) {
      return addModalAlert("You must fill the owner field");
    }
    if (!$("#beginDowntimeReason").val()) {
      return addModalAlert("You must fill the reason field");
    }
    var uri = "/api/begin-downtime/" + node.Key.Hostname + "/" + node.Key.Port + "/" + $("#beginDowntimeOwner").val() + "/" + $("#beginDowntimeReason").val() + "/" + $("#beginDowntimeDuration").val();
    apiCommand(uri);
  });
  $('#node_modal button[data-btn=refresh-instance]').click(function() {
    apiCommand("/api/refresh/" + node.Key.Hostname + "/" + node.Key.Port, "refresh");
  });
  $('#node_modal button[data-btn=skip-query]').click(function() {
    apiCommand("/api/skip-query/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=start-slave]').click(function() {
    apiCommand("/api/start-slave/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=restart-slave]').click(function() {
    apiCommand("/api/restart-slave/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal [data-btn=stop-slave]').click(function() {
    apiCommand("/api/stop-slave/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal [data-btn=stop-slave-nice]').click(function() {
    apiCommand("/api/stop-slave-nice/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=detach-slave]').click(function() {
    apiCommand("/api/detach-slave/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=reattach-slave]').click(function() {
    apiCommand("/api/reattach-slave/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=reattach-slave-master-host]').click(function() {
    apiCommand("/api/reattach-slave-master-host/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=reset-slave]').click(function() {
    var message = "<p>Are you sure you wish to reset <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
      "</strong></code>?" +
      "<p>This will stop and break the replication." +
      "<p>FYI, this is a destructive operation that cannot be easily reverted";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        apiCommand("/api/reset-slave/" + node.Key.Hostname + "/" + node.Key.Port);
      }
    });
    return false;
  });
  $('#node_modal button[data-btn=set-read-only]').click(function() {
    apiCommand("/api/set-read-only/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=set-writeable]').click(function() {
    apiCommand("/api/set-writeable/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=enable-gtid]').click(function() {
    var message = "<p>Are you sure you wish to enable GTID on <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
      "</strong></code>?" +
      "<p>Replication <i>might</i> break as consequence";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        apiCommand("/api/enable-gtid/" + node.Key.Hostname + "/" + node.Key.Port);
      }
    });
  });
  $('#node_modal button[data-btn=disable-gtid]').click(function() {
    var message = "<p>Are you sure you wish to disable GTID on <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
      "</strong></code>?" +
      "<p>Replication <i>might</i> break as consequence";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        apiCommand("/api/disable-gtid/" + node.Key.Hostname + "/" + node.Key.Port);
      }
    });
  });
  $('#node_modal button[data-btn=forget-instance]').click(function() {
    var message = "<p>Are you sure you wish to forget <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
      "</strong></code>?" +
      "<p>It may be re-discovered if accessible from an existing instance through replication topology.";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        apiCommand("/api/forget/" + node.Key.Hostname + "/" + node.Key.Port);
      }
    });
    return false;
  });

  $("body").on("click", "#node_modal a[data-btn=move-equivalent]", function(event) {
    var targetHostname = $(event.target).attr("data-hostname");
    var targetPort = $(event.target).attr("data-port");
    apiCommand("/api/move-equivalent/" + node.Key.Hostname + "/" + node.Key.Port + "/" + targetHostname + "/" + targetPort);
  });

  if (node.IsDowntimed) {
    $('#node_modal [data-panel-type=downtime]').html("Downtimed by <strong>" + node.DowntimeOwner + "</strong> until " + node.DowntimeEndTimestamp);
    $('#node_modal [data-description=downtime-status]').html(
      node.DowntimeReason
    );
    $('#node_modal [data-panel-type=begin-downtime]').hide();
    $('#node_modal button[data-btn=begin-downtime]').hide();
    $('#node_modal [data-panel-type=end-downtime]').show();
  } else {
    $('#node_modal [data-panel-type=downtime]').html("Downtime");
    $('#node_modal [data-panel-type=begin-downtime]').show();
    $('#node_modal [data-panel-type=end-downtime]').hide();
    $('#node_modal button[data-btn=end-downtime]').hide();
  }
  $('#node_modal button[data-btn=skip-query]').hide();
  $('#node_modal button[data-btn=start-slave]').hide();
  $('#node_modal button[data-btn=restart-slave]').hide();
  $('#node_modal [data-btn-group=stop-slave]').hide();

  if (node.MasterKey.Hostname) {
    if (node.replicationRunning || node.replicationAttemptingToRun) {
      $('#node_modal [data-btn-group=stop-slave]').show();
      $('#node_modal button[data-btn=restart-slave]').show();
    } else if (!node.replicationRunning) {
      $('#node_modal button[data-btn=start-slave]').show();
    }
    if (!node.Slave_SQL_Running && node.LastSQLError) {
      $('#node_modal button[data-btn=skip-query]').show();
    }
  }

  $('#node_modal button[data-btn=set-read-only]').hide();
  $('#node_modal button[data-btn=set-writeable]').hide();
  if (node.ReadOnly) {
    $('#node_modal button[data-btn=set-writeable]').show();
  } else {
    $('#node_modal button[data-btn=set-read-only]').show();
  }

  $('#node_modal button[data-btn=enable-gtid]').hide();
  $('#node_modal button[data-btn=disable-gtid]').hide();
  if (node.usingGTID) {
    $('#node_modal button[data-btn=disable-gtid]').show();
  } else {
    $('#node_modal button[data-btn=enable-gtid]').show();
  }

  $('#node_modal button[data-btn=regroup-slaves]').hide();
  if (node.SlaveHosts.length > 1) {
    $('#node_modal button[data-btn=regroup-slaves]').show();
  }
  $('#node_modal button[data-btn=regroup-slaves]').click(function() {
    var message = "<p>Are you sure you wish to regroup slaves of <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
      "</strong></code>?" +
      "<p>This will attempt to promote one slave over its siblings";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        apiCommand("/api/regroup-slaves/" + node.Key.Hostname + "/" + node.Key.Port);
      }
    });
  });

  $('#node_modal button[data-btn=enslave-siblings]').hide();
  if (node.LogBinEnabled && node.LogSlaveUpdatesEnabled) {
    $('#node_modal button[data-btn=enslave-siblings]').show();
  }
  $('#node_modal button[data-btn=enslave-siblings]').click(function() {
    var message = "<p>Are you sure you want <code><strong>" + node.Key.Hostname + ":" + node.Key.Port +
      "</strong></code> to enslave its siblings?" +
      "<p>This will stop replication on this slave and on its siblings throughout the operation";
    bootbox.confirm(message, function(confirm) {
      if (confirm) {
        apiCommand("/api/enslave-siblings/" + node.Key.Hostname + "/" + node.Key.Port);
      }
    });
  });
  $('#node_modal button[data-btn=end-downtime]').click(function() {
    apiCommand("/api/end-downtime/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=recover]').hide();
  if (node.lastCheckInvalidProblem() && node.children && node.children.length > 0) {
    $('#node_modal button[data-btn=recover]').show();
  }
  $('#node_modal button[data-btn=recover]').click(function() {
    apiCommand("/api/recover/" + node.Key.Hostname + "/" + node.Key.Port);
  });
  $('#node_modal button[data-btn=end-maintenance]').hide();
  if (node.inMaintenance) {
    $('#node_modal button[data-btn=end-maintenance]').show();
  }
  $('#node_modal button[data-btn=end-maintenance]').click(function() {
    apiCommand("/api/end-maintenance/" + node.Key.Hostname + "/" + node.Key.Port);
  });


  if (!isAuthorizedForAction()) {
    $('#node_modal button[data-btn]').hide();
    $('#node_modal [data-btn-group]').hide();
  }

  $('#node_modal').modal({})
  $('#node_modal').unbind('hidden.bs.modal');
  $('#node_modal').on('hidden.bs.modal', function() {
    nodeModalVisible = false;
  })
}


function normalizeInstance(instance) {
  instance.id = getInstanceId(instance.Key.Hostname, instance.Key.Port);
  instance.title = instance.Key.Hostname + ':' + instance.Key.Port;
  instance.canonicalTitle = instance.title;
  instance.masterTitle = instance.MasterKey.Hostname + ":" + instance.MasterKey.Port;
  instance.masterId = getInstanceId(instance.MasterKey.Hostname,
    instance.MasterKey.Port);

  instance.replicationRunning = instance.Slave_SQL_Running && instance.Slave_IO_Running;
  instance.replicationAttemptingToRun = instance.Slave_SQL_Running || instance.Slave_IO_Running;
  instance.replicationLagReasonable = Math.abs(instance.SlaveLagSeconds.Int64 - instance.SQLDelay) <= 10;
  instance.isSeenRecently = instance.SecondsSinceLastSeen.Valid && instance.SecondsSinceLastSeen.Int64 <= 3600;
  instance.usingGTID = instance.UsingOracleGTID || instance.UsingMariaDBGTID;
  instance.isMaxScale = (instance.Version.indexOf("maxscale") >= 0);

  // used by cluster-tree
  instance.children = [];
  instance.parent = null;
  instance.hasMaster = true;
  instance.masterNode = null;
  instance.inMaintenance = false;
  instance.maintenanceEntry = null;
  instance.isFirstChildInDisplay = false

  instance.isMaster = (instance.title == instance.ClusterName);
  instance.isCoMaster = false;
  instance.isCandidateMaster = false;
  instance.isMostAdvancedOfSiblings = false;
  instance.isVirtual = false;
  instance.isAnchor = false;
  instance.isAggregate = false;

  instance.renderHint = "";
}

function normalizeInstanceProblem(instance) {
  instance.inMaintenanceProblem = function() {
    return instance.inMaintenance;
  }
  instance.lastCheckInvalidProblem = function() {
    return !instance.IsLastCheckValid;
  }
  instance.notRecentlyCheckedProblem = function() {
    return !instance.IsRecentlyChecked;
  }
  instance.notReplicatingProblem = function() {
    return !instance.replicationRunning && !(instance.isMaster && !instance.isCoMaster);
  }
  instance.replicationLagProblem = function() {
    return !instance.replicationLagReasonable;
  }

  instance.problem = null;
  instance.problemOrder = 0;
  if (instance.inMaintenanceProblem()) {
    instance.problem = "in_maintenance";
    instance.problemDescription = "This instance is now under maintenance due to some pending operation.\nSee audit page";
    instance.problemOrder = 1;
  } else if (instance.lastCheckInvalidProblem()) {
    instance.problem = "last_check_invalid";
    instance.problemDescription = "Instance cannot be reached by orchestrator.\nIt might be dead or there may be a network problem";
    instance.problemOrder = 2;
  } else if (instance.notRecentlyCheckedProblem()) {
    instance.problem = "not_recently_checked";
    instance.problemDescription = "Orchestrator has not made an attempt to reach this instance for a while now.\nThis should generally not happen; consider refreshing or re-discovering this instance";
    instance.problemOrder = 3;
  } else if (instance.notReplicatingProblem()) {
    // check slaves only; where not replicating
    instance.problem = "not_replicating";
    instance.problemDescription = "Replication is not running.\nEither stopped manually or is failing on I/O or SQL error.";
    instance.problemOrder = 4;
  } else if (instance.replicationLagProblem()) {
    instance.problem = "replication_lag";
    instance.problemDescription = "Slave is lagging in replication.\nThis diagnostic is based on either Seconds_behind_master or configured SlaveLagQuery";
    instance.problemOrder = 5;
  }
  instance.hasProblem = (instance.problem != null);
  instance.hasConnectivityProblem = (!instance.IsLastCheckValid || !instance.IsRecentlyChecked);
}

var virtualInstanceCounter = 0;

function createVirtualInstance() {
  var virtualInstance = {
    id: "orchestrator-virtual-instance-" + (virtualInstanceCounter++),
    children: [],
    parent: null,
    hasMaster: false,
    inMaintenance: false,
    maintenanceEntry: null,
    isMaster: false,
    isCoMaster: false,
    isVirtual: true,
    SlaveLagSeconds: 0,
    SecondsSinceLastSeen: 0
  }
  normalizeInstanceProblem(virtualInstance);
  return virtualInstance;
}

function normalizeInstances(instances, maintenanceList) {
  if (!instances) {
    instances = [];
  }
  instances.forEach(function(instance) {
    normalizeInstance(instance);
  });
  // Take canonical host name: strip down longest common suffix of all hosts
  // (experimental; you may not like it)
  var hostNames = instances.map(function(instance) {
    return instance.title
  });
  var suffixLength = commonSuffixLength(hostNames);
  instances.forEach(function(instance) {
    instance.canonicalTitle = canonizeInstanceTitle(instance.title)
    if (instance.canonicalTitle == instance.title) {
      instance.canonicalTitle = instance.title.substring(0, instance.title.length - suffixLength);
    }
  });
  var instancesMap = {};
  instances.forEach(function(instance) {
    instancesMap[instance.id] = instance;
  });
  // mark maintenance instances
  maintenanceList.forEach(function(maintenanceEntry) {
    var instanceId = getInstanceId(maintenanceEntry.Key.Hostname, maintenanceEntry.Key.Port)
    if (instanceId in instancesMap) {
      instancesMap[instanceId].inMaintenance = true;
      instancesMap[instanceId].maintenanceEntry = maintenanceEntry;
    }
  });
  instances.forEach(function(instance) {
    // Now that we also know about maintenance
    normalizeInstanceProblem(instance);
  });
  // create the tree array
  instances.forEach(function(instance) {
    // add to parent
    var parent = instancesMap[instance.masterId];
    if (parent) {
      instance.parent = parent;
      instance.masterNode = parent;
      // create child array if it doesn't exist
      parent.children.push(instance);
      // (parent.contents || (parent.contents = [])).push(instance);
    } else {
      // parent is null or missing
      instance.hasMaster = false;
      instance.parent = null;
      instance.masterNode = null;
    }
  });

  instances.forEach(function(instance) {
    if (instance.masterNode != null) {
      instance.isSQLThreadCaughtUpWithIOThread = (instance.ExecBinlogCoordinates.LogFile == instance.ReadBinlogCoordinates.LogFile &&
        instance.ExecBinlogCoordinates.LogPos == instance.ReadBinlogCoordinates.LogPos);
    } else {
      instance.isSQLThreadCaughtUpWithIOThread = false;
    }
  });

  instances.forEach(function(instance) {
    if (instance.isMaster && instance.parent != null && instance.parent.parent != null && instance.parent.parent.id == instance.id) {
      // In case there's a master-master setup, introduce a virtual node
      // that is parent of both.
      // This is for visualization purposes...
      var virtualCoMastersRoot = createVirtualInstance();
      coMaster = instance.parent;

      function setAsCoMaster(instance, coMaster) {
        instance.isCoMaster = true;
        instance.hasMaster = true;
        instance.masterId = coMaster.id;
        instance.masterNode = coMaster;

        var index = coMaster.children.indexOf(instance);
        if (index >= 0)
          coMaster.children.splice(index, 1);

        instance.parent = virtualCoMastersRoot;
        virtualCoMastersRoot.children.push(instance);
      }
      setAsCoMaster(instance, coMaster);
      setAsCoMaster(coMaster, instance);

      instancesMap[virtualCoMastersRoot.id] = virtualCoMastersRoot;
    }
  });
  return instancesMap;
}


function renderInstanceElement(popoverElement, instance, renderType) {
  popoverElement.attr("data-nodeid", instance.id);
  popoverElement.find("h3").attr('title', instance.title);
  popoverElement.find("h3").html('&nbsp;<div class="pull-left">' +
    instance.canonicalTitle + '</div><div class="pull-right"><a href="#"><span class="glyphicon glyphicon-cog" title="Open config dialog"></span></a></div>');
  var indicateLastSeenInStatus = false;

  if (instance.isAggregate) {
    popoverElement.find("h3 div.pull-right span").remove();
    popoverElement.find(".instance-content").append('<div>Instances: <div class="pull-right"></div></div>');

    function addInstancesBadge(count, badgeClass, title) {
      popoverElement.find(".instance-content .pull-right").append('<span class="badge ' + badgeClass + '" data-toggle="tooltip" data-placement="bottom" data-html="true" title="' + title + '">' + count + '</span> ');
      popoverElement.find('[data-toggle="tooltip"]').tooltip();
    }
    var instancesHint = instance.aggregatedProblems[""].join("<br>");
    addInstancesBadge(instance.aggregatedInstances.length, "label-primary", "Aggregated instances<br>" + instancesHint);

    for (var problemType in instance.aggregatedProblems) {
      if (errorMapping[problemType]) {
        var description = errorMapping[problemType]["description"];
        var instancesHint = instance.aggregatedProblems[problemType].join("<br>");
        addInstancesBadge(instance.aggregatedProblems[problemType].length, errorMapping[problemType]["badge"], description + "<br>" + instancesHint);
      }
    }
  }
  if (!instance.isAggregate) {
    if (instance.isFirstChildInDisplay) {
      popoverElement.addClass("first-child-in-display");
      popoverElement.attr("data-first-child-in-display", "true");
    }
    if (instance.usingGTID) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-globe" title="Using GTID"></span> ');
    }
    if (instance.UsingPseudoGTID) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-globe" title="Using Pseudo GTID"></span> ');
    }
    if (!instance.ReadOnly) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-pencil" title="Writeable"></span> ');
    }
    if (instance.isMostAdvancedOfSiblings) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-star" title="Most advanced slave"></span> ');
    }
    if (instance.CountMySQLSnapshots > 0) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-camera" title="' + instance.CountMySQLSnapshots + ' snapshots"></span> ');
    }
    if (instance.HasReplicationFilters) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-filter" title="Using replication filters"></span> ');
    }
    if (instance.LogBinEnabled && instance.LogSlaveUpdatesEnabled && !(instance.isMaster && !instance.isCoMaster)) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-forward" title="Logs slave updates"></span> ');
    }
    if (instance.IsCandidate) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-heart" title="Candidate"></span> ');
    }
    if (instance.inMaintenanceProblem()) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-wrench" title="In maintenance"></span> ');
    }
    if (instance.IsDetached) {
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-remove-sign" title="Replication forcibly detached"></span> ');
    }
    if (instance.IsDowntimed) {
      var downtimeMessage = 'Downtimed by ' + instance.DowntimeOwner + ': ' + instance.DowntimeReason + '.\nEnds: ' + instance.DowntimeEndTimestamp;
      popoverElement.find("h3 div.pull-right").prepend('<span class="glyphicon glyphicon-volume-off" title="' + downtimeMessage + '"></span> ');
    }

    if (instance.lastCheckInvalidProblem()) {
      instance.renderHint = "fatal";
      indicateLastSeenInStatus = true;
    } else if (instance.notRecentlyCheckedProblem()) {
      instance.renderHint = "stale";
      indicateLastSeenInStatus = true;
    } else if (instance.notReplicatingProblem()) {
      // check slaves only; check master only if it's co-master where not
      // replicating
      instance.renderHint = "danger";
    } else if (instance.replicationLagProblem()) {
      instance.renderHint = "warning";
    }
    if (instance.renderHint != "") {
      popoverElement.find("h3").addClass("label-" + instance.renderHint);
    }
    var statusMessage = instance.SlaveLagSeconds.Int64 + ' seconds lag';
    if (indicateLastSeenInStatus) {
      statusMessage = 'seen ' + instance.SecondsSinceLastSeen.Int64 + ' seconds ago';
    }
    var contentHtml = '' + instance.Version;
    if (instance.LogBinEnabled) {
      contentHtml += " " + instance.Binlog_format;
    }
    contentHtml = '' + '<div class="pull-right">' + statusMessage + ' </div>' + '<p class="instance-basic-info">' + contentHtml + '</p>';
    if (instance.isCoMaster) {
      contentHtml += '<p><strong>Co master</strong></p>';
    } else if (instance.isMaster) {
      contentHtml += '<p><strong>Master</strong></p>';
    }
    if (renderType == "search") {
      contentHtml += '<p>' + 'Cluster: <a href="' + appUrl('/web/cluster/' + instance.ClusterName) + '">' + instance.ClusterName + '</a>' + '</p>';
    }
    if (renderType == "problems") {
      contentHtml += '<p>' + 'Problem: <strong title="' + instance.problemDescription + '">' + instance.problem.replace(/_/g, ' ') + '</strong>' + '</p>';
    }
    popoverElement.find(".instance-content").html(contentHtml);
  }
  // if (instance.isCandidateMaster) {
  // popoverElement.append('<h4 class="popover-footer"><strong>Master
  // candidate</strong><div class="pull-right"><button class="btn btn-xs
  // btn-default" data-command="make-master"><span class="glyphicon
  // glyphicon-play"></span> Make master</button></div></h4>');
  // } else if (instance.isMostAdvancedOfSiblings) {
  // popoverElement.append('<h4
  // class="popover-footer"><strong>Candidate</strong><div
  // class="pull-right"><button class="btn btn-xs btn-default"
  // data-command="make-local-master"><span class="glyphicon
  // glyphicon-play"></span> Make local master</button></div></h4>');
  // }

  popoverElement.find("h3 a").click(function() {
    openNodeModal(instance);
    return false;
  });
}

var onClustersListeners = [];

function onClusters(func) {
  onClustersListeners.push(func);
}


function getParameterByName(name) {
  name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
  var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
    results = regex.exec(location.search);
  return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
}


$(document).ready(function() {
  visualizeBrand();

  $('body').css('background-image', 'url(' + appUrl('/images/tile.png') + ')');

  $(".navbar-nav li").removeClass("active");
  $(".navbar-nav li[data-nav-page='" + activePage() + "']").addClass("active");

  $.get(appUrl("/api/clusters-info"), function(clusters) {
    clusters.forEach(function(cluster) {
      var title = '<span class="small">' + cluster.ClusterName + '</span>';
      title = ((cluster.ClusterAlias != "") ? '<strong>' + cluster.ClusterAlias + '</strong>, ' + title : title);
      $("#dropdown-clusters").append('<li><a href="' + appUrl('/web/cluster/' + cluster.ClusterName) + '">' + title + '</a></li>');
    });
    onClustersListeners.forEach(function(func) {
      func(clusters);
    });
  }, "json");
  $(".ajaxLoader").click(function() {
    return false;
  });
  $("#refreshCountdown").click(function() {
    if ($.cookie("auto-refresh") == "true") {
      $.cookie("auto-refresh", "false", {
        path: '/',
        expires: 1
      });
    } else {
      $.cookie("auto-refresh", "true", {
        path: '/',
        expires: 1
      });
    }
    updateCountdownDisplay();
  });
  if (agentsHttpActive() == "true") {
    $("#nav_agents").show();
  }
  if (contextMenuVisible() == "true") {
    showContextMenu();
  }
  if (!isAuthorizedForAction()) {
    $("[data-nav-page=read-only]").css('display', 'inline-block');
  }
  if (getUserId() != "") {
    $("[data-nav-page=user-id]").css('display', 'inline-block');
    $("[data-nav-page=user-id] a").html(" " + getUserId());
  }
  var orchestratorMsg = getParameterByName("orchestrator-msg")
  if (orchestratorMsg) {
    addInfo(orchestratorMsg)

    reloadPageHint = {
      hint: getParameterByName("hint"),
      hostname: getParameterByName("hostname"),
      port: getParameterByName("port")
    }
    history.pushState(null, document.title, location.href.split("?orchestrator-msg=")[0])
  }
  if (typeof($.cookie("auto-refresh")) === 'undefined') {
    $.cookie("auto-refresh", "true", {
      path: '/',
      expires: 1
    });
  }
  $("#searchInput").focus();
});
