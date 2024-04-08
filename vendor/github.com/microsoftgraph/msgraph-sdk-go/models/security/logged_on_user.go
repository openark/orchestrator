package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LoggedOnUser 
type LoggedOnUser struct {
    // User account name of the logged-on user.
    accountName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // User account domain of the logged-on user.
    domainName *string
    // The OdataType property
    odataType *string
}
// NewLoggedOnUser instantiates a new loggedOnUser and sets the default values.
func NewLoggedOnUser()(*LoggedOnUser) {
    m := &LoggedOnUser{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLoggedOnUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLoggedOnUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLoggedOnUser(), nil
}
// GetAccountName gets the accountName property value. User account name of the logged-on user.
func (m *LoggedOnUser) GetAccountName()(*string) {
    return m.accountName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoggedOnUser) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDomainName gets the domainName property value. User account domain of the logged-on user.
func (m *LoggedOnUser) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LoggedOnUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
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
func (m *LoggedOnUser) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LoggedOnUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
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
// SetAccountName sets the accountName property value. User account name of the logged-on user.
func (m *LoggedOnUser) SetAccountName(value *string)() {
    m.accountName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoggedOnUser) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDomainName sets the domainName property value. User account domain of the logged-on user.
func (m *LoggedOnUser) SetDomainName(value *string)() {
    m.domainName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LoggedOnUser) SetOdataType(value *string)() {
    m.odataType = value
}
