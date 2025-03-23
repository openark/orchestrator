$(document).ready(function () {
    $("body").on("click", ".ack-cluster-message", function (event) {
        var messageId = $(event.target).attr("data-message-id");
        bootbox.confirm(
            "Acknowledge the message if it is not actual anymore",
            function (result) {
                if (result === true) {
                    apiCommand("/api/ack-cluster-message/" + messageId);
                }
            }
        );
    });
});

$(document).on("orchestrator:postRenderCluster", function () {
    const clusterName = currentClusterName();
    $.get(appUrl("/api/cluster-messages/" + clusterName), function (messages) {
        messages = messages || [];
        for (const message of messages) {
            addClusterMessage(message)
        }
    }, "json");
});

function addClusterMessage(message) {
    if (isAnonymized()) {
        return false;
    }
    var ackButton = $('<button type="button" class="btn btn-xs btn-primary pull-right ack-cluster-message">Acknowledge</button>');
    ackButton.attr("data-message-id", message.ID);

    $("#alerts_container").append(
        $('<div class="alert alert-' + message.Level + ' alert-dismissable">' + '<button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>' + message.Message + '</div>').append(ackButton)
    );
    $(".alert").alert();
}
