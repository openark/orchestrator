package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse 
type ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
    // The value property
    value []string
}
// NewItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse instantiates a new ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse and sets the default values.
func NewItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse()(*ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse) {
    m := &ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsResponse) SetValue(value []string)() {
    m.value = value
}
