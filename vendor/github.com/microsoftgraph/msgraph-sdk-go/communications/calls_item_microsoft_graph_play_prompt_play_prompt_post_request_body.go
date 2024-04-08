package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody 
type CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The clientContext property
    clientContext *string
    // The prompts property
    prompts []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable
}
// NewCallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody instantiates a new CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody and sets the default values.
func NewCallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody()(*CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) {
    m := &CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["prompts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePromptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable)
            }
            m.SetPrompts(res)
        }
        return nil
    }
    return res
}
// GetPrompts gets the prompts property value. The prompts property
func (m *CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) GetPrompts()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable) {
    return m.prompts
}
// Serialize serializes information the current object
func (m *CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    if m.GetPrompts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrompts()))
        for i, v := range m.GetPrompts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("prompts", cast)
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
func (m *CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetPrompts sets the prompts property value. The prompts property
func (m *CallsItemMicrosoftGraphPlayPromptPlayPromptPostRequestBody) SetPrompts(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable)() {
    m.prompts = value
}
