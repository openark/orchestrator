package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityProtectionRoot 
type IdentityProtectionRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Risk detection in Azure AD Identity Protection and the associated information about the detection.
    riskDetections []RiskDetectionable
    // Azure AD service principals that are at risk.
    riskyServicePrincipals []RiskyServicePrincipalable
    // Users that are flagged as at-risk by Azure AD Identity Protection.
    riskyUsers []RiskyUserable
    // Represents information about detected at-risk service principals in an Azure AD tenant.
    servicePrincipalRiskDetections []ServicePrincipalRiskDetectionable
}
// NewIdentityProtectionRoot instantiates a new IdentityProtectionRoot and sets the default values.
func NewIdentityProtectionRoot()(*IdentityProtectionRoot) {
    m := &IdentityProtectionRoot{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIdentityProtectionRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityProtectionRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityProtectionRoot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityProtectionRoot) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityProtectionRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["riskDetections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRiskDetectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskDetectionable, len(val))
            for i, v := range val {
                res[i] = v.(RiskDetectionable)
            }
            m.SetRiskDetections(res)
        }
        return nil
    }
    res["riskyServicePrincipals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRiskyServicePrincipalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskyServicePrincipalable, len(val))
            for i, v := range val {
                res[i] = v.(RiskyServicePrincipalable)
            }
            m.SetRiskyServicePrincipals(res)
        }
        return nil
    }
    res["riskyUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRiskyUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskyUserable, len(val))
            for i, v := range val {
                res[i] = v.(RiskyUserable)
            }
            m.SetRiskyUsers(res)
        }
        return nil
    }
    res["servicePrincipalRiskDetections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalRiskDetectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalRiskDetectionable, len(val))
            for i, v := range val {
                res[i] = v.(ServicePrincipalRiskDetectionable)
            }
            m.SetServicePrincipalRiskDetections(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IdentityProtectionRoot) GetOdataType()(*string) {
    return m.odataType
}
// GetRiskDetections gets the riskDetections property value. Risk detection in Azure AD Identity Protection and the associated information about the detection.
func (m *IdentityProtectionRoot) GetRiskDetections()([]RiskDetectionable) {
    return m.riskDetections
}
// GetRiskyServicePrincipals gets the riskyServicePrincipals property value. Azure AD service principals that are at risk.
func (m *IdentityProtectionRoot) GetRiskyServicePrincipals()([]RiskyServicePrincipalable) {
    return m.riskyServicePrincipals
}
// GetRiskyUsers gets the riskyUsers property value. Users that are flagged as at-risk by Azure AD Identity Protection.
func (m *IdentityProtectionRoot) GetRiskyUsers()([]RiskyUserable) {
    return m.riskyUsers
}
// GetServicePrincipalRiskDetections gets the servicePrincipalRiskDetections property value. Represents information about detected at-risk service principals in an Azure AD tenant.
func (m *IdentityProtectionRoot) GetServicePrincipalRiskDetections()([]ServicePrincipalRiskDetectionable) {
    return m.servicePrincipalRiskDetections
}
// Serialize serializes information the current object
func (m *IdentityProtectionRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRiskDetections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRiskDetections()))
        for i, v := range m.GetRiskDetections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("riskDetections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskyServicePrincipals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRiskyServicePrincipals()))
        for i, v := range m.GetRiskyServicePrincipals() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("riskyServicePrincipals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskyUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRiskyUsers()))
        for i, v := range m.GetRiskyUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("riskyUsers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalRiskDetections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePrincipalRiskDetections()))
        for i, v := range m.GetServicePrincipalRiskDetections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("servicePrincipalRiskDetections", cast)
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
func (m *IdentityProtectionRoot) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IdentityProtectionRoot) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRiskDetections sets the riskDetections property value. Risk detection in Azure AD Identity Protection and the associated information about the detection.
func (m *IdentityProtectionRoot) SetRiskDetections(value []RiskDetectionable)() {
    m.riskDetections = value
}
// SetRiskyServicePrincipals sets the riskyServicePrincipals property value. Azure AD service principals that are at risk.
func (m *IdentityProtectionRoot) SetRiskyServicePrincipals(value []RiskyServicePrincipalable)() {
    m.riskyServicePrincipals = value
}
// SetRiskyUsers sets the riskyUsers property value. Users that are flagged as at-risk by Azure AD Identity Protection.
func (m *IdentityProtectionRoot) SetRiskyUsers(value []RiskyUserable)() {
    m.riskyUsers = value
}
// SetServicePrincipalRiskDetections sets the servicePrincipalRiskDetections property value. Represents information about detected at-risk service principals in an Azure AD tenant.
func (m *IdentityProtectionRoot) SetServicePrincipalRiskDetections(value []ServicePrincipalRiskDetectionable)() {
    m.servicePrincipalRiskDetections = value
}
