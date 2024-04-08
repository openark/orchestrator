package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Website 
type Website struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URL of the website.
    address *string
    // The display name of the web site.
    displayName *string
    // The OdataType property
    odataType *string
    // The possible values are: other, home, work, blog, profile.
    typeEscaped *WebsiteType
}
// NewWebsite instantiates a new website and sets the default values.
func NewWebsite()(*Website) {
    m := &Website{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebsiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebsiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebsite(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Website) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The URL of the website.
func (m *Website) GetAddress()(*string) {
    return m.address
}
// GetDisplayName gets the displayName property value. The display name of the web site.
func (m *Website) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Website) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWebsiteType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*WebsiteType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Website) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. The possible values are: other, home, work, blog, profile.
func (m *Website) GetType()(*WebsiteType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Website) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *Website) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The URL of the website.
func (m *Website) SetAddress(value *string)() {
    m.address = value
}
// SetDisplayName sets the displayName property value. The display name of the web site.
func (m *Website) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Website) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. The possible values are: other, home, work, blog, profile.
func (m *Website) SetType(value *WebsiteType)() {
    m.typeEscaped = value
}
