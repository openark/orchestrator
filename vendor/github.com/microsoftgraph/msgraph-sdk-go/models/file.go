package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// File 
type File struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Hashes of the file's binary content, if available. Read-only.
    hashes Hashesable
    // The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
    mimeType *string
    // The OdataType property
    odataType *string
    // The processingMetadata property
    processingMetadata *bool
}
// NewFile instantiates a new file and sets the default values.
func NewFile()(*File) {
    m := &File{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *File) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *File) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hashes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHashesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashes(val.(Hashesable))
        }
        return nil
    }
    res["mimeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMimeType(val)
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
    res["processingMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingMetadata(val)
        }
        return nil
    }
    return res
}
// GetHashes gets the hashes property value. Hashes of the file's binary content, if available. Read-only.
func (m *File) GetHashes()(Hashesable) {
    return m.hashes
}
// GetMimeType gets the mimeType property value. The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
func (m *File) GetMimeType()(*string) {
    return m.mimeType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *File) GetOdataType()(*string) {
    return m.odataType
}
// GetProcessingMetadata gets the processingMetadata property value. The processingMetadata property
func (m *File) GetProcessingMetadata()(*bool) {
    return m.processingMetadata
}
// Serialize serializes information the current object
func (m *File) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("hashes", m.GetHashes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mimeType", m.GetMimeType())
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
        err := writer.WriteBoolValue("processingMetadata", m.GetProcessingMetadata())
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
func (m *File) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHashes sets the hashes property value. Hashes of the file's binary content, if available. Read-only.
func (m *File) SetHashes(value Hashesable)() {
    m.hashes = value
}
// SetMimeType sets the mimeType property value. The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
func (m *File) SetMimeType(value *string)() {
    m.mimeType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *File) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProcessingMetadata sets the processingMetadata property value. The processingMetadata property
func (m *File) SetProcessingMetadata(value *bool)() {
    m.processingMetadata = value
}
