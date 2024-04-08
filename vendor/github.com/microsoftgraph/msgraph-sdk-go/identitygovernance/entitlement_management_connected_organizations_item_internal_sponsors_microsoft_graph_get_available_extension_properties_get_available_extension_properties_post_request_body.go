package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody 
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The isSyncedFromOnPremises property
    isSyncedFromOnPremises *bool
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody instantiates a new EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody()(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody) {
    m := &EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isSyncedFromOnPremises"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSyncedFromOnPremises(val)
        }
        return nil
    }
    return res
}
// GetIsSyncedFromOnPremises gets the isSyncedFromOnPremises property value. The isSyncedFromOnPremises property
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody) GetIsSyncedFromOnPremises()(*bool) {
    return m.isSyncedFromOnPremises
}
// Serialize serializes information the current object
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isSyncedFromOnPremises", m.GetIsSyncedFromOnPremises())
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
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsSyncedFromOnPremises sets the isSyncedFromOnPremises property value. The isSyncedFromOnPremises property
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBody) SetIsSyncedFromOnPremises(value *bool)() {
    m.isSyncedFromOnPremises = value
}
