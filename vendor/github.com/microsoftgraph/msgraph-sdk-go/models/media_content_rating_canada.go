package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaContentRatingCanada 
type MediaContentRatingCanada struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Movies rating labels in Canada
    movieRating *RatingCanadaMoviesType
    // The OdataType property
    odataType *string
    // TV content rating labels in Canada
    tvRating *RatingCanadaTelevisionType
}
// NewMediaContentRatingCanada instantiates a new mediaContentRatingCanada and sets the default values.
func NewMediaContentRatingCanada()(*MediaContentRatingCanada) {
    m := &MediaContentRatingCanada{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMediaContentRatingCanadaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaContentRatingCanadaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaContentRatingCanada(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaContentRatingCanada) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaContentRatingCanada) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["movieRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingCanadaMoviesType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMovieRating(val.(*RatingCanadaMoviesType))
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
        val, err := n.GetEnumValue(ParseRatingCanadaTelevisionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTvRating(val.(*RatingCanadaTelevisionType))
        }
        return nil
    }
    return res
}
// GetMovieRating gets the movieRating property value. Movies rating labels in Canada
func (m *MediaContentRatingCanada) GetMovieRating()(*RatingCanadaMoviesType) {
    return m.movieRating
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MediaContentRatingCanada) GetOdataType()(*string) {
    return m.odataType
}
// GetTvRating gets the tvRating property value. TV content rating labels in Canada
func (m *MediaContentRatingCanada) GetTvRating()(*RatingCanadaTelevisionType) {
    return m.tvRating
}
// Serialize serializes information the current object
func (m *MediaContentRatingCanada) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MediaContentRatingCanada) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMovieRating sets the movieRating property value. Movies rating labels in Canada
func (m *MediaContentRatingCanada) SetMovieRating(value *RatingCanadaMoviesType)() {
    m.movieRating = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MediaContentRatingCanada) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTvRating sets the tvRating property value. TV content rating labels in Canada
func (m *MediaContentRatingCanada) SetTvRating(value *RatingCanadaTelevisionType)() {
    m.tvRating = value
}
