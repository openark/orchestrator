package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestCollectionResponse 
type SubjectRightsRequestCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []SubjectRightsRequestable
}
// NewSubjectRightsRequestCollectionResponse instantiates a new SubjectRightsRequestCollectionResponse and sets the default values.
func NewSubjectRightsRequestCollectionResponse()(*SubjectRightsRequestCollectionResponse) {
    m := &SubjectRightsRequestCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSubjectRightsRequestCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectRightsRequestable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SubjectRightsRequestCollectionResponse) GetValue()([]SubjectRightsRequestable) {
    return m.value
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *SubjectRightsRequestCollectionResponse) SetValue(value []SubjectRightsRequestable)() {
    m.value = value
}
