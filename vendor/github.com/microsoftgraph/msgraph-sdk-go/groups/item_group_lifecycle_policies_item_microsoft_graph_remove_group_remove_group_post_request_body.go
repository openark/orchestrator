package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody 
type ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The groupId property
    groupId *string
}
// NewItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody instantiates a new ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody and sets the default values.
func NewItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody()(*ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody) {
    m := &ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The groupId property
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody) GetGroupId()(*string) {
    return m.groupId
}
// Serialize serializes information the current object
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
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
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphRemoveGroupRemoveGroupPostRequestBody) SetGroupId(value *string)() {
    m.groupId = value
}
