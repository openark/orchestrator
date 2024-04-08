package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ThreatAssessmentRequest 
type ThreatAssessmentRequest struct {
    Entity
    // The category property
    category *ThreatCategory
    // The content type of threat assessment. Possible values are: mail, url, file.
    contentType *ThreatAssessmentContentType
    // The threat assessment request creator.
    createdBy IdentitySetable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The expectedAssessment property
    expectedAssessment *ThreatExpectedAssessment
    // The source of the threat assessment request. Possible values are: administrator.
    requestSource *ThreatAssessmentRequestSource
    // A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
    results []ThreatAssessmentResultable
    // The assessment process status. Possible values are: pending, completed.
    status *ThreatAssessmentStatus
}
// NewThreatAssessmentRequest instantiates a new ThreatAssessmentRequest and sets the default values.
func NewThreatAssessmentRequest()(*ThreatAssessmentRequest) {
    m := &ThreatAssessmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateThreatAssessmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThreatAssessmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.emailFileAssessmentRequest":
                        return NewEmailFileAssessmentRequest(), nil
                    case "#microsoft.graph.fileAssessmentRequest":
                        return NewFileAssessmentRequest(), nil
                    case "#microsoft.graph.mailAssessmentRequest":
                        return NewMailAssessmentRequest(), nil
                    case "#microsoft.graph.urlAssessmentRequest":
                        return NewUrlAssessmentRequest(), nil
                }
            }
        }
    }
    return NewThreatAssessmentRequest(), nil
}
// GetCategory gets the category property value. The category property
func (m *ThreatAssessmentRequest) GetCategory()(*ThreatCategory) {
    return m.category
}
// GetContentType gets the contentType property value. The content type of threat assessment. Possible values are: mail, url, file.
func (m *ThreatAssessmentRequest) GetContentType()(*ThreatAssessmentContentType) {
    return m.contentType
}
// GetCreatedBy gets the createdBy property value. The threat assessment request creator.
func (m *ThreatAssessmentRequest) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ThreatAssessmentRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetExpectedAssessment gets the expectedAssessment property value. The expectedAssessment property
func (m *ThreatAssessmentRequest) GetExpectedAssessment()(*ThreatExpectedAssessment) {
    return m.expectedAssessment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThreatAssessmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ThreatCategory))
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentContentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(*ThreatAssessmentContentType))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
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
    res["expectedAssessment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatExpectedAssessment)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpectedAssessment(val.(*ThreatExpectedAssessment))
        }
        return nil
    }
    res["requestSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentRequestSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestSource(val.(*ThreatAssessmentRequestSource))
        }
        return nil
    }
    res["results"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateThreatAssessmentResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThreatAssessmentResultable, len(val))
            for i, v := range val {
                res[i] = v.(ThreatAssessmentResultable)
            }
            m.SetResults(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ThreatAssessmentStatus))
        }
        return nil
    }
    return res
}
// GetRequestSource gets the requestSource property value. The source of the threat assessment request. Possible values are: administrator.
func (m *ThreatAssessmentRequest) GetRequestSource()(*ThreatAssessmentRequestSource) {
    return m.requestSource
}
// GetResults gets the results property value. A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
func (m *ThreatAssessmentRequest) GetResults()([]ThreatAssessmentResultable) {
    return m.results
}
// GetStatus gets the status property value. The assessment process status. Possible values are: pending, completed.
func (m *ThreatAssessmentRequest) GetStatus()(*ThreatAssessmentStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ThreatAssessmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentType() != nil {
        cast := (*m.GetContentType()).String()
        err = writer.WriteStringValue("contentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
    if m.GetExpectedAssessment() != nil {
        cast := (*m.GetExpectedAssessment()).String()
        err = writer.WriteStringValue("expectedAssessment", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequestSource() != nil {
        cast := (*m.GetRequestSource()).String()
        err = writer.WriteStringValue("requestSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResults()))
        for i, v := range m.GetResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("results", cast)
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
// SetCategory sets the category property value. The category property
func (m *ThreatAssessmentRequest) SetCategory(value *ThreatCategory)() {
    m.category = value
}
// SetContentType sets the contentType property value. The content type of threat assessment. Possible values are: mail, url, file.
func (m *ThreatAssessmentRequest) SetContentType(value *ThreatAssessmentContentType)() {
    m.contentType = value
}
// SetCreatedBy sets the createdBy property value. The threat assessment request creator.
func (m *ThreatAssessmentRequest) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ThreatAssessmentRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetExpectedAssessment sets the expectedAssessment property value. The expectedAssessment property
func (m *ThreatAssessmentRequest) SetExpectedAssessment(value *ThreatExpectedAssessment)() {
    m.expectedAssessment = value
}
// SetRequestSource sets the requestSource property value. The source of the threat assessment request. Possible values are: administrator.
func (m *ThreatAssessmentRequest) SetRequestSource(value *ThreatAssessmentRequestSource)() {
    m.requestSource = value
}
// SetResults sets the results property value. A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
func (m *ThreatAssessmentRequest) SetResults(value []ThreatAssessmentResultable)() {
    m.results = value
}
// SetStatus sets the status property value. The assessment process status. Possible values are: pending, completed.
func (m *ThreatAssessmentRequest) SetStatus(value *ThreatAssessmentStatus)() {
    m.status = value
}
