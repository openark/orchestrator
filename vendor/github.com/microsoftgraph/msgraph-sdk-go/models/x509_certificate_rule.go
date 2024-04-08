package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateRule 
type X509CertificateRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The identifier of the X.509 certificate. Required.
    identifier *string
    // The OdataType property
    odataType *string
    // The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue. Required.
    x509CertificateAuthenticationMode *X509CertificateAuthenticationMode
    // The type of the X.509 certificate mode configuration rule. The possible values are: issuerSubject, policyOID, unknownFutureValue. Required.
    x509CertificateRuleType *X509CertificateRuleType
}
// NewX509CertificateRule instantiates a new x509CertificateRule and sets the default values.
func NewX509CertificateRule()(*X509CertificateRule) {
    m := &X509CertificateRule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateX509CertificateRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateX509CertificateRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *X509CertificateRule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *X509CertificateRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
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
    res["x509CertificateAuthenticationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateAuthenticationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateAuthenticationMode(val.(*X509CertificateAuthenticationMode))
        }
        return nil
    }
    res["x509CertificateRuleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseX509CertificateRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX509CertificateRuleType(val.(*X509CertificateRuleType))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier of the X.509 certificate. Required.
func (m *X509CertificateRule) GetIdentifier()(*string) {
    return m.identifier
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *X509CertificateRule) GetOdataType()(*string) {
    return m.odataType
}
// GetX509CertificateAuthenticationMode gets the x509CertificateAuthenticationMode property value. The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue. Required.
func (m *X509CertificateRule) GetX509CertificateAuthenticationMode()(*X509CertificateAuthenticationMode) {
    return m.x509CertificateAuthenticationMode
}
// GetX509CertificateRuleType gets the x509CertificateRuleType property value. The type of the X.509 certificate mode configuration rule. The possible values are: issuerSubject, policyOID, unknownFutureValue. Required.
func (m *X509CertificateRule) GetX509CertificateRuleType()(*X509CertificateRuleType) {
    return m.x509CertificateRuleType
}
// Serialize serializes information the current object
func (m *X509CertificateRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
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
    if m.GetX509CertificateAuthenticationMode() != nil {
        cast := (*m.GetX509CertificateAuthenticationMode()).String()
        err := writer.WriteStringValue("x509CertificateAuthenticationMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetX509CertificateRuleType() != nil {
        cast := (*m.GetX509CertificateRuleType()).String()
        err := writer.WriteStringValue("x509CertificateRuleType", &cast)
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
func (m *X509CertificateRule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIdentifier sets the identifier property value. The identifier of the X.509 certificate. Required.
func (m *X509CertificateRule) SetIdentifier(value *string)() {
    m.identifier = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *X509CertificateRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetX509CertificateAuthenticationMode sets the x509CertificateAuthenticationMode property value. The type of strong authentication mode. The possible values are: x509CertificateSingleFactor, x509CertificateMultiFactor, unknownFutureValue. Required.
func (m *X509CertificateRule) SetX509CertificateAuthenticationMode(value *X509CertificateAuthenticationMode)() {
    m.x509CertificateAuthenticationMode = value
}
// SetX509CertificateRuleType sets the x509CertificateRuleType property value. The type of the X.509 certificate mode configuration rule. The possible values are: issuerSubject, policyOID, unknownFutureValue. Required.
func (m *X509CertificateRule) SetX509CertificateRuleType(value *X509CertificateRuleType)() {
    m.x509CertificateRuleType = value
}
