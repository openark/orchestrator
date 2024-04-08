package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Shared 
type Shared struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The identity of the owner of the shared item. Read-only.
    owner IdentitySetable
    // Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only.
    scope *string
    // The identity of the user who shared the item. Read-only.
    sharedBy IdentitySetable
    // The UTC date and time when the item was shared. Read-only.
    sharedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewShared instantiates a new shared and sets the default values.
func NewShared()(*Shared) {
    m := &Shared{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSharedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShared(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Shared) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Shared) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(IdentitySetable))
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
    res["sharedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["sharedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedDateTime(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Shared) GetOdataType()(*string) {
    return m.odataType
}
// GetOwner gets the owner property value. The identity of the owner of the shared item. Read-only.
func (m *Shared) GetOwner()(IdentitySetable) {
    return m.owner
}
// GetScope gets the scope property value. Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only.
func (m *Shared) GetScope()(*string) {
    return m.scope
}
// GetSharedBy gets the sharedBy property value. The identity of the user who shared the item. Read-only.
func (m *Shared) GetSharedBy()(IdentitySetable) {
    return m.sharedBy
}
// GetSharedDateTime gets the sharedDateTime property value. The UTC date and time when the item was shared. Read-only.
func (m *Shared) GetSharedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.sharedDateTime
}
// Serialize serializes information the current object
func (m *Shared) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
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
        err := writer.WriteObjectValue("sharedBy", m.GetSharedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("sharedDateTime", m.GetSharedDateTime())
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
func (m *Shared) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Shared) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOwner sets the owner property value. The identity of the owner of the shared item. Read-only.
func (m *Shared) SetOwner(value IdentitySetable)() {
    m.owner = value
}
// SetScope sets the scope property value. Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only.
func (m *Shared) SetScope(value *string)() {
    m.scope = value
}
// SetSharedBy sets the sharedBy property value. The identity of the user who shared the item. Read-only.
func (m *Shared) SetSharedBy(value IdentitySetable)() {
    m.sharedBy = value
}
// SetSharedDateTime sets the sharedDateTime property value. The UTC date and time when the item was shared. Read-only.
func (m *Shared) SetSharedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sharedDateTime = value
}
