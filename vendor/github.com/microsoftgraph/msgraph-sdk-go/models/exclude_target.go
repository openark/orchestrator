package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExcludeTarget 
type ExcludeTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The object identifier of an Azure Active Directory user or group.
    id *string
    // The OdataType property
    odataType *string
    // The targetType property
    targetType *AuthenticationMethodTargetType
}
// NewExcludeTarget instantiates a new excludeTarget and sets the default values.
func NewExcludeTarget()(*ExcludeTarget) {
    m := &ExcludeTarget{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExcludeTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExcludeTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExcludeTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExcludeTarget) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExcludeTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["targetType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodTargetType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*AuthenticationMethodTargetType))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The object identifier of an Azure Active Directory user or group.
func (m *ExcludeTarget) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExcludeTarget) GetOdataType()(*string) {
    return m.odataType
}
// GetTargetType gets the targetType property value. The targetType property
func (m *ExcludeTarget) GetTargetType()(*AuthenticationMethodTargetType) {
    return m.targetType
}
// Serialize serializes information the current object
func (m *ExcludeTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err := writer.WriteStringValue("targetType", &cast)
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
func (m *ExcludeTarget) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The object identifier of an Azure Active Directory user or group.
func (m *ExcludeTarget) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExcludeTarget) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTargetType sets the targetType property value. The targetType property
func (m *ExcludeTarget) SetTargetType(value *AuthenticationMethodTargetType)() {
    m.targetType = value
}
