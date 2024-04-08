package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharingDetail 
type SharingDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The user who shared the document.
    sharedBy InsightIdentityable
    // The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    sharedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The sharingReference property
    sharingReference ResourceReferenceable
    // The subject with which the document was shared.
    sharingSubject *string
    // Determines the way the document was shared, can be by a 'Link', 'Attachment', 'Group', 'Site'.
    sharingType *string
}
// NewSharingDetail instantiates a new sharingDetail and sets the default values.
func NewSharingDetail()(*SharingDetail) {
    m := &SharingDetail{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSharingDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharingDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharingDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharingDetail) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharingDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["sharedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedBy(val.(InsightIdentityable))
        }
        return nil
    }
    res["sharedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedDateTime(val)
        }
        return nil
    }
    res["sharingReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharingReference(val.(ResourceReferenceable))
        }
        return nil
    }
    res["sharingSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharingSubject(val)
        }
        return nil
    }
    res["sharingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharingType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SharingDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetSharedBy gets the sharedBy property value. The user who shared the document.
func (m *SharingDetail) GetSharedBy()(InsightIdentityable) {
    return m.sharedBy
}
// GetSharedDateTime gets the sharedDateTime property value. The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *SharingDetail) GetSharedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.sharedDateTime
}
// GetSharingReference gets the sharingReference property value. The sharingReference property
func (m *SharingDetail) GetSharingReference()(ResourceReferenceable) {
    return m.sharingReference
}
// GetSharingSubject gets the sharingSubject property value. The subject with which the document was shared.
func (m *SharingDetail) GetSharingSubject()(*string) {
    return m.sharingSubject
}
// GetSharingType gets the sharingType property value. Determines the way the document was shared, can be by a 'Link', 'Attachment', 'Group', 'Site'.
func (m *SharingDetail) GetSharingType()(*string) {
    return m.sharingType
}
// Serialize serializes information the current object
func (m *SharingDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sharedBy", m.GetSharedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("sharedDateTime", m.GetSharedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sharingSubject", m.GetSharingSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sharingType", m.GetSharingType())
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
func (m *SharingDetail) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SharingDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSharedBy sets the sharedBy property value. The user who shared the document.
func (m *SharingDetail) SetSharedBy(value InsightIdentityable)() {
    m.sharedBy = value
}
// SetSharedDateTime sets the sharedDateTime property value. The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *SharingDetail) SetSharedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sharedDateTime = value
}
// SetSharingReference sets the sharingReference property value. The sharingReference property
func (m *SharingDetail) SetSharingReference(value ResourceReferenceable)() {
    m.sharingReference = value
}
// SetSharingSubject sets the sharingSubject property value. The subject with which the document was shared.
func (m *SharingDetail) SetSharingSubject(value *string)() {
    m.sharingSubject = value
}
// SetSharingType sets the sharingType property value. Determines the way the document was shared, can be by a 'Link', 'Attachment', 'Group', 'Site'.
func (m *SharingDetail) SetSharingType(value *string)() {
    m.sharingType = value
}
