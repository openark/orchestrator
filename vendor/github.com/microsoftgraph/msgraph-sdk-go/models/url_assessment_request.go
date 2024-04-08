package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UrlAssessmentRequest 
type UrlAssessmentRequest struct {
    ThreatAssessmentRequest
    // The URL string.
    url *string
}
// NewUrlAssessmentRequest instantiates a new UrlAssessmentRequest and sets the default values.
func NewUrlAssessmentRequest()(*UrlAssessmentRequest) {
    m := &UrlAssessmentRequest{
        ThreatAssessmentRequest: *NewThreatAssessmentRequest(),
    }
    odataTypeValue := "#microsoft.graph.urlAssessmentRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUrlAssessmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUrlAssessmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUrlAssessmentRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UrlAssessmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ThreatAssessmentRequest.GetFieldDeserializers()
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. The URL string.
func (m *UrlAssessmentRequest) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *UrlAssessmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ThreatAssessmentRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUrl sets the url property value. The URL string.
func (m *UrlAssessmentRequest) SetUrl(value *string)() {
    m.url = value
}
