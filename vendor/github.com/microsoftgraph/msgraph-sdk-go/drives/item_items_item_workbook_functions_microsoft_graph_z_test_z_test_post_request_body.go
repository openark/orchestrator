package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The array property
    array iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The sigma property
    sigma iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The x property
    x iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArray gets the array property value. The array property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) GetArray()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.array
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["array"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArray(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["sigma"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigma(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
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
// GetSigma gets the sigma property value. The sigma property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) GetSigma()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.sigma
}
// GetX gets the x property value. The x property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) GetX()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.x
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("array", m.GetArray())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sigma", m.GetSigma())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArray sets the array property value. The array property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) SetArray(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.array = value
}
// SetSigma sets the sigma property value. The sigma property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) SetSigma(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.sigma = value
}
// SetX sets the x property value. The x property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphZ_TestZ_TestPostRequestBody) SetX(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.x = value
}
