package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppContentFile contains properties for a single installer file that is associated with a given mobileAppContent version.
type MobileAppContentFile struct {
    Entity
    // The Azure Storage URI.
    azureStorageUri *string
    // The time the Azure storage Uri expires.
    azureStorageUriExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time the file was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A value indicating whether the file is committed.
    isCommitted *bool
    // The manifest information.
    manifest []byte
    // the file name.
    name *string
    // The size of the file prior to encryption.
    size *int64
    // The size of the file after encryption.
    sizeEncrypted *int64
    // Contains properties for upload request states.
    uploadState *MobileAppContentFileUploadState
}
// NewMobileAppContentFile instantiates a new mobileAppContentFile and sets the default values.
func NewMobileAppContentFile()(*MobileAppContentFile) {
    m := &MobileAppContentFile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppContentFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppContentFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppContentFile(), nil
}
// GetAzureStorageUri gets the azureStorageUri property value. The Azure Storage URI.
func (m *MobileAppContentFile) GetAzureStorageUri()(*string) {
    return m.azureStorageUri
}
// GetAzureStorageUriExpirationDateTime gets the azureStorageUriExpirationDateTime property value. The time the Azure storage Uri expires.
func (m *MobileAppContentFile) GetAzureStorageUriExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.azureStorageUriExpirationDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. The time the file was created.
func (m *MobileAppContentFile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppContentFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureStorageUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageUri(val)
        }
        return nil
    }
    res["azureStorageUriExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageUriExpirationDateTime(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["isCommitted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCommitted(val)
        }
        return nil
    }
    res["manifest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifest(val)
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
    res["sizeEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeEncrypted(val)
        }
        return nil
    }
    res["uploadState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppContentFileUploadState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadState(val.(*MobileAppContentFileUploadState))
        }
        return nil
    }
    return res
}
// GetIsCommitted gets the isCommitted property value. A value indicating whether the file is committed.
func (m *MobileAppContentFile) GetIsCommitted()(*bool) {
    return m.isCommitted
}
// GetManifest gets the manifest property value. The manifest information.
func (m *MobileAppContentFile) GetManifest()([]byte) {
    return m.manifest
}
// GetName gets the name property value. the file name.
func (m *MobileAppContentFile) GetName()(*string) {
    return m.name
}
// GetSize gets the size property value. The size of the file prior to encryption.
func (m *MobileAppContentFile) GetSize()(*int64) {
    return m.size
}
// GetSizeEncrypted gets the sizeEncrypted property value. The size of the file after encryption.
func (m *MobileAppContentFile) GetSizeEncrypted()(*int64) {
    return m.sizeEncrypted
}
// GetUploadState gets the uploadState property value. Contains properties for upload request states.
func (m *MobileAppContentFile) GetUploadState()(*MobileAppContentFileUploadState) {
    return m.uploadState
}
// Serialize serializes information the current object
func (m *MobileAppContentFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureStorageUri", m.GetAzureStorageUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("azureStorageUriExpirationDateTime", m.GetAzureStorageUriExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCommitted", m.GetIsCommitted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("manifest", m.GetManifest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
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
    {
        err = writer.WriteInt64Value("sizeEncrypted", m.GetSizeEncrypted())
        if err != nil {
            return err
        }
    }
    if m.GetUploadState() != nil {
        cast := (*m.GetUploadState()).String()
        err = writer.WriteStringValue("uploadState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureStorageUri sets the azureStorageUri property value. The Azure Storage URI.
func (m *MobileAppContentFile) SetAzureStorageUri(value *string)() {
    m.azureStorageUri = value
}
// SetAzureStorageUriExpirationDateTime sets the azureStorageUriExpirationDateTime property value. The time the Azure storage Uri expires.
func (m *MobileAppContentFile) SetAzureStorageUriExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.azureStorageUriExpirationDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. The time the file was created.
func (m *MobileAppContentFile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetIsCommitted sets the isCommitted property value. A value indicating whether the file is committed.
func (m *MobileAppContentFile) SetIsCommitted(value *bool)() {
    m.isCommitted = value
}
// SetManifest sets the manifest property value. The manifest information.
func (m *MobileAppContentFile) SetManifest(value []byte)() {
    m.manifest = value
}
// SetName sets the name property value. the file name.
func (m *MobileAppContentFile) SetName(value *string)() {
    m.name = value
}
// SetSize sets the size property value. The size of the file prior to encryption.
func (m *MobileAppContentFile) SetSize(value *int64)() {
    m.size = value
}
// SetSizeEncrypted sets the sizeEncrypted property value. The size of the file after encryption.
func (m *MobileAppContentFile) SetSizeEncrypted(value *int64)() {
    m.sizeEncrypted = value
}
// SetUploadState sets the uploadState property value. Contains properties for upload request states.
func (m *MobileAppContentFile) SetUploadState(value *MobileAppContentFileUploadState)() {
    m.uploadState = value
}
