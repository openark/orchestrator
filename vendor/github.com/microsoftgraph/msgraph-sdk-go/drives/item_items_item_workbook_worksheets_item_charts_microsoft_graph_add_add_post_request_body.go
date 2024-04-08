package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody 
type ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The seriesBy property
    seriesBy *string
    // The sourceData property
    sourceData iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The type property
    typeEscaped *string
}
// NewItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody()(*ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["seriesBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesBy(val)
        }
        return nil
    }
    res["sourceData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceData(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetSeriesBy gets the seriesBy property value. The seriesBy property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) GetSeriesBy()(*string) {
    return m.seriesBy
}
// GetSourceData gets the sourceData property value. The sourceData property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) GetSourceData()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.sourceData
}
// GetType gets the type property value. The type property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("seriesBy", m.GetSeriesBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceData", m.GetSourceData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSeriesBy sets the seriesBy property value. The seriesBy property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) SetSeriesBy(value *string)() {
    m.seriesBy = value
}
// SetSourceData sets the sourceData property value. The sourceData property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) SetSourceData(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.sourceData = value
}
// SetType sets the type property value. The type property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsMicrosoftGraphAddAddPostRequestBody) SetType(value *string)() {
    m.typeEscaped = value
}
