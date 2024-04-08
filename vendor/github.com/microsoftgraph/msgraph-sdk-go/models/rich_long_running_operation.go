package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RichLongRunningOperation 
type RichLongRunningOperation struct {
    LongRunningOperation
    // Error that caused the operation to fail.
    error PublicErrorable
    // A value between 0 and 100 that indicates the progress of the operation.
    percentageComplete *int32
    // The unique identifier for the result.
    resourceId *string
    // The type of the operation.
    typeEscaped *string
}
// NewRichLongRunningOperation instantiates a new RichLongRunningOperation and sets the default values.
func NewRichLongRunningOperation()(*RichLongRunningOperation) {
    m := &RichLongRunningOperation{
        LongRunningOperation: *NewLongRunningOperation(),
    }
    return m
}
// CreateRichLongRunningOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRichLongRunningOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRichLongRunningOperation(), nil
}
// GetError gets the error property value. Error that caused the operation to fail.
func (m *RichLongRunningOperation) GetError()(PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RichLongRunningOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LongRunningOperation.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(PublicErrorable))
        }
        return nil
    }
    res["percentageComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentageComplete(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetPercentageComplete gets the percentageComplete property value. A value between 0 and 100 that indicates the progress of the operation.
func (m *RichLongRunningOperation) GetPercentageComplete()(*int32) {
    return m.percentageComplete
}
// GetResourceId gets the resourceId property value. The unique identifier for the result.
func (m *RichLongRunningOperation) GetResourceId()(*string) {
    return m.resourceId
}
// GetType gets the type property value. The type of the operation.
func (m *RichLongRunningOperation) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RichLongRunningOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LongRunningOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentageComplete", m.GetPercentageComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. Error that caused the operation to fail.
func (m *RichLongRunningOperation) SetError(value PublicErrorable)() {
    m.error = value
}
// SetPercentageComplete sets the percentageComplete property value. A value between 0 and 100 that indicates the progress of the operation.
func (m *RichLongRunningOperation) SetPercentageComplete(value *int32)() {
    m.percentageComplete = value
}
// SetResourceId sets the resourceId property value. The unique identifier for the result.
func (m *RichLongRunningOperation) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetType sets the type property value. The type of the operation.
func (m *RichLongRunningOperation) SetType(value *string)() {
    m.typeEscaped = value
}
