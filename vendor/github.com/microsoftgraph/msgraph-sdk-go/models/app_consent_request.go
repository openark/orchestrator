package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppConsentRequest 
type AppConsentRequest struct {
    Entity
    // The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
    appDisplayName *string
    // The identifier of the application. Required. Supports $filter (eq only) and $orderby.
    appId *string
    // A list of pending scopes waiting for approval. Required.
    pendingScopes []AppConsentRequestScopeable
    // A list of pending user consent requests. Supports $filter (eq).
    userConsentRequests []UserConsentRequestable
}
// NewAppConsentRequest instantiates a new appConsentRequest and sets the default values.
func NewAppConsentRequest()(*AppConsentRequest) {
    m := &AppConsentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppConsentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppConsentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppConsentRequest(), nil
}
// GetAppDisplayName gets the appDisplayName property value. The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
func (m *AppConsentRequest) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppId gets the appId property value. The identifier of the application. Required. Supports $filter (eq only) and $orderby.
func (m *AppConsentRequest) GetAppId()(*string) {
    return m.appId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppConsentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["pendingScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppConsentRequestScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppConsentRequestScopeable, len(val))
            for i, v := range val {
                res[i] = v.(AppConsentRequestScopeable)
            }
            m.SetPendingScopes(res)
        }
        return nil
    }
    res["userConsentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserConsentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserConsentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(UserConsentRequestable)
            }
            m.SetUserConsentRequests(res)
        }
        return nil
    }
    return res
}
// GetPendingScopes gets the pendingScopes property value. A list of pending scopes waiting for approval. Required.
func (m *AppConsentRequest) GetPendingScopes()([]AppConsentRequestScopeable) {
    return m.pendingScopes
}
// GetUserConsentRequests gets the userConsentRequests property value. A list of pending user consent requests. Supports $filter (eq).
func (m *AppConsentRequest) GetUserConsentRequests()([]UserConsentRequestable) {
    return m.userConsentRequests
}
// Serialize serializes information the current object
func (m *AppConsentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetPendingScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPendingScopes()))
        for i, v := range m.GetPendingScopes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("pendingScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserConsentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserConsentRequests()))
        for i, v := range m.GetUserConsentRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userConsentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
func (m *AppConsentRequest) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppId sets the appId property value. The identifier of the application. Required. Supports $filter (eq only) and $orderby.
func (m *AppConsentRequest) SetAppId(value *string)() {
    m.appId = value
}
// SetPendingScopes sets the pendingScopes property value. A list of pending scopes waiting for approval. Required.
func (m *AppConsentRequest) SetPendingScopes(value []AppConsentRequestScopeable)() {
    m.pendingScopes = value
}
// SetUserConsentRequests sets the userConsentRequests property value. A list of pending user consent requests. Supports $filter (eq).
func (m *AppConsentRequest) SetUserConsentRequests(value []UserConsentRequestable)() {
    m.userConsentRequests = value
}
