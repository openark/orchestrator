package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationOrganization 
type EducationOrganization struct {
    Entity
    // Organization description.
    description *string
    // Organization display name.
    displayName *string
    // Source where this organization was created from. Possible values are: sis, manual.
    externalSource *EducationExternalSource
    // The name of the external source this resources was generated from.
    externalSourceDetail *string
}
// NewEducationOrganization instantiates a new educationOrganization and sets the default values.
func NewEducationOrganization()(*EducationOrganization) {
    m := &EducationOrganization{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationOrganizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationOrganizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.educationSchool":
                        return NewEducationSchool(), nil
                }
            }
        }
    }
    return NewEducationOrganization(), nil
}
// GetDescription gets the description property value. Organization description.
func (m *EducationOrganization) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Organization display name.
func (m *EducationOrganization) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalSource gets the externalSource property value. Source where this organization was created from. Possible values are: sis, manual.
func (m *EducationOrganization) GetExternalSource()(*EducationExternalSource) {
    return m.externalSource
}
// GetExternalSourceDetail gets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationOrganization) GetExternalSourceDetail()(*string) {
    return m.externalSourceDetail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationOrganization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["externalSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationExternalSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSource(val.(*EducationExternalSource))
        }
        return nil
    }
    res["externalSourceDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSourceDetail(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EducationOrganization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetExternalSource() != nil {
        cast := (*m.GetExternalSource()).String()
        err = writer.WriteStringValue("externalSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalSourceDetail", m.GetExternalSourceDetail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Organization description.
func (m *EducationOrganization) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Organization display name.
func (m *EducationOrganization) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalSource sets the externalSource property value. Source where this organization was created from. Possible values are: sis, manual.
func (m *EducationOrganization) SetExternalSource(value *EducationExternalSource)() {
    m.externalSource = value
}
// SetExternalSourceDetail sets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationOrganization) SetExternalSourceDetail(value *string)() {
    m.externalSourceDetail = value
}
