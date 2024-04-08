package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangeNotification 
type ChangeNotification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The changeType property
    changeType *ChangeType
    // Value of the clientState property sent in the subscription request (if any). The maximum length is 255 characters. The client can check whether the change notification came from the service by comparing the values of the clientState property. The value of the clientState property sent with the subscription is compared with the value of the clientState property received with each change notification. Optional.
    clientState *string
    // (Preview) Encrypted content attached with the change notification. Only provided if encryptionCertificate and includeResourceData were defined during the subscription request and if the resource supports it. Optional.
    encryptedContent ChangeNotificationEncryptedContentable
    // Unique ID for the notification. Optional.
    id *string
    // The type of lifecycle notification if the current notification is a lifecycle notification. Optional. Supported values are missed, subscriptionRemoved, reauthorizationRequired. Optional.
    lifecycleEvent *LifecycleEventType
    // The OdataType property
    odataType *string
    // The URI of the resource that emitted the change notification relative to https://graph.microsoft.com. Required.
    resource *string
    // The content of this property depends on the type of resource being subscribed to. Optional.
    resourceData ResourceDataable
    // The expiration time for the subscription. Required.
    subscriptionExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the subscription that generated the notification.Required.
    subscriptionId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The unique identifier of the tenant from which the change notification originated. Required.
    tenantId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewChangeNotification instantiates a new changeNotification and sets the default values.
func NewChangeNotification()(*ChangeNotification) {
    m := &ChangeNotification{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChangeNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChangeNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChangeNotification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeNotification) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChangeType gets the changeType property value. The changeType property
func (m *ChangeNotification) GetChangeType()(*ChangeType) {
    return m.changeType
}
// GetClientState gets the clientState property value. Value of the clientState property sent in the subscription request (if any). The maximum length is 255 characters. The client can check whether the change notification came from the service by comparing the values of the clientState property. The value of the clientState property sent with the subscription is compared with the value of the clientState property received with each change notification. Optional.
func (m *ChangeNotification) GetClientState()(*string) {
    return m.clientState
}
// GetEncryptedContent gets the encryptedContent property value. (Preview) Encrypted content attached with the change notification. Only provided if encryptionCertificate and includeResourceData were defined during the subscription request and if the resource supports it. Optional.
func (m *ChangeNotification) GetEncryptedContent()(ChangeNotificationEncryptedContentable) {
    return m.encryptedContent
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangeNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["changeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChangeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeType(val.(*ChangeType))
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
    res["encryptedContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChangeNotificationEncryptedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedContent(val.(ChangeNotificationEncryptedContentable))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["lifecycleEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLifecycleEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifecycleEvent(val.(*LifecycleEventType))
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
    res["resourceData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceData(val.(ResourceDataable))
        }
        return nil
    }
    res["subscriptionExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionExpirationDateTime(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique ID for the notification. Optional.
func (m *ChangeNotification) GetId()(*string) {
    return m.id
}
// GetLifecycleEvent gets the lifecycleEvent property value. The type of lifecycle notification if the current notification is a lifecycle notification. Optional. Supported values are missed, subscriptionRemoved, reauthorizationRequired. Optional.
func (m *ChangeNotification) GetLifecycleEvent()(*LifecycleEventType) {
    return m.lifecycleEvent
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ChangeNotification) GetOdataType()(*string) {
    return m.odataType
}
// GetResource gets the resource property value. The URI of the resource that emitted the change notification relative to https://graph.microsoft.com. Required.
func (m *ChangeNotification) GetResource()(*string) {
    return m.resource
}
// GetResourceData gets the resourceData property value. The content of this property depends on the type of resource being subscribed to. Optional.
func (m *ChangeNotification) GetResourceData()(ResourceDataable) {
    return m.resourceData
}
// GetSubscriptionExpirationDateTime gets the subscriptionExpirationDateTime property value. The expiration time for the subscription. Required.
func (m *ChangeNotification) GetSubscriptionExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.subscriptionExpirationDateTime
}
// GetSubscriptionId gets the subscriptionId property value. The unique identifier of the subscription that generated the notification.Required.
func (m *ChangeNotification) GetSubscriptionId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.subscriptionId
}
// GetTenantId gets the tenantId property value. The unique identifier of the tenant from which the change notification originated. Required.
func (m *ChangeNotification) GetTenantId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ChangeNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetChangeType() != nil {
        cast := (*m.GetChangeType()).String()
        err := writer.WriteStringValue("changeType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientState", m.GetClientState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("encryptedContent", m.GetEncryptedContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetLifecycleEvent() != nil {
        cast := (*m.GetLifecycleEvent()).String()
        err := writer.WriteStringValue("lifecycleEvent", &cast)
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
        err := writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resourceData", m.GetResourceData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("subscriptionExpirationDateTime", m.GetSubscriptionExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("tenantId", m.GetTenantId())
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
func (m *ChangeNotification) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChangeType sets the changeType property value. The changeType property
func (m *ChangeNotification) SetChangeType(value *ChangeType)() {
    m.changeType = value
}
// SetClientState sets the clientState property value. Value of the clientState property sent in the subscription request (if any). The maximum length is 255 characters. The client can check whether the change notification came from the service by comparing the values of the clientState property. The value of the clientState property sent with the subscription is compared with the value of the clientState property received with each change notification. Optional.
func (m *ChangeNotification) SetClientState(value *string)() {
    m.clientState = value
}
// SetEncryptedContent sets the encryptedContent property value. (Preview) Encrypted content attached with the change notification. Only provided if encryptionCertificate and includeResourceData were defined during the subscription request and if the resource supports it. Optional.
func (m *ChangeNotification) SetEncryptedContent(value ChangeNotificationEncryptedContentable)() {
    m.encryptedContent = value
}
// SetId sets the id property value. Unique ID for the notification. Optional.
func (m *ChangeNotification) SetId(value *string)() {
    m.id = value
}
// SetLifecycleEvent sets the lifecycleEvent property value. The type of lifecycle notification if the current notification is a lifecycle notification. Optional. Supported values are missed, subscriptionRemoved, reauthorizationRequired. Optional.
func (m *ChangeNotification) SetLifecycleEvent(value *LifecycleEventType)() {
    m.lifecycleEvent = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChangeNotification) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResource sets the resource property value. The URI of the resource that emitted the change notification relative to https://graph.microsoft.com. Required.
func (m *ChangeNotification) SetResource(value *string)() {
    m.resource = value
}
// SetResourceData sets the resourceData property value. The content of this property depends on the type of resource being subscribed to. Optional.
func (m *ChangeNotification) SetResourceData(value ResourceDataable)() {
    m.resourceData = value
}
// SetSubscriptionExpirationDateTime sets the subscriptionExpirationDateTime property value. The expiration time for the subscription. Required.
func (m *ChangeNotification) SetSubscriptionExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.subscriptionExpirationDateTime = value
}
// SetSubscriptionId sets the subscriptionId property value. The unique identifier of the subscription that generated the notification.Required.
func (m *ChangeNotification) SetSubscriptionId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.subscriptionId = value
}
// SetTenantId sets the tenantId property value. The unique identifier of the tenant from which the change notification originated. Required.
func (m *ChangeNotification) SetTenantId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.tenantId = value
}
