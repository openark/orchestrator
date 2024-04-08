package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The degFreedom1 property
    degFreedom1 iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The degFreedom2 property
    degFreedom2 iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The x property
    x iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDegFreedom1 gets the degFreedom1 property value. The degFreedom1 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) GetDegFreedom1()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.degFreedom1
}
// GetDegFreedom2 gets the degFreedom2 property value. The degFreedom2 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) GetDegFreedom2()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.degFreedom2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) GetX()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.x
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDegFreedom1 sets the degFreedom1 property value. The degFreedom1 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) SetDegFreedom1(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.degFreedom1 = value
}
// SetDegFreedom2 sets the degFreedom2 property value. The degFreedom2 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) SetDegFreedom2(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.degFreedom2 = value
}
// SetX sets the x property value. The x property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphF_Dist_RTF_Dist_RTPostRequestBody) SetX(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.x = value
}
