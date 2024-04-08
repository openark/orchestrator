package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RelatedContact 
type RelatedContact struct {
    // Indicates whether the user has been consented to access student data.
    accessConsent *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the contact. Required.
    displayName *string
    // Primary email address of the contact. Required.
    emailAddress *string
    // Mobile phone number of the contact.
    mobilePhone *string
    // The OdataType property
    odataType *string
    // The relationship property
    relationship *ContactRelationship
}
// NewRelatedContact instantiates a new relatedContact and sets the default values.
func NewRelatedContact()(*RelatedContact) {
    m := &RelatedContact{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRelatedContactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRelatedContactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedContact(), nil
}
// GetAccessConsent gets the accessConsent property value. Indicates whether the user has been consented to access student data.
func (m *RelatedContact) GetAccessConsent()(*bool) {
    return m.accessConsent
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RelatedContact) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Name of the contact. Required.
func (m *RelatedContact) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmailAddress gets the emailAddress property value. Primary email address of the contact. Required.
func (m *RelatedContact) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RelatedContact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessConsent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessConsent(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["mobilePhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
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
    res["relationship"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContactRelationship)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationship(val.(*ContactRelationship))
        }
        return nil
    }
    return res
}
// GetMobilePhone gets the mobilePhone property value. Mobile phone number of the contact.
func (m *RelatedContact) GetMobilePhone()(*string) {
    return m.mobilePhone
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RelatedContact) GetOdataType()(*string) {
    return m.odataType
}
// GetRelationship gets the relationship property value. The relationship property
func (m *RelatedContact) GetRelationship()(*ContactRelationship) {
    return m.relationship
}
// Serialize serializes information the current object
func (m *RelatedContact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("accessConsent", m.GetAccessConsent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
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
    if m.GetRelationship() != nil {
        cast := (*m.GetRelationship()).String()
        err := writer.WriteStringValue("relationship", &cast)
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
// SetAccessConsent sets the accessConsent property value. Indicates whether the user has been consented to access student data.
func (m *RelatedContact) SetAccessConsent(value *bool)() {
    m.accessConsent = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RelatedContact) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Name of the contact. Required.
func (m *RelatedContact) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmailAddress sets the emailAddress property value. Primary email address of the contact. Required.
func (m *RelatedContact) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetMobilePhone sets the mobilePhone property value. Mobile phone number of the contact.
func (m *RelatedContact) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RelatedContact) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRelationship sets the relationship property value. The relationship property
func (m *RelatedContact) SetRelationship(value *ContactRelationship)() {
    m.relationship = value
}
