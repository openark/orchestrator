package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryCustodian 
type EdiscoveryCustodian struct {
    DataSourceContainer
    // Date and time the custodian acknowledged a hold notification.
    acknowledgedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Email address of the custodian.
    email *string
    // Operation entity that represents the latest indexing for the custodian.
    lastIndexOperation EdiscoveryIndexOperationable
    // Data source entity for SharePoint sites associated with the custodian.
    siteSources []SiteSourceable
    // Data source entity for groups associated with the custodian.
    unifiedGroupSources []UnifiedGroupSourceable
    // Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
    userSources []UserSourceable
}
// NewEdiscoveryCustodian instantiates a new EdiscoveryCustodian and sets the default values.
func NewEdiscoveryCustodian()(*EdiscoveryCustodian) {
    m := &EdiscoveryCustodian{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    odataTypeValue := "#microsoft.graph.security.ediscoveryCustodian"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEdiscoveryCustodianFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryCustodianFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCustodian(), nil
}
// GetAcknowledgedDateTime gets the acknowledgedDateTime property value. Date and time the custodian acknowledged a hold notification.
func (m *EdiscoveryCustodian) GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.acknowledgedDateTime
}
// GetEmail gets the email property value. Email address of the custodian.
func (m *EdiscoveryCustodian) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryCustodian) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["acknowledgedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcknowledgedDateTime(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["lastIndexOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryIndexOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastIndexOperation(val.(EdiscoveryIndexOperationable))
        }
        return nil
    }
    res["siteSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSiteSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SiteSourceable, len(val))
            for i, v := range val {
                res[i] = v.(SiteSourceable)
            }
            m.SetSiteSources(res)
        }
        return nil
    }
    res["unifiedGroupSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedGroupSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedGroupSourceable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedGroupSourceable)
            }
            m.SetUnifiedGroupSources(res)
        }
        return nil
    }
    res["userSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSourceable, len(val))
            for i, v := range val {
                res[i] = v.(UserSourceable)
            }
            m.SetUserSources(res)
        }
        return nil
    }
    return res
}
// GetLastIndexOperation gets the lastIndexOperation property value. Operation entity that represents the latest indexing for the custodian.
func (m *EdiscoveryCustodian) GetLastIndexOperation()(EdiscoveryIndexOperationable) {
    return m.lastIndexOperation
}
// GetSiteSources gets the siteSources property value. Data source entity for SharePoint sites associated with the custodian.
func (m *EdiscoveryCustodian) GetSiteSources()([]SiteSourceable) {
    return m.siteSources
}
// GetUnifiedGroupSources gets the unifiedGroupSources property value. Data source entity for groups associated with the custodian.
func (m *EdiscoveryCustodian) GetUnifiedGroupSources()([]UnifiedGroupSourceable) {
    return m.unifiedGroupSources
}
// GetUserSources gets the userSources property value. Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
func (m *EdiscoveryCustodian) GetUserSources()([]UserSourceable) {
    return m.userSources
}
// Serialize serializes information the current object
func (m *EdiscoveryCustodian) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acknowledgedDateTime", m.GetAcknowledgedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastIndexOperation", m.GetLastIndexOperation())
        if err != nil {
            return err
        }
    }
    if m.GetSiteSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSiteSources()))
        for i, v := range m.GetSiteSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("siteSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUnifiedGroupSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnifiedGroupSources()))
        for i, v := range m.GetUnifiedGroupSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("unifiedGroupSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSources()))
        for i, v := range m.GetUserSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcknowledgedDateTime sets the acknowledgedDateTime property value. Date and time the custodian acknowledged a hold notification.
func (m *EdiscoveryCustodian) SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acknowledgedDateTime = value
}
// SetEmail sets the email property value. Email address of the custodian.
func (m *EdiscoveryCustodian) SetEmail(value *string)() {
    m.email = value
}
// SetLastIndexOperation sets the lastIndexOperation property value. Operation entity that represents the latest indexing for the custodian.
func (m *EdiscoveryCustodian) SetLastIndexOperation(value EdiscoveryIndexOperationable)() {
    m.lastIndexOperation = value
}
// SetSiteSources sets the siteSources property value. Data source entity for SharePoint sites associated with the custodian.
func (m *EdiscoveryCustodian) SetSiteSources(value []SiteSourceable)() {
    m.siteSources = value
}
// SetUnifiedGroupSources sets the unifiedGroupSources property value. Data source entity for groups associated with the custodian.
func (m *EdiscoveryCustodian) SetUnifiedGroupSources(value []UnifiedGroupSourceable)() {
    m.unifiedGroupSources = value
}
// SetUserSources sets the userSources property value. Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
func (m *EdiscoveryCustodian) SetUserSources(value []UserSourceable)() {
    m.userSources = value
}
