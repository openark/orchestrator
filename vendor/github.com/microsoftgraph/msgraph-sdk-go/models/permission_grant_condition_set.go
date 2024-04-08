package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PermissionGrantConditionSet 
type PermissionGrantConditionSet struct {
    Entity
    // A list of appId values for the client applications to match with, or a list with the single value all to match any client application. Default is the single value all.
    clientApplicationIds []string
    // A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the single value all to match with client apps from any publisher. Default is the single value all.
    clientApplicationPublisherIds []string
    // Set to true to only match on client applications with a verified publisher. Set to false to match on any client app, even if it does not have a verified publisher. Default is false.
    clientApplicationsFromVerifiedPublisherOnly *bool
    // A list of Azure Active Directory tenant IDs in which the client application is registered, or a list with the single value all to match with client apps registered in any tenant. Default is the single value all.
    clientApplicationTenantIds []string
    // The permission classification for the permission being granted, or all to match with any permission classification (including permissions which are not classified). Default is all.
    permissionClassification *string
    // The list of id values for the specific permissions to match with, or a list with the single value all to match with any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API's **servicePrincipal** object. The id of application permissions can be found in the appRoles property of the API's **servicePrincipal** object. The id of resource-specific application permissions can be found in the resourceSpecificApplicationPermissions property of the API's **servicePrincipal** object. Default is the single value all.
    permissions []string
    // The permission type of the permission being granted. Possible values: application for application permissions (e.g. app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated permissions which have not been configured by the API publisher to require admin consent—this value may be used in built-in permission grant policies, but cannot be used in custom permission grant policies. Required.
    permissionType *PermissionType
    // The appId of the resource application (e.g. the API) for which a permission is being granted, or any to match with any resource application or API. Default is any.
    resourceApplication *string
}
// NewPermissionGrantConditionSet instantiates a new permissionGrantConditionSet and sets the default values.
func NewPermissionGrantConditionSet()(*PermissionGrantConditionSet) {
    m := &PermissionGrantConditionSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePermissionGrantConditionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionGrantConditionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionGrantConditionSet(), nil
}
// GetClientApplicationIds gets the clientApplicationIds property value. A list of appId values for the client applications to match with, or a list with the single value all to match any client application. Default is the single value all.
func (m *PermissionGrantConditionSet) GetClientApplicationIds()([]string) {
    return m.clientApplicationIds
}
// GetClientApplicationPublisherIds gets the clientApplicationPublisherIds property value. A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the single value all to match with client apps from any publisher. Default is the single value all.
func (m *PermissionGrantConditionSet) GetClientApplicationPublisherIds()([]string) {
    return m.clientApplicationPublisherIds
}
// GetClientApplicationsFromVerifiedPublisherOnly gets the clientApplicationsFromVerifiedPublisherOnly property value. Set to true to only match on client applications with a verified publisher. Set to false to match on any client app, even if it does not have a verified publisher. Default is false.
func (m *PermissionGrantConditionSet) GetClientApplicationsFromVerifiedPublisherOnly()(*bool) {
    return m.clientApplicationsFromVerifiedPublisherOnly
}
// GetClientApplicationTenantIds gets the clientApplicationTenantIds property value. A list of Azure Active Directory tenant IDs in which the client application is registered, or a list with the single value all to match with client apps registered in any tenant. Default is the single value all.
func (m *PermissionGrantConditionSet) GetClientApplicationTenantIds()([]string) {
    return m.clientApplicationTenantIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionGrantConditionSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientApplicationIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetClientApplicationIds(res)
        }
        return nil
    }
    res["clientApplicationPublisherIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetClientApplicationPublisherIds(res)
        }
        return nil
    }
    res["clientApplicationsFromVerifiedPublisherOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientApplicationsFromVerifiedPublisherOnly(val)
        }
        return nil
    }
    res["clientApplicationTenantIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetClientApplicationTenantIds(res)
        }
        return nil
    }
    res["permissionClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionClassification(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPermissions(res)
        }
        return nil
    }
    res["permissionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionType(val.(*PermissionType))
        }
        return nil
    }
    res["resourceApplication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceApplication(val)
        }
        return nil
    }
    return res
}
// GetPermissionClassification gets the permissionClassification property value. The permission classification for the permission being granted, or all to match with any permission classification (including permissions which are not classified). Default is all.
func (m *PermissionGrantConditionSet) GetPermissionClassification()(*string) {
    return m.permissionClassification
}
// GetPermissions gets the permissions property value. The list of id values for the specific permissions to match with, or a list with the single value all to match with any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API's **servicePrincipal** object. The id of application permissions can be found in the appRoles property of the API's **servicePrincipal** object. The id of resource-specific application permissions can be found in the resourceSpecificApplicationPermissions property of the API's **servicePrincipal** object. Default is the single value all.
func (m *PermissionGrantConditionSet) GetPermissions()([]string) {
    return m.permissions
}
// GetPermissionType gets the permissionType property value. The permission type of the permission being granted. Possible values: application for application permissions (e.g. app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated permissions which have not been configured by the API publisher to require admin consent—this value may be used in built-in permission grant policies, but cannot be used in custom permission grant policies. Required.
func (m *PermissionGrantConditionSet) GetPermissionType()(*PermissionType) {
    return m.permissionType
}
// GetResourceApplication gets the resourceApplication property value. The appId of the resource application (e.g. the API) for which a permission is being granted, or any to match with any resource application or API. Default is any.
func (m *PermissionGrantConditionSet) GetResourceApplication()(*string) {
    return m.resourceApplication
}
// Serialize serializes information the current object
func (m *PermissionGrantConditionSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClientApplicationIds() != nil {
        err = writer.WriteCollectionOfStringValues("clientApplicationIds", m.GetClientApplicationIds())
        if err != nil {
            return err
        }
    }
    if m.GetClientApplicationPublisherIds() != nil {
        err = writer.WriteCollectionOfStringValues("clientApplicationPublisherIds", m.GetClientApplicationPublisherIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("clientApplicationsFromVerifiedPublisherOnly", m.GetClientApplicationsFromVerifiedPublisherOnly())
        if err != nil {
            return err
        }
    }
    if m.GetClientApplicationTenantIds() != nil {
        err = writer.WriteCollectionOfStringValues("clientApplicationTenantIds", m.GetClientApplicationTenantIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionClassification", m.GetPermissionClassification())
        if err != nil {
            return err
        }
    }
    if m.GetPermissions() != nil {
        err = writer.WriteCollectionOfStringValues("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionType() != nil {
        cast := (*m.GetPermissionType()).String()
        err = writer.WriteStringValue("permissionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceApplication", m.GetResourceApplication())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientApplicationIds sets the clientApplicationIds property value. A list of appId values for the client applications to match with, or a list with the single value all to match any client application. Default is the single value all.
func (m *PermissionGrantConditionSet) SetClientApplicationIds(value []string)() {
    m.clientApplicationIds = value
}
// SetClientApplicationPublisherIds sets the clientApplicationPublisherIds property value. A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the single value all to match with client apps from any publisher. Default is the single value all.
func (m *PermissionGrantConditionSet) SetClientApplicationPublisherIds(value []string)() {
    m.clientApplicationPublisherIds = value
}
// SetClientApplicationsFromVerifiedPublisherOnly sets the clientApplicationsFromVerifiedPublisherOnly property value. Set to true to only match on client applications with a verified publisher. Set to false to match on any client app, even if it does not have a verified publisher. Default is false.
func (m *PermissionGrantConditionSet) SetClientApplicationsFromVerifiedPublisherOnly(value *bool)() {
    m.clientApplicationsFromVerifiedPublisherOnly = value
}
// SetClientApplicationTenantIds sets the clientApplicationTenantIds property value. A list of Azure Active Directory tenant IDs in which the client application is registered, or a list with the single value all to match with client apps registered in any tenant. Default is the single value all.
func (m *PermissionGrantConditionSet) SetClientApplicationTenantIds(value []string)() {
    m.clientApplicationTenantIds = value
}
// SetPermissionClassification sets the permissionClassification property value. The permission classification for the permission being granted, or all to match with any permission classification (including permissions which are not classified). Default is all.
func (m *PermissionGrantConditionSet) SetPermissionClassification(value *string)() {
    m.permissionClassification = value
}
// SetPermissions sets the permissions property value. The list of id values for the specific permissions to match with, or a list with the single value all to match with any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API's **servicePrincipal** object. The id of application permissions can be found in the appRoles property of the API's **servicePrincipal** object. The id of resource-specific application permissions can be found in the resourceSpecificApplicationPermissions property of the API's **servicePrincipal** object. Default is the single value all.
func (m *PermissionGrantConditionSet) SetPermissions(value []string)() {
    m.permissions = value
}
// SetPermissionType sets the permissionType property value. The permission type of the permission being granted. Possible values: application for application permissions (e.g. app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated permissions which have not been configured by the API publisher to require admin consent—this value may be used in built-in permission grant policies, but cannot be used in custom permission grant policies. Required.
func (m *PermissionGrantConditionSet) SetPermissionType(value *PermissionType)() {
    m.permissionType = value
}
// SetResourceApplication sets the resourceApplication property value. The appId of the resource application (e.g. the API) for which a permission is being granted, or any to match with any resource application or API. Default is any.
func (m *PermissionGrantConditionSet) SetResourceApplication(value *string)() {
    m.resourceApplication = value
}
