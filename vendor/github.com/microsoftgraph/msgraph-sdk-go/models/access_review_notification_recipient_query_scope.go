package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewNotificationRecipientQueryScope 
type AccessReviewNotificationRecipientQueryScope struct {
    AccessReviewNotificationRecipientScope
    // Represents the query for who the recipients are. For example, /groups/{group id}/members for group members and /users/{user id} for a specific user.
    query *string
    // In the scenario where reviewers need to be specified dynamically, indicates the relative source of the query. This property is only required if a relative query (that is, ./manager) is specified.
    queryRoot *string
    // Indicates the type of query. Allowed value is MicrosoftGraph.
    queryType *string
}
// NewAccessReviewNotificationRecipientQueryScope instantiates a new AccessReviewNotificationRecipientQueryScope and sets the default values.
func NewAccessReviewNotificationRecipientQueryScope()(*AccessReviewNotificationRecipientQueryScope) {
    m := &AccessReviewNotificationRecipientQueryScope{
        AccessReviewNotificationRecipientScope: *NewAccessReviewNotificationRecipientScope(),
    }
    odataTypeValue := "#microsoft.graph.accessReviewNotificationRecipientQueryScope"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessReviewNotificationRecipientQueryScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewNotificationRecipientQueryScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewNotificationRecipientQueryScope(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewNotificationRecipientQueryScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessReviewNotificationRecipientScope.GetFieldDeserializers()
    res["query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    res["queryRoot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryRoot(val)
        }
        return nil
    }
    res["queryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryType(val)
        }
        return nil
    }
    return res
}
// GetQuery gets the query property value. Represents the query for who the recipients are. For example, /groups/{group id}/members for group members and /users/{user id} for a specific user.
func (m *AccessReviewNotificationRecipientQueryScope) GetQuery()(*string) {
    return m.query
}
// GetQueryRoot gets the queryRoot property value. In the scenario where reviewers need to be specified dynamically, indicates the relative source of the query. This property is only required if a relative query (that is, ./manager) is specified.
func (m *AccessReviewNotificationRecipientQueryScope) GetQueryRoot()(*string) {
    return m.queryRoot
}
// GetQueryType gets the queryType property value. Indicates the type of query. Allowed value is MicrosoftGraph.
func (m *AccessReviewNotificationRecipientQueryScope) GetQueryType()(*string) {
    return m.queryType
}
// Serialize serializes information the current object
func (m *AccessReviewNotificationRecipientQueryScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessReviewNotificationRecipientScope.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("queryRoot", m.GetQueryRoot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("queryType", m.GetQueryType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetQuery sets the query property value. Represents the query for who the recipients are. For example, /groups/{group id}/members for group members and /users/{user id} for a specific user.
func (m *AccessReviewNotificationRecipientQueryScope) SetQuery(value *string)() {
    m.query = value
}
// SetQueryRoot sets the queryRoot property value. In the scenario where reviewers need to be specified dynamically, indicates the relative source of the query. This property is only required if a relative query (that is, ./manager) is specified.
func (m *AccessReviewNotificationRecipientQueryScope) SetQueryRoot(value *string)() {
    m.queryRoot = value
}
// SetQueryType sets the queryType property value. Indicates the type of query. Allowed value is MicrosoftGraph.
func (m *AccessReviewNotificationRecipientQueryScope) SetQueryType(value *string)() {
    m.queryType = value
}
