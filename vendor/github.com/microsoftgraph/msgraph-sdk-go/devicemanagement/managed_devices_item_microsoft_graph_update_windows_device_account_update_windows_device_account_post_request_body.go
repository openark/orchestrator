package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody 
type ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The updateWindowsDeviceAccountActionParameter property
    updateWindowsDeviceAccountActionParameter iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateWindowsDeviceAccountActionParameterable
}
// NewManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody instantiates a new ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody and sets the default values.
func NewManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody()(*ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody) {
    m := &ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["updateWindowsDeviceAccountActionParameter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUpdateWindowsDeviceAccountActionParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindowsDeviceAccountActionParameter(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateWindowsDeviceAccountActionParameterable))
        }
        return nil
    }
    return res
}
// GetUpdateWindowsDeviceAccountActionParameter gets the updateWindowsDeviceAccountActionParameter property value. The updateWindowsDeviceAccountActionParameter property
func (m *ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody) GetUpdateWindowsDeviceAccountActionParameter()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateWindowsDeviceAccountActionParameterable) {
    return m.updateWindowsDeviceAccountActionParameter
}
// Serialize serializes information the current object
func (m *ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("updateWindowsDeviceAccountActionParameter", m.GetUpdateWindowsDeviceAccountActionParameter())
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
func (m *ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUpdateWindowsDeviceAccountActionParameter sets the updateWindowsDeviceAccountActionParameter property value. The updateWindowsDeviceAccountActionParameter property
func (m *ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountPostRequestBody) SetUpdateWindowsDeviceAccountActionParameter(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UpdateWindowsDeviceAccountActionParameterable)() {
    m.updateWindowsDeviceAccountActionParameter = value
}
