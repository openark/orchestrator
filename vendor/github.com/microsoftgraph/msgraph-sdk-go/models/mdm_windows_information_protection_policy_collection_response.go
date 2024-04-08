package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MdmWindowsInformationProtectionPolicyCollectionResponse 
type MdmWindowsInformationProtectionPolicyCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []MdmWindowsInformationProtectionPolicyable
}
// NewMdmWindowsInformationProtectionPolicyCollectionResponse instantiates a new MdmWindowsInformationProtectionPolicyCollectionResponse and sets the default values.
func NewMdmWindowsInformationProtectionPolicyCollectionResponse()(*MdmWindowsInformationProtectionPolicyCollectionResponse) {
    m := &MdmWindowsInformationProtectionPolicyCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMdmWindowsInformationProtectionPolicyCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMdmWindowsInformationProtectionPolicyCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMdmWindowsInformationProtectionPolicyCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MdmWindowsInformationProtectionPolicyCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MdmWindowsInformationProtectionPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(MdmWindowsInformationProtectionPolicyable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MdmWindowsInformationProtectionPolicyCollectionResponse) GetValue()([]MdmWindowsInformationProtectionPolicyable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MdmWindowsInformationProtectionPolicyCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MdmWindowsInformationProtectionPolicyCollectionResponse) SetValue(value []MdmWindowsInformationProtectionPolicyable)() {
    m.value = value
}
