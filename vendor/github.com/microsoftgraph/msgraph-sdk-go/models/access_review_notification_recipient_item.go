package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewNotificationRecipientItem 
type AccessReviewNotificationRecipientItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Determines the recipient of the notification email.
    notificationRecipientScope AccessReviewNotificationRecipientScopeable
    // Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients, which sends review completion notifications to the recipients.
    notificationTemplateType *string
    // The OdataType property
    odataType *string
}
// NewAccessReviewNotificationRecipientItem instantiates a new accessReviewNotificationRecipientItem and sets the default values.
func NewAccessReviewNotificationRecipientItem()(*AccessReviewNotificationRecipientItem) {
    m := &AccessReviewNotificationRecipientItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessReviewNotificationRecipientItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewNotificationRecipientItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewNotificationRecipientItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewNotificationRecipientItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewNotificationRecipientItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notificationRecipientScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewNotificationRecipientScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationRecipientScope(val.(AccessReviewNotificationRecipientScopeable))
        }
        return nil
    }
    res["notificationTemplateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTemplateType(val)
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
    return res
}
// GetNotificationRecipientScope gets the notificationRecipientScope property value. Determines the recipient of the notification email.
func (m *AccessReviewNotificationRecipientItem) GetNotificationRecipientScope()(AccessReviewNotificationRecipientScopeable) {
    return m.notificationRecipientScope
}
// GetNotificationTemplateType gets the notificationTemplateType property value. Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients, which sends review completion notifications to the recipients.
func (m *AccessReviewNotificationRecipientItem) GetNotificationTemplateType()(*string) {
    return m.notificationTemplateType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewNotificationRecipientItem) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AccessReviewNotificationRecipientItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("notificationRecipientScope", m.GetNotificationRecipientScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTemplateType", m.GetNotificationTemplateType())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewNotificationRecipientItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNotificationRecipientScope sets the notificationRecipientScope property value. Determines the recipient of the notification email.
func (m *AccessReviewNotificationRecipientItem) SetNotificationRecipientScope(value AccessReviewNotificationRecipientScopeable)() {
    m.notificationRecipientScope = value
}
// SetNotificationTemplateType sets the notificationTemplateType property value. Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients, which sends review completion notifications to the recipients.
func (m *AccessReviewNotificationRecipientItem) SetNotificationTemplateType(value *string)() {
    m.notificationTemplateType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewNotificationRecipientItem) SetOdataType(value *string)() {
    m.odataType = value
}
