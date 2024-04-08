package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Endpoint 
type Endpoint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // User-agent reported by this endpoint.
    userAgent UserAgentable
}
// NewEndpoint instantiates a new endpoint and sets the default values.
func NewEndpoint()(*Endpoint) {
    m := &Endpoint{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.callRecords.participantEndpoint":
                        return NewParticipantEndpoint(), nil
                    case "#microsoft.graph.callRecords.serviceEndpoint":
                        return NewServiceEndpoint(), nil
                }
            }
        }
    }
    return NewEndpoint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Endpoint) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Endpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["userAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserAgentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAgent(val.(UserAgentable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Endpoint) GetOdataType()(*string) {
    return m.odataType
}
// GetUserAgent gets the userAgent property value. User-agent reported by this endpoint.
func (m *Endpoint) GetUserAgent()(UserAgentable) {
    return m.userAgent
}
// Serialize serializes information the current object
func (m *Endpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("userAgent", m.GetUserAgent())
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
func (m *Endpoint) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Endpoint) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserAgent sets the userAgent property value. User-agent reported by this endpoint.
func (m *Endpoint) SetUserAgent(value UserAgentable)() {
    m.userAgent = value
}
