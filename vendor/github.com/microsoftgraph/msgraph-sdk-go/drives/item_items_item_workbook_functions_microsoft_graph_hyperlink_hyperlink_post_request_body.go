package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The friendlyName property
    friendlyName iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The linkLocation property
    linkLocation iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["friendlyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFriendlyName(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["linkLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkLocation(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetFriendlyName gets the friendlyName property value. The friendlyName property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) GetFriendlyName()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.friendlyName
}
// GetLinkLocation gets the linkLocation property value. The linkLocation property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) GetLinkLocation()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.linkLocation
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("friendlyName", m.GetFriendlyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("linkLocation", m.GetLinkLocation())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFriendlyName sets the friendlyName property value. The friendlyName property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) SetFriendlyName(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.friendlyName = value
}
// SetLinkLocation sets the linkLocation property value. The linkLocation property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHyperlinkHyperlinkPostRequestBody) SetLinkLocation(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.linkLocation = value
}
