package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UrlEvidence 
type UrlEvidence struct {
    AlertEvidence
    // The Unique Resource Locator (URL).
    url *string
}
// NewUrlEvidence instantiates a new UrlEvidence and sets the default values.
func NewUrlEvidence()(*UrlEvidence) {
    m := &UrlEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateUrlEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUrlEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUrlEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UrlEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
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
// GetUrl gets the url property value. The Unique Resource Locator (URL).
func (m *UrlEvidence) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *UrlEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
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
// SetUrl sets the url property value. The Unique Resource Locator (URL).
func (m *UrlEvidence) SetUrl(value *string)() {
    m.url = value
}
