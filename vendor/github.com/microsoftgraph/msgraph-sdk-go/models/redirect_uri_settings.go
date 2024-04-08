package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RedirectUriSettings 
type RedirectUriSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The index property
    index *int32
    // The OdataType property
    odataType *string
    // The uri property
    uri *string
}
// NewRedirectUriSettings instantiates a new redirectUriSettings and sets the default values.
func NewRedirectUriSettings()(*RedirectUriSettings) {
    m := &RedirectUriSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRedirectUriSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRedirectUriSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRedirectUriSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectUriSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedirectUriSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
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
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. The index property
func (m *RedirectUriSettings) GetIndex()(*int32) {
    return m.index
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RedirectUriSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetUri gets the uri property value. The uri property
func (m *RedirectUriSettings) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *RedirectUriSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
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
        err := writer.WriteStringValue("uri", m.GetUri())
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
func (m *RedirectUriSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIndex sets the index property value. The index property
func (m *RedirectUriSettings) SetIndex(value *int32)() {
    m.index = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RedirectUriSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUri sets the uri property value. The uri property
func (m *RedirectUriSettings) SetUri(value *string)() {
    m.uri = value
}
