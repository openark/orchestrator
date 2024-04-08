package drives

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody 
type ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The expirationDateTime property
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The message property
    message *string
    // The password property
    password *string
    // The retainInheritedPermissions property
    retainInheritedPermissions *bool
    // The scope property
    scope *string
    // The type property
    typeEscaped *string
}
// NewItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody instantiates a new ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody and sets the default values.
func NewItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody()(*ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) {
    m := &ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["retainInheritedPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetainInheritedPermissions(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) GetMessage()(*string) {
    return m.message
}
// GetPassword gets the password property value. The password property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetRetainInheritedPermissions gets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) GetRetainInheritedPermissions()(*bool) {
    return m.retainInheritedPermissions
}
// GetScope gets the scope property value. The scope property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) GetScope()(*string) {
    return m.scope
}
// GetType gets the type property value. The type property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("retainInheritedPermissions", m.GetRetainInheritedPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetMessage sets the message property value. The message property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetPassword sets the password property value. The password property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetRetainInheritedPermissions sets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) SetRetainInheritedPermissions(value *bool)() {
    m.retainInheritedPermissions = value
}
// SetScope sets the scope property value. The scope property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) SetScope(value *string)() {
    m.scope = value
}
// SetType sets the type property value. The type property
func (m *ItemItemsItemMicrosoftGraphCreateLinkCreateLinkPostRequestBody) SetType(value *string)() {
    m.typeEscaped = value
}
