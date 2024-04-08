package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadSession 
type UploadSession struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date and time in UTC that the upload session will expire. The complete file must be uploaded before this expiration time is reached.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of byte ranges that the server is missing for the file. These ranges are zero indexed and of the format 'start-end' (e.g. '0-26' to indicate the first 27 bytes of the file). When uploading files as Outlook attachments, instead of a collection of ranges, this property always indicates a single value '{start}', the location in the file where the next upload should begin.
    nextExpectedRanges []string
    // The OdataType property
    odataType *string
    // The URL endpoint that accepts PUT requests for byte ranges of the file.
    uploadUrl *string
}
// NewUploadSession instantiates a new uploadSession and sets the default values.
func NewUploadSession()(*UploadSession) {
    m := &UploadSession{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUploadSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUploadSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUploadSession(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadSession) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpirationDateTime gets the expirationDateTime property value. The date and time in UTC that the upload session will expire. The complete file must be uploaded before this expiration time is reached.
func (m *UploadSession) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UploadSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["nextExpectedRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNextExpectedRanges(res)
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
    res["uploadUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadUrl(val)
        }
        return nil
    }
    return res
}
// GetNextExpectedRanges gets the nextExpectedRanges property value. A collection of byte ranges that the server is missing for the file. These ranges are zero indexed and of the format 'start-end' (e.g. '0-26' to indicate the first 27 bytes of the file). When uploading files as Outlook attachments, instead of a collection of ranges, this property always indicates a single value '{start}', the location in the file where the next upload should begin.
func (m *UploadSession) GetNextExpectedRanges()([]string) {
    return m.nextExpectedRanges
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UploadSession) GetOdataType()(*string) {
    return m.odataType
}
// GetUploadUrl gets the uploadUrl property value. The URL endpoint that accepts PUT requests for byte ranges of the file.
func (m *UploadSession) GetUploadUrl()(*string) {
    return m.uploadUrl
}
// Serialize serializes information the current object
func (m *UploadSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetNextExpectedRanges() != nil {
        err := writer.WriteCollectionOfStringValues("nextExpectedRanges", m.GetNextExpectedRanges())
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
        err := writer.WriteStringValue("uploadUrl", m.GetUploadUrl())
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
func (m *UploadSession) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The date and time in UTC that the upload session will expire. The complete file must be uploaded before this expiration time is reached.
func (m *UploadSession) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetNextExpectedRanges sets the nextExpectedRanges property value. A collection of byte ranges that the server is missing for the file. These ranges are zero indexed and of the format 'start-end' (e.g. '0-26' to indicate the first 27 bytes of the file). When uploading files as Outlook attachments, instead of a collection of ranges, this property always indicates a single value '{start}', the location in the file where the next upload should begin.
func (m *UploadSession) SetNextExpectedRanges(value []string)() {
    m.nextExpectedRanges = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UploadSession) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUploadUrl sets the uploadUrl property value. The URL endpoint that accepts PUT requests for byte ranges of the file.
func (m *UploadSession) SetUploadUrl(value *string)() {
    m.uploadUrl = value
}
