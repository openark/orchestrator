package azure

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/google/uuid"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/me"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/openark/orchestrator/go/config"
)

type GraphHelper struct {
	usernamePasswordCredential *azidentity.UsernamePasswordCredential
	graphUserScopes            []string
	appClient                  *msgraphsdkgo.GraphServiceClient
}

func NewGraphHelper() GraphHelper {
	g := GraphHelper{}
	return g
}

func (g *GraphHelper) InitializeGraphForAppAuth(email string, password string) error {
	clientId := config.Config.AzureClientID
	tenantId := config.Config.AzureTenantID
	scope := config.Config.AzureGraphUserScope
	g.graphUserScopes = strings.Split(scope, ",")
	cred, err := azidentity.NewUsernamePasswordCredential(tenantId, clientId, email, password, nil)
	if err != nil {
		log.Fatal("New Browser Cred Error: ", err)
	}
	g.usernamePasswordCredential = cred

	client, err := msgraphsdkgo.NewGraphServiceClientWithCredentials(cred, g.graphUserScopes)
	g.appClient = client

	return err

}

func (g *GraphHelper) GetUser() (models.Userable, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := me.MeRequestBuilderGetQueryParameters{
		// Only request specific properties
		Select: []string{"displayName", "id"},
	}
	user, err := g.appClient.Me().Get(ctx, &me.MeRequestBuilderGetRequestConfiguration{
		QueryParameters: &query,
	})
	return user, err
}

func (g *GraphHelper) GetAppRoleAssignments(userId string, appName string, appID string) (string, error) {
	var hasUserRole bool
	var hasGroupRole bool
	var userAppRoleId uuid.UUID
	var groupAppRoleId uuid.UUID
	var roleToCheck uuid.UUID
	var roleString string

	ctxAppRoleAssignements, cancelAppRoleAssignements := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelAppRoleAssignements()

	// Get User RoleID for appName
	appRoleAssignements, err := g.appClient.UsersById(userId).AppRoleAssignments().Get(ctxAppRoleAssignements, nil)

	values := appRoleAssignements.GetValue()
	for _, value := range values {
		if *value.GetResourceDisplayName() == appName {
			if *value.GetPrincipalType() == "User" {
				hasUserRole = true
				userAppRoleId = *value.GetAppRoleId()
			} else {
				hasGroupRole = true
				groupAppRoleId = *value.GetAppRoleId()
			}
		}
	}

	if hasUserRole {
		roleToCheck = userAppRoleId
	} else if !hasUserRole && hasGroupRole {
		roleToCheck = groupAppRoleId
	} else {
		roleToCheck = userAppRoleId
	}

	ctxAppRoles, cancelAppRoles := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelAppRoles()

	// Get Application Roles
	appRoles, err := g.appClient.ApplicationsById(appID).Get(ctxAppRoles, nil)

	roles := appRoles.GetAppRoles()
	for _, role := range roles {
		if *role.GetId() == roleToCheck {
			roleString = *role.GetValue()
			break
		}
	}

	return roleString, err
}

func (g *GraphHelper) Clear() {
	g.usernamePasswordCredential = nil
	g.graphUserScopes = nil
	g.appClient = nil
}
