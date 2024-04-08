package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainState 
type DomainState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Timestamp for when the last activity occurred. The value is updated when an operation is scheduled, the asynchronous task starts, and when the operation completes.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Type of asynchronous operation. The values can be ForceDelete or Verification
    operation *string
    // Current status of the operation.  Scheduled - Operation has been scheduled but has not started.  InProgress - Task has started and is in progress.  Failed - Operation has failed.
    status *string
}
// NewDomainState instantiates a new domainState and sets the default values.
func NewDomainState()(*DomainState) {
    m := &DomainState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDomainStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DomainState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
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
    res["operation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperation(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetLastActionDateTime gets the lastActionDateTime property value. Timestamp for when the last activity occurred. The value is updated when an operation is scheduled, the asynchronous task starts, and when the operation completes.
func (m *DomainState) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DomainState) GetOdataType()(*string) {
    return m.odataType
}
// GetOperation gets the operation property value. Type of asynchronous operation. The values can be ForceDelete or Verification
func (m *DomainState) GetOperation()(*string) {
    return m.operation
}
// GetStatus gets the status property value. Current status of the operation.  Scheduled - Operation has been scheduled but has not started.  InProgress - Task has started and is in progress.  Failed - Operation has failed.
func (m *DomainState) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *DomainState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
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
        err := writer.WriteStringValue("operation", m.GetOperation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
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
func (m *DomainState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. Timestamp for when the last activity occurred. The value is updated when an operation is scheduled, the asynchronous task starts, and when the operation completes.
func (m *DomainState) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DomainState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperation sets the operation property value. Type of asynchronous operation. The values can be ForceDelete or Verification
func (m *DomainState) SetOperation(value *string)() {
    m.operation = value
}
// SetStatus sets the status property value. Current status of the operation.  Scheduled - Operation has been scheduled but has not started.  InProgress - Task has started and is in progress.  Failed - Operation has failed.
func (m *DomainState) SetStatus(value *string)() {
    m.status = value
}
