package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PendingOperations 
type PendingOperations struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // A property that indicates that an operation that might update the binary content of a file is pending completion.
    pendingContentUpdate PendingContentUpdateable
}
// NewPendingOperations instantiates a new pendingOperations and sets the default values.
func NewPendingOperations()(*PendingOperations) {
    m := &PendingOperations{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePendingOperationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePendingOperationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPendingOperations(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PendingOperations) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PendingOperations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["pendingContentUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePendingContentUpdateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingContentUpdate(val.(PendingContentUpdateable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PendingOperations) GetOdataType()(*string) {
    return m.odataType
}
// GetPendingContentUpdate gets the pendingContentUpdate property value. A property that indicates that an operation that might update the binary content of a file is pending completion.
func (m *PendingOperations) GetPendingContentUpdate()(PendingContentUpdateable) {
    return m.pendingContentUpdate
}
// Serialize serializes information the current object
func (m *PendingOperations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pendingContentUpdate", m.GetPendingContentUpdate())
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
func (m *PendingOperations) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PendingOperations) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPendingContentUpdate sets the pendingContentUpdate property value. A property that indicates that an operation that might update the binary content of a file is pending completion.
func (m *PendingOperations) SetPendingContentUpdate(value PendingContentUpdateable)() {
    m.pendingContentUpdate = value
}
