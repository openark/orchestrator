package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InferenceClassificationOverride 
type InferenceClassificationOverride struct {
    Entity
    // Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
    classifyAs *InferenceClassificationType
    // The email address information of the sender for whom the override is created.
    senderEmailAddress EmailAddressable
}
// NewInferenceClassificationOverride instantiates a new inferenceClassificationOverride and sets the default values.
func NewInferenceClassificationOverride()(*InferenceClassificationOverride) {
    m := &InferenceClassificationOverride{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInferenceClassificationOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInferenceClassificationOverrideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInferenceClassificationOverride(), nil
}
// GetClassifyAs gets the classifyAs property value. Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
func (m *InferenceClassificationOverride) GetClassifyAs()(*InferenceClassificationType) {
    return m.classifyAs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InferenceClassificationOverride) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classifyAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInferenceClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassifyAs(val.(*InferenceClassificationType))
        }
        return nil
    }
    res["senderEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderEmailAddress(val.(EmailAddressable))
        }
        return nil
    }
    return res
}
// GetSenderEmailAddress gets the senderEmailAddress property value. The email address information of the sender for whom the override is created.
func (m *InferenceClassificationOverride) GetSenderEmailAddress()(EmailAddressable) {
    return m.senderEmailAddress
}
// Serialize serializes information the current object
func (m *InferenceClassificationOverride) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassifyAs() != nil {
        cast := (*m.GetClassifyAs()).String()
        err = writer.WriteStringValue("classifyAs", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("senderEmailAddress", m.GetSenderEmailAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassifyAs sets the classifyAs property value. Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
func (m *InferenceClassificationOverride) SetClassifyAs(value *InferenceClassificationType)() {
    m.classifyAs = value
}
// SetSenderEmailAddress sets the senderEmailAddress property value. The email address information of the sender for whom the override is created.
func (m *InferenceClassificationOverride) SetSenderEmailAddress(value EmailAddressable)() {
    m.senderEmailAddress = value
}
