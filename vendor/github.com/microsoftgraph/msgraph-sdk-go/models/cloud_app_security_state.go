package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudAppSecurityState 
type CloudAppSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Destination IP Address of the connection to the cloud application/service.
    destinationServiceIp *string
    // Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
    destinationServiceName *string
    // The OdataType property
    odataType *string
    // Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage.
    riskScore *string
}
// NewCloudAppSecurityState instantiates a new cloudAppSecurityState and sets the default values.
func NewCloudAppSecurityState()(*CloudAppSecurityState) {
    m := &CloudAppSecurityState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudAppSecurityStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudAppSecurityStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudAppSecurityState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudAppSecurityState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestinationServiceIp gets the destinationServiceIp property value. Destination IP Address of the connection to the cloud application/service.
func (m *CloudAppSecurityState) GetDestinationServiceIp()(*string) {
    return m.destinationServiceIp
}
// GetDestinationServiceName gets the destinationServiceName property value. Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
func (m *CloudAppSecurityState) GetDestinationServiceName()(*string) {
    return m.destinationServiceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudAppSecurityState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationServiceIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationServiceIp(val)
        }
        return nil
    }
    res["destinationServiceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationServiceName(val)
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
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudAppSecurityState) GetOdataType()(*string) {
    return m.odataType
}
// GetRiskScore gets the riskScore property value. Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage.
func (m *CloudAppSecurityState) GetRiskScore()(*string) {
    return m.riskScore
}
// Serialize serializes information the current object
func (m *CloudAppSecurityState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationServiceIp", m.GetDestinationServiceIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationServiceName", m.GetDestinationServiceName())
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
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
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
func (m *CloudAppSecurityState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestinationServiceIp sets the destinationServiceIp property value. Destination IP Address of the connection to the cloud application/service.
func (m *CloudAppSecurityState) SetDestinationServiceIp(value *string)() {
    m.destinationServiceIp = value
}
// SetDestinationServiceName sets the destinationServiceName property value. Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
func (m *CloudAppSecurityState) SetDestinationServiceName(value *string)() {
    m.destinationServiceName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudAppSecurityState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRiskScore sets the riskScore property value. Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage.
func (m *CloudAppSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
