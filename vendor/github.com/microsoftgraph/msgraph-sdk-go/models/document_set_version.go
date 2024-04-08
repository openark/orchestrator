package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentSetVersion 
type DocumentSetVersion struct {
    ListItemVersion
    // Comment about the captured version.
    comment *string
    // User who captured the version.
    createdBy IdentitySetable
    // Date and time when this version was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Items within the document set that are captured as part of this version.
    items []DocumentSetVersionItemable
    // If true, minor versions of items are also captured; otherwise, only major versions will be captured. Default value is false.
    shouldCaptureMinorVersion *bool
}
// NewDocumentSetVersion instantiates a new DocumentSetVersion and sets the default values.
func NewDocumentSetVersion()(*DocumentSetVersion) {
    m := &DocumentSetVersion{
        ListItemVersion: *NewListItemVersion(),
    }
    odataTypeValue := "#microsoft.graph.documentSetVersion"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDocumentSetVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentSetVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentSetVersion(), nil
}
// GetComment gets the comment property value. Comment about the captured version.
func (m *DocumentSetVersion) GetComment()(*string) {
    return m.comment
}
// GetCreatedBy gets the createdBy property value. User who captured the version.
func (m *DocumentSetVersion) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time when this version was created.
func (m *DocumentSetVersion) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentSetVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ListItemVersion.GetFieldDeserializers()
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
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
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDocumentSetVersionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DocumentSetVersionItemable, len(val))
            for i, v := range val {
                res[i] = v.(DocumentSetVersionItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["shouldCaptureMinorVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldCaptureMinorVersion(val)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. Items within the document set that are captured as part of this version.
func (m *DocumentSetVersion) GetItems()([]DocumentSetVersionItemable) {
    return m.items
}
// GetShouldCaptureMinorVersion gets the shouldCaptureMinorVersion property value. If true, minor versions of items are also captured; otherwise, only major versions will be captured. Default value is false.
func (m *DocumentSetVersion) GetShouldCaptureMinorVersion()(*bool) {
    return m.shouldCaptureMinorVersion
}
// Serialize serializes information the current object
func (m *DocumentSetVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ListItemVersion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shouldCaptureMinorVersion", m.GetShouldCaptureMinorVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComment sets the comment property value. Comment about the captured version.
func (m *DocumentSetVersion) SetComment(value *string)() {
    m.comment = value
}
// SetCreatedBy sets the createdBy property value. User who captured the version.
func (m *DocumentSetVersion) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when this version was created.
func (m *DocumentSetVersion) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetItems sets the items property value. Items within the document set that are captured as part of this version.
func (m *DocumentSetVersion) SetItems(value []DocumentSetVersionItemable)() {
    m.items = value
}
// SetShouldCaptureMinorVersion sets the shouldCaptureMinorVersion property value. If true, minor versions of items are also captured; otherwise, only major versions will be captured. Default value is false.
func (m *DocumentSetVersion) SetShouldCaptureMinorVersion(value *bool)() {
    m.shouldCaptureMinorVersion = value
}
