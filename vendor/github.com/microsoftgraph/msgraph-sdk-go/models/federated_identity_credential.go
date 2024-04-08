package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FederatedIdentityCredential 
type FederatedIdentityCredential struct {
    Entity
    // The audience that can appear in the external token. This field is mandatory and should be set to api://AzureADTokenExchange for Azure AD. It says what Microsoft identity platform should accept in the aud claim in the incoming token. This value represents Azure AD in your external identity provider and has no fixed value across identity providers - you may need to create a new application registration in your identity provider to serve as the audience of this token. This field can only accept a single value and has a limit of 600 characters. Required.
    audiences []string
    // The un-validated, user-provided description of the federated identity credential. It has a limit of 600 characters. Optional.
    description *string
    // The URL of the external identity provider and must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app. It has a limit of 600 characters. Required.
    issuer *string
    // is the unique identifier for the federated identity credential, which has a limit of 120 characters and must be URL friendly. It is immutable once created. Required. Not nullable. Supports $filter (eq).
    name *string
    // Required. The identifier of the external software workload within the external identity provider. Like the audience value, it has no fixed format, as each identity provider uses their own - sometimes a GUID, sometimes a colon delimited identifier, sometimes arbitrary strings. The value here must match the sub claim within the token presented to Azure AD. The combination of issuer and subject must be unique on the app. It has a limit of 600 characters. Supports $filter (eq).
    subject *string
}
// NewFederatedIdentityCredential instantiates a new federatedIdentityCredential and sets the default values.
func NewFederatedIdentityCredential()(*FederatedIdentityCredential) {
    m := &FederatedIdentityCredential{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFederatedIdentityCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFederatedIdentityCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFederatedIdentityCredential(), nil
}
// GetAudiences gets the audiences property value. The audience that can appear in the external token. This field is mandatory and should be set to api://AzureADTokenExchange for Azure AD. It says what Microsoft identity platform should accept in the aud claim in the incoming token. This value represents Azure AD in your external identity provider and has no fixed value across identity providers - you may need to create a new application registration in your identity provider to serve as the audience of this token. This field can only accept a single value and has a limit of 600 characters. Required.
func (m *FederatedIdentityCredential) GetAudiences()([]string) {
    return m.audiences
}
// GetDescription gets the description property value. The un-validated, user-provided description of the federated identity credential. It has a limit of 600 characters. Optional.
func (m *FederatedIdentityCredential) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FederatedIdentityCredential) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audiences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAudiences(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
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
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetIssuer gets the issuer property value. The URL of the external identity provider and must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app. It has a limit of 600 characters. Required.
func (m *FederatedIdentityCredential) GetIssuer()(*string) {
    return m.issuer
}
// GetName gets the name property value. is the unique identifier for the federated identity credential, which has a limit of 120 characters and must be URL friendly. It is immutable once created. Required. Not nullable. Supports $filter (eq).
func (m *FederatedIdentityCredential) GetName()(*string) {
    return m.name
}
// GetSubject gets the subject property value. Required. The identifier of the external software workload within the external identity provider. Like the audience value, it has no fixed format, as each identity provider uses their own - sometimes a GUID, sometimes a colon delimited identifier, sometimes arbitrary strings. The value here must match the sub claim within the token presented to Azure AD. The combination of issuer and subject must be unique on the app. It has a limit of 600 characters. Supports $filter (eq).
func (m *FederatedIdentityCredential) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *FederatedIdentityCredential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAudiences() != nil {
        err = writer.WriteCollectionOfStringValues("audiences", m.GetAudiences())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudiences sets the audiences property value. The audience that can appear in the external token. This field is mandatory and should be set to api://AzureADTokenExchange for Azure AD. It says what Microsoft identity platform should accept in the aud claim in the incoming token. This value represents Azure AD in your external identity provider and has no fixed value across identity providers - you may need to create a new application registration in your identity provider to serve as the audience of this token. This field can only accept a single value and has a limit of 600 characters. Required.
func (m *FederatedIdentityCredential) SetAudiences(value []string)() {
    m.audiences = value
}
// SetDescription sets the description property value. The un-validated, user-provided description of the federated identity credential. It has a limit of 600 characters. Optional.
func (m *FederatedIdentityCredential) SetDescription(value *string)() {
    m.description = value
}
// SetIssuer sets the issuer property value. The URL of the external identity provider and must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app. It has a limit of 600 characters. Required.
func (m *FederatedIdentityCredential) SetIssuer(value *string)() {
    m.issuer = value
}
// SetName sets the name property value. is the unique identifier for the federated identity credential, which has a limit of 120 characters and must be URL friendly. It is immutable once created. Required. Not nullable. Supports $filter (eq).
func (m *FederatedIdentityCredential) SetName(value *string)() {
    m.name = value
}
// SetSubject sets the subject property value. Required. The identifier of the external software workload within the external identity provider. Like the audience value, it has no fixed format, as each identity provider uses their own - sometimes a GUID, sometimes a colon delimited identifier, sometimes arbitrary strings. The value here must match the sub claim within the token presented to Azure AD. The combination of issuer and subject must be unique on the app. It has a limit of 600 characters. Supports $filter (eq).
func (m *FederatedIdentityCredential) SetSubject(value *string)() {
    m.subject = value
}
