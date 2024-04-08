package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionProxiedDomainCollection windows Information Protection Proxied Domain Collection
type WindowsInformationProtectionProxiedDomainCollection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name
    displayName *string
    // The OdataType property
    odataType *string
    // Collection of proxied domains
    proxiedDomains []ProxiedDomainable
}
// NewWindowsInformationProtectionProxiedDomainCollection instantiates a new windowsInformationProtectionProxiedDomainCollection and sets the default values.
func NewWindowsInformationProtectionProxiedDomainCollection()(*WindowsInformationProtectionProxiedDomainCollection) {
    m := &WindowsInformationProtectionProxiedDomainCollection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsInformationProtectionProxiedDomainCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionProxiedDomainCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionProxiedDomainCollection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionProxiedDomainCollection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Display name
func (m *WindowsInformationProtectionProxiedDomainCollection) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionProxiedDomainCollection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["proxiedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProxiedDomainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProxiedDomainable, len(val))
            for i, v := range val {
                res[i] = v.(ProxiedDomainable)
            }
            m.SetProxiedDomains(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsInformationProtectionProxiedDomainCollection) GetOdataType()(*string) {
    return m.odataType
}
// GetProxiedDomains gets the proxiedDomains property value. Collection of proxied domains
func (m *WindowsInformationProtectionProxiedDomainCollection) GetProxiedDomains()([]ProxiedDomainable) {
    return m.proxiedDomains
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionProxiedDomainCollection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetProxiedDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProxiedDomains()))
        for i, v := range m.GetProxiedDomains() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("proxiedDomains", cast)
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
func (m *WindowsInformationProtectionProxiedDomainCollection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Display name
func (m *WindowsInformationProtectionProxiedDomainCollection) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsInformationProtectionProxiedDomainCollection) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProxiedDomains sets the proxiedDomains property value. Collection of proxied domains
func (m *WindowsInformationProtectionProxiedDomainCollection) SetProxiedDomains(value []ProxiedDomainable)() {
    m.proxiedDomains = value
}
