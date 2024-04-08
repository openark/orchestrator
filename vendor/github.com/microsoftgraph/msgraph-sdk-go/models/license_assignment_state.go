package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseAssignmentState 
type LicenseAssignmentState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The assignedByGroup property
    assignedByGroup *string
    // The disabledPlans property
    disabledPlans []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The error property
    error *string
    // The lastUpdatedDateTime property
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The skuId property
    skuId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The state property
    state *string
}
// NewLicenseAssignmentState instantiates a new licenseAssignmentState and sets the default values.
func NewLicenseAssignmentState()(*LicenseAssignmentState) {
    m := &LicenseAssignmentState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLicenseAssignmentStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLicenseAssignmentStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicenseAssignmentState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LicenseAssignmentState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignedByGroup gets the assignedByGroup property value. The assignedByGroup property
func (m *LicenseAssignmentState) GetAssignedByGroup()(*string) {
    return m.assignedByGroup
}
// GetDisabledPlans gets the disabledPlans property value. The disabledPlans property
func (m *LicenseAssignmentState) GetDisabledPlans()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.disabledPlans
}
// GetError gets the error property value. The error property
func (m *LicenseAssignmentState) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LicenseAssignmentState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedByGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedByGroup(val)
        }
        return nil
    }
    res["disabledPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID, len(val))
            for i, v := range val {
                res[i] = *(v.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID))
            }
            m.SetDisabledPlans(res)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
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
    res["skuId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *LicenseAssignmentState) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LicenseAssignmentState) GetOdataType()(*string) {
    return m.odataType
}
// GetSkuId gets the skuId property value. The skuId property
func (m *LicenseAssignmentState) GetSkuId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.skuId
}
// GetState gets the state property value. The state property
func (m *LicenseAssignmentState) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *LicenseAssignmentState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignedByGroup", m.GetAssignedByGroup())
        if err != nil {
            return err
        }
    }
    if m.GetDisabledPlans() != nil {
        err := writer.WriteCollectionOfUUIDValues("disabledPlans", m.GetDisabledPlans())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
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
        err := writer.WriteUUIDValue("skuId", m.GetSkuId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *LicenseAssignmentState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignedByGroup sets the assignedByGroup property value. The assignedByGroup property
func (m *LicenseAssignmentState) SetAssignedByGroup(value *string)() {
    m.assignedByGroup = value
}
// SetDisabledPlans sets the disabledPlans property value. The disabledPlans property
func (m *LicenseAssignmentState) SetDisabledPlans(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.disabledPlans = value
}
// SetError sets the error property value. The error property
func (m *LicenseAssignmentState) SetError(value *string)() {
    m.error = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *LicenseAssignmentState) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LicenseAssignmentState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSkuId sets the skuId property value. The skuId property
func (m *LicenseAssignmentState) SetSkuId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.skuId = value
}
// SetState sets the state property value. The state property
func (m *LicenseAssignmentState) SetState(value *string)() {
    m.state = value
}
