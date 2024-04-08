package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppDiagnosticStatus represents diagnostics status.
type ManagedAppDiagnosticStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Instruction on how to mitigate a failed validation
    mitigationInstruction *string
    // The OdataType property
    odataType *string
    // The state of the operation
    state *string
    // The validation friendly name
    validationName *string
}
// NewManagedAppDiagnosticStatus instantiates a new managedAppDiagnosticStatus and sets the default values.
func NewManagedAppDiagnosticStatus()(*ManagedAppDiagnosticStatus) {
    m := &ManagedAppDiagnosticStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedAppDiagnosticStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppDiagnosticStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAppDiagnosticStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedAppDiagnosticStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppDiagnosticStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mitigationInstruction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMitigationInstruction(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["validationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidationName(val)
        }
        return nil
    }
    return res
}
// GetMitigationInstruction gets the mitigationInstruction property value. Instruction on how to mitigate a failed validation
func (m *ManagedAppDiagnosticStatus) GetMitigationInstruction()(*string) {
    return m.mitigationInstruction
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedAppDiagnosticStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. The state of the operation
func (m *ManagedAppDiagnosticStatus) GetState()(*string) {
    return m.state
}
// GetValidationName gets the validationName property value. The validation friendly name
func (m *ManagedAppDiagnosticStatus) GetValidationName()(*string) {
    return m.validationName
}
// Serialize serializes information the current object
func (m *ManagedAppDiagnosticStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("mitigationInstruction", m.GetMitigationInstruction())
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
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("validationName", m.GetValidationName())
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
func (m *ManagedAppDiagnosticStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMitigationInstruction sets the mitigationInstruction property value. Instruction on how to mitigate a failed validation
func (m *ManagedAppDiagnosticStatus) SetMitigationInstruction(value *string)() {
    m.mitigationInstruction = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedAppDiagnosticStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. The state of the operation
func (m *ManagedAppDiagnosticStatus) SetState(value *string)() {
    m.state = value
}
// SetValidationName sets the validationName property value. The validation friendly name
func (m *ManagedAppDiagnosticStatus) SetValidationName(value *string)() {
    m.validationName = value
}
