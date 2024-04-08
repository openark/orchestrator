package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedMobileLobApp 
type ManagedMobileLobApp struct {
    ManagedApp
    // The internal committed content version.
    committedContentVersion *string
    // The list of content versions for this app.
    contentVersions []MobileAppContentable
    // The name of the main Lob application file.
    fileName *string
    // The total size, including all uploaded files.
    size *int64
}
// NewManagedMobileLobApp instantiates a new ManagedMobileLobApp and sets the default values.
func NewManagedMobileLobApp()(*ManagedMobileLobApp) {
    m := &ManagedMobileLobApp{
        ManagedApp: *NewManagedApp(),
    }
    odataTypeValue := "#microsoft.graph.managedMobileLobApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateManagedMobileLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedMobileLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.managedAndroidLobApp":
                        return NewManagedAndroidLobApp(), nil
                    case "#microsoft.graph.managedIOSLobApp":
                        return NewManagedIOSLobApp(), nil
                }
            }
        }
    }
    return NewManagedMobileLobApp(), nil
}
// GetCommittedContentVersion gets the committedContentVersion property value. The internal committed content version.
func (m *ManagedMobileLobApp) GetCommittedContentVersion()(*string) {
    return m.committedContentVersion
}
// GetContentVersions gets the contentVersions property value. The list of content versions for this app.
func (m *ManagedMobileLobApp) GetContentVersions()([]MobileAppContentable) {
    return m.contentVersions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedMobileLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedApp.GetFieldDeserializers()
    res["committedContentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommittedContentVersion(val)
        }
        return nil
    }
    res["contentVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppContentable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppContentable)
            }
            m.SetContentVersions(res)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The name of the main Lob application file.
func (m *ManagedMobileLobApp) GetFileName()(*string) {
    return m.fileName
}
// GetSize gets the size property value. The total size, including all uploaded files.
func (m *ManagedMobileLobApp) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *ManagedMobileLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("committedContentVersion", m.GetCommittedContentVersion())
        if err != nil {
            return err
        }
    }
    if m.GetContentVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentVersions()))
        for i, v := range m.GetContentVersions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("contentVersions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCommittedContentVersion sets the committedContentVersion property value. The internal committed content version.
func (m *ManagedMobileLobApp) SetCommittedContentVersion(value *string)() {
    m.committedContentVersion = value
}
// SetContentVersions sets the contentVersions property value. The list of content versions for this app.
func (m *ManagedMobileLobApp) SetContentVersions(value []MobileAppContentable)() {
    m.contentVersions = value
}
// SetFileName sets the fileName property value. The name of the main Lob application file.
func (m *ManagedMobileLobApp) SetFileName(value *string)() {
    m.fileName = value
}
// SetSize sets the size property value. The total size, including all uploaded files.
func (m *ManagedMobileLobApp) SetSize(value *int64)() {
    m.size = value
}
