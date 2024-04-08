package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OptionalClaim 
type OptionalClaim struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the optional claim specified in the name property.
    additionalProperties []string
    // If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience for the specific task requested by the end user. The default value is false.
    essential *bool
    // The name of the optional claim.
    name *string
    // The OdataType property
    odataType *string
    // The source (directory object) of the claim. There are predefined claims and user-defined claims from extension properties. If the source value is null, the claim is a predefined optional claim. If the source value is user, the value in the name property is the extension property from the user object.
    source *string
}
// NewOptionalClaim instantiates a new optionalClaim and sets the default values.
func NewOptionalClaim()(*OptionalClaim) {
    m := &OptionalClaim{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOptionalClaimFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOptionalClaimFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOptionalClaim(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OptionalClaim) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdditionalProperties gets the additionalProperties property value. Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the optional claim specified in the name property.
func (m *OptionalClaim) GetAdditionalProperties()([]string) {
    return m.additionalProperties
}
// GetEssential gets the essential property value. If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience for the specific task requested by the end user. The default value is false.
func (m *OptionalClaim) GetEssential()(*bool) {
    return m.essential
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OptionalClaim) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAdditionalProperties(res)
        }
        return nil
    }
    res["essential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEssential(val)
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the optional claim.
func (m *OptionalClaim) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OptionalClaim) GetOdataType()(*string) {
    return m.odataType
}
// GetSource gets the source property value. The source (directory object) of the claim. There are predefined claims and user-defined claims from extension properties. If the source value is null, the claim is a predefined optional claim. If the source value is user, the value in the name property is the extension property from the user object.
func (m *OptionalClaim) GetSource()(*string) {
    return m.source
}
// Serialize serializes information the current object
func (m *OptionalClaim) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalProperties() != nil {
        err := writer.WriteCollectionOfStringValues("additionalProperties", m.GetAdditionalProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("essential", m.GetEssential())
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
    {
        err := writer.WriteStringValue("source", m.GetSource())
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
func (m *OptionalClaim) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdditionalProperties sets the additionalProperties property value. Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the optional claim specified in the name property.
func (m *OptionalClaim) SetAdditionalProperties(value []string)() {
    m.additionalProperties = value
}
// SetEssential sets the essential property value. If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience for the specific task requested by the end user. The default value is false.
func (m *OptionalClaim) SetEssential(value *bool)() {
    m.essential = value
}
// SetName sets the name property value. The name of the optional claim.
func (m *OptionalClaim) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OptionalClaim) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSource sets the source property value. The source (directory object) of the claim. There are predefined claims and user-defined claims from extension properties. If the source value is null, the claim is a predefined optional claim. If the source value is user, the value in the name property is the extension property from the user object.
func (m *OptionalClaim) SetSource(value *string)() {
    m.source = value
}
