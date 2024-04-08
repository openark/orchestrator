package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DisplayNameLocalization 
type DisplayNameLocalization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If present, the value of this field contains the displayName string that has been set for the language present in the languageTag field.
    displayName *string
    // Provides the language culture-code and friendly name of the language that the displayName field has been provided in.
    languageTag *string
    // The OdataType property
    odataType *string
}
// NewDisplayNameLocalization instantiates a new displayNameLocalization and sets the default values.
func NewDisplayNameLocalization()(*DisplayNameLocalization) {
    m := &DisplayNameLocalization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDisplayNameLocalizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDisplayNameLocalizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDisplayNameLocalization(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DisplayNameLocalization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. If present, the value of this field contains the displayName string that has been set for the language present in the languageTag field.
func (m *DisplayNameLocalization) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DisplayNameLocalization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
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
// GetLanguageTag gets the languageTag property value. Provides the language culture-code and friendly name of the language that the displayName field has been provided in.
func (m *DisplayNameLocalization) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DisplayNameLocalization) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DisplayNameLocalization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
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
func (m *DisplayNameLocalization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. If present, the value of this field contains the displayName string that has been set for the language present in the languageTag field.
func (m *DisplayNameLocalization) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLanguageTag sets the languageTag property value. Provides the language culture-code and friendly name of the language that the displayName field has been provided in.
func (m *DisplayNameLocalization) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DisplayNameLocalization) SetOdataType(value *string)() {
    m.odataType = value
}
