package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosHomeScreenFolderPage a page for a folder containing apps and web clips on the Home Screen.
type IosHomeScreenFolderPage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A list of apps and web clips to appear on a page within a folder. This collection can contain a maximum of 500 elements.
    apps []IosHomeScreenAppable
    // Name of the folder page
    displayName *string
    // The OdataType property
    odataType *string
}
// NewIosHomeScreenFolderPage instantiates a new iosHomeScreenFolderPage and sets the default values.
func NewIosHomeScreenFolderPage()(*IosHomeScreenFolderPage) {
    m := &IosHomeScreenFolderPage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIosHomeScreenFolderPageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosHomeScreenFolderPageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosHomeScreenFolderPage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosHomeScreenFolderPage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApps gets the apps property value. A list of apps and web clips to appear on a page within a folder. This collection can contain a maximum of 500 elements.
func (m *IosHomeScreenFolderPage) GetApps()([]IosHomeScreenAppable) {
    return m.apps
}
// GetDisplayName gets the displayName property value. Name of the folder page
func (m *IosHomeScreenFolderPage) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosHomeScreenFolderPage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosHomeScreenAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosHomeScreenAppable, len(val))
            for i, v := range val {
                res[i] = v.(IosHomeScreenAppable)
            }
            m.SetApps(res)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosHomeScreenFolderPage) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *IosHomeScreenFolderPage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("apps", cast)
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
func (m *IosHomeScreenFolderPage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApps sets the apps property value. A list of apps and web clips to appear on a page within a folder. This collection can contain a maximum of 500 elements.
func (m *IosHomeScreenFolderPage) SetApps(value []IosHomeScreenAppable)() {
    m.apps = value
}
// SetDisplayName sets the displayName property value. Name of the folder page
func (m *IosHomeScreenFolderPage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosHomeScreenFolderPage) SetOdataType(value *string)() {
    m.odataType = value
}
