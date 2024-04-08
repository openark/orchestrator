package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchAlteration 
type SearchAlteration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Defines the altered highlighted query string with spelling correction. The annotation around the corrected segment is: /ue000, /ue001.
    alteredHighlightedQueryString *string
    // Defines the altered query string with spelling correction.
    alteredQueryString *string
    // Represents changed segments related to an original user query.
    alteredQueryTokens []AlteredQueryTokenable
    // The OdataType property
    odataType *string
}
// NewSearchAlteration instantiates a new searchAlteration and sets the default values.
func NewSearchAlteration()(*SearchAlteration) {
    m := &SearchAlteration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchAlterationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchAlterationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchAlteration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAlteration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlteredHighlightedQueryString gets the alteredHighlightedQueryString property value. Defines the altered highlighted query string with spelling correction. The annotation around the corrected segment is: /ue000, /ue001.
func (m *SearchAlteration) GetAlteredHighlightedQueryString()(*string) {
    return m.alteredHighlightedQueryString
}
// GetAlteredQueryString gets the alteredQueryString property value. Defines the altered query string with spelling correction.
func (m *SearchAlteration) GetAlteredQueryString()(*string) {
    return m.alteredQueryString
}
// GetAlteredQueryTokens gets the alteredQueryTokens property value. Represents changed segments related to an original user query.
func (m *SearchAlteration) GetAlteredQueryTokens()([]AlteredQueryTokenable) {
    return m.alteredQueryTokens
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchAlteration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alteredHighlightedQueryString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlteredHighlightedQueryString(val)
        }
        return nil
    }
    res["alteredQueryString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlteredQueryString(val)
        }
        return nil
    }
    res["alteredQueryTokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlteredQueryTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlteredQueryTokenable, len(val))
            for i, v := range val {
                res[i] = v.(AlteredQueryTokenable)
            }
            m.SetAlteredQueryTokens(res)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchAlteration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SearchAlteration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alteredHighlightedQueryString", m.GetAlteredHighlightedQueryString())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alteredQueryString", m.GetAlteredQueryString())
        if err != nil {
            return err
        }
    }
    if m.GetAlteredQueryTokens() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlteredQueryTokens()))
        for i, v := range m.GetAlteredQueryTokens() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("alteredQueryTokens", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAlteration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlteredHighlightedQueryString sets the alteredHighlightedQueryString property value. Defines the altered highlighted query string with spelling correction. The annotation around the corrected segment is: /ue000, /ue001.
func (m *SearchAlteration) SetAlteredHighlightedQueryString(value *string)() {
    m.alteredHighlightedQueryString = value
}
// SetAlteredQueryString sets the alteredQueryString property value. Defines the altered query string with spelling correction.
func (m *SearchAlteration) SetAlteredQueryString(value *string)() {
    m.alteredQueryString = value
}
// SetAlteredQueryTokens sets the alteredQueryTokens property value. Represents changed segments related to an original user query.
func (m *SearchAlteration) SetAlteredQueryTokens(value []AlteredQueryTokenable)() {
    m.alteredQueryTokens = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchAlteration) SetOdataType(value *string)() {
    m.odataType = value
}
