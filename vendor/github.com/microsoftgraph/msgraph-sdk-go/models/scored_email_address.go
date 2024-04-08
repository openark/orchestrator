package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScoredEmailAddress 
type ScoredEmailAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email address.
    address *string
    // The itemId property
    itemId *string
    // The OdataType property
    odataType *string
    // The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
    relevanceScore *float64
    // The selectionLikelihood property
    selectionLikelihood *SelectionLikelihoodInfo
}
// NewScoredEmailAddress instantiates a new scoredEmailAddress and sets the default values.
func NewScoredEmailAddress()(*ScoredEmailAddress) {
    m := &ScoredEmailAddress{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScoredEmailAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScoredEmailAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScoredEmailAddress(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScoredEmailAddress) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The email address.
func (m *ScoredEmailAddress) GetAddress()(*string) {
    return m.address
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScoredEmailAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["itemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
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
    res["relevanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelevanceScore(val)
        }
        return nil
    }
    res["selectionLikelihood"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSelectionLikelihoodInfo)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelectionLikelihood(val.(*SelectionLikelihoodInfo))
        }
        return nil
    }
    return res
}
// GetItemId gets the itemId property value. The itemId property
func (m *ScoredEmailAddress) GetItemId()(*string) {
    return m.itemId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ScoredEmailAddress) GetOdataType()(*string) {
    return m.odataType
}
// GetRelevanceScore gets the relevanceScore property value. The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
func (m *ScoredEmailAddress) GetRelevanceScore()(*float64) {
    return m.relevanceScore
}
// GetSelectionLikelihood gets the selectionLikelihood property value. The selectionLikelihood property
func (m *ScoredEmailAddress) GetSelectionLikelihood()(*SelectionLikelihoodInfo) {
    return m.selectionLikelihood
}
// Serialize serializes information the current object
func (m *ScoredEmailAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("itemId", m.GetItemId())
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
        err := writer.WriteFloat64Value("relevanceScore", m.GetRelevanceScore())
        if err != nil {
            return err
        }
    }
    if m.GetSelectionLikelihood() != nil {
        cast := (*m.GetSelectionLikelihood()).String()
        err := writer.WriteStringValue("selectionLikelihood", &cast)
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
func (m *ScoredEmailAddress) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The email address.
func (m *ScoredEmailAddress) SetAddress(value *string)() {
    m.address = value
}
// SetItemId sets the itemId property value. The itemId property
func (m *ScoredEmailAddress) SetItemId(value *string)() {
    m.itemId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ScoredEmailAddress) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRelevanceScore sets the relevanceScore property value. The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
func (m *ScoredEmailAddress) SetRelevanceScore(value *float64)() {
    m.relevanceScore = value
}
// SetSelectionLikelihood sets the selectionLikelihood property value. The selectionLikelihood property
func (m *ScoredEmailAddress) SetSelectionLikelihood(value *SelectionLikelihoodInfo)() {
    m.selectionLikelihood = value
}
