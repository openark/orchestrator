package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocaleInfo 
type LocaleInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A name representing the user's locale in natural language, for example, 'English (United States)'.
    displayName *string
    // A locale representation for the user, which includes the user's preferred language and country/region. For example, 'en-us'. The language component follows 2-letter codes as defined in ISO 639-1, and the country component follows 2-letter codes as defined in ISO 3166-1 alpha-2.
    locale *string
    // The OdataType property
    odataType *string
}
// NewLocaleInfo instantiates a new localeInfo and sets the default values.
func NewLocaleInfo()(*LocaleInfo) {
    m := &LocaleInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLocaleInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocaleInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocaleInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocaleInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. A name representing the user's locale in natural language, for example, 'English (United States)'.
func (m *LocaleInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocaleInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
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
    return res
}
// GetLocale gets the locale property value. A locale representation for the user, which includes the user's preferred language and country/region. For example, 'en-us'. The language component follows 2-letter codes as defined in ISO 639-1, and the country component follows 2-letter codes as defined in ISO 3166-1 alpha-2.
func (m *LocaleInfo) GetLocale()(*string) {
    return m.locale
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LocaleInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LocaleInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocaleInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. A name representing the user's locale in natural language, for example, 'English (United States)'.
func (m *LocaleInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLocale sets the locale property value. A locale representation for the user, which includes the user's preferred language and country/region. For example, 'en-us'. The language component follows 2-letter codes as defined in ISO 639-1, and the country component follows 2-letter codes as defined in ISO 3166-1 alpha-2.
func (m *LocaleInfo) SetLocale(value *string)() {
    m.locale = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LocaleInfo) SetOdataType(value *string)() {
    m.odataType = value
}
