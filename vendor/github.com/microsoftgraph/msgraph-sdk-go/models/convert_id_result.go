package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConvertIdResult 
type ConvertIdResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
    errorDetails GenericErrorable
    // The OdataType property
    odataType *string
    // The identifier that was converted. This value is the original, un-converted identifier.
    sourceId *string
    // The converted identifier. This value is not present if the conversion failed.
    targetId *string
}
// NewConvertIdResult instantiates a new convertIdResult and sets the default values.
func NewConvertIdResult()(*ConvertIdResult) {
    m := &ConvertIdResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConvertIdResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConvertIdResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConvertIdResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConvertIdResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetErrorDetails gets the errorDetails property value. An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
func (m *ConvertIdResult) GetErrorDetails()(GenericErrorable) {
    return m.errorDetails
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConvertIdResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGenericErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDetails(val.(GenericErrorable))
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
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["targetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConvertIdResult) GetOdataType()(*string) {
    return m.odataType
}
// GetSourceId gets the sourceId property value. The identifier that was converted. This value is the original, un-converted identifier.
func (m *ConvertIdResult) GetSourceId()(*string) {
    return m.sourceId
}
// GetTargetId gets the targetId property value. The converted identifier. This value is not present if the conversion failed.
func (m *ConvertIdResult) GetTargetId()(*string) {
    return m.targetId
}
// Serialize serializes information the current object
func (m *ConvertIdResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("errorDetails", m.GetErrorDetails())
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
        err := writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetId", m.GetTargetId())
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
func (m *ConvertIdResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetErrorDetails sets the errorDetails property value. An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded.
func (m *ConvertIdResult) SetErrorDetails(value GenericErrorable)() {
    m.errorDetails = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConvertIdResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSourceId sets the sourceId property value. The identifier that was converted. This value is the original, un-converted identifier.
func (m *ConvertIdResult) SetSourceId(value *string)() {
    m.sourceId = value
}
// SetTargetId sets the targetId property value. The converted identifier. This value is not present if the conversion failed.
func (m *ConvertIdResult) SetTargetId(value *string)() {
    m.targetId = value
}
