package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SiteCollection 
type SiteCollection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The geographic region code for where this site collection resides. Read-only.
    dataLocationCode *string
    // The hostname for the site collection. Read-only.
    hostname *string
    // The OdataType property
    odataType *string
    // If present, indicates that this is a root site collection in SharePoint. Read-only.
    root Rootable
}
// NewSiteCollection instantiates a new siteCollection and sets the default values.
func NewSiteCollection()(*SiteCollection) {
    m := &SiteCollection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSiteCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSiteCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSiteCollection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SiteCollection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDataLocationCode gets the dataLocationCode property value. The geographic region code for where this site collection resides. Read-only.
func (m *SiteCollection) GetDataLocationCode()(*string) {
    return m.dataLocationCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SiteCollection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dataLocationCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataLocationCode(val)
        }
        return nil
    }
    res["hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val)
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
    res["root"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(Rootable))
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. The hostname for the site collection. Read-only.
func (m *SiteCollection) GetHostname()(*string) {
    return m.hostname
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SiteCollection) GetOdataType()(*string) {
    return m.odataType
}
// GetRoot gets the root property value. If present, indicates that this is a root site collection in SharePoint. Read-only.
func (m *SiteCollection) GetRoot()(Rootable) {
    return m.root
}
// Serialize serializes information the current object
func (m *SiteCollection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dataLocationCode", m.GetDataLocationCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
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
        err := writer.WriteObjectValue("root", m.GetRoot())
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
func (m *SiteCollection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDataLocationCode sets the dataLocationCode property value. The geographic region code for where this site collection resides. Read-only.
func (m *SiteCollection) SetDataLocationCode(value *string)() {
    m.dataLocationCode = value
}
// SetHostname sets the hostname property value. The hostname for the site collection. Read-only.
func (m *SiteCollection) SetHostname(value *string)() {
    m.hostname = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SiteCollection) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRoot sets the root property value. If present, indicates that this is a root site collection in SharePoint. Read-only.
func (m *SiteCollection) SetRoot(value Rootable)() {
    m.root = value
}
