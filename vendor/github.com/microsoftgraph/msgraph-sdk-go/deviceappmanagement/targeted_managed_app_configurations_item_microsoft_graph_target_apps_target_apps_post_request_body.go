package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody 
type TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The appGroupType property
    appGroupType *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppGroupType
    // The apps property
    apps []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileAppable
}
// NewTargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody instantiates a new TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody and sets the default values.
func NewTargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody()(*TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) {
    m := &TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppGroupType gets the appGroupType property value. The appGroupType property
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetAppGroupType()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppGroupType) {
    return m.appGroupType
}
// GetApps gets the apps property value. The apps property
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetApps()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileAppable) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appGroupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseTargetedManagedAppGroupType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppGroupType(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppGroupType))
        }
        return nil
    }
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppGroupType() != nil {
        cast := (*m.GetAppGroupType()).String()
        err := writer.WriteStringValue("appGroupType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("apps", cast)
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
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppGroupType sets the appGroupType property value. The appGroupType property
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) SetAppGroupType(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TargetedManagedAppGroupType)() {
    m.appGroupType = value
}
// SetApps sets the apps property value. The apps property
func (m *TargetedManagedAppConfigurationsItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) SetApps(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileAppable)() {
    m.apps = value
}
