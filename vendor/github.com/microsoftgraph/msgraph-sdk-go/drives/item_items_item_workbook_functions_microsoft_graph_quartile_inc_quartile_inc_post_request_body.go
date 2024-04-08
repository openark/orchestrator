package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The array property
    array iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The quart property
    quart iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArray gets the array property value. The array property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) GetArray()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.array
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["quart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuart(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetQuart gets the quart property value. The quart property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) GetQuart()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.quart
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("array", m.GetArray())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("quart", m.GetQuart())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArray sets the array property value. The array property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) SetArray(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.array = value
}
// SetQuart sets the quart property value. The quart property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuartile_IncQuartile_IncPostRequestBody) SetQuart(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.quart = value
}
