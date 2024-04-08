package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamFunSettings 
type TeamFunSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If set to true, enables users to include custom memes.
    allowCustomMemes *bool
    // If set to true, enables Giphy use.
    allowGiphy *bool
    // If set to true, enables users to include stickers and memes.
    allowStickersAndMemes *bool
    // Giphy content rating. Possible values are: moderate, strict.
    giphyContentRating *GiphyRatingType
    // The OdataType property
    odataType *string
}
// NewTeamFunSettings instantiates a new teamFunSettings and sets the default values.
func NewTeamFunSettings()(*TeamFunSettings) {
    m := &TeamFunSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamFunSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamFunSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamFunSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamFunSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowCustomMemes gets the allowCustomMemes property value. If set to true, enables users to include custom memes.
func (m *TeamFunSettings) GetAllowCustomMemes()(*bool) {
    return m.allowCustomMemes
}
// GetAllowGiphy gets the allowGiphy property value. If set to true, enables Giphy use.
func (m *TeamFunSettings) GetAllowGiphy()(*bool) {
    return m.allowGiphy
}
// GetAllowStickersAndMemes gets the allowStickersAndMemes property value. If set to true, enables users to include stickers and memes.
func (m *TeamFunSettings) GetAllowStickersAndMemes()(*bool) {
    return m.allowStickersAndMemes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamFunSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowCustomMemes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCustomMemes(val)
        }
        return nil
    }
    res["allowGiphy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowGiphy(val)
        }
        return nil
    }
    res["allowStickersAndMemes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowStickersAndMemes(val)
        }
        return nil
    }
    res["giphyContentRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGiphyRatingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGiphyContentRating(val.(*GiphyRatingType))
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
// GetGiphyContentRating gets the giphyContentRating property value. Giphy content rating. Possible values are: moderate, strict.
func (m *TeamFunSettings) GetGiphyContentRating()(*GiphyRatingType) {
    return m.giphyContentRating
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamFunSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamFunSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowCustomMemes", m.GetAllowCustomMemes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowGiphy", m.GetAllowGiphy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowStickersAndMemes", m.GetAllowStickersAndMemes())
        if err != nil {
            return err
        }
    }
    if m.GetGiphyContentRating() != nil {
        cast := (*m.GetGiphyContentRating()).String()
        err := writer.WriteStringValue("giphyContentRating", &cast)
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
func (m *TeamFunSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowCustomMemes sets the allowCustomMemes property value. If set to true, enables users to include custom memes.
func (m *TeamFunSettings) SetAllowCustomMemes(value *bool)() {
    m.allowCustomMemes = value
}
// SetAllowGiphy sets the allowGiphy property value. If set to true, enables Giphy use.
func (m *TeamFunSettings) SetAllowGiphy(value *bool)() {
    m.allowGiphy = value
}
// SetAllowStickersAndMemes sets the allowStickersAndMemes property value. If set to true, enables users to include stickers and memes.
func (m *TeamFunSettings) SetAllowStickersAndMemes(value *bool)() {
    m.allowStickersAndMemes = value
}
// SetGiphyContentRating sets the giphyContentRating property value. Giphy content rating. Possible values are: moderate, strict.
func (m *TeamFunSettings) SetGiphyContentRating(value *GiphyRatingType)() {
    m.giphyContentRating = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamFunSettings) SetOdataType(value *string)() {
    m.odataType = value
}
