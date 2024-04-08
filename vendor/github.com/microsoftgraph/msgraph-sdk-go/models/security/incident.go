package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Incident 
type Incident struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
    // The list of related alerts. Supports $expand.
    alerts []Alertable
    // Owner of the incident, or null if no owner is assigned. Free editable text.
    assignedTo *string
    // The specification for the incident. Possible values are: unknown, falsePositive, truePositive, informationalExpectedActivity, unknownFutureValue.
    classification *AlertClassification
    // Array of comments created by the Security Operations (SecOps) team when the incident is managed.
    comments []AlertCommentable
    // Time when the incident was first created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Array of custom tags associated with an incident.
    customTags []string
    // Specifies the determination of the incident. Possible values are: unknown, apt, malware, securityPersonnel, securityTesting, unwantedSoftware, other, multiStagedAttack, compromisedUser, phishing, maliciousUserActivity, clean, insufficientData, confirmedUserActivity, lineOfBusinessApplication, unknownFutureValue.
    determination *AlertDetermination
    // The incident name.
    displayName *string
    // The URL for the incident page in the Microsoft 365 Defender portal.
    incidentWebUrl *string
    // Time when the incident was last updated.
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only populated in case an incident is grouped together with another incident, as part of the logic that processes incidents. In such a case, the status property is redirected.
    redirectIncidentId *string
    // The severity property
    severity *AlertSeverity
    // The status property
    status *IncidentStatus
    // The Azure Active Directory tenant in which the alert was created.
    tenantId *string
}
// NewIncident instantiates a new incident and sets the default values.
func NewIncident()(*Incident) {
    m := &Incident{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIncident(), nil
}
// GetAlerts gets the alerts property value. The list of related alerts. Supports $expand.
func (m *Incident) GetAlerts()([]Alertable) {
    return m.alerts
}
// GetAssignedTo gets the assignedTo property value. Owner of the incident, or null if no owner is assigned. Free editable text.
func (m *Incident) GetAssignedTo()(*string) {
    return m.assignedTo
}
// GetClassification gets the classification property value. The specification for the incident. Possible values are: unknown, falsePositive, truePositive, informationalExpectedActivity, unknownFutureValue.
func (m *Incident) GetClassification()(*AlertClassification) {
    return m.classification
}
// GetComments gets the comments property value. Array of comments created by the Security Operations (SecOps) team when the incident is managed.
func (m *Incident) GetComments()([]AlertCommentable) {
    return m.comments
}
// GetCreatedDateTime gets the createdDateTime property value. Time when the incident was first created.
func (m *Incident) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomTags gets the customTags property value. Array of custom tags associated with an incident.
func (m *Incident) GetCustomTags()([]string) {
    return m.customTags
}
// GetDetermination gets the determination property value. Specifies the determination of the incident. Possible values are: unknown, apt, malware, securityPersonnel, securityTesting, unwantedSoftware, other, multiStagedAttack, compromisedUser, phishing, maliciousUserActivity, clean, insufficientData, confirmedUserActivity, lineOfBusinessApplication, unknownFutureValue.
func (m *Incident) GetDetermination()(*AlertDetermination) {
    return m.determination
}
// GetDisplayName gets the displayName property value. The incident name.
func (m *Incident) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Incident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alertable, len(val))
            for i, v := range val {
                res[i] = v.(Alertable)
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertClassification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*AlertClassification))
        }
        return nil
    }
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertCommentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertCommentable, len(val))
            for i, v := range val {
                res[i] = v.(AlertCommentable)
            }
            m.SetComments(res)
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
    res["customTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCustomTags(res)
        }
        return nil
    }
    res["determination"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertDetermination)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetermination(val.(*AlertDetermination))
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
    res["incidentWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncidentWebUrl(val)
        }
        return nil
    }
    res["lastUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
        }
        return nil
    }
    res["redirectIncidentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectIncidentId(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*AlertSeverity))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncidentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*IncidentStatus))
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetIncidentWebUrl gets the incidentWebUrl property value. The URL for the incident page in the Microsoft 365 Defender portal.
func (m *Incident) GetIncidentWebUrl()(*string) {
    return m.incidentWebUrl
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. Time when the incident was last updated.
func (m *Incident) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdateDateTime
}
// GetRedirectIncidentId gets the redirectIncidentId property value. Only populated in case an incident is grouped together with another incident, as part of the logic that processes incidents. In such a case, the status property is redirected.
func (m *Incident) GetRedirectIncidentId()(*string) {
    return m.redirectIncidentId
}
// GetSeverity gets the severity property value. The severity property
func (m *Incident) GetSeverity()(*AlertSeverity) {
    return m.severity
}
// GetStatus gets the status property value. The status property
func (m *Incident) GetStatus()(*IncidentStatus) {
    return m.status
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant in which the alert was created.
func (m *Incident) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *Incident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetComments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComments()))
        for i, v := range m.GetComments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("comments", cast)
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
    if m.GetCustomTags() != nil {
        err = writer.WriteCollectionOfStringValues("customTags", m.GetCustomTags())
        if err != nil {
            return err
        }
    }
    if m.GetDetermination() != nil {
        cast := (*m.GetDetermination()).String()
        err = writer.WriteStringValue("determination", &cast)
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
        err = writer.WriteStringValue("incidentWebUrl", m.GetIncidentWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectIncidentId", m.GetRedirectIncidentId())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
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
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlerts sets the alerts property value. The list of related alerts. Supports $expand.
func (m *Incident) SetAlerts(value []Alertable)() {
    m.alerts = value
}
// SetAssignedTo sets the assignedTo property value. Owner of the incident, or null if no owner is assigned. Free editable text.
func (m *Incident) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// SetClassification sets the classification property value. The specification for the incident. Possible values are: unknown, falsePositive, truePositive, informationalExpectedActivity, unknownFutureValue.
func (m *Incident) SetClassification(value *AlertClassification)() {
    m.classification = value
}
// SetComments sets the comments property value. Array of comments created by the Security Operations (SecOps) team when the incident is managed.
func (m *Incident) SetComments(value []AlertCommentable)() {
    m.comments = value
}
// SetCreatedDateTime sets the createdDateTime property value. Time when the incident was first created.
func (m *Incident) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomTags sets the customTags property value. Array of custom tags associated with an incident.
func (m *Incident) SetCustomTags(value []string)() {
    m.customTags = value
}
// SetDetermination sets the determination property value. Specifies the determination of the incident. Possible values are: unknown, apt, malware, securityPersonnel, securityTesting, unwantedSoftware, other, multiStagedAttack, compromisedUser, phishing, maliciousUserActivity, clean, insufficientData, confirmedUserActivity, lineOfBusinessApplication, unknownFutureValue.
func (m *Incident) SetDetermination(value *AlertDetermination)() {
    m.determination = value
}
// SetDisplayName sets the displayName property value. The incident name.
func (m *Incident) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIncidentWebUrl sets the incidentWebUrl property value. The URL for the incident page in the Microsoft 365 Defender portal.
func (m *Incident) SetIncidentWebUrl(value *string)() {
    m.incidentWebUrl = value
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. Time when the incident was last updated.
func (m *Incident) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
// SetRedirectIncidentId sets the redirectIncidentId property value. Only populated in case an incident is grouped together with another incident, as part of the logic that processes incidents. In such a case, the status property is redirected.
func (m *Incident) SetRedirectIncidentId(value *string)() {
    m.redirectIncidentId = value
}
// SetSeverity sets the severity property value. The severity property
func (m *Incident) SetSeverity(value *AlertSeverity)() {
    m.severity = value
}
// SetStatus sets the status property value. The status property
func (m *Incident) SetStatus(value *IncidentStatus)() {
    m.status = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant in which the alert was created.
func (m *Incident) SetTenantId(value *string)() {
    m.tenantId = value
}
