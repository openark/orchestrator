package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageTextInputQuestionCollectionResponse 
type AccessPackageTextInputQuestionCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []AccessPackageTextInputQuestionable
}
// NewAccessPackageTextInputQuestionCollectionResponse instantiates a new AccessPackageTextInputQuestionCollectionResponse and sets the default values.
func NewAccessPackageTextInputQuestionCollectionResponse()(*AccessPackageTextInputQuestionCollectionResponse) {
    m := &AccessPackageTextInputQuestionCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAccessPackageTextInputQuestionCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageTextInputQuestionCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageTextInputQuestionCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageTextInputQuestionCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageTextInputQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageTextInputQuestionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageTextInputQuestionable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *AccessPackageTextInputQuestionCollectionResponse) GetValue()([]AccessPackageTextInputQuestionable) {
    return m.value
}
// Serialize serializes information the current object
func (m *AccessPackageTextInputQuestionCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AccessPackageTextInputQuestionCollectionResponse) SetValue(value []AccessPackageTextInputQuestionable)() {
    m.value = value
}
