package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidMobileAppIdentifier 
type AndroidMobileAppIdentifier struct {
    MobileAppIdentifier
    // The identifier for an app, as specified in the play store.
    packageId *string
}
// NewAndroidMobileAppIdentifier instantiates a new AndroidMobileAppIdentifier and sets the default values.
func NewAndroidMobileAppIdentifier()(*AndroidMobileAppIdentifier) {
    m := &AndroidMobileAppIdentifier{
        MobileAppIdentifier: *NewMobileAppIdentifier(),
    }
    odataTypeValue := "#microsoft.graph.androidMobileAppIdentifier"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidMobileAppIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidMobileAppIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidMobileAppIdentifier(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidMobileAppIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppIdentifier.GetFieldDeserializers()
    res["packageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageId(val)
        }
        return nil
    }
    return res
}
// GetPackageId gets the packageId property value. The identifier for an app, as specified in the play store.
func (m *AndroidMobileAppIdentifier) GetPackageId()(*string) {
    return m.packageId
}
// Serialize serializes information the current object
func (m *AndroidMobileAppIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppIdentifier.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("packageId", m.GetPackageId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPackageId sets the packageId property value. The identifier for an app, as specified in the play store.
func (m *AndroidMobileAppIdentifier) SetPackageId(value *string)() {
    m.packageId = value
}
