package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestSignatureVerification 
type RequestSignatureVerification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allowedWeakAlgorithms property
    allowedWeakAlgorithms *WeakAlgorithms
    // The isSignedRequestRequired property
    isSignedRequestRequired *bool
    // The OdataType property
    odataType *string
}
// NewRequestSignatureVerification instantiates a new requestSignatureVerification and sets the default values.
func NewRequestSignatureVerification()(*RequestSignatureVerification) {
    m := &RequestSignatureVerification{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRequestSignatureVerificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestSignatureVerificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestSignatureVerification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestSignatureVerification) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedWeakAlgorithms gets the allowedWeakAlgorithms property value. The allowedWeakAlgorithms property
func (m *RequestSignatureVerification) GetAllowedWeakAlgorithms()(*WeakAlgorithms) {
    return m.allowedWeakAlgorithms
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestSignatureVerification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedWeakAlgorithms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeakAlgorithms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedWeakAlgorithms(val.(*WeakAlgorithms))
        }
        return nil
    }
    res["isSignedRequestRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSignedRequestRequired(val)
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
// GetIsSignedRequestRequired gets the isSignedRequestRequired property value. The isSignedRequestRequired property
func (m *RequestSignatureVerification) GetIsSignedRequestRequired()(*bool) {
    return m.isSignedRequestRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RequestSignatureVerification) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *RequestSignatureVerification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedWeakAlgorithms() != nil {
        cast := (*m.GetAllowedWeakAlgorithms()).String()
        err := writer.WriteStringValue("allowedWeakAlgorithms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSignedRequestRequired", m.GetIsSignedRequestRequired())
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
func (m *RequestSignatureVerification) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedWeakAlgorithms sets the allowedWeakAlgorithms property value. The allowedWeakAlgorithms property
func (m *RequestSignatureVerification) SetAllowedWeakAlgorithms(value *WeakAlgorithms)() {
    m.allowedWeakAlgorithms = value
}
// SetIsSignedRequestRequired sets the isSignedRequestRequired property value. The isSignedRequestRequired property
func (m *RequestSignatureVerification) SetIsSignedRequestRequired(value *bool)() {
    m.isSignedRequestRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RequestSignatureVerification) SetOdataType(value *string)() {
    m.odataType = value
}
