package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Recipient 
type Recipient struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The recipient's email address.
    emailAddress EmailAddressable
    // The OdataType property
    odataType *string
}
// NewRecipient instantiates a new recipient and sets the default values.
func NewRecipient()(*Recipient) {
    m := &Recipient{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.attendee":
                        return NewAttendee(), nil
                    case "#microsoft.graph.attendeeBase":
                        return NewAttendeeBase(), nil
                }
            }
        }
    }
    return NewRecipient(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Recipient) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmailAddress gets the emailAddress property value. The recipient's email address.
func (m *Recipient) GetEmailAddress()(EmailAddressable) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Recipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val.(EmailAddressable))
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Recipient) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Recipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("emailAddress", m.GetEmailAddress())
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
func (m *Recipient) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmailAddress sets the emailAddress property value. The recipient's email address.
func (m *Recipient) SetEmailAddress(value EmailAddressable)() {
    m.emailAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Recipient) SetOdataType(value *string)() {
    m.odataType = value
}
