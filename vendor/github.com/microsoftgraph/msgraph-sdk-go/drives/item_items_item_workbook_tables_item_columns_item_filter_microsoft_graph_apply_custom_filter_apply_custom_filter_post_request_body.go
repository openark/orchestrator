package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody 
type ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The criteria1 property
    criteria1 *string
    // The criteria2 property
    criteria2 *string
    // The oper property
    oper *string
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody()(*ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) {
    m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCriteria1 gets the criteria1 property value. The criteria1 property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) GetCriteria1()(*string) {
    return m.criteria1
}
// GetCriteria2 gets the criteria2 property value. The criteria2 property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) GetCriteria2()(*string) {
    return m.criteria2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["criteria1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria1(val)
        }
        return nil
    }
    res["criteria2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria2(val)
        }
        return nil
    }
    res["oper"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOper(val)
        }
        return nil
    }
    return res
}
// GetOper gets the oper property value. The oper property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) GetOper()(*string) {
    return m.oper
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("criteria1", m.GetCriteria1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criteria2", m.GetCriteria2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oper", m.GetOper())
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCriteria1 sets the criteria1 property value. The criteria1 property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) SetCriteria1(value *string)() {
    m.criteria1 = value
}
// SetCriteria2 sets the criteria2 property value. The criteria2 property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) SetCriteria2(value *string)() {
    m.criteria2 = value
}
// SetOper sets the oper property value. The oper property
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterMicrosoftGraphApplyCustomFilterApplyCustomFilterPostRequestBody) SetOper(value *string)() {
    m.oper = value
}
