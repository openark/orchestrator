package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19 "github.com/microsoftgraph/msgraph-sdk-go/models/callrecords"
)

// CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse 
type CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
    // The value property
    value []iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.DirectRoutingLogRowable
}
// NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse instantiates a new CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse and sets the default values.
func NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse()(*CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse) {
    m := &CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreateDirectRoutingLogRowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.DirectRoutingLogRowable, len(val))
            for i, v := range val {
                res[i] = v.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.DirectRoutingLogRowable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse) GetValue()([]iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.DirectRoutingLogRowable) {
    return m.value
}
// Serialize serializes information the current object
func (m *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *CallRecordsMicrosoftGraphCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse) SetValue(value []iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.DirectRoutingLogRowable)() {
    m.value = value
}
