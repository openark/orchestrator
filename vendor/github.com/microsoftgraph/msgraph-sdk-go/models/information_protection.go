package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtection 
type InformationProtection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The bitlocker property
    bitlocker Bitlockerable
    // The OdataType property
    odataType *string
    // The threatAssessmentRequests property
    threatAssessmentRequests []ThreatAssessmentRequestable
}
// NewInformationProtection instantiates a new InformationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInformationProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBitlocker gets the bitlocker property value. The bitlocker property
func (m *InformationProtection) GetBitlocker()(Bitlockerable) {
    return m.bitlocker
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bitlocker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitlockerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitlocker(val.(Bitlockerable))
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
    res["threatAssessmentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateThreatAssessmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThreatAssessmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(ThreatAssessmentRequestable)
            }
            m.SetThreatAssessmentRequests(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InformationProtection) GetOdataType()(*string) {
    return m.odataType
}
// GetThreatAssessmentRequests gets the threatAssessmentRequests property value. The threatAssessmentRequests property
func (m *InformationProtection) GetThreatAssessmentRequests()([]ThreatAssessmentRequestable) {
    return m.threatAssessmentRequests
}
// Serialize serializes information the current object
func (m *InformationProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bitlocker", m.GetBitlocker())
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
    if m.GetThreatAssessmentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetThreatAssessmentRequests()))
        for i, v := range m.GetThreatAssessmentRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("threatAssessmentRequests", cast)
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
func (m *InformationProtection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBitlocker sets the bitlocker property value. The bitlocker property
func (m *InformationProtection) SetBitlocker(value Bitlockerable)() {
    m.bitlocker = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InformationProtection) SetOdataType(value *string)() {
    m.odataType = value
}
// SetThreatAssessmentRequests sets the threatAssessmentRequests property value. The threatAssessmentRequests property
func (m *InformationProtection) SetThreatAssessmentRequests(value []ThreatAssessmentRequestable)() {
    m.threatAssessmentRequests = value
}
