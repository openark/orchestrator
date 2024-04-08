package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsApp 
type TeamsApp struct {
    Entity
    // The details for each version of the app.
    appDefinitions []TeamsAppDefinitionable
    // The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
    displayName *string
    // The method of distribution for the app. Read-only.
    distributionMethod *TeamsAppDistributionMethod
    // The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
    externalId *string
}
// NewTeamsApp instantiates a new teamsApp and sets the default values.
func NewTeamsApp()(*TeamsApp) {
    m := &TeamsApp{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsApp(), nil
}
// GetAppDefinitions gets the appDefinitions property value. The details for each version of the app.
func (m *TeamsApp) GetAppDefinitions()([]TeamsAppDefinitionable) {
    return m.appDefinitions
}
// GetDisplayName gets the displayName property value. The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
func (m *TeamsApp) GetDisplayName()(*string) {
    return m.displayName
}
// GetDistributionMethod gets the distributionMethod property value. The method of distribution for the app. Read-only.
func (m *TeamsApp) GetDistributionMethod()(*TeamsAppDistributionMethod) {
    return m.distributionMethod
}
// GetExternalId gets the externalId property value. The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
func (m *TeamsApp) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamsAppDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAppDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(TeamsAppDefinitionable)
            }
            m.SetAppDefinitions(res)
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
    res["distributionMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppDistributionMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDistributionMethod(val.(*TeamsAppDistributionMethod))
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeamsApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppDefinitions()))
        for i, v := range m.GetAppDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appDefinitions", cast)
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
    if m.GetDistributionMethod() != nil {
        cast := (*m.GetDistributionMethod()).String()
        err = writer.WriteStringValue("distributionMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDefinitions sets the appDefinitions property value. The details for each version of the app.
func (m *TeamsApp) SetAppDefinitions(value []TeamsAppDefinitionable)() {
    m.appDefinitions = value
}
// SetDisplayName sets the displayName property value. The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
func (m *TeamsApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDistributionMethod sets the distributionMethod property value. The method of distribution for the app. Read-only.
func (m *TeamsApp) SetDistributionMethod(value *TeamsAppDistributionMethod)() {
    m.distributionMethod = value
}
// SetExternalId sets the externalId property value. The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
func (m *TeamsApp) SetExternalId(value *string)() {
    m.externalId = value
}
