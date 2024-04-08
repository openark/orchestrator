package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseCollectionPaginationCountResponse 
type BaseCollectionPaginationCountResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataCount property
    odataCount *int64
    // The OdataNextLink property
    odataNextLink *string
}
// NewBaseCollectionPaginationCountResponse instantiates a new BaseCollectionPaginationCountResponse and sets the default values.
func NewBaseCollectionPaginationCountResponse()(*BaseCollectionPaginationCountResponse) {
    m := &BaseCollectionPaginationCountResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBaseCollectionPaginationCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBaseCollectionPaginationCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBaseCollectionPaginationCountResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BaseCollectionPaginationCountResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseCollectionPaginationCountResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataCount(val)
        }
        return nil
    }
    res["@odata.nextLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataNextLink(val)
        }
        return nil
    }
    return res
}
// GetOdataCount gets the @odata.count property value. The OdataCount property
func (m *BaseCollectionPaginationCountResponse) GetOdataCount()(*int64) {
    return m.odataCount
}
// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *BaseCollectionPaginationCountResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// Serialize serializes information the current object
func (m *BaseCollectionPaginationCountResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("@odata.count", m.GetOdataCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetOdataNextLink())
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
func (m *BaseCollectionPaginationCountResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataCount sets the @odata.count property value. The OdataCount property
func (m *BaseCollectionPaginationCountResponse) SetOdataCount(value *int64)() {
    m.odataCount = value
}
// SetOdataNextLink sets the @odata.nextLink property value. The OdataNextLink property
func (m *BaseCollectionPaginationCountResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
