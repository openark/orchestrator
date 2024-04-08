package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The coefficients property
    coefficients iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The m property
    m iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The n property
    n iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The x property
    x iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCoefficients gets the coefficients property value. The coefficients property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) GetCoefficients()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.coefficients
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["coefficients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoefficients(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["m"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetM(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["n"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetN(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["x"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetM gets the m property value. The m property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) GetM()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.m
}
// GetN gets the n property value. The n property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) GetN()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.n
}
// GetX gets the x property value. The x property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) GetX()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.x
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("coefficients", m.GetCoefficients())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("m", m.GetM())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("n", m.GetN())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("x", m.GetX())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCoefficients sets the coefficients property value. The coefficients property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) SetCoefficients(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.coefficients = value
}
// SetM sets the m property value. The m property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) SetM(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.m = value
}
// SetN sets the n property value. The n property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) SetN(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.n = value
}
// SetX sets the x property value. The x property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSeriesSumSeriesSumPostRequestBody) SetX(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.x = value
}
