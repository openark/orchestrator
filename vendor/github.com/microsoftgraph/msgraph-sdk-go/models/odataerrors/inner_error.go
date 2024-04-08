package odataerrors

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InnerError 
type InnerError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Client request Id as sent by the client application.
    clientRequestId *string
    // Date when the error occured.
    date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Request Id as tracked internally by the service
    requestId *string
}
// NewInnerError instantiates a new InnerError and sets the default values.
func NewInnerError()(*InnerError) {
    m := &InnerError{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInnerErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInnerErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInnerError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InnerError) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientRequestId gets the client-request-id property value. Client request Id as sent by the client application.
func (m *InnerError) GetClientRequestId()(*string) {
    return m.clientRequestId
}
// GetDate gets the date property value. Date when the error occured.
func (m *InnerError) GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InnerError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["client-request-id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientRequestId(val)
        }
        return nil
    }
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
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
    res["request-id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InnerError) GetOdataType()(*string) {
    return m.odataType
}
// GetRequestId gets the request-id property value. Request Id as tracked internally by the service
func (m *InnerError) GetRequestId()(*string) {
    return m.requestId
}
// Serialize serializes information the current object
func (m *InnerError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client-request-id", m.GetClientRequestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("date", m.GetDate())
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
        err := writer.WriteStringValue("request-id", m.GetRequestId())
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
func (m *InnerError) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientRequestId sets the client-request-id property value. Client request Id as sent by the client application.
func (m *InnerError) SetClientRequestId(value *string)() {
    m.clientRequestId = value
}
// SetDate sets the date property value. Date when the error occured.
func (m *InnerError) SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InnerError) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRequestId sets the request-id property value. Request Id as tracked internally by the service
func (m *InnerError) SetRequestId(value *string)() {
    m.requestId = value
}
