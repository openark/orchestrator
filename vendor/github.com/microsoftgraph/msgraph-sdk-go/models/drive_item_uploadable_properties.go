package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveItemUploadableProperties 
type DriveItemUploadableProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
    description *string
    // Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
    fileSize *int64
    // File system information on client. Read-write.
    fileSystemInfo FileSystemInfoable
    // The name of the item (filename and extension). Read-write.
    name *string
    // The OdataType property
    odataType *string
}
// NewDriveItemUploadableProperties instantiates a new driveItemUploadableProperties and sets the default values.
func NewDriveItemUploadableProperties()(*DriveItemUploadableProperties) {
    m := &DriveItemUploadableProperties{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDriveItemUploadablePropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveItemUploadablePropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveItemUploadableProperties(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveItemUploadableProperties) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItemUploadableProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["fileSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSize(val)
        }
        return nil
    }
    res["fileSystemInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSystemInfo(val.(FileSystemInfoable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetFileSize gets the fileSize property value. Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) GetFileSize()(*int64) {
    return m.fileSize
}
// GetFileSystemInfo gets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItemUploadableProperties) GetFileSystemInfo()(FileSystemInfoable) {
    return m.fileSystemInfo
}
// GetName gets the name property value. The name of the item (filename and extension). Read-write.
func (m *DriveItemUploadableProperties) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DriveItemUploadableProperties) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DriveItemUploadableProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("fileSize", m.GetFileSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fileSystemInfo", m.GetFileSystemInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *DriveItemUploadableProperties) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) SetDescription(value *string)() {
    m.description = value
}
// SetFileSize sets the fileSize property value. Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) SetFileSize(value *int64)() {
    m.fileSize = value
}
// SetFileSystemInfo sets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItemUploadableProperties) SetFileSystemInfo(value FileSystemInfoable)() {
    m.fileSystemInfo = value
}
// SetName sets the name property value. The name of the item (filename and extension). Read-write.
func (m *DriveItemUploadableProperties) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DriveItemUploadableProperties) SetOdataType(value *string)() {
    m.odataType = value
}
