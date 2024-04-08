package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditEvent 
type AuditEvent struct {
    Entity
    // Friendly name of the activity.
    activity *string
    // The date time in UTC when the activity was performed.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The HTTP operation type of the activity.
    activityOperationType *string
    // The result of the activity.
    activityResult *string
    // The type of activity that was being performed.
    activityType *string
    // AAD user and application that are associated with the audit event.
    actor AuditActorable
    // Audit category.
    category *string
    // Component name.
    componentName *string
    // The client request Id that is used to correlate activity within the system.
    correlationId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Event display name.
    displayName *string
    // Resources being modified.
    resources []AuditResourceable
}
// NewAuditEvent instantiates a new AuditEvent and sets the default values.
func NewAuditEvent()(*AuditEvent) {
    m := &AuditEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuditEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditEvent(), nil
}
// GetActivity gets the activity property value. Friendly name of the activity.
func (m *AuditEvent) GetActivity()(*string) {
    return m.activity
}
// GetActivityDateTime gets the activityDateTime property value. The date time in UTC when the activity was performed.
func (m *AuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activityDateTime
}
// GetActivityOperationType gets the activityOperationType property value. The HTTP operation type of the activity.
func (m *AuditEvent) GetActivityOperationType()(*string) {
    return m.activityOperationType
}
// GetActivityResult gets the activityResult property value. The result of the activity.
func (m *AuditEvent) GetActivityResult()(*string) {
    return m.activityResult
}
// GetActivityType gets the activityType property value. The type of activity that was being performed.
func (m *AuditEvent) GetActivityType()(*string) {
    return m.activityType
}
// GetActor gets the actor property value. AAD user and application that are associated with the audit event.
func (m *AuditEvent) GetActor()(AuditActorable) {
    return m.actor
}
// GetCategory gets the category property value. Audit category.
func (m *AuditEvent) GetCategory()(*string) {
    return m.category
}
// GetComponentName gets the componentName property value. Component name.
func (m *AuditEvent) GetComponentName()(*string) {
    return m.componentName
}
// GetCorrelationId gets the correlationId property value. The client request Id that is used to correlate activity within the system.
func (m *AuditEvent) GetCorrelationId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.correlationId
}
// GetDisplayName gets the displayName property value. Event display name.
func (m *AuditEvent) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["activityDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["activityOperationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityOperationType(val)
        }
        return nil
    }
    res["activityResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityResult(val)
        }
        return nil
    }
    res["activityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuditActorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(AuditActorable))
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["componentName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComponentName(val)
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
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
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditResourceable, len(val))
            for i, v := range val {
                res[i] = v.(AuditResourceable)
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetResources gets the resources property value. Resources being modified.
func (m *AuditEvent) GetResources()([]AuditResourceable) {
    return m.resources
}
// Serialize serializes information the current object
func (m *AuditEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityOperationType", m.GetActivityOperationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityResult", m.GetActivityResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("componentName", m.GetComponentName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("correlationId", m.GetCorrelationId())
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
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. Friendly name of the activity.
func (m *AuditEvent) SetActivity(value *string)() {
    m.activity = value
}
// SetActivityDateTime sets the activityDateTime property value. The date time in UTC when the activity was performed.
func (m *AuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// SetActivityOperationType sets the activityOperationType property value. The HTTP operation type of the activity.
func (m *AuditEvent) SetActivityOperationType(value *string)() {
    m.activityOperationType = value
}
// SetActivityResult sets the activityResult property value. The result of the activity.
func (m *AuditEvent) SetActivityResult(value *string)() {
    m.activityResult = value
}
// SetActivityType sets the activityType property value. The type of activity that was being performed.
func (m *AuditEvent) SetActivityType(value *string)() {
    m.activityType = value
}
// SetActor sets the actor property value. AAD user and application that are associated with the audit event.
func (m *AuditEvent) SetActor(value AuditActorable)() {
    m.actor = value
}
// SetCategory sets the category property value. Audit category.
func (m *AuditEvent) SetCategory(value *string)() {
    m.category = value
}
// SetComponentName sets the componentName property value. Component name.
func (m *AuditEvent) SetComponentName(value *string)() {
    m.componentName = value
}
// SetCorrelationId sets the correlationId property value. The client request Id that is used to correlate activity within the system.
func (m *AuditEvent) SetCorrelationId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.correlationId = value
}
// SetDisplayName sets the displayName property value. Event display name.
func (m *AuditEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetResources sets the resources property value. Resources being modified.
func (m *AuditEvent) SetResources(value []AuditResourceable)() {
    m.resources = value
}
