package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecordingInfo 
type RecordingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The identities of the recording initiator.
    initiator IdentitySetable
    // The OdataType property
    odataType *string
    // The recordingStatus property
    recordingStatus *RecordingStatus
}
// NewRecordingInfo instantiates a new recordingInfo and sets the default values.
func NewRecordingInfo()(*RecordingInfo) {
    m := &RecordingInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecordingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecordingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecordingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordingInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecordingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["initiator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
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
    res["recordingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecordingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingStatus(val.(*RecordingStatus))
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. The identities of the recording initiator.
func (m *RecordingInfo) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RecordingInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetRecordingStatus gets the recordingStatus property value. The recordingStatus property
func (m *RecordingInfo) GetRecordingStatus()(*RecordingStatus) {
    return m.recordingStatus
}
// Serialize serializes information the current object
func (m *RecordingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("initiator", m.GetInitiator())
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
    if m.GetRecordingStatus() != nil {
        cast := (*m.GetRecordingStatus()).String()
        err := writer.WriteStringValue("recordingStatus", &cast)
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
func (m *RecordingInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInitiator sets the initiator property value. The identities of the recording initiator.
func (m *RecordingInfo) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RecordingInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecordingStatus sets the recordingStatus property value. The recordingStatus property
func (m *RecordingInfo) SetRecordingStatus(value *RecordingStatus)() {
    m.recordingStatus = value
}
