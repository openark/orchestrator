package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody 
type MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The storageLocation property
    storageLocation *string
}
// NewMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody instantiates a new MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody and sets the default values.
func NewMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody()(*MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) {
    m := &MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["storageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageLocation(val)
        }
        return nil
    }
    return res
}
// GetStorageLocation gets the storageLocation property value. The storageLocation property
func (m *MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) GetStorageLocation()(*string) {
    return m.storageLocation
}
// Serialize serializes information the current object
func (m *MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("storageLocation", m.GetStorageLocation())
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
func (m *MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStorageLocation sets the storageLocation property value. The storageLocation property
func (m *MicrosoftGraphExportPersonalDataExportPersonalDataPostRequestBody) SetStorageLocation(value *string)() {
    m.storageLocation = value
}
