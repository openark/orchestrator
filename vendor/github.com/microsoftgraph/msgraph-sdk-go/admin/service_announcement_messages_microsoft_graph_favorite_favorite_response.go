package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse 
type ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The value property
    value *bool
}
// NewServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse instantiates a new ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse and sets the default values.
func NewServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse()(*ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse) {
    m := &ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse) GetValue()(*bool) {
    return m.value
}
// Serialize serializes information the current object
func (m *ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("value", m.GetValue())
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
func (m *ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *ServiceAnnouncementMessagesMicrosoftGraphFavoriteFavoriteResponse) SetValue(value *bool)() {
    m.value = value
}
