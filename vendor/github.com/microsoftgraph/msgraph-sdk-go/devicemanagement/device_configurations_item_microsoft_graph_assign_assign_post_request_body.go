package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody 
type DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The assignments property
    assignments []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationAssignmentable
}
// NewDeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody instantiates a new DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody and sets the default values.
func NewDeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody()(*DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody) {
    m := &DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignments gets the assignments property value. The assignments property
func (m *DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody) GetAssignments()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationAssignmentable) {
    return m.assignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignments", cast)
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
func (m *DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignments sets the assignments property value. The assignments property
func (m *DeviceConfigurationsItemMicrosoftGraphAssignAssignPostRequestBody) SetAssignments(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationAssignmentable)() {
    m.assignments = value
}
