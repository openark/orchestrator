package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OcrSettings 
type OcrSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether or not OCR is enabled for the case.
    isEnabled *bool
    // Maximum image size that will be processed in KB).
    maxImageSize *int32
    // The OdataType property
    odataType *string
    // The timeout duration for the OCR engine. A longer timeout might increase success of OCR, but might add to the total processing time.
    timeout *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewOcrSettings instantiates a new ocrSettings and sets the default values.
func NewOcrSettings()(*OcrSettings) {
    m := &OcrSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOcrSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOcrSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOcrSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OcrSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OcrSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["maxImageSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxImageSize(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["timeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeout(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Indicates whether or not OCR is enabled for the case.
func (m *OcrSettings) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetMaxImageSize gets the maxImageSize property value. Maximum image size that will be processed in KB).
func (m *OcrSettings) GetMaxImageSize()(*int32) {
    return m.maxImageSize
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OcrSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetTimeout gets the timeout property value. The timeout duration for the OCR engine. A longer timeout might increase success of OCR, but might add to the total processing time.
func (m *OcrSettings) GetTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.timeout
}
// Serialize serializes information the current object
func (m *OcrSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxImageSize", m.GetMaxImageSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("timeout", m.GetTimeout())
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
func (m *OcrSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether or not OCR is enabled for the case.
func (m *OcrSettings) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetMaxImageSize sets the maxImageSize property value. Maximum image size that will be processed in KB).
func (m *OcrSettings) SetMaxImageSize(value *int32)() {
    m.maxImageSize = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OcrSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTimeout sets the timeout property value. The timeout duration for the OCR engine. A longer timeout might increase success of OCR, but might add to the total processing time.
func (m *OcrSettings) SetTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.timeout = value
}
