package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InsightIdentity 
type InsightIdentity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email address of the user who shared the item.
    address *string
    // The display name of the user who shared the item.
    displayName *string
    // The id of the user who shared the item.
    id *string
    // The OdataType property
    odataType *string
}
// NewInsightIdentity instantiates a new insightIdentity and sets the default values.
func NewInsightIdentity()(*InsightIdentity) {
    m := &InsightIdentity{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInsightIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInsightIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInsightIdentity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InsightIdentity) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The email address of the user who shared the item.
func (m *InsightIdentity) GetAddress()(*string) {
    return m.address
}
// GetDisplayName gets the displayName property value. The display name of the user who shared the item.
func (m *InsightIdentity) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InsightIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
// GetId gets the id property value. The id of the user who shared the item.
func (m *InsightIdentity) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InsightIdentity) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *InsightIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
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
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *InsightIdentity) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The email address of the user who shared the item.
func (m *InsightIdentity) SetAddress(value *string)() {
    m.address = value
}
// SetDisplayName sets the displayName property value. The display name of the user who shared the item.
func (m *InsightIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The id of the user who shared the item.
func (m *InsightIdentity) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InsightIdentity) SetOdataType(value *string)() {
    m.odataType = value
}
