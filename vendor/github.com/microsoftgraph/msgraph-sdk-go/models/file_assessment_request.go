package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileAssessmentRequest 
type FileAssessmentRequest struct {
    ThreatAssessmentRequest
    // Base64 encoded file content. The file content cannot fetch back because it isn't stored.
    contentData *string
    // The file name.
    fileName *string
}
// NewFileAssessmentRequest instantiates a new FileAssessmentRequest and sets the default values.
func NewFileAssessmentRequest()(*FileAssessmentRequest) {
    m := &FileAssessmentRequest{
        ThreatAssessmentRequest: *NewThreatAssessmentRequest(),
    }
    odataTypeValue := "#microsoft.graph.fileAssessmentRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFileAssessmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileAssessmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileAssessmentRequest(), nil
}
// GetContentData gets the contentData property value. Base64 encoded file content. The file content cannot fetch back because it isn't stored.
func (m *FileAssessmentRequest) GetContentData()(*string) {
    return m.contentData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileAssessmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ThreatAssessmentRequest.GetFieldDeserializers()
    res["contentData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentData(val)
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
    return res
}
// GetFileName gets the fileName property value. The file name.
func (m *FileAssessmentRequest) GetFileName()(*string) {
    return m.fileName
}
// Serialize serializes information the current object
func (m *FileAssessmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ThreatAssessmentRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentData", m.GetContentData())
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
    return nil
}
// SetContentData sets the contentData property value. Base64 encoded file content. The file content cannot fetch back because it isn't stored.
func (m *FileAssessmentRequest) SetContentData(value *string)() {
    m.contentData = value
}
// SetFileName sets the fileName property value. The file name.
func (m *FileAssessmentRequest) SetFileName(value *string)() {
    m.fileName = value
}
