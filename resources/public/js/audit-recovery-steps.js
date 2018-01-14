$(document).ready(function() {
  showLoader();

  auditRecoverySteps(recoveryUID(), $('#audit_recovery_steps tbody'))
  var recoveryUri = "/api/audit-recovery/uid/" + recoveryUID();

  $.get(appUrl(recoveryUri), function(recoveryAudits) {
    recoveryAudits = recoveryAudits || [];
    displayRecoveryAudit(recoveryAudits);
  }, "json");

  function displayRecoveryAudit(recoveryAudits) {
    hideLoader();

    recoveryAudits.forEach(function(audit) {
      // There's really just the one audit.
      var analyzedInstanceDisplay = audit.AnalysisEntry.AnalyzedInstanceKey.Hostname + ":" + audit.AnalysisEntry.AnalyzedInstanceKey.Port;
      var row = jQuery('<tr/>');
      var moreInfoElement = $('<span class="more-detection-info pull-right glyphicon glyphicon-info-sign text-primary" title="More info"></span>');
      moreInfoElement.attr("data-detection-id", audit.Id);

      $('<td/>', {
        text: audit.AnalysisEntry.Analysis
      }).prepend(moreInfoElement).appendTo(row);
      $('<a/>', {
        text: analyzedInstanceDisplay,
        href: appUrl("/web/search/" + analyzedInstanceDisplay)
      }).wrap($("<td/>")).parent().appendTo(row);
      $('<td/>', {
        text: audit.AnalysisEntry.CountReplicas
      }).appendTo(row);
      $('<a/>', {
        text: audit.AnalysisEntry.ClusterDetails.ClusterName,
        href: appUrl("/web/cluster/" + audit.AnalysisEntry.ClusterDetails.ClusterName)
      }).wrap($("<td/>")).parent().appendTo(row);
      $('<a/>', {
        text: audit.AnalysisEntry.ClusterDetails.ClusterAlias,
        href: appUrl("/web/cluster/alias/" + audit.AnalysisEntry.ClusterDetails.ClusterAlias)
      }).wrap($("<td/>")).parent().appendTo(row);
      $('<td/>', {
        text: audit.RecoveryStartTimestamp
      }).appendTo(row);
      row.appendTo('#audit_recovery_table tbody');
    });
  }
});
