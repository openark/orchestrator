package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody 
type ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comment property
    comment *string
    // The name property
    name *string
    // The reference property
    reference iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody instantiates a new ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody()(*ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) {
    m := &ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The comment property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["reference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReference(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) GetName()(*string) {
    return m.name
}
// GetReference gets the reference property value. The reference property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) GetReference()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.reference
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reference", m.GetReference())
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
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The comment property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetName sets the name property value. The name property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetReference sets the reference property value. The reference property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddAddPostRequestBody) SetReference(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.reference = value
}
