package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The numberS property
    numberS iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The numberS2 property
    numberS2 iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The probabilityS property
    probabilityS iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The trials property
    trials iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["numberS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberS(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["numberS2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberS2(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
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
// GetNumberS gets the numberS property value. The numberS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetNumberS()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.numberS
}
// GetNumberS2 gets the numberS2 property value. The numberS2 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetNumberS2()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.numberS2
}
// GetProbabilityS gets the probabilityS property value. The probabilityS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetProbabilityS()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.probabilityS
}
// GetTrials gets the trials property value. The trials property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetTrials()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.trials
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("numberS", m.GetNumberS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numberS2", m.GetNumberS2())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumberS sets the numberS property value. The numberS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetNumberS(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.numberS = value
}
// SetNumberS2 sets the numberS2 property value. The numberS2 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetNumberS2(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.numberS2 = value
}
// SetProbabilityS sets the probabilityS property value. The probabilityS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetProbabilityS(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.probabilityS = value
}
// SetTrials sets the trials property value. The trials property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetTrials(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.trials = value
}
