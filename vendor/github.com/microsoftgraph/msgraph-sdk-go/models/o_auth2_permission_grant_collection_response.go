package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OAuth2PermissionGrantCollectionResponse 
type OAuth2PermissionGrantCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []OAuth2PermissionGrantable
}
// NewOAuth2PermissionGrantCollectionResponse instantiates a new OAuth2PermissionGrantCollectionResponse and sets the default values.
func NewOAuth2PermissionGrantCollectionResponse()(*OAuth2PermissionGrantCollectionResponse) {
    m := &OAuth2PermissionGrantCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateOAuth2PermissionGrantCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOAuth2PermissionGrantCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOAuth2PermissionGrantCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OAuth2PermissionGrantCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOAuth2PermissionGrantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OAuth2PermissionGrantable, len(val))
            for i, v := range val {
                res[i] = v.(OAuth2PermissionGrantable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *OAuth2PermissionGrantCollectionResponse) GetValue()([]OAuth2PermissionGrantable) {
    return m.value
}
// Serialize serializes information the current object
func (m *OAuth2PermissionGrantCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OAuth2PermissionGrantCollectionResponse) SetValue(value []OAuth2PermissionGrantable)() {
    m.value = value
}
