package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalBranding 
type OrganizationalBranding struct {
    OrganizationalBrandingProperties
    // Add different branding based on a locale.
    localizations []OrganizationalBrandingLocalizationable
}
// NewOrganizationalBranding instantiates a new OrganizationalBranding and sets the default values.
func NewOrganizationalBranding()(*OrganizationalBranding) {
    m := &OrganizationalBranding{
        OrganizationalBrandingProperties: *NewOrganizationalBrandingProperties(),
    }
    odataTypeValue := "#microsoft.graph.organizationalBranding"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOrganizationalBrandingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalBrandingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalBranding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalBranding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OrganizationalBrandingProperties.GetFieldDeserializers()
    res["localizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationalBrandingLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationalBrandingLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(OrganizationalBrandingLocalizationable)
            }
            m.SetLocalizations(res)
        }
        return nil
    }
    return res
}
// GetLocalizations gets the localizations property value. Add different branding based on a locale.
func (m *OrganizationalBranding) GetLocalizations()([]OrganizationalBrandingLocalizationable) {
    return m.localizations
}
// Serialize serializes information the current object
func (m *OrganizationalBranding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OrganizationalBrandingProperties.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLocalizations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("localizations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLocalizations sets the localizations property value. Add different branding based on a locale.
func (m *OrganizationalBranding) SetLocalizations(value []OrganizationalBrandingLocalizationable)() {
    m.localizations = value
}
