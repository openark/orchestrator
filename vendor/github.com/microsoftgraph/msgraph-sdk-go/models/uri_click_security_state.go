package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UriClickSecurityState 
type UriClickSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The clickAction property
    clickAction *string
    // The clickDateTime property
    clickDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *string
    // The OdataType property
    odataType *string
    // The sourceId property
    sourceId *string
    // The uriDomain property
    uriDomain *string
    // The verdict property
    verdict *string
}
// NewUriClickSecurityState instantiates a new uriClickSecurityState and sets the default values.
func NewUriClickSecurityState()(*UriClickSecurityState) {
    m := &UriClickSecurityState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUriClickSecurityStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUriClickSecurityStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUriClickSecurityState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UriClickSecurityState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClickAction gets the clickAction property value. The clickAction property
func (m *UriClickSecurityState) GetClickAction()(*string) {
    return m.clickAction
}
// GetClickDateTime gets the clickDateTime property value. The clickDateTime property
func (m *UriClickSecurityState) GetClickDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.clickDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UriClickSecurityState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clickAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClickAction(val)
        }
        return nil
    }
    res["clickDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClickDateTime(val)
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
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["uriDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUriDomain(val)
        }
        return nil
    }
    res["verdict"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerdict(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *UriClickSecurityState) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UriClickSecurityState) GetOdataType()(*string) {
    return m.odataType
}
// GetSourceId gets the sourceId property value. The sourceId property
func (m *UriClickSecurityState) GetSourceId()(*string) {
    return m.sourceId
}
// GetUriDomain gets the uriDomain property value. The uriDomain property
func (m *UriClickSecurityState) GetUriDomain()(*string) {
    return m.uriDomain
}
// GetVerdict gets the verdict property value. The verdict property
func (m *UriClickSecurityState) GetVerdict()(*string) {
    return m.verdict
}
// Serialize serializes information the current object
func (m *UriClickSecurityState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clickAction", m.GetClickAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("clickDateTime", m.GetClickDateTime())
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
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uriDomain", m.GetUriDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("verdict", m.GetVerdict())
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
func (m *UriClickSecurityState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClickAction sets the clickAction property value. The clickAction property
func (m *UriClickSecurityState) SetClickAction(value *string)() {
    m.clickAction = value
}
// SetClickDateTime sets the clickDateTime property value. The clickDateTime property
func (m *UriClickSecurityState) SetClickDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.clickDateTime = value
}
// SetId sets the id property value. The id property
func (m *UriClickSecurityState) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UriClickSecurityState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSourceId sets the sourceId property value. The sourceId property
func (m *UriClickSecurityState) SetSourceId(value *string)() {
    m.sourceId = value
}
// SetUriDomain sets the uriDomain property value. The uriDomain property
func (m *UriClickSecurityState) SetUriDomain(value *string)() {
    m.uriDomain = value
}
// SetVerdict sets the verdict property value. The verdict property
func (m *UriClickSecurityState) SetVerdict(value *string)() {
    m.verdict = value
}
