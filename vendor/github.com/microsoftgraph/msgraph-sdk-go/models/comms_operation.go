package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommsOperation 
type CommsOperation struct {
    Entity
    // Unique Client Context string. Max limit is 256 chars.
    clientContext *string
    // The result information. Read-only.
    resultInfo ResultInfoable
    // The status property
    status *OperationStatus
}
// NewCommsOperation instantiates a new commsOperation and sets the default values.
func NewCommsOperation()(*CommsOperation) {
    m := &CommsOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCommsOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommsOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.addLargeGalleryViewOperation":
                        return NewAddLargeGalleryViewOperation(), nil
                    case "#microsoft.graph.cancelMediaProcessingOperation":
                        return NewCancelMediaProcessingOperation(), nil
                    case "#microsoft.graph.inviteParticipantsOperation":
                        return NewInviteParticipantsOperation(), nil
                    case "#microsoft.graph.muteParticipantOperation":
                        return NewMuteParticipantOperation(), nil
                    case "#microsoft.graph.playPromptOperation":
                        return NewPlayPromptOperation(), nil
                    case "#microsoft.graph.recordOperation":
                        return NewRecordOperation(), nil
                    case "#microsoft.graph.startHoldMusicOperation":
                        return NewStartHoldMusicOperation(), nil
                    case "#microsoft.graph.stopHoldMusicOperation":
                        return NewStopHoldMusicOperation(), nil
                    case "#microsoft.graph.subscribeToToneOperation":
                        return NewSubscribeToToneOperation(), nil
                    case "#microsoft.graph.unmuteParticipantOperation":
                        return NewUnmuteParticipantOperation(), nil
                    case "#microsoft.graph.updateRecordingStatusOperation":
                        return NewUpdateRecordingStatusOperation(), nil
                }
            }
        }
    }
    return NewCommsOperation(), nil
}
// GetClientContext gets the clientContext property value. Unique Client Context string. Max limit is 256 chars.
func (m *CommsOperation) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommsOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["resultInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResultInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultInfo(val.(ResultInfoable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*OperationStatus))
        }
        return nil
    }
    return res
}
// GetResultInfo gets the resultInfo property value. The result information. Read-only.
func (m *CommsOperation) GetResultInfo()(ResultInfoable) {
    return m.resultInfo
}
// GetStatus gets the status property value. The status property
func (m *CommsOperation) GetStatus()(*OperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CommsOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resultInfo", m.GetResultInfo())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientContext sets the clientContext property value. Unique Client Context string. Max limit is 256 chars.
func (m *CommsOperation) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetResultInfo sets the resultInfo property value. The result information. Read-only.
func (m *CommsOperation) SetResultInfo(value ResultInfoable)() {
    m.resultInfo = value
}
// SetStatus sets the status property value. The status property
func (m *CommsOperation) SetStatus(value *OperationStatus)() {
    m.status = value
}
