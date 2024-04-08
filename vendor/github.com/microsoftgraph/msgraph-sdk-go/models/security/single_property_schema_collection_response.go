package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// SinglePropertySchemaCollectionResponse 
type SinglePropertySchemaCollectionResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
    // The value property
    value []SinglePropertySchemaable
}
// NewSinglePropertySchemaCollectionResponse instantiates a new SinglePropertySchemaCollectionResponse and sets the default values.
func NewSinglePropertySchemaCollectionResponse()(*SinglePropertySchemaCollectionResponse) {
    m := &SinglePropertySchemaCollectionResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSinglePropertySchemaCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSinglePropertySchemaCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSinglePropertySchemaCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SinglePropertySchemaCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSinglePropertySchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SinglePropertySchemaable, len(val))
            for i, v := range val {
                res[i] = v.(SinglePropertySchemaable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SinglePropertySchemaCollectionResponse) GetValue()([]SinglePropertySchemaable) {
    return m.value
}
// Serialize serializes information the current object
func (m *SinglePropertySchemaCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SinglePropertySchemaCollectionResponse) SetValue(value []SinglePropertySchemaable)() {
    m.value = value
}
