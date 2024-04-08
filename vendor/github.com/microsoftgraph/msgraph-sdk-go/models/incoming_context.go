package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IncomingContext 
type IncomingContext struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the participant that is under observation. Read-only.
    observedParticipantId *string
    // The OdataType property
    odataType *string
    // The identity that the call is happening on behalf of.
    onBehalfOf IdentitySetable
    // The ID of the participant that triggered the incoming call. Read-only.
    sourceParticipantId *string
    // The identity that transferred the call.
    transferor IdentitySetable
}
// NewIncomingContext instantiates a new incomingContext and sets the default values.
func NewIncomingContext()(*IncomingContext) {
    m := &IncomingContext{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIncomingContextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIncomingContextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIncomingContext(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IncomingContext) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IncomingContext) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["observedParticipantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObservedParticipantId(val)
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
    res["onBehalfOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBehalfOf(val.(IdentitySetable))
        }
        return nil
    }
    res["sourceParticipantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceParticipantId(val)
        }
        return nil
    }
    res["transferor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferor(val.(IdentitySetable))
        }
        return nil
    }
    return res
}
// GetObservedParticipantId gets the observedParticipantId property value. The ID of the participant that is under observation. Read-only.
func (m *IncomingContext) GetObservedParticipantId()(*string) {
    return m.observedParticipantId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IncomingContext) GetOdataType()(*string) {
    return m.odataType
}
// GetOnBehalfOf gets the onBehalfOf property value. The identity that the call is happening on behalf of.
func (m *IncomingContext) GetOnBehalfOf()(IdentitySetable) {
    return m.onBehalfOf
}
// GetSourceParticipantId gets the sourceParticipantId property value. The ID of the participant that triggered the incoming call. Read-only.
func (m *IncomingContext) GetSourceParticipantId()(*string) {
    return m.sourceParticipantId
}
// GetTransferor gets the transferor property value. The identity that transferred the call.
func (m *IncomingContext) GetTransferor()(IdentitySetable) {
    return m.transferor
}
// Serialize serializes information the current object
func (m *IncomingContext) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("observedParticipantId", m.GetObservedParticipantId())
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
        err := writer.WriteObjectValue("onBehalfOf", m.GetOnBehalfOf())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceParticipantId", m.GetSourceParticipantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transferor", m.GetTransferor())
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
func (m *IncomingContext) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetObservedParticipantId sets the observedParticipantId property value. The ID of the participant that is under observation. Read-only.
func (m *IncomingContext) SetObservedParticipantId(value *string)() {
    m.observedParticipantId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IncomingContext) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnBehalfOf sets the onBehalfOf property value. The identity that the call is happening on behalf of.
func (m *IncomingContext) SetOnBehalfOf(value IdentitySetable)() {
    m.onBehalfOf = value
}
// SetSourceParticipantId sets the sourceParticipantId property value. The ID of the participant that triggered the incoming call. Read-only.
func (m *IncomingContext) SetSourceParticipantId(value *string)() {
    m.sourceParticipantId = value
}
// SetTransferor sets the transferor property value. The identity that transferred the call.
func (m *IncomingContext) SetTransferor(value IdentitySetable)() {
    m.transferor = value
}
