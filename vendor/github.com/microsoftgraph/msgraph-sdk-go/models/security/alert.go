package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Alert 
type Alert struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
    // The adversary or activity group that is associated with this alert.
    actorDisplayName *string
    // URL for the alert page in the Microsoft 365 Defender portal.
    alertWebUrl *string
    // Owner of the alert, or null if no owner is assigned.
    assignedTo *string
    // The attack kill-chain category that the alert belongs to. Aligned with the MITRE ATT&CK framework.
    category *string
    // Specifies whether the alert represents a true threat. Possible values are: unknown, falsePositive, truePositive, benignPositive, unknownFutureValue.
    classification *AlertClassification
    // Array of comments created by the Security Operations (SecOps) team during the alert management process.
    comments []AlertCommentable
    // Time when Microsoft 365 Defender created the alert.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // String value describing each alert.
    description *string
    // Detection technology or sensor that identified the notable component or activity.
    detectionSource *DetectionSource
    // The ID of the detector that triggered the alert.
    detectorId *string
    // Specifies the result of the investigation, whether the alert represents a true attack and if so, the nature of the attack. Possible values are: unknown, apt, malware, securityPersonnel, securityTesting, unwantedSoftware, other, multiStagedAttack, compromisedUser, phishing, maliciousUserActivity, clean, insufficientData, confirmedUserActivity, lineOfBusinessApplication, unknownFutureValue.
    determination *AlertDetermination
    // Collection of evidence related to the alert.
    evidence []AlertEvidenceable
    // The earliest activity associated with the alert.
    firstActivityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique identifier to represent the incident this alert resource is associated with.
    incidentId *string
    // URL for the incident page in the Microsoft 365 Defender portal.
    incidentWebUrl *string
    // The oldest activity associated with the alert.
    lastActivityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Time when the alert was last updated at Microsoft 365 Defender.
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The attack techniques, as aligned with the MITRE ATT&CK framework.
    mitreTechniques []string
    // The ID of the alert as it appears in the security provider product that generated the alert.
    providerAlertId *string
    // Recommended response and remediation actions to take in the event this alert was generated.
    recommendedActions *string
    // Time when the alert was resolved.
    resolvedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The serviceSource property
    serviceSource *ServiceSource
    // The severity property
    severity *AlertSeverity
    // The status property
    status *AlertStatus
    // The Azure Active Directory tenant the alert was created in.
    tenantId *string
    // The threat associated with this alert.
    threatDisplayName *string
    // Threat family associated with this alert.
    threatFamilyName *string
    // Brief identifying string value describing the alert.
    title *string
}
// NewAlert instantiates a new alert and sets the default values.
func NewAlert()(*Alert) {
    m := &Alert{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlert(), nil
}
// GetActorDisplayName gets the actorDisplayName property value. The adversary or activity group that is associated with this alert.
func (m *Alert) GetActorDisplayName()(*string) {
    return m.actorDisplayName
}
// GetAlertWebUrl gets the alertWebUrl property value. URL for the alert page in the Microsoft 365 Defender portal.
func (m *Alert) GetAlertWebUrl()(*string) {
    return m.alertWebUrl
}
// GetAssignedTo gets the assignedTo property value. Owner of the alert, or null if no owner is assigned.
func (m *Alert) GetAssignedTo()(*string) {
    return m.assignedTo
}
// GetCategory gets the category property value. The attack kill-chain category that the alert belongs to. Aligned with the MITRE ATT&CK framework.
func (m *Alert) GetCategory()(*string) {
    return m.category
}
// GetClassification gets the classification property value. Specifies whether the alert represents a true threat. Possible values are: unknown, falsePositive, truePositive, benignPositive, unknownFutureValue.
func (m *Alert) GetClassification()(*AlertClassification) {
    return m.classification
}
// GetComments gets the comments property value. Array of comments created by the Security Operations (SecOps) team during the alert management process.
func (m *Alert) GetComments()([]AlertCommentable) {
    return m.comments
}
// GetCreatedDateTime gets the createdDateTime property value. Time when Microsoft 365 Defender created the alert.
func (m *Alert) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. String value describing each alert.
func (m *Alert) GetDescription()(*string) {
    return m.description
}
// GetDetectionSource gets the detectionSource property value. Detection technology or sensor that identified the notable component or activity.
func (m *Alert) GetDetectionSource()(*DetectionSource) {
    return m.detectionSource
}
// GetDetectorId gets the detectorId property value. The ID of the detector that triggered the alert.
func (m *Alert) GetDetectorId()(*string) {
    return m.detectorId
}
// GetDetermination gets the determination property value. Specifies the result of the investigation, whether the alert represents a true attack and if so, the nature of the attack. Possible values are: unknown, apt, malware, securityPersonnel, securityTesting, unwantedSoftware, other, multiStagedAttack, compromisedUser, phishing, maliciousUserActivity, clean, insufficientData, confirmedUserActivity, lineOfBusinessApplication, unknownFutureValue.
func (m *Alert) GetDetermination()(*AlertDetermination) {
    return m.determination
}
// GetEvidence gets the evidence property value. Collection of evidence related to the alert.
func (m *Alert) GetEvidence()([]AlertEvidenceable) {
    return m.evidence
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Alert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actorDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorDisplayName(val)
        }
        return nil
    }
    res["alertWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertWebUrl(val)
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
    res["detectionSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDetectionSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionSource(val.(*DetectionSource))
        }
        return nil
    }
    res["detectorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectorId(val)
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
    res["evidence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertEvidenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertEvidenceable, len(val))
            for i, v := range val {
                res[i] = v.(AlertEvidenceable)
            }
            m.SetEvidence(res)
        }
        return nil
    }
    res["firstActivityDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstActivityDateTime(val)
        }
        return nil
    }
    res["incidentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncidentId(val)
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
    res["lastActivityDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDateTime(val)
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
    res["mitreTechniques"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMitreTechniques(res)
        }
        return nil
    }
    res["providerAlertId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderAlertId(val)
        }
        return nil
    }
    res["recommendedActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedActions(val)
        }
        return nil
    }
    res["resolvedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedDateTime(val)
        }
        return nil
    }
    res["serviceSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceSource(val.(*ServiceSource))
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
        val, err := n.GetEnumValue(ParseAlertStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AlertStatus))
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
    res["threatDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatDisplayName(val)
        }
        return nil
    }
    res["threatFamilyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatFamilyName(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetFirstActivityDateTime gets the firstActivityDateTime property value. The earliest activity associated with the alert.
func (m *Alert) GetFirstActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstActivityDateTime
}
// GetIncidentId gets the incidentId property value. Unique identifier to represent the incident this alert resource is associated with.
func (m *Alert) GetIncidentId()(*string) {
    return m.incidentId
}
// GetIncidentWebUrl gets the incidentWebUrl property value. URL for the incident page in the Microsoft 365 Defender portal.
func (m *Alert) GetIncidentWebUrl()(*string) {
    return m.incidentWebUrl
}
// GetLastActivityDateTime gets the lastActivityDateTime property value. The oldest activity associated with the alert.
func (m *Alert) GetLastActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActivityDateTime
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. Time when the alert was last updated at Microsoft 365 Defender.
func (m *Alert) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdateDateTime
}
// GetMitreTechniques gets the mitreTechniques property value. The attack techniques, as aligned with the MITRE ATT&CK framework.
func (m *Alert) GetMitreTechniques()([]string) {
    return m.mitreTechniques
}
// GetProviderAlertId gets the providerAlertId property value. The ID of the alert as it appears in the security provider product that generated the alert.
func (m *Alert) GetProviderAlertId()(*string) {
    return m.providerAlertId
}
// GetRecommendedActions gets the recommendedActions property value. Recommended response and remediation actions to take in the event this alert was generated.
func (m *Alert) GetRecommendedActions()(*string) {
    return m.recommendedActions
}
// GetResolvedDateTime gets the resolvedDateTime property value. Time when the alert was resolved.
func (m *Alert) GetResolvedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.resolvedDateTime
}
// GetServiceSource gets the serviceSource property value. The serviceSource property
func (m *Alert) GetServiceSource()(*ServiceSource) {
    return m.serviceSource
}
// GetSeverity gets the severity property value. The severity property
func (m *Alert) GetSeverity()(*AlertSeverity) {
    return m.severity
}
// GetStatus gets the status property value. The status property
func (m *Alert) GetStatus()(*AlertStatus) {
    return m.status
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant the alert was created in.
func (m *Alert) GetTenantId()(*string) {
    return m.tenantId
}
// GetThreatDisplayName gets the threatDisplayName property value. The threat associated with this alert.
func (m *Alert) GetThreatDisplayName()(*string) {
    return m.threatDisplayName
}
// GetThreatFamilyName gets the threatFamilyName property value. Threat family associated with this alert.
func (m *Alert) GetThreatFamilyName()(*string) {
    return m.threatFamilyName
}
// GetTitle gets the title property value. Brief identifying string value describing the alert.
func (m *Alert) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *Alert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actorDisplayName", m.GetActorDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertWebUrl", m.GetAlertWebUrl())
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
    {
        err = writer.WriteStringValue("category", m.GetCategory())
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
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionSource() != nil {
        cast := (*m.GetDetectionSource()).String()
        err = writer.WriteStringValue("detectionSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("detectorId", m.GetDetectorId())
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
    if m.GetEvidence() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvidence()))
        for i, v := range m.GetEvidence() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("evidence", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstActivityDateTime", m.GetFirstActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("incidentId", m.GetIncidentId())
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
        err = writer.WriteTimeValue("lastActivityDateTime", m.GetLastActivityDateTime())
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
    if m.GetMitreTechniques() != nil {
        err = writer.WriteCollectionOfStringValues("mitreTechniques", m.GetMitreTechniques())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerAlertId", m.GetProviderAlertId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendedActions", m.GetRecommendedActions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("resolvedDateTime", m.GetResolvedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetServiceSource() != nil {
        cast := (*m.GetServiceSource()).String()
        err = writer.WriteStringValue("serviceSource", &cast)
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
    {
        err = writer.WriteStringValue("threatDisplayName", m.GetThreatDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("threatFamilyName", m.GetThreatFamilyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActorDisplayName sets the actorDisplayName property value. The adversary or activity group that is associated with this alert.
func (m *Alert) SetActorDisplayName(value *string)() {
    m.actorDisplayName = value
}
// SetAlertWebUrl sets the alertWebUrl property value. URL for the alert page in the Microsoft 365 Defender portal.
func (m *Alert) SetAlertWebUrl(value *string)() {
    m.alertWebUrl = value
}
// SetAssignedTo sets the assignedTo property value. Owner of the alert, or null if no owner is assigned.
func (m *Alert) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// SetCategory sets the category property value. The attack kill-chain category that the alert belongs to. Aligned with the MITRE ATT&CK framework.
func (m *Alert) SetCategory(value *string)() {
    m.category = value
}
// SetClassification sets the classification property value. Specifies whether the alert represents a true threat. Possible values are: unknown, falsePositive, truePositive, benignPositive, unknownFutureValue.
func (m *Alert) SetClassification(value *AlertClassification)() {
    m.classification = value
}
// SetComments sets the comments property value. Array of comments created by the Security Operations (SecOps) team during the alert management process.
func (m *Alert) SetComments(value []AlertCommentable)() {
    m.comments = value
}
// SetCreatedDateTime sets the createdDateTime property value. Time when Microsoft 365 Defender created the alert.
func (m *Alert) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. String value describing each alert.
func (m *Alert) SetDescription(value *string)() {
    m.description = value
}
// SetDetectionSource sets the detectionSource property value. Detection technology or sensor that identified the notable component or activity.
func (m *Alert) SetDetectionSource(value *DetectionSource)() {
    m.detectionSource = value
}
// SetDetectorId sets the detectorId property value. The ID of the detector that triggered the alert.
func (m *Alert) SetDetectorId(value *string)() {
    m.detectorId = value
}
// SetDetermination sets the determination property value. Specifies the result of the investigation, whether the alert represents a true attack and if so, the nature of the attack. Possible values are: unknown, apt, malware, securityPersonnel, securityTesting, unwantedSoftware, other, multiStagedAttack, compromisedUser, phishing, maliciousUserActivity, clean, insufficientData, confirmedUserActivity, lineOfBusinessApplication, unknownFutureValue.
func (m *Alert) SetDetermination(value *AlertDetermination)() {
    m.determination = value
}
// SetEvidence sets the evidence property value. Collection of evidence related to the alert.
func (m *Alert) SetEvidence(value []AlertEvidenceable)() {
    m.evidence = value
}
// SetFirstActivityDateTime sets the firstActivityDateTime property value. The earliest activity associated with the alert.
func (m *Alert) SetFirstActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstActivityDateTime = value
}
// SetIncidentId sets the incidentId property value. Unique identifier to represent the incident this alert resource is associated with.
func (m *Alert) SetIncidentId(value *string)() {
    m.incidentId = value
}
// SetIncidentWebUrl sets the incidentWebUrl property value. URL for the incident page in the Microsoft 365 Defender portal.
func (m *Alert) SetIncidentWebUrl(value *string)() {
    m.incidentWebUrl = value
}
// SetLastActivityDateTime sets the lastActivityDateTime property value. The oldest activity associated with the alert.
func (m *Alert) SetLastActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActivityDateTime = value
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. Time when the alert was last updated at Microsoft 365 Defender.
func (m *Alert) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
// SetMitreTechniques sets the mitreTechniques property value. The attack techniques, as aligned with the MITRE ATT&CK framework.
func (m *Alert) SetMitreTechniques(value []string)() {
    m.mitreTechniques = value
}
// SetProviderAlertId sets the providerAlertId property value. The ID of the alert as it appears in the security provider product that generated the alert.
func (m *Alert) SetProviderAlertId(value *string)() {
    m.providerAlertId = value
}
// SetRecommendedActions sets the recommendedActions property value. Recommended response and remediation actions to take in the event this alert was generated.
func (m *Alert) SetRecommendedActions(value *string)() {
    m.recommendedActions = value
}
// SetResolvedDateTime sets the resolvedDateTime property value. Time when the alert was resolved.
func (m *Alert) SetResolvedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.resolvedDateTime = value
}
// SetServiceSource sets the serviceSource property value. The serviceSource property
func (m *Alert) SetServiceSource(value *ServiceSource)() {
    m.serviceSource = value
}
// SetSeverity sets the severity property value. The severity property
func (m *Alert) SetSeverity(value *AlertSeverity)() {
    m.severity = value
}
// SetStatus sets the status property value. The status property
func (m *Alert) SetStatus(value *AlertStatus)() {
    m.status = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant the alert was created in.
func (m *Alert) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetThreatDisplayName sets the threatDisplayName property value. The threat associated with this alert.
func (m *Alert) SetThreatDisplayName(value *string)() {
    m.threatDisplayName = value
}
// SetThreatFamilyName sets the threatFamilyName property value. Threat family associated with this alert.
func (m *Alert) SetThreatFamilyName(value *string)() {
    m.threatFamilyName = value
}
// SetTitle sets the title property value. Brief identifying string value describing the alert.
func (m *Alert) SetTitle(value *string)() {
    m.title = value
}
