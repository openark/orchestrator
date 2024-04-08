package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody 
type ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The apps property
    apps []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileAppable
}
// NewManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody instantiates a new ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody and sets the default values.
func NewManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody()(*ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) {
    m := &ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApps gets the apps property value. The apps property
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetApps()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileAppable) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApps sets the apps property value. The apps property
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) SetApps(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedMobileAppable)() {
    m.apps = value
}
