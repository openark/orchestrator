package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkTag 
type TeamworkTag struct {
    Entity
    // The description of the tag as it will appear to the user in Microsoft Teams.
    description *string
    // The name of the tag as it will appear to the user in Microsoft Teams.
    displayName *string
    // The number of users assigned to the tag.
    memberCount *int32
    // Users assigned to the tag.
    members []TeamworkTagMemberable
    // The type of the tag. Default is standard.
    tagType *TeamworkTagType
    // ID of the team in which the tag is defined.
    teamId *string
}
// NewTeamworkTag instantiates a new teamworkTag and sets the default values.
func NewTeamworkTag()(*TeamworkTag) {
    m := &TeamworkTag{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkTagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkTag(), nil
}
// GetDescription gets the description property value. The description of the tag as it will appear to the user in Microsoft Teams.
func (m *TeamworkTag) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the tag as it will appear to the user in Microsoft Teams.
func (m *TeamworkTag) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["memberCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberCount(val)
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkTagMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkTagMemberable, len(val))
            for i, v := range val {
                res[i] = v.(TeamworkTagMemberable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["tagType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkTagType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTagType(val.(*TeamworkTagType))
        }
        return nil
    }
    res["teamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    return res
}
// GetMemberCount gets the memberCount property value. The number of users assigned to the tag.
func (m *TeamworkTag) GetMemberCount()(*int32) {
    return m.memberCount
}
// GetMembers gets the members property value. Users assigned to the tag.
func (m *TeamworkTag) GetMembers()([]TeamworkTagMemberable) {
    return m.members
}
// GetTagType gets the tagType property value. The type of the tag. Default is standard.
func (m *TeamworkTag) GetTagType()(*TeamworkTagType) {
    return m.tagType
}
// GetTeamId gets the teamId property value. ID of the team in which the tag is defined.
func (m *TeamworkTag) GetTeamId()(*string) {
    return m.teamId
}
// Serialize serializes information the current object
func (m *TeamworkTag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteInt32Value("memberCount", m.GetMemberCount())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTagType() != nil {
        cast := (*m.GetTagType()).String()
        err = writer.WriteStringValue("tagType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the tag as it will appear to the user in Microsoft Teams.
func (m *TeamworkTag) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the tag as it will appear to the user in Microsoft Teams.
func (m *TeamworkTag) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMemberCount sets the memberCount property value. The number of users assigned to the tag.
func (m *TeamworkTag) SetMemberCount(value *int32)() {
    m.memberCount = value
}
// SetMembers sets the members property value. Users assigned to the tag.
func (m *TeamworkTag) SetMembers(value []TeamworkTagMemberable)() {
    m.members = value
}
// SetTagType sets the tagType property value. The type of the tag. Default is standard.
func (m *TeamworkTag) SetTagType(value *TeamworkTagType)() {
    m.tagType = value
}
// SetTeamId sets the teamId property value. ID of the team in which the tag is defined.
func (m *TeamworkTag) SetTeamId(value *string)() {
    m.teamId = value
}
