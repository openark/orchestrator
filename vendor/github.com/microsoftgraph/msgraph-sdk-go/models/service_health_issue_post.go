package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceHealthIssuePost 
type ServiceHealthIssuePost struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The published time of the post.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The content of the service issue post. The supported value for the contentType property is html.
    description ItemBodyable
    // The OdataType property
    odataType *string
    // The post type of the service issue historical post. Possible values are: regular, quick, strategic, unknownFutureValue.
    postType *PostType
}
// NewServiceHealthIssuePost instantiates a new serviceHealthIssuePost and sets the default values.
func NewServiceHealthIssuePost()(*ServiceHealthIssuePost) {
    m := &ServiceHealthIssuePost{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceHealthIssuePostFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceHealthIssuePostFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceHealthIssuePost(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceHealthIssuePost) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedDateTime gets the createdDateTime property value. The published time of the post.
func (m *ServiceHealthIssuePost) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The content of the service issue post. The supported value for the contentType property is html.
func (m *ServiceHealthIssuePost) GetDescription()(ItemBodyable) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceHealthIssuePost) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(ItemBodyable))
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
    res["postType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePostType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostType(val.(*PostType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ServiceHealthIssuePost) GetOdataType()(*string) {
    return m.odataType
}
// GetPostType gets the postType property value. The post type of the service issue historical post. Possible values are: regular, quick, strategic, unknownFutureValue.
func (m *ServiceHealthIssuePost) GetPostType()(*PostType) {
    return m.postType
}
// Serialize serializes information the current object
func (m *ServiceHealthIssuePost) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("description", m.GetDescription())
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
    if m.GetPostType() != nil {
        cast := (*m.GetPostType()).String()
        err := writer.WriteStringValue("postType", &cast)
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
func (m *ServiceHealthIssuePost) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedDateTime sets the createdDateTime property value. The published time of the post.
func (m *ServiceHealthIssuePost) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The content of the service issue post. The supported value for the contentType property is html.
func (m *ServiceHealthIssuePost) SetDescription(value ItemBodyable)() {
    m.description = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ServiceHealthIssuePost) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPostType sets the postType property value. The post type of the service issue historical post. Possible values are: regular, quick, strategic, unknownFutureValue.
func (m *ServiceHealthIssuePost) SetPostType(value *PostType)() {
    m.postType = value
}
