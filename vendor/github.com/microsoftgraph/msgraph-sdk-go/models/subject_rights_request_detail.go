package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestDetail 
type SubjectRightsRequestDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Count of items that are excluded from the request.
    excludedItemCount *int64
    // Count of items per insight.
    insightCounts []KeyValuePairable
    // Count of items found.
    itemCount *int64
    // Count of item that need review.
    itemNeedReview *int64
    // The OdataType property
    odataType *string
    // Count of items per product, such as Exchange, SharePoint, OneDrive, and Teams.
    productItemCounts []KeyValuePairable
    // Count of items signed off by the administrator.
    signedOffItemCount *int64
    // Total item size in bytes.
    totalItemSize *int64
}
// NewSubjectRightsRequestDetail instantiates a new subjectRightsRequestDetail and sets the default values.
func NewSubjectRightsRequestDetail()(*SubjectRightsRequestDetail) {
    m := &SubjectRightsRequestDetail{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSubjectRightsRequestDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubjectRightsRequestDetail) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExcludedItemCount gets the excludedItemCount property value. Count of items that are excluded from the request.
func (m *SubjectRightsRequestDetail) GetExcludedItemCount()(*int64) {
    return m.excludedItemCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludedItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedItemCount(val)
        }
        return nil
    }
    res["insightCounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetInsightCounts(res)
        }
        return nil
    }
    res["itemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCount(val)
        }
        return nil
    }
    res["itemNeedReview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemNeedReview(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["productItemCounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetProductItemCounts(res)
        }
        return nil
    }
    res["signedOffItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignedOffItemCount(val)
        }
        return nil
    }
    res["totalItemSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalItemSize(val)
        }
        return nil
    }
    return res
}
// GetInsightCounts gets the insightCounts property value. Count of items per insight.
func (m *SubjectRightsRequestDetail) GetInsightCounts()([]KeyValuePairable) {
    return m.insightCounts
}
// GetItemCount gets the itemCount property value. Count of items found.
func (m *SubjectRightsRequestDetail) GetItemCount()(*int64) {
    return m.itemCount
}
// GetItemNeedReview gets the itemNeedReview property value. Count of item that need review.
func (m *SubjectRightsRequestDetail) GetItemNeedReview()(*int64) {
    return m.itemNeedReview
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SubjectRightsRequestDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetProductItemCounts gets the productItemCounts property value. Count of items per product, such as Exchange, SharePoint, OneDrive, and Teams.
func (m *SubjectRightsRequestDetail) GetProductItemCounts()([]KeyValuePairable) {
    return m.productItemCounts
}
// GetSignedOffItemCount gets the signedOffItemCount property value. Count of items signed off by the administrator.
func (m *SubjectRightsRequestDetail) GetSignedOffItemCount()(*int64) {
    return m.signedOffItemCount
}
// GetTotalItemSize gets the totalItemSize property value. Total item size in bytes.
func (m *SubjectRightsRequestDetail) GetTotalItemSize()(*int64) {
    return m.totalItemSize
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("excludedItemCount", m.GetExcludedItemCount())
        if err != nil {
            return err
        }
    }
    if m.GetInsightCounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInsightCounts()))
        for i, v := range m.GetInsightCounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("insightCounts", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("itemCount", m.GetItemCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("itemNeedReview", m.GetItemNeedReview())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetProductItemCounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProductItemCounts()))
        for i, v := range m.GetProductItemCounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("productItemCounts", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("signedOffItemCount", m.GetSignedOffItemCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalItemSize", m.GetTotalItemSize())
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
func (m *SubjectRightsRequestDetail) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExcludedItemCount sets the excludedItemCount property value. Count of items that are excluded from the request.
func (m *SubjectRightsRequestDetail) SetExcludedItemCount(value *int64)() {
    m.excludedItemCount = value
}
// SetInsightCounts sets the insightCounts property value. Count of items per insight.
func (m *SubjectRightsRequestDetail) SetInsightCounts(value []KeyValuePairable)() {
    m.insightCounts = value
}
// SetItemCount sets the itemCount property value. Count of items found.
func (m *SubjectRightsRequestDetail) SetItemCount(value *int64)() {
    m.itemCount = value
}
// SetItemNeedReview sets the itemNeedReview property value. Count of item that need review.
func (m *SubjectRightsRequestDetail) SetItemNeedReview(value *int64)() {
    m.itemNeedReview = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SubjectRightsRequestDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProductItemCounts sets the productItemCounts property value. Count of items per product, such as Exchange, SharePoint, OneDrive, and Teams.
func (m *SubjectRightsRequestDetail) SetProductItemCounts(value []KeyValuePairable)() {
    m.productItemCounts = value
}
// SetSignedOffItemCount sets the signedOffItemCount property value. Count of items signed off by the administrator.
func (m *SubjectRightsRequestDetail) SetSignedOffItemCount(value *int64)() {
    m.signedOffItemCount = value
}
// SetTotalItemSize sets the totalItemSize property value. Total item size in bytes.
func (m *SubjectRightsRequestDetail) SetTotalItemSize(value *int64)() {
    m.totalItemSize = value
}
