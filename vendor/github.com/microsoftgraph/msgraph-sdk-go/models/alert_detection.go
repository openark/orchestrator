package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertDetection 
type AlertDetection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The detectionType property
    detectionType *string
    // The method property
    method *string
    // The name property
    name *string
    // The OdataType property
    odataType *string
}
// NewAlertDetection instantiates a new alertDetection and sets the default values.
func NewAlertDetection()(*AlertDetection) {
    m := &AlertDetection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAlertDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertDetection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertDetection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetectionType gets the detectionType property value. The detectionType property
func (m *AlertDetection) GetDetectionType()(*string) {
    return m.detectionType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionType(val)
        }
        return nil
    }
    res["method"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethod(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetMethod gets the method property value. The method property
func (m *AlertDetection) GetMethod()(*string) {
    return m.method
}
// GetName gets the name property value. The name property
func (m *AlertDetection) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AlertDetection) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AlertDetection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detectionType", m.GetDetectionType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("method", m.GetMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *AlertDetection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetectionType sets the detectionType property value. The detectionType property
func (m *AlertDetection) SetDetectionType(value *string)() {
    m.detectionType = value
}
// SetMethod sets the method property value. The method property
func (m *AlertDetection) SetMethod(value *string)() {
    m.method = value
}
// SetName sets the name property value. The name property
func (m *AlertDetection) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlertDetection) SetOdataType(value *string)() {
    m.odataType = value
}
