package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// DataSourceContainer 
type DataSourceContainer struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
    // Created date and time of the dataSourceContainer entity.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Display name of the dataSourceContainer entity.
    displayName *string
    // The hold status of the dataSourceContainer. The possible values are: notApplied, applied, applying, removing, partial
    holdStatus *DataSourceHoldStatus
    // Last modified date and time of the dataSourceContainer.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time that the dataSourceContainer was released from the case.
    releasedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Latest status of the dataSourceContainer. Possible values are: Active, Released.
    status *DataSourceContainerStatus
}
// NewDataSourceContainer instantiates a new dataSourceContainer and sets the default values.
func NewDataSourceContainer()(*DataSourceContainer) {
    m := &DataSourceContainer{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateDataSourceContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataSourceContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.ediscoveryCustodian":
                        return NewEdiscoveryCustodian(), nil
                    case "#microsoft.graph.security.ediscoveryNoncustodialDataSource":
                        return NewEdiscoveryNoncustodialDataSource(), nil
                }
            }
        }
    }
    return NewDataSourceContainer(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Created date and time of the dataSourceContainer entity.
func (m *DataSourceContainer) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. Display name of the dataSourceContainer entity.
func (m *DataSourceContainer) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataSourceContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["holdStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSourceHoldStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHoldStatus(val.(*DataSourceHoldStatus))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["releasedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasedDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSourceContainerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DataSourceContainerStatus))
        }
        return nil
    }
    return res
}
// GetHoldStatus gets the holdStatus property value. The hold status of the dataSourceContainer. The possible values are: notApplied, applied, applying, removing, partial
func (m *DataSourceContainer) GetHoldStatus()(*DataSourceHoldStatus) {
    return m.holdStatus
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified date and time of the dataSourceContainer.
func (m *DataSourceContainer) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetReleasedDateTime gets the releasedDateTime property value. Date and time that the dataSourceContainer was released from the case.
func (m *DataSourceContainer) GetReleasedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.releasedDateTime
}
// GetStatus gets the status property value. Latest status of the dataSourceContainer. Possible values are: Active, Released.
func (m *DataSourceContainer) GetStatus()(*DataSourceContainerStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DataSourceContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    if m.GetHoldStatus() != nil {
        cast := (*m.GetHoldStatus()).String()
        err = writer.WriteStringValue("holdStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releasedDateTime", m.GetReleasedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Created date and time of the dataSourceContainer entity.
func (m *DataSourceContainer) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. Display name of the dataSourceContainer entity.
func (m *DataSourceContainer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHoldStatus sets the holdStatus property value. The hold status of the dataSourceContainer. The possible values are: notApplied, applied, applying, removing, partial
func (m *DataSourceContainer) SetHoldStatus(value *DataSourceHoldStatus)() {
    m.holdStatus = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified date and time of the dataSourceContainer.
func (m *DataSourceContainer) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetReleasedDateTime sets the releasedDateTime property value. Date and time that the dataSourceContainer was released from the case.
func (m *DataSourceContainer) SetReleasedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.releasedDateTime = value
}
// SetStatus sets the status property value. Latest status of the dataSourceContainer. Possible values are: Active, Released.
func (m *DataSourceContainer) SetStatus(value *DataSourceContainerStatus)() {
    m.status = value
}
