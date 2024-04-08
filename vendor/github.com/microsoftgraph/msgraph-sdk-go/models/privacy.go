package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Privacy 
type Privacy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The subjectRightsRequests property
    subjectRightsRequests []SubjectRightsRequestable
}
// NewPrivacy instantiates a new Privacy and sets the default values.
func NewPrivacy()(*Privacy) {
    m := &Privacy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrivacyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivacyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivacy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Privacy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Privacy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["subjectRightsRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectRightsRequestable)
            }
            m.SetSubjectRightsRequests(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Privacy) GetOdataType()(*string) {
    return m.odataType
}
// GetSubjectRightsRequests gets the subjectRightsRequests property value. The subjectRightsRequests property
func (m *Privacy) GetSubjectRightsRequests()([]SubjectRightsRequestable) {
    return m.subjectRightsRequests
}
// Serialize serializes information the current object
func (m *Privacy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectRightsRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjectRightsRequests()))
        for i, v := range m.GetSubjectRightsRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("subjectRightsRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Privacy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Privacy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSubjectRightsRequests sets the subjectRightsRequests property value. The subjectRightsRequests property
func (m *Privacy) SetSubjectRightsRequests(value []SubjectRightsRequestable)() {
    m.subjectRightsRequests = value
}
