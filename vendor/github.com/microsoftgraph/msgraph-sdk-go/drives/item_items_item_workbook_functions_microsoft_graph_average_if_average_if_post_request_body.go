package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The averageRange property
    averageRange iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The criteria property
    criteria iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The range property
    rangeEscaped iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAverageRange gets the averageRange property value. The averageRange property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) GetAverageRange()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.averageRange
}
// GetCriteria gets the criteria property value. The criteria property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) GetCriteria()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.criteria
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["averageRange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageRange(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["criteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["range"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRange(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetRange gets the range property value. The range property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) GetRange()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.rangeEscaped
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("averageRange", m.GetAverageRange())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("criteria", m.GetCriteria())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("range", m.GetRange())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAverageRange sets the averageRange property value. The averageRange property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) SetAverageRange(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.averageRange = value
}
// SetCriteria sets the criteria property value. The criteria property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) SetCriteria(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.criteria = value
}
// SetRange sets the range property value. The range property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAverageIfAverageIfPostRequestBody) SetRange(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.rangeEscaped = value
}
