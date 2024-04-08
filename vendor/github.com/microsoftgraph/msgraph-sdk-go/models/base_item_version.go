package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseItemVersion 
type BaseItemVersion struct {
    Entity
    // Identity of the user which last modified the version. Read-only.
    lastModifiedBy IdentitySetable
    // Date and time the version was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates the publication status of this particular version. Read-only.
    publication PublicationFacetable
}
// NewBaseItemVersion instantiates a new baseItemVersion and sets the default values.
func NewBaseItemVersion()(*BaseItemVersion) {
    m := &BaseItemVersion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBaseItemVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBaseItemVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.documentSetVersion":
                        return NewDocumentSetVersion(), nil
                    case "#microsoft.graph.driveItemVersion":
                        return NewDriveItemVersion(), nil
                    case "#microsoft.graph.listItemVersion":
                        return NewListItemVersion(), nil
                }
            }
        }
    }
    return NewBaseItemVersion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseItemVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
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
    res["publication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicationFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublication(val.(PublicationFacetable))
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user which last modified the version. Read-only.
func (m *BaseItemVersion) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time the version was last modified. Read-only.
func (m *BaseItemVersion) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPublication gets the publication property value. Indicates the publication status of this particular version. Read-only.
func (m *BaseItemVersion) GetPublication()(PublicationFacetable) {
    return m.publication
}
// Serialize serializes information the current object
func (m *BaseItemVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
        err = writer.WriteObjectValue("publication", m.GetPublication())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user which last modified the version. Read-only.
func (m *BaseItemVersion) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time the version was last modified. Read-only.
func (m *BaseItemVersion) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPublication sets the publication property value. Indicates the publication status of this particular version. Read-only.
func (m *BaseItemVersion) SetPublication(value PublicationFacetable)() {
    m.publication = value
}
