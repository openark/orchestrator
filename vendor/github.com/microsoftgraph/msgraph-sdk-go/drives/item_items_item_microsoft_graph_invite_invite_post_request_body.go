package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody 
type ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The expirationDateTime property
    expirationDateTime *string
    // The message property
    message *string
    // The password property
    password *string
    // The recipients property
    recipients []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveRecipientable
    // The requireSignIn property
    requireSignIn *bool
    // The retainInheritedPermissions property
    retainInheritedPermissions *bool
    // The roles property
    roles []string
    // The sendInvitation property
    sendInvitation *bool
}
// NewItemItemsItemMicrosoftGraphInviteInvitePostRequestBody instantiates a new ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody and sets the default values.
func NewItemItemsItemMicrosoftGraphInviteInvitePostRequestBody()(*ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) {
    m := &ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemMicrosoftGraphInviteInvitePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemMicrosoftGraphInviteInvitePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemMicrosoftGraphInviteInvitePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetExpirationDateTime()(*string) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["requireSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireSignIn(val)
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
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["sendInvitation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendInvitation(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetMessage()(*string) {
    return m.message
}
// GetPassword gets the password property value. The password property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetRecipients gets the recipients property value. The recipients property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetRecipients()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveRecipientable) {
    return m.recipients
}
// GetRequireSignIn gets the requireSignIn property value. The requireSignIn property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetRequireSignIn()(*bool) {
    return m.requireSignIn
}
// GetRetainInheritedPermissions gets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetRetainInheritedPermissions()(*bool) {
    return m.retainInheritedPermissions
}
// GetRoles gets the roles property value. The roles property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetRoles()([]string) {
    return m.roles
}
// GetSendInvitation gets the sendInvitation property value. The sendInvitation property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) GetSendInvitation()(*bool) {
    return m.sendInvitation
}
// Serialize serializes information the current object
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expirationDateTime", m.GetExpirationDateTime())
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
    if m.GetRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requireSignIn", m.GetRequireSignIn())
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
    if m.GetRoles() != nil {
        err := writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendInvitation", m.GetSendInvitation())
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
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetExpirationDateTime(value *string)() {
    m.expirationDateTime = value
}
// SetMessage sets the message property value. The message property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetPassword sets the password property value. The password property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetRecipients(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveRecipientable)() {
    m.recipients = value
}
// SetRequireSignIn sets the requireSignIn property value. The requireSignIn property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetRequireSignIn(value *bool)() {
    m.requireSignIn = value
}
// SetRetainInheritedPermissions sets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetRetainInheritedPermissions(value *bool)() {
    m.retainInheritedPermissions = value
}
// SetRoles sets the roles property value. The roles property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetRoles(value []string)() {
    m.roles = value
}
// SetSendInvitation sets the sendInvitation property value. The sendInvitation property
func (m *ItemItemsItemMicrosoftGraphInviteInvitePostRequestBody) SetSendInvitation(value *bool)() {
    m.sendInvitation = value
}
