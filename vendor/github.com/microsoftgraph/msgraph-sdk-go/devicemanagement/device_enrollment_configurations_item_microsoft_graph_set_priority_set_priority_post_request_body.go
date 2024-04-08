package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody 
type DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The priority property
    priority *int32
}
// NewDeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody instantiates a new DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody and sets the default values.
func NewDeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody()(*DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody) {
    m := &DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. The priority property
func (m *DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
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
func (m *DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPriority sets the priority property value. The priority property
func (m *DeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityPostRequestBody) SetPriority(value *int32)() {
    m.priority = value
}
