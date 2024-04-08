package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Fido2KeyRestrictions 
type Fido2KeyRestrictions struct {
    // A collection of Authenticator Attestation GUIDs. AADGUIDs define key types and manufacturers.
    aaGuids []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Enforcement type. Possible values are: allow, block.
    enforcementType *Fido2RestrictionEnforcementType
    // Determines if the configured key enforcement is enabled.
    isEnforced *bool
    // The OdataType property
    odataType *string
}
// NewFido2KeyRestrictions instantiates a new fido2KeyRestrictions and sets the default values.
func NewFido2KeyRestrictions()(*Fido2KeyRestrictions) {
    m := &Fido2KeyRestrictions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFido2KeyRestrictionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFido2KeyRestrictionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFido2KeyRestrictions(), nil
}
// GetAaGuids gets the aaGuids property value. A collection of Authenticator Attestation GUIDs. AADGUIDs define key types and manufacturers.
func (m *Fido2KeyRestrictions) GetAaGuids()([]string) {
    return m.aaGuids
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Fido2KeyRestrictions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnforcementType gets the enforcementType property value. Enforcement type. Possible values are: allow, block.
func (m *Fido2KeyRestrictions) GetEnforcementType()(*Fido2RestrictionEnforcementType) {
    return m.enforcementType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Fido2KeyRestrictions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aaGuids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAaGuids(res)
        }
        return nil
    }
    res["enforcementType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFido2RestrictionEnforcementType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcementType(val.(*Fido2RestrictionEnforcementType))
        }
        return nil
    }
    res["isEnforced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnforced(val)
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
    return res
}
// GetIsEnforced gets the isEnforced property value. Determines if the configured key enforcement is enabled.
func (m *Fido2KeyRestrictions) GetIsEnforced()(*bool) {
    return m.isEnforced
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Fido2KeyRestrictions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Fido2KeyRestrictions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAaGuids() != nil {
        err := writer.WriteCollectionOfStringValues("aaGuids", m.GetAaGuids())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcementType() != nil {
        cast := (*m.GetEnforcementType()).String()
        err := writer.WriteStringValue("enforcementType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnforced", m.GetIsEnforced())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAaGuids sets the aaGuids property value. A collection of Authenticator Attestation GUIDs. AADGUIDs define key types and manufacturers.
func (m *Fido2KeyRestrictions) SetAaGuids(value []string)() {
    m.aaGuids = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Fido2KeyRestrictions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnforcementType sets the enforcementType property value. Enforcement type. Possible values are: allow, block.
func (m *Fido2KeyRestrictions) SetEnforcementType(value *Fido2RestrictionEnforcementType)() {
    m.enforcementType = value
}
// SetIsEnforced sets the isEnforced property value. Determines if the configured key enforcement is enabled.
func (m *Fido2KeyRestrictions) SetIsEnforced(value *bool)() {
    m.isEnforced = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Fido2KeyRestrictions) SetOdataType(value *string)() {
    m.odataType = value
}
