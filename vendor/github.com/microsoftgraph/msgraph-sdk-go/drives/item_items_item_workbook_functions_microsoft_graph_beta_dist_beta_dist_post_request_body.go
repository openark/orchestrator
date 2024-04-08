package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody struct {
    // The A property
    a iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The alpha property
    alpha iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The B property
    b iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The beta property
    beta iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The cumulative property
    cumulative iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The x property
    x iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody(), nil
}
// GetA gets the a property value. The A property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) GetA()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.a
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlpha gets the alpha property value. The alpha property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) GetAlpha()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.alpha
}
// GetB gets the b property value. The B property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) GetB()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.b
}
// GetBeta gets the beta property value. The beta property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) GetBeta()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.beta
}
// GetCumulative gets the cumulative property value. The cumulative property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) GetCumulative()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.cumulative
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["a"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetA(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
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
    res["b"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["beta"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBeta(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
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
// GetX gets the x property value. The x property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) GetX()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.x
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("a", m.GetA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("alpha", m.GetAlpha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("b", m.GetB())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("beta", m.GetBeta())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cumulative", m.GetCumulative())
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
// SetA sets the a property value. The A property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) SetA(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.a = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlpha sets the alpha property value. The alpha property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) SetAlpha(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.alpha = value
}
// SetB sets the b property value. The B property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) SetB(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.b = value
}
// SetBeta sets the beta property value. The beta property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) SetBeta(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.beta = value
}
// SetCumulative sets the cumulative property value. The cumulative property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) SetCumulative(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.cumulative = value
}
// SetX sets the x property value. The x property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBeta_DistBeta_DistPostRequestBody) SetX(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.x = value
}
