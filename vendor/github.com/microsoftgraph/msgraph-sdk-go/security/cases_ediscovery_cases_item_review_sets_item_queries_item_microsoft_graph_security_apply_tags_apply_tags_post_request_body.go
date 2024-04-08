package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody 
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tagsToAdd property
    tagsToAdd []idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable
    // The tagsToRemove property
    tagsToRemove []idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody instantiates a new CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody()(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tagsToAdd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryReviewTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable, len(val))
            for i, v := range val {
                res[i] = v.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable)
            }
            m.SetTagsToAdd(res)
        }
        return nil
    }
    res["tagsToRemove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryReviewTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable, len(val))
            for i, v := range val {
                res[i] = v.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable)
            }
            m.SetTagsToRemove(res)
        }
        return nil
    }
    return res
}
// GetTagsToAdd gets the tagsToAdd property value. The tagsToAdd property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) GetTagsToAdd()([]idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable) {
    return m.tagsToAdd
}
// GetTagsToRemove gets the tagsToRemove property value. The tagsToRemove property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) GetTagsToRemove()([]idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable) {
    return m.tagsToRemove
}
// Serialize serializes information the current object
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTagsToAdd() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTagsToAdd()))
        for i, v := range m.GetTagsToAdd() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tagsToAdd", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTagsToRemove() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTagsToRemove()))
        for i, v := range m.GetTagsToRemove() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tagsToRemove", cast)
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
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTagsToAdd sets the tagsToAdd property value. The tagsToAdd property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) SetTagsToAdd(value []idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable)() {
    m.tagsToAdd = value
}
// SetTagsToRemove sets the tagsToRemove property value. The tagsToRemove property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityApplyTagsApplyTagsPostRequestBody) SetTagsToRemove(value []idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable)() {
    m.tagsToRemove = value
}
