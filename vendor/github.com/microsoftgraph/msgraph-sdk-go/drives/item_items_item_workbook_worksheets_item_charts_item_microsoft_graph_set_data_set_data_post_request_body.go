package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody 
type ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The seriesBy property
    seriesBy *string
    // The sourceData property
    sourceData iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody()(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetSeriesBy gets the seriesBy property value. The seriesBy property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) GetSeriesBy()(*string) {
    return m.seriesBy
}
// GetSourceData gets the sourceData property value. The sourceData property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) GetSourceData()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.sourceData
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSeriesBy sets the seriesBy property value. The seriesBy property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) SetSeriesBy(value *string)() {
    m.seriesBy = value
}
// SetSourceData sets the sourceData property value. The sourceData property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetDataSetDataPostRequestBody) SetSourceData(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.sourceData = value
}
