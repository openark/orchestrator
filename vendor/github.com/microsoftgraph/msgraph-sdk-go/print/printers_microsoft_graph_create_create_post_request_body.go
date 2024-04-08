package print

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// PrintersMicrosoftGraphCreateCreatePostRequestBody 
type PrintersMicrosoftGraphCreateCreatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The certificateSigningRequest property
    certificateSigningRequest iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintCertificateSigningRequestable
    // The connectorId property
    connectorId *string
    // The displayName property
    displayName *string
    // The hasPhysicalDevice property
    hasPhysicalDevice *bool
    // The manufacturer property
    manufacturer *string
    // The model property
    model *string
    // The physicalDeviceId property
    physicalDeviceId *string
}
// NewPrintersMicrosoftGraphCreateCreatePostRequestBody instantiates a new PrintersMicrosoftGraphCreateCreatePostRequestBody and sets the default values.
func NewPrintersMicrosoftGraphCreateCreatePostRequestBody()(*PrintersMicrosoftGraphCreateCreatePostRequestBody) {
    m := &PrintersMicrosoftGraphCreateCreatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrintersMicrosoftGraphCreateCreatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintersMicrosoftGraphCreateCreatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintersMicrosoftGraphCreateCreatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificateSigningRequest gets the certificateSigningRequest property value. The certificateSigningRequest property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetCertificateSigningRequest()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintCertificateSigningRequestable) {
    return m.certificateSigningRequest
}
// GetConnectorId gets the connectorId property value. The connectorId property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetConnectorId()(*string) {
    return m.connectorId
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificateSigningRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintCertificateSigningRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSigningRequest(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintCertificateSigningRequestable))
        }
        return nil
    }
    res["connectorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["hasPhysicalDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasPhysicalDevice(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["physicalDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalDeviceId(val)
        }
        return nil
    }
    return res
}
// GetHasPhysicalDevice gets the hasPhysicalDevice property value. The hasPhysicalDevice property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetHasPhysicalDevice()(*bool) {
    return m.hasPhysicalDevice
}
// GetManufacturer gets the manufacturer property value. The manufacturer property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The model property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetModel()(*string) {
    return m.model
}
// GetPhysicalDeviceId gets the physicalDeviceId property value. The physicalDeviceId property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) GetPhysicalDeviceId()(*string) {
    return m.physicalDeviceId
}
// Serialize serializes information the current object
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("certificateSigningRequest", m.GetCertificateSigningRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectorId", m.GetConnectorId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasPhysicalDevice", m.GetHasPhysicalDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("physicalDeviceId", m.GetPhysicalDeviceId())
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
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificateSigningRequest sets the certificateSigningRequest property value. The certificateSigningRequest property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) SetCertificateSigningRequest(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintCertificateSigningRequestable)() {
    m.certificateSigningRequest = value
}
// SetConnectorId sets the connectorId property value. The connectorId property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) SetConnectorId(value *string)() {
    m.connectorId = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHasPhysicalDevice sets the hasPhysicalDevice property value. The hasPhysicalDevice property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) SetHasPhysicalDevice(value *bool)() {
    m.hasPhysicalDevice = value
}
// SetManufacturer sets the manufacturer property value. The manufacturer property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The model property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) SetModel(value *string)() {
    m.model = value
}
// SetPhysicalDeviceId sets the physicalDeviceId property value. The physicalDeviceId property
func (m *PrintersMicrosoftGraphCreateCreatePostRequestBody) SetPhysicalDeviceId(value *string)() {
    m.physicalDeviceId = value
}
