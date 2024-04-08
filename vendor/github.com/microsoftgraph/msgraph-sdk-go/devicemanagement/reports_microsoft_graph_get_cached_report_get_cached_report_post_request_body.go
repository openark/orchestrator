package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody 
type ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The groupBy property
    groupBy []string
    // The id property
    id *string
    // The orderBy property
    orderBy []string
    // The search property
    search *string
    // The select property
    selectEscaped []string
    // The skip property
    skip *int32
    // The top property
    top *int32
}
// NewReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody instantiates a new ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody and sets the default values.
func NewReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody()(*ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) {
    m := &ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupBy(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["orderBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrderBy(res)
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val)
        }
        return nil
    }
    res["select"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelect(res)
        }
        return nil
    }
    res["skip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkip(val)
        }
        return nil
    }
    res["top"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    return res
}
// GetGroupBy gets the groupBy property value. The groupBy property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetGroupBy()([]string) {
    return m.groupBy
}
// GetId gets the id property value. The id property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetId()(*string) {
    return m.id
}
// GetOrderBy gets the orderBy property value. The orderBy property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetOrderBy()([]string) {
    return m.orderBy
}
// GetSearch gets the search property value. The search property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetSearch()(*string) {
    return m.search
}
// GetSelect gets the select property value. The select property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetSelect()([]string) {
    return m.selectEscaped
}
// GetSkip gets the skip property value. The skip property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetSkip()(*int32) {
    return m.skip
}
// GetTop gets the top property value. The top property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) GetTop()(*int32) {
    return m.top
}
// Serialize serializes information the current object
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroupBy() != nil {
        err := writer.WriteCollectionOfStringValues("groupBy", m.GetGroupBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetOrderBy() != nil {
        err := writer.WriteCollectionOfStringValues("orderBy", m.GetOrderBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    if m.GetSelect() != nil {
        err := writer.WriteCollectionOfStringValues("select", m.GetSelect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
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
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupBy sets the groupBy property value. The groupBy property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) SetGroupBy(value []string)() {
    m.groupBy = value
}
// SetId sets the id property value. The id property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) SetId(value *string)() {
    m.id = value
}
// SetOrderBy sets the orderBy property value. The orderBy property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) SetOrderBy(value []string)() {
    m.orderBy = value
}
// SetSearch sets the search property value. The search property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) SetSearch(value *string)() {
    m.search = value
}
// SetSelect sets the select property value. The select property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) SetSelect(value []string)() {
    m.selectEscaped = value
}
// SetSkip sets the skip property value. The skip property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) SetSkip(value *int32)() {
    m.skip = value
}
// SetTop sets the top property value. The top property
func (m *ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBody) SetTop(value *int32)() {
    m.top = value
}
