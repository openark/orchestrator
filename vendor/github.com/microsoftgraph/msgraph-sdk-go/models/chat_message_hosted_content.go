package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageHostedContent 
type ChatMessageHostedContent struct {
    TeamworkHostedContent
}
// NewChatMessageHostedContent instantiates a new ChatMessageHostedContent and sets the default values.
func NewChatMessageHostedContent()(*ChatMessageHostedContent) {
    m := &ChatMessageHostedContent{
        TeamworkHostedContent: *NewTeamworkHostedContent(),
    }
    return m
}
// CreateChatMessageHostedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageHostedContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageHostedContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageHostedContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamworkHostedContent.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ChatMessageHostedContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamworkHostedContent.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
