package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosHomeScreenPage a page containing apps, folders, and web clips on the Home Screen.
type IosHomeScreenPage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the page
    displayName *string
    // A list of apps, folders, and web clips to appear on a page. This collection can contain a maximum of 500 elements.
    icons []IosHomeScreenItemable
    // The OdataType property
    odataType *string
}
// NewIosHomeScreenPage instantiates a new iosHomeScreenPage and sets the default values.
func NewIosHomeScreenPage()(*IosHomeScreenPage) {
    m := &IosHomeScreenPage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIosHomeScreenPageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosHomeScreenPageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosHomeScreenPage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosHomeScreenPage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Name of the page
func (m *IosHomeScreenPage) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosHomeScreenPage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["icons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosHomeScreenItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosHomeScreenItemable, len(val))
            for i, v := range val {
                res[i] = v.(IosHomeScreenItemable)
            }
            m.SetIcons(res)
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
// GetIcons gets the icons property value. A list of apps, folders, and web clips to appear on a page. This collection can contain a maximum of 500 elements.
func (m *IosHomeScreenPage) GetIcons()([]IosHomeScreenItemable) {
    return m.icons
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosHomeScreenPage) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *IosHomeScreenPage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetIcons() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIcons()))
        for i, v := range m.GetIcons() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("icons", cast)
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
func (m *IosHomeScreenPage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Name of the page
func (m *IosHomeScreenPage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIcons sets the icons property value. A list of apps, folders, and web clips to appear on a page. This collection can contain a maximum of 500 elements.
func (m *IosHomeScreenPage) SetIcons(value []IosHomeScreenItemable)() {
    m.icons = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosHomeScreenPage) SetOdataType(value *string)() {
    m.odataType = value
}
