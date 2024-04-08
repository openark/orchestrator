package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChecklistItem 
type ChecklistItem struct {
    Entity
    // The date and time when the checklistItem was finished.
    checkedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time when the checklistItem was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Field indicating the title of checklistItem.
    displayName *string
    // State indicating whether the item is checked off or not.
    isChecked *bool
}
// NewChecklistItem instantiates a new checklistItem and sets the default values.
func NewChecklistItem()(*ChecklistItem) {
    m := &ChecklistItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateChecklistItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChecklistItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChecklistItem(), nil
}
// GetCheckedDateTime gets the checkedDateTime property value. The date and time when the checklistItem was finished.
func (m *ChecklistItem) GetCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.checkedDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the checklistItem was created.
func (m *ChecklistItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. Field indicating the title of checklistItem.
func (m *ChecklistItem) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChecklistItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["checkedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckedDateTime(val)
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
    res["isChecked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsChecked(val)
        }
        return nil
    }
    return res
}
// GetIsChecked gets the isChecked property value. State indicating whether the item is checked off or not.
func (m *ChecklistItem) GetIsChecked()(*bool) {
    return m.isChecked
}
// Serialize serializes information the current object
func (m *ChecklistItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("checkedDateTime", m.GetCheckedDateTime())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isChecked", m.GetIsChecked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCheckedDateTime sets the checkedDateTime property value. The date and time when the checklistItem was finished.
func (m *ChecklistItem) SetCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.checkedDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the checklistItem was created.
func (m *ChecklistItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. Field indicating the title of checklistItem.
func (m *ChecklistItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsChecked sets the isChecked property value. State indicating whether the item is checked off or not.
func (m *ChecklistItem) SetIsChecked(value *bool)() {
    m.isChecked = value
}
