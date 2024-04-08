package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cumulative property
    cumulative iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The numberS property
    numberS iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The probabilityS property
    probabilityS iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The trials property
    trials iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCumulative gets the cumulative property value. The cumulative property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) GetCumulative()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.cumulative
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cumulative"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCumulative(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) GetNumberS()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.numberS
}
// GetProbabilityS gets the probabilityS property value. The probabilityS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) GetProbabilityS()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.probabilityS
}
// GetTrials gets the trials property value. The trials property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) GetTrials()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.trials
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cumulative", m.GetCumulative())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numberS", m.GetNumberS())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCumulative sets the cumulative property value. The cumulative property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) SetCumulative(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.cumulative = value
}
// SetNumberS sets the numberS property value. The numberS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) SetNumberS(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.numberS = value
}
// SetProbabilityS sets the probabilityS property value. The probabilityS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) SetProbabilityS(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.probabilityS = value
}
// SetTrials sets the trials property value. The trials property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_DistBinom_DistPostRequestBody) SetTrials(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.trials = value
}
