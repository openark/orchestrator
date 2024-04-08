package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Subscription 
type Subscription struct {
    Entity
    // Optional. Identifier of the application used to create the subscription. Read-only.
    applicationId *string
    // Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType. Use updated to receive notifications when user or group is created, updated or soft deleted.  Use deleted to receive notifications when user or group is permanently deleted.
    changeType *string
    // Required. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
    clientState *string
    // Optional. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only.
    creatorId *string
    // Optional. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
    encryptionCertificate *string
    // Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data.
    encryptionCertificateId *string
    // Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Optional. When set to true, change notifications include resource data (such as content of a chat message).
    includeResourceData *bool
    // Optional. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
    latestSupportedTlsVersion *string
    // Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved, reauthorizationRequired, and missed notifications. This URL must make use of the HTTPS protocol.
    lifecycleNotificationUrl *string
    // Optional. OData query options for specifying value for the targeting resource. Clients receive notifications when resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.  Supported only for Universal Print Service. For more information, see Subscribe to change notifications from cloud printing APIs using Microsoft Graph.
    notificationQueryOptions *string
    // Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol.
    notificationUrl *string
    // Optional. The app ID that the subscription service can use to generate the validation token. This allows the client to validate the authenticity of the notification received.
    notificationUrlAppId *string
    // Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource.
    resource *string
}
// NewSubscription instantiates a new subscription and sets the default values.
func NewSubscription()(*Subscription) {
    m := &Subscription{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSubscriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubscriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubscription(), nil
}
// GetApplicationId gets the applicationId property value. Optional. Identifier of the application used to create the subscription. Read-only.
func (m *Subscription) GetApplicationId()(*string) {
    return m.applicationId
}
// GetChangeType gets the changeType property value. Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType. Use updated to receive notifications when user or group is created, updated or soft deleted.  Use deleted to receive notifications when user or group is permanently deleted.
func (m *Subscription) GetChangeType()(*string) {
    return m.changeType
}
// GetClientState gets the clientState property value. Required. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
func (m *Subscription) GetClientState()(*string) {
    return m.clientState
}
// GetCreatorId gets the creatorId property value. Optional. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only.
func (m *Subscription) GetCreatorId()(*string) {
    return m.creatorId
}
// GetEncryptionCertificate gets the encryptionCertificate property value. Optional. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
func (m *Subscription) GetEncryptionCertificate()(*string) {
    return m.encryptionCertificate
}
// GetEncryptionCertificateId gets the encryptionCertificateId property value. Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data.
func (m *Subscription) GetEncryptionCertificateId()(*string) {
    return m.encryptionCertificateId
}
// GetExpirationDateTime gets the expirationDateTime property value. Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
func (m *Subscription) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Subscription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["changeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeType(val)
        }
        return nil
    }
    res["clientState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientState(val)
        }
        return nil
    }
    res["creatorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatorId(val)
        }
        return nil
    }
    res["encryptionCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionCertificate(val)
        }
        return nil
    }
    res["encryptionCertificateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionCertificateId(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["includeResourceData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeResourceData(val)
        }
        return nil
    }
    res["latestSupportedTlsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestSupportedTlsVersion(val)
        }
        return nil
    }
    res["lifecycleNotificationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifecycleNotificationUrl(val)
        }
        return nil
    }
    res["notificationQueryOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationQueryOptions(val)
        }
        return nil
    }
    res["notificationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationUrl(val)
        }
        return nil
    }
    res["notificationUrlAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationUrlAppId(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    return res
}
// GetIncludeResourceData gets the includeResourceData property value. Optional. When set to true, change notifications include resource data (such as content of a chat message).
func (m *Subscription) GetIncludeResourceData()(*bool) {
    return m.includeResourceData
}
// GetLatestSupportedTlsVersion gets the latestSupportedTlsVersion property value. Optional. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
func (m *Subscription) GetLatestSupportedTlsVersion()(*string) {
    return m.latestSupportedTlsVersion
}
// GetLifecycleNotificationUrl gets the lifecycleNotificationUrl property value. Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved, reauthorizationRequired, and missed notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) GetLifecycleNotificationUrl()(*string) {
    return m.lifecycleNotificationUrl
}
// GetNotificationQueryOptions gets the notificationQueryOptions property value. Optional. OData query options for specifying value for the targeting resource. Clients receive notifications when resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.  Supported only for Universal Print Service. For more information, see Subscribe to change notifications from cloud printing APIs using Microsoft Graph.
func (m *Subscription) GetNotificationQueryOptions()(*string) {
    return m.notificationQueryOptions
}
// GetNotificationUrl gets the notificationUrl property value. Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) GetNotificationUrl()(*string) {
    return m.notificationUrl
}
// GetNotificationUrlAppId gets the notificationUrlAppId property value. Optional. The app ID that the subscription service can use to generate the validation token. This allows the client to validate the authenticity of the notification received.
func (m *Subscription) GetNotificationUrlAppId()(*string) {
    return m.notificationUrlAppId
}
// GetResource gets the resource property value. Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource.
func (m *Subscription) GetResource()(*string) {
    return m.resource
}
// Serialize serializes information the current object
func (m *Subscription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("changeType", m.GetChangeType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientState", m.GetClientState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creatorId", m.GetCreatorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("encryptionCertificate", m.GetEncryptionCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("encryptionCertificateId", m.GetEncryptionCertificateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeResourceData", m.GetIncludeResourceData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("latestSupportedTlsVersion", m.GetLatestSupportedTlsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lifecycleNotificationUrl", m.GetLifecycleNotificationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationQueryOptions", m.GetNotificationQueryOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationUrl", m.GetNotificationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationUrlAppId", m.GetNotificationUrlAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationId sets the applicationId property value. Optional. Identifier of the application used to create the subscription. Read-only.
func (m *Subscription) SetApplicationId(value *string)() {
    m.applicationId = value
}
// SetChangeType sets the changeType property value. Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list. Note:  Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType. Use updated to receive notifications when user or group is created, updated or soft deleted.  Use deleted to receive notifications when user or group is permanently deleted.
func (m *Subscription) SetChangeType(value *string)() {
    m.changeType = value
}
// SetClientState sets the clientState property value. Required. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
func (m *Subscription) SetClientState(value *string)() {
    m.clientState = value
}
// SetCreatorId sets the creatorId property value. Optional. Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only.
func (m *Subscription) SetCreatorId(value *string)() {
    m.creatorId = value
}
// SetEncryptionCertificate sets the encryptionCertificate property value. Optional. A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional but required when includeResourceData is true.
func (m *Subscription) SetEncryptionCertificate(value *string)() {
    m.encryptionCertificate = value
}
// SetEncryptionCertificateId sets the encryptionCertificateId property value. Optional. A custom app-provided identifier to help identify the certificate needed to decrypt resource data.
func (m *Subscription) SetEncryptionCertificateId(value *string)() {
    m.encryptionCertificateId = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to. For the maximum supported subscription length of time, see the table below.
func (m *Subscription) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetIncludeResourceData sets the includeResourceData property value. Optional. When set to true, change notifications include resource data (such as content of a chat message).
func (m *Subscription) SetIncludeResourceData(value *bool)() {
    m.includeResourceData = value
}
// SetLatestSupportedTlsVersion sets the latestSupportedTlsVersion property value. Optional. Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
func (m *Subscription) SetLatestSupportedTlsVersion(value *string)() {
    m.latestSupportedTlsVersion = value
}
// SetLifecycleNotificationUrl sets the lifecycleNotificationUrl property value. Optional. The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved, reauthorizationRequired, and missed notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) SetLifecycleNotificationUrl(value *string)() {
    m.lifecycleNotificationUrl = value
}
// SetNotificationQueryOptions sets the notificationQueryOptions property value. Optional. OData query options for specifying value for the targeting resource. Clients receive notifications when resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property. For example, when the print job is completed or when a print job resource isFetchable property value becomes true etc.  Supported only for Universal Print Service. For more information, see Subscribe to change notifications from cloud printing APIs using Microsoft Graph.
func (m *Subscription) SetNotificationQueryOptions(value *string)() {
    m.notificationQueryOptions = value
}
// SetNotificationUrl sets the notificationUrl property value. Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol.
func (m *Subscription) SetNotificationUrl(value *string)() {
    m.notificationUrl = value
}
// SetNotificationUrlAppId sets the notificationUrlAppId property value. Optional. The app ID that the subscription service can use to generate the validation token. This allows the client to validate the authenticity of the notification received.
func (m *Subscription) SetNotificationUrlAppId(value *string)() {
    m.notificationUrlAppId = value
}
// SetResource sets the resource property value. Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource.
func (m *Subscription) SetResource(value *string)() {
    m.resource = value
}
