package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveItemVersion 
type DriveItemVersion struct {
    BaseItemVersion
    // The content stream for this version of the item.
    content []byte
    // Indicates the size of the content stream for this version of the item.
    size *int64
}
// NewDriveItemVersion instantiates a new DriveItemVersion and sets the default values.
func NewDriveItemVersion()(*DriveItemVersion) {
    m := &DriveItemVersion{
        BaseItemVersion: *NewBaseItemVersion(),
    }
    odataTypeValue := "#microsoft.graph.driveItemVersion"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDriveItemVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveItemVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveItemVersion(), nil
}
// GetContent gets the content property value. The content stream for this version of the item.
func (m *DriveItemVersion) GetContent()([]byte) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItemVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItemVersion.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
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
// GetSize gets the size property value. Indicates the size of the content stream for this version of the item.
func (m *DriveItemVersion) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *DriveItemVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItemVersion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
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
// SetContent sets the content property value. The content stream for this version of the item.
func (m *DriveItemVersion) SetContent(value []byte)() {
    m.content = value
}
// SetSize sets the size property value. Indicates the size of the content stream for this version of the item.
func (m *DriveItemVersion) SetSize(value *int64)() {
    m.size = value
}
