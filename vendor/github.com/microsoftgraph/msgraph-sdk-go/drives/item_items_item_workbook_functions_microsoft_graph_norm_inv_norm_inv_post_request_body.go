package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The mean property
    mean iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The probability property
    probability iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The standardDev property
    standardDev iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mean"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMean(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["probability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProbability(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["standardDev"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStandardDev(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetMean gets the mean property value. The mean property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) GetMean()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.mean
}
// GetProbability gets the probability property value. The probability property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) GetProbability()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.probability
}
// GetStandardDev gets the standardDev property value. The standardDev property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) GetStandardDev()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.standardDev
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("mean", m.GetMean())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("probability", m.GetProbability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("standardDev", m.GetStandardDev())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMean sets the mean property value. The mean property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) SetMean(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.mean = value
}
// SetProbability sets the probability property value. The probability property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) SetProbability(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.probability = value
}
// SetStandardDev sets the standardDev property value. The standardDev property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNorm_InvNorm_InvPostRequestBody) SetStandardDev(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.standardDev = value
}
