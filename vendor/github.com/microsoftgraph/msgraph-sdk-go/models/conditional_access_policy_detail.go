package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessPolicyDetail 
type ConditionalAccessPolicyDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The conditions property
    conditions ConditionalAccessConditionSetable
    // Represents grant controls that must be fulfilled for the policy.
    grantControls ConditionalAccessGrantControlsable
    // The OdataType property
    odataType *string
    // Represents a complex type of session controls that is enforced after sign-in.
    sessionControls ConditionalAccessSessionControlsable
}
// NewConditionalAccessPolicyDetail instantiates a new conditionalAccessPolicyDetail and sets the default values.
func NewConditionalAccessPolicyDetail()(*ConditionalAccessPolicyDetail) {
    m := &ConditionalAccessPolicyDetail{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessPolicyDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessPolicyDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessPolicyDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessPolicyDetail) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConditions gets the conditions property value. The conditions property
func (m *ConditionalAccessPolicyDetail) GetConditions()(ConditionalAccessConditionSetable) {
    return m.conditions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessPolicyDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessConditionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(ConditionalAccessConditionSetable))
        }
        return nil
    }
    res["grantControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessGrantControlsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrantControls(val.(ConditionalAccessGrantControlsable))
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
    res["sessionControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessSessionControlsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionControls(val.(ConditionalAccessSessionControlsable))
        }
        return nil
    }
    return res
}
// GetGrantControls gets the grantControls property value. Represents grant controls that must be fulfilled for the policy.
func (m *ConditionalAccessPolicyDetail) GetGrantControls()(ConditionalAccessGrantControlsable) {
    return m.grantControls
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessPolicyDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetSessionControls gets the sessionControls property value. Represents a complex type of session controls that is enforced after sign-in.
func (m *ConditionalAccessPolicyDetail) GetSessionControls()(ConditionalAccessSessionControlsable) {
    return m.sessionControls
}
// Serialize serializes information the current object
func (m *ConditionalAccessPolicyDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("grantControls", m.GetGrantControls())
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
        err := writer.WriteObjectValue("sessionControls", m.GetSessionControls())
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
func (m *ConditionalAccessPolicyDetail) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConditions sets the conditions property value. The conditions property
func (m *ConditionalAccessPolicyDetail) SetConditions(value ConditionalAccessConditionSetable)() {
    m.conditions = value
}
// SetGrantControls sets the grantControls property value. Represents grant controls that must be fulfilled for the policy.
func (m *ConditionalAccessPolicyDetail) SetGrantControls(value ConditionalAccessGrantControlsable)() {
    m.grantControls = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessPolicyDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSessionControls sets the sessionControls property value. Represents a complex type of session controls that is enforced after sign-in.
func (m *ConditionalAccessPolicyDetail) SetSessionControls(value ConditionalAccessSessionControlsable)() {
    m.sessionControls = value
}
