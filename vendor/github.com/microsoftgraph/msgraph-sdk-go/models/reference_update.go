package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferenceUpdate 
type ReferenceUpdate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataId property
    odataId *string
    // The OdataType property
    odataType *string
}
// NewReferenceUpdate instantiates a new ReferenceUpdate and sets the default values.
func NewReferenceUpdate()(*ReferenceUpdate) {
    m := &ReferenceUpdate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReferenceUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReferenceUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReferenceUpdate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReferenceUpdate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReferenceUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataId(val)
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
// GetOdataId gets the @odata.id property value. The OdataId property
func (m *ReferenceUpdate) GetOdataId()(*string) {
    return m.odataId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ReferenceUpdate) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ReferenceUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.id", m.GetOdataId())
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
func (m *ReferenceUpdate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataId sets the @odata.id property value. The OdataId property
func (m *ReferenceUpdate) SetOdataId(value *string)() {
    m.odataId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ReferenceUpdate) SetOdataType(value *string)() {
    m.odataType = value
}
