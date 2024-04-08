package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageCatalog 
type AccessPackageCatalog struct {
    Entity
    // The access packages in this catalog. Read-only. Nullable.
    accessPackages []AccessPackageable
    // Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
    catalogType *AccessPackageCatalogType
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of the access package catalog.
    description *string
    // The display name of the access package catalog.
    displayName *string
    // Whether the access packages in this catalog can be requested by users outside of the tenant.
    isExternallyVisible *bool
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Has the value published if the access packages are available for management. The possible values are: unpublished, published, unknownFutureValue.
    state *AccessPackageCatalogState
}
// NewAccessPackageCatalog instantiates a new accessPackageCatalog and sets the default values.
func NewAccessPackageCatalog()(*AccessPackageCatalog) {
    m := &AccessPackageCatalog{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageCatalogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageCatalog(), nil
}
// GetAccessPackages gets the accessPackages property value. The access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageCatalog) GetAccessPackages()([]AccessPackageable) {
    return m.accessPackages
}
// GetCatalogType gets the catalogType property value. Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
func (m *AccessPackageCatalog) GetCatalogType()(*AccessPackageCatalogType) {
    return m.catalogType
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the access package catalog.
func (m *AccessPackageCatalog) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageCatalog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageable)
            }
            m.SetAccessPackages(res)
        }
        return nil
    }
    res["catalogType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCatalogType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogType(val.(*AccessPackageCatalogType))
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["isExternallyVisible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternallyVisible(val)
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCatalogState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AccessPackageCatalogState))
        }
        return nil
    }
    return res
}
// GetIsExternallyVisible gets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) GetIsExternallyVisible()(*bool) {
    return m.isExternallyVisible
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetState gets the state property value. Has the value published if the access packages are available for management. The possible values are: unpublished, published, unknownFutureValue.
func (m *AccessPackageCatalog) GetState()(*AccessPackageCatalogState) {
    return m.state
}
// Serialize serializes information the current object
func (m *AccessPackageCatalog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCatalogType() != nil {
        cast := (*m.GetCatalogType()).String()
        err = writer.WriteStringValue("catalogType", &cast)
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
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteBoolValue("isExternallyVisible", m.GetIsExternallyVisible())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackages sets the accessPackages property value. The access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageCatalog) SetAccessPackages(value []AccessPackageable)() {
    m.accessPackages = value
}
// SetCatalogType sets the catalogType property value. Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue.
func (m *AccessPackageCatalog) SetCatalogType(value *AccessPackageCatalogType)() {
    m.catalogType = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of the access package catalog.
func (m *AccessPackageCatalog) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the access package catalog.
func (m *AccessPackageCatalog) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsExternallyVisible sets the isExternallyVisible property value. Whether the access packages in this catalog can be requested by users outside of the tenant.
func (m *AccessPackageCatalog) SetIsExternallyVisible(value *bool)() {
    m.isExternallyVisible = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageCatalog) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetState sets the state property value. Has the value published if the access packages are available for management. The possible values are: unpublished, published, unknownFutureValue.
func (m *AccessPackageCatalog) SetState(value *AccessPackageCatalogState)() {
    m.state = value
}
