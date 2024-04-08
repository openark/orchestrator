package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cumulative property
    cumulative iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The degFreedom1 property
    degFreedom1 iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The degFreedom2 property
    degFreedom2 iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The x property
    x iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCumulative gets the cumulative property value. The cumulative property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) GetCumulative()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.cumulative
}
// GetDegFreedom1 gets the degFreedom1 property value. The degFreedom1 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) GetDegFreedom1()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.degFreedom1
}
// GetDegFreedom2 gets the degFreedom2 property value. The degFreedom2 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) GetDegFreedom2()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.degFreedom2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["degFreedom1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDegFreedom1(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["degFreedom2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDegFreedom2(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) GetX()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.x
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cumulative", m.GetCumulative())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("degFreedom1", m.GetDegFreedom1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("degFreedom2", m.GetDegFreedom2())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCumulative sets the cumulative property value. The cumulative property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) SetCumulative(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.cumulative = value
}
// SetDegFreedom1 sets the degFreedom1 property value. The degFreedom1 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) SetDegFreedom1(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.degFreedom1 = value
}
// SetDegFreedom2 sets the degFreedom2 property value. The degFreedom2 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) SetDegFreedom2(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.degFreedom2 = value
}
// SetX sets the x property value. The x property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_DistF_DistPostRequestBody) SetX(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.x = value
}
