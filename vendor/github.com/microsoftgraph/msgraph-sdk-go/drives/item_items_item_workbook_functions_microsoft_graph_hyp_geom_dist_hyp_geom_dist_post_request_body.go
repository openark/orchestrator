package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cumulative property
    cumulative iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The numberPop property
    numberPop iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The numberSample property
    numberSample iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The populationS property
    populationS iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The sampleS property
    sampleS iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCumulative gets the cumulative property value. The cumulative property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) GetCumulative()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.cumulative
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["numberPop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberPop(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["numberSample"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberSample(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["populationS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPopulationS(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["sampleS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSampleS(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetNumberPop gets the numberPop property value. The numberPop property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) GetNumberPop()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.numberPop
}
// GetNumberSample gets the numberSample property value. The numberSample property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) GetNumberSample()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.numberSample
}
// GetPopulationS gets the populationS property value. The populationS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) GetPopulationS()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.populationS
}
// GetSampleS gets the sampleS property value. The sampleS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) GetSampleS()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.sampleS
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cumulative", m.GetCumulative())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numberPop", m.GetNumberPop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numberSample", m.GetNumberSample())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("populationS", m.GetPopulationS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sampleS", m.GetSampleS())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCumulative sets the cumulative property value. The cumulative property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) SetCumulative(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.cumulative = value
}
// SetNumberPop sets the numberPop property value. The numberPop property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) SetNumberPop(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.numberPop = value
}
// SetNumberSample sets the numberSample property value. The numberSample property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) SetNumberSample(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.numberSample = value
}
// SetPopulationS sets the populationS property value. The populationS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) SetPopulationS(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.populationS = value
}
// SetSampleS sets the sampleS property value. The sampleS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHypGeom_DistHypGeom_DistPostRequestBody) SetSampleS(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.sampleS = value
}
