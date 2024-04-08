package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentRecipient 
type EducationAssignmentRecipient struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
}
// NewEducationAssignmentRecipient instantiates a new educationAssignmentRecipient and sets the default values.
func NewEducationAssignmentRecipient()(*EducationAssignmentRecipient) {
    m := &EducationAssignmentRecipient{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationAssignmentRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.educationAssignmentClassRecipient":
                        return NewEducationAssignmentClassRecipient(), nil
                    case "#microsoft.graph.educationAssignmentGroupRecipient":
                        return NewEducationAssignmentGroupRecipient(), nil
                    case "#microsoft.graph.educationAssignmentIndividualRecipient":
                        return NewEducationAssignmentIndividualRecipient(), nil
                }
            }
        }
    }
    return NewEducationAssignmentRecipient(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationAssignmentRecipient) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentRecipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationAssignmentRecipient) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *EducationAssignmentRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *EducationAssignmentRecipient) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationAssignmentRecipient) SetOdataType(value *string)() {
    m.odataType = value
}
