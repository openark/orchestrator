package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentSetContent 
type DocumentSetContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Content type information of the file.
    contentType ContentTypeInfoable
    // Name of the file in resource folder that should be added as a default content or a template in the document set.
    fileName *string
    // Folder name in which the file will be placed when a new document set is created in the library.
    folderName *string
    // The OdataType property
    odataType *string
}
// NewDocumentSetContent instantiates a new documentSetContent and sets the default values.
func NewDocumentSetContent()(*DocumentSetContent) {
    m := &DocumentSetContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDocumentSetContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentSetContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentSetContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DocumentSetContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentType gets the contentType property value. Content type information of the file.
func (m *DocumentSetContent) GetContentType()(ContentTypeInfoable) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentSetContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentTypeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(ContentTypeInfoable))
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
    res["folderName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderName(val)
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
// GetFileName gets the fileName property value. Name of the file in resource folder that should be added as a default content or a template in the document set.
func (m *DocumentSetContent) GetFileName()(*string) {
    return m.fileName
}
// GetFolderName gets the folderName property value. Folder name in which the file will be placed when a new document set is created in the library.
func (m *DocumentSetContent) GetFolderName()(*string) {
    return m.folderName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DocumentSetContent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DocumentSetContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("folderName", m.GetFolderName())
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
func (m *DocumentSetContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentType sets the contentType property value. Content type information of the file.
func (m *DocumentSetContent) SetContentType(value ContentTypeInfoable)() {
    m.contentType = value
}
// SetFileName sets the fileName property value. Name of the file in resource folder that should be added as a default content or a template in the document set.
func (m *DocumentSetContent) SetFileName(value *string)() {
    m.fileName = value
}
// SetFolderName sets the folderName property value. Folder name in which the file will be placed when a new document set is created in the library.
func (m *DocumentSetContent) SetFolderName(value *string)() {
    m.folderName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DocumentSetContent) SetOdataType(value *string)() {
    m.odataType = value
}
