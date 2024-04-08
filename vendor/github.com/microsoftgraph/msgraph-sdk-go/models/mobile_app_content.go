package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppContent contains content properties for a specific app version. Each mobileAppContent can have multiple mobileAppContentFile.
type MobileAppContent struct {
    Entity
    // The collection of contained apps in a MobileLobApp acting as a package.
    containedApps []MobileContainedAppable
    // The list of files for this app content version.
    files []MobileAppContentFileable
}
// NewMobileAppContent instantiates a new mobileAppContent and sets the default values.
func NewMobileAppContent()(*MobileAppContent) {
    m := &MobileAppContent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppContent(), nil
}
// GetContainedApps gets the containedApps property value. The collection of contained apps in a MobileLobApp acting as a package.
func (m *MobileAppContent) GetContainedApps()([]MobileContainedAppable) {
    return m.containedApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["containedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileContainedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileContainedAppable, len(val))
            for i, v := range val {
                res[i] = v.(MobileContainedAppable)
            }
            m.SetContainedApps(res)
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppContentFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppContentFileable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppContentFileable)
            }
            m.SetFiles(res)
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. The list of files for this app content version.
func (m *MobileAppContent) GetFiles()([]MobileAppContentFileable) {
    return m.files
}
// Serialize serializes information the current object
func (m *MobileAppContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContainedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContainedApps()))
        for i, v := range m.GetContainedApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("containedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("files", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainedApps sets the containedApps property value. The collection of contained apps in a MobileLobApp acting as a package.
func (m *MobileAppContent) SetContainedApps(value []MobileContainedAppable)() {
    m.containedApps = value
}
// SetFiles sets the files property value. The list of files for this app content version.
func (m *MobileAppContent) SetFiles(value []MobileAppContentFileable)() {
    m.files = value
}
