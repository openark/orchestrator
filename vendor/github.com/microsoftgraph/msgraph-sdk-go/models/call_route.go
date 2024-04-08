package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallRoute 
type CallRoute struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The final property
    final IdentitySetable
    // The OdataType property
    odataType *string
    // The original property
    original IdentitySetable
    // The routingType property
    routingType *RoutingType
}
// NewCallRoute instantiates a new callRoute and sets the default values.
func NewCallRoute()(*CallRoute) {
    m := &CallRoute{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallRouteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallRouteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallRoute(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallRoute) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallRoute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["final"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinal(val.(IdentitySetable))
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
    res["original"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginal(val.(IdentitySetable))
        }
        return nil
    }
    res["routingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoutingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoutingType(val.(*RoutingType))
        }
        return nil
    }
    return res
}
// GetFinal gets the final property value. The final property
func (m *CallRoute) GetFinal()(IdentitySetable) {
    return m.final
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CallRoute) GetOdataType()(*string) {
    return m.odataType
}
// GetOriginal gets the original property value. The original property
func (m *CallRoute) GetOriginal()(IdentitySetable) {
    return m.original
}
// GetRoutingType gets the routingType property value. The routingType property
func (m *CallRoute) GetRoutingType()(*RoutingType) {
    return m.routingType
}
// Serialize serializes information the current object
func (m *CallRoute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("final", m.GetFinal())
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
        err := writer.WriteObjectValue("original", m.GetOriginal())
        if err != nil {
            return err
        }
    }
    if m.GetRoutingType() != nil {
        cast := (*m.GetRoutingType()).String()
        err := writer.WriteStringValue("routingType", &cast)
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
func (m *CallRoute) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFinal sets the final property value. The final property
func (m *CallRoute) SetFinal(value IdentitySetable)() {
    m.final = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CallRoute) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOriginal sets the original property value. The original property
func (m *CallRoute) SetOriginal(value IdentitySetable)() {
    m.original = value
}
// SetRoutingType sets the routingType property value. The routingType property
func (m *CallRoute) SetRoutingType(value *RoutingType)() {
    m.routingType = value
}
