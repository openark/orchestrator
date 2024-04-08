package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationalUrl 
type InformationalUrl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // CDN URL to the application's logo, Read-only.
    logoUrl *string
    // Link to the application's marketing page. For example, https://www.contoso.com/app/marketing
    marketingUrl *string
    // The OdataType property
    odataType *string
    // Link to the application's privacy statement. For example, https://www.contoso.com/app/privacy
    privacyStatementUrl *string
    // Link to the application's support page. For example, https://www.contoso.com/app/support
    supportUrl *string
    // Link to the application's terms of service statement. For example, https://www.contoso.com/app/termsofservice
    termsOfServiceUrl *string
}
// NewInformationalUrl instantiates a new informationalUrl and sets the default values.
func NewInformationalUrl()(*InformationalUrl) {
    m := &InformationalUrl{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInformationalUrlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationalUrlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationalUrl(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationalUrl) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationalUrl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["logoUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoUrl(val)
        }
        return nil
    }
    res["marketingUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarketingUrl(val)
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
    res["privacyStatementUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyStatementUrl(val)
        }
        return nil
    }
    res["supportUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportUrl(val)
        }
        return nil
    }
    res["termsOfServiceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsOfServiceUrl(val)
        }
        return nil
    }
    return res
}
// GetLogoUrl gets the logoUrl property value. CDN URL to the application's logo, Read-only.
func (m *InformationalUrl) GetLogoUrl()(*string) {
    return m.logoUrl
}
// GetMarketingUrl gets the marketingUrl property value. Link to the application's marketing page. For example, https://www.contoso.com/app/marketing
func (m *InformationalUrl) GetMarketingUrl()(*string) {
    return m.marketingUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InformationalUrl) GetOdataType()(*string) {
    return m.odataType
}
// GetPrivacyStatementUrl gets the privacyStatementUrl property value. Link to the application's privacy statement. For example, https://www.contoso.com/app/privacy
func (m *InformationalUrl) GetPrivacyStatementUrl()(*string) {
    return m.privacyStatementUrl
}
// GetSupportUrl gets the supportUrl property value. Link to the application's support page. For example, https://www.contoso.com/app/support
func (m *InformationalUrl) GetSupportUrl()(*string) {
    return m.supportUrl
}
// GetTermsOfServiceUrl gets the termsOfServiceUrl property value. Link to the application's terms of service statement. For example, https://www.contoso.com/app/termsofservice
func (m *InformationalUrl) GetTermsOfServiceUrl()(*string) {
    return m.termsOfServiceUrl
}
// Serialize serializes information the current object
func (m *InformationalUrl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("logoUrl", m.GetLogoUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("marketingUrl", m.GetMarketingUrl())
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
        err := writer.WriteStringValue("privacyStatementUrl", m.GetPrivacyStatementUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("supportUrl", m.GetSupportUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("termsOfServiceUrl", m.GetTermsOfServiceUrl())
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
func (m *InformationalUrl) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLogoUrl sets the logoUrl property value. CDN URL to the application's logo, Read-only.
func (m *InformationalUrl) SetLogoUrl(value *string)() {
    m.logoUrl = value
}
// SetMarketingUrl sets the marketingUrl property value. Link to the application's marketing page. For example, https://www.contoso.com/app/marketing
func (m *InformationalUrl) SetMarketingUrl(value *string)() {
    m.marketingUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InformationalUrl) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrivacyStatementUrl sets the privacyStatementUrl property value. Link to the application's privacy statement. For example, https://www.contoso.com/app/privacy
func (m *InformationalUrl) SetPrivacyStatementUrl(value *string)() {
    m.privacyStatementUrl = value
}
// SetSupportUrl sets the supportUrl property value. Link to the application's support page. For example, https://www.contoso.com/app/support
func (m *InformationalUrl) SetSupportUrl(value *string)() {
    m.supportUrl = value
}
// SetTermsOfServiceUrl sets the termsOfServiceUrl property value. Link to the application's terms of service statement. For example, https://www.contoso.com/app/termsofservice
func (m *InformationalUrl) SetTermsOfServiceUrl(value *string)() {
    m.termsOfServiceUrl = value
}
