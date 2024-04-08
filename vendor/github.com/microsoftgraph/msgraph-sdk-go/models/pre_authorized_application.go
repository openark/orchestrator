package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PreAuthorizedApplication 
type PreAuthorizedApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier for the application.
    appId *string
    // The unique identifier for the oauth2PermissionScopes the application requires.
    delegatedPermissionIds []string
    // The OdataType property
    odataType *string
}
// NewPreAuthorizedApplication instantiates a new preAuthorizedApplication and sets the default values.
func NewPreAuthorizedApplication()(*PreAuthorizedApplication) {
    m := &PreAuthorizedApplication{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePreAuthorizedApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePreAuthorizedApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreAuthorizedApplication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreAuthorizedApplication) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppId gets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) GetAppId()(*string) {
    return m.appId
}
// GetDelegatedPermissionIds gets the delegatedPermissionIds property value. The unique identifier for the oauth2PermissionScopes the application requires.
func (m *PreAuthorizedApplication) GetDelegatedPermissionIds()([]string) {
    return m.delegatedPermissionIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PreAuthorizedApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["delegatedPermissionIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDelegatedPermissionIds(res)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PreAuthorizedApplication) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PreAuthorizedApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedPermissionIds() != nil {
        err := writer.WriteCollectionOfStringValues("delegatedPermissionIds", m.GetDelegatedPermissionIds())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreAuthorizedApplication) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppId sets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) SetAppId(value *string)() {
    m.appId = value
}
// SetDelegatedPermissionIds sets the delegatedPermissionIds property value. The unique identifier for the oauth2PermissionScopes the application requires.
func (m *PreAuthorizedApplication) SetDelegatedPermissionIds(value []string)() {
    m.delegatedPermissionIds = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PreAuthorizedApplication) SetOdataType(value *string)() {
    m.odataType = value
}
