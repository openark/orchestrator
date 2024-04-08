package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppReturnCodeCollectionResponse 
type Win32LobAppReturnCodeCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Win32LobAppReturnCodeable
}
// NewWin32LobAppReturnCodeCollectionResponse instantiates a new Win32LobAppReturnCodeCollectionResponse and sets the default values.
func NewWin32LobAppReturnCodeCollectionResponse()(*Win32LobAppReturnCodeCollectionResponse) {
    m := &Win32LobAppReturnCodeCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWin32LobAppReturnCodeCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppReturnCodeCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppReturnCodeCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppReturnCodeCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppReturnCodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppReturnCodeable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppReturnCodeable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Win32LobAppReturnCodeCollectionResponse) GetValue()([]Win32LobAppReturnCodeable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Win32LobAppReturnCodeCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Win32LobAppReturnCodeCollectionResponse) SetValue(value []Win32LobAppReturnCodeable)() {
    m.value = value
}
