package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintDocument 
type PrintDocument struct {
    Entity
    // The document's content (MIME) type. Read-only.
    contentType *string
    // The document's name. Read-only.
    displayName *string
    // The document's size in bytes. Read-only.
    size *int64
}
// NewPrintDocument instantiates a new printDocument and sets the default values.
func NewPrintDocument()(*PrintDocument) {
    m := &PrintDocument{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintDocumentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintDocumentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintDocument(), nil
}
// GetContentType gets the contentType property value. The document's content (MIME) type. Read-only.
func (m *PrintDocument) GetContentType()(*string) {
    return m.contentType
}
// GetDisplayName gets the displayName property value. The document's name. Read-only.
func (m *PrintDocument) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintDocument) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
// GetSize gets the size property value. The document's size in bytes. Read-only.
func (m *PrintDocument) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *PrintDocument) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
// SetContentType sets the contentType property value. The document's content (MIME) type. Read-only.
func (m *PrintDocument) SetContentType(value *string)() {
    m.contentType = value
}
// SetDisplayName sets the displayName property value. The document's name. Read-only.
func (m *PrintDocument) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetSize sets the size property value. The document's size in bytes. Read-only.
func (m *PrintDocument) SetSize(value *int64)() {
    m.size = value
}
