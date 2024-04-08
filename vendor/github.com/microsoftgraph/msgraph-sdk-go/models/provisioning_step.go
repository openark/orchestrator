package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningStep 
type ProvisioningStep struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Summary of what occurred during the step.
    description *string
    // Details of what occurred during the step.
    details DetailsInfoable
    // Name of the step.
    name *string
    // The OdataType property
    odataType *string
    // Type of step. Possible values are: import, scoping, matching, processing, referenceResolution, export, unknownFutureValue.
    provisioningStepType *ProvisioningStepType
    // Status of the step. Possible values are: success, warning,  failure, skipped, unknownFutureValue.
    status *ProvisioningResult
}
// NewProvisioningStep instantiates a new provisioningStep and sets the default values.
func NewProvisioningStep()(*ProvisioningStep) {
    m := &ProvisioningStep{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProvisioningStepFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningStepFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningStep(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningStep) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Summary of what occurred during the step.
func (m *ProvisioningStep) GetDescription()(*string) {
    return m.description
}
// GetDetails gets the details property value. Details of what occurred during the step.
func (m *ProvisioningStep) GetDetails()(DetailsInfoable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningStep) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetailsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(DetailsInfoable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
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
    res["provisioningStepType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningStepType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStepType(val.(*ProvisioningStepType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ProvisioningResult))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the step.
func (m *ProvisioningStep) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ProvisioningStep) GetOdataType()(*string) {
    return m.odataType
}
// GetProvisioningStepType gets the provisioningStepType property value. Type of step. Possible values are: import, scoping, matching, processing, referenceResolution, export, unknownFutureValue.
func (m *ProvisioningStep) GetProvisioningStepType()(*ProvisioningStepType) {
    return m.provisioningStepType
}
// GetStatus gets the status property value. Status of the step. Possible values are: success, warning,  failure, skipped, unknownFutureValue.
func (m *ProvisioningStep) GetStatus()(*ProvisioningResult) {
    return m.status
}
// Serialize serializes information the current object
func (m *ProvisioningStep) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningStepType() != nil {
        cast := (*m.GetProvisioningStepType()).String()
        err := writer.WriteStringValue("provisioningStepType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *ProvisioningStep) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Summary of what occurred during the step.
func (m *ProvisioningStep) SetDescription(value *string)() {
    m.description = value
}
// SetDetails sets the details property value. Details of what occurred during the step.
func (m *ProvisioningStep) SetDetails(value DetailsInfoable)() {
    m.details = value
}
// SetName sets the name property value. Name of the step.
func (m *ProvisioningStep) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ProvisioningStep) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProvisioningStepType sets the provisioningStepType property value. Type of step. Possible values are: import, scoping, matching, processing, referenceResolution, export, unknownFutureValue.
func (m *ProvisioningStep) SetProvisioningStepType(value *ProvisioningStepType)() {
    m.provisioningStepType = value
}
// SetStatus sets the status property value. Status of the step. Possible values are: success, warning,  failure, skipped, unknownFutureValue.
func (m *ProvisioningStep) SetStatus(value *ProvisioningResult)() {
    m.status = value
}
