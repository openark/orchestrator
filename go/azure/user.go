package azure

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/openark/golib/log"
	"github.com/openark/orchestrator/go/config"
)

func GetAzureUser(g *GraphHelper) models.Userable {
	guser, err := g.GetUser()

	if err != nil {
		log.Error("Error getting user: ", err)
		return nil
	}

	return guser
}

func GetAzureUserAppRoleAssignments(guser models.Userable, g *GraphHelper) string {
	userRole, err := g.GetAppRoleAssignments(*guser.GetId(), config.Config.AzureApplicationName, config.Config.AzureApplicationID)
	if err != nil {
		log.Error("Error getting user information: ", err)
	}

	return userRole
}
