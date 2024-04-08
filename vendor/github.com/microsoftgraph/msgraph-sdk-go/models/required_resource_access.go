package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequiredResourceAccess 
type RequiredResourceAccess struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
    resourceAccess []ResourceAccessable
    // The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
    resourceAppId *string
}
// NewRequiredResourceAccess instantiates a new requiredResourceAccess and sets the default values.
func NewRequiredResourceAccess()(*RequiredResourceAccess) {
    m := &RequiredResourceAccess{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRequiredResourceAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequiredResourceAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequiredResourceAccess(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequiredResourceAccess) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequiredResourceAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["resourceAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceAccessable, len(val))
            for i, v := range val {
                res[i] = v.(ResourceAccessable)
            }
            m.SetResourceAccess(res)
        }
        return nil
    }
    res["resourceAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAppId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RequiredResourceAccess) GetOdataType()(*string) {
    return m.odataType
}
// GetResourceAccess gets the resourceAccess property value. The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
func (m *RequiredResourceAccess) GetResourceAccess()([]ResourceAccessable) {
    return m.resourceAccess
}
// GetResourceAppId gets the resourceAppId property value. The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
func (m *RequiredResourceAccess) GetResourceAppId()(*string) {
    return m.resourceAppId
}
// Serialize serializes information the current object
func (m *RequiredResourceAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetResourceAccess() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceAccess()))
        for i, v := range m.GetResourceAccess() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("resourceAccess", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceAppId", m.GetResourceAppId())
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
func (m *RequiredResourceAccess) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RequiredResourceAccess) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResourceAccess sets the resourceAccess property value. The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
func (m *RequiredResourceAccess) SetResourceAccess(value []ResourceAccessable)() {
    m.resourceAccess = value
}
// SetResourceAppId sets the resourceAppId property value. The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
func (m *RequiredResourceAccess) SetResourceAppId(value *string)() {
    m.resourceAppId = value
}
