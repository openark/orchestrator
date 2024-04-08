package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The alpha property
    alpha iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The probabilityS property
    probabilityS iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The trials property
    trials iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlpha gets the alpha property value. The alpha property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) GetAlpha()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.alpha
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alpha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlpha(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["probabilityS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProbabilityS(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["trials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrials(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetProbabilityS gets the probabilityS property value. The probabilityS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) GetProbabilityS()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.probabilityS
}
// GetTrials gets the trials property value. The trials property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) GetTrials()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.trials
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("alpha", m.GetAlpha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("probabilityS", m.GetProbabilityS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trials", m.GetTrials())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlpha sets the alpha property value. The alpha property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) SetAlpha(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.alpha = value
}
// SetProbabilityS sets the probabilityS property value. The probabilityS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) SetProbabilityS(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.probabilityS = value
}
// SetTrials sets the trials property value. The trials property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_InvBinom_InvPostRequestBody) SetTrials(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.trials = value
}
