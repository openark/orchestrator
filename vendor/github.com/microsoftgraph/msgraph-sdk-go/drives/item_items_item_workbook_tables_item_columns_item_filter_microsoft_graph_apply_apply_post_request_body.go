package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody 
type ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The criteria property
    criteria iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterCriteriaable
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody) {
    m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCriteria gets the criteria property value. The criteria property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody) GetCriteria()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterCriteriaable) {
    return m.criteria
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["criteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFilterCriteriaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterCriteriaable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("criteria", m.GetCriteria())
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCriteria sets the criteria property value. The criteria property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyApplyPostRequestBody) SetCriteria(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFilterCriteriaable)() {
    m.criteria = value
}
