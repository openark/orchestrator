package identitygovernance

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody 
type EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The displayName property
    displayName *string
    // The entityType property
    entityType *string
    // The mailNickname property
    mailNickname *string
    // The onBehalfOfUserId property
    onBehalfOfUserId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody instantiates a new EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody()(*EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) {
    m := &EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetEntityType gets the entityType property value. The entityType property
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) GetEntityType()(*string) {
    return m.entityType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["entityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityType(val)
        }
        return nil
    }
    res["mailNickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["onBehalfOfUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBehalfOfUserId(val)
        }
        return nil
    }
    return res
}
// GetMailNickname gets the mailNickname property value. The mailNickname property
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetOnBehalfOfUserId gets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) GetOnBehalfOfUserId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.onBehalfOfUserId
}
// Serialize serializes information the current object
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityType", m.GetEntityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("onBehalfOfUserId", m.GetOnBehalfOfUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEntityType sets the entityType property value. The entityType property
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) SetEntityType(value *string)() {
    m.entityType = value
}
// SetMailNickname sets the mailNickname property value. The mailNickname property
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetOnBehalfOfUserId sets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsMicrosoftGraphValidatePropertiesValidatePropertiesPostRequestBody) SetOnBehalfOfUserId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.onBehalfOfUserId = value
}
