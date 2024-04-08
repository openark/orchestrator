package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppIdentity 
type AppIdentity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Refers to the Unique GUID representing Application Id in the Azure Active Directory.
    appId *string
    // Refers to the Application Name displayed in the Azure Portal.
    displayName *string
    // The OdataType property
    odataType *string
    // Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App.
    servicePrincipalId *string
    // Refers to the Service Principal Name is the Application name in the tenant.
    servicePrincipalName *string
}
// NewAppIdentity instantiates a new appIdentity and sets the default values.
func NewAppIdentity()(*AppIdentity) {
    m := &AppIdentity{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppIdentity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppIdentity) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppId gets the appId property value. Refers to the Unique GUID representing Application Id in the Azure Active Directory.
func (m *AppIdentity) GetAppId()(*string) {
    return m.appId
}
// GetDisplayName gets the displayName property value. Refers to the Application Name displayed in the Azure Portal.
func (m *AppIdentity) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AppIdentity) GetOdataType()(*string) {
    return m.odataType
}
// GetServicePrincipalId gets the servicePrincipalId property value. Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App.
func (m *AppIdentity) GetServicePrincipalId()(*string) {
    return m.servicePrincipalId
}
// GetServicePrincipalName gets the servicePrincipalName property value. Refers to the Service Principal Name is the Application name in the tenant.
func (m *AppIdentity) GetServicePrincipalName()(*string) {
    return m.servicePrincipalName
}
// Serialize serializes information the current object
func (m *AppIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
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
func (m *AppIdentity) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppId sets the appId property value. Refers to the Unique GUID representing Application Id in the Azure Active Directory.
func (m *AppIdentity) SetAppId(value *string)() {
    m.appId = value
}
// SetDisplayName sets the displayName property value. Refers to the Application Name displayed in the Azure Portal.
func (m *AppIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppIdentity) SetOdataType(value *string)() {
    m.odataType = value
}
// SetServicePrincipalId sets the servicePrincipalId property value. Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App.
func (m *AppIdentity) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
// SetServicePrincipalName sets the servicePrincipalName property value. Refers to the Service Principal Name is the Application name in the tenant.
func (m *AppIdentity) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
