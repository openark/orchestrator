package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaContentRatingNewZealand 
type MediaContentRatingNewZealand struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Movies rating labels in New Zealand
    movieRating *RatingNewZealandMoviesType
    // The OdataType property
    odataType *string
    // TV content rating labels in New Zealand
    tvRating *RatingNewZealandTelevisionType
}
// NewMediaContentRatingNewZealand instantiates a new mediaContentRatingNewZealand and sets the default values.
func NewMediaContentRatingNewZealand()(*MediaContentRatingNewZealand) {
    m := &MediaContentRatingNewZealand{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMediaContentRatingNewZealandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaContentRatingNewZealandFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaContentRatingNewZealand(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaContentRatingNewZealand) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaContentRatingNewZealand) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["movieRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingNewZealandMoviesType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMovieRating(val.(*RatingNewZealandMoviesType))
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
    res["tvRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingNewZealandTelevisionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTvRating(val.(*RatingNewZealandTelevisionType))
        }
        return nil
    }
    return res
}
// GetMovieRating gets the movieRating property value. Movies rating labels in New Zealand
func (m *MediaContentRatingNewZealand) GetMovieRating()(*RatingNewZealandMoviesType) {
    return m.movieRating
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MediaContentRatingNewZealand) GetOdataType()(*string) {
    return m.odataType
}
// GetTvRating gets the tvRating property value. TV content rating labels in New Zealand
func (m *MediaContentRatingNewZealand) GetTvRating()(*RatingNewZealandTelevisionType) {
    return m.tvRating
}
// Serialize serializes information the current object
func (m *MediaContentRatingNewZealand) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMovieRating() != nil {
        cast := (*m.GetMovieRating()).String()
        err := writer.WriteStringValue("movieRating", &cast)
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
    if m.GetTvRating() != nil {
        cast := (*m.GetTvRating()).String()
        err := writer.WriteStringValue("tvRating", &cast)
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
func (m *MediaContentRatingNewZealand) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMovieRating sets the movieRating property value. Movies rating labels in New Zealand
func (m *MediaContentRatingNewZealand) SetMovieRating(value *RatingNewZealandMoviesType)() {
    m.movieRating = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MediaContentRatingNewZealand) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTvRating sets the tvRating property value. TV content rating labels in New Zealand
func (m *MediaContentRatingNewZealand) SetTvRating(value *RatingNewZealandTelevisionType)() {
    m.tvRating = value
}
