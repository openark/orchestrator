package identityprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody 
type RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The servicePrincipalIds property
    servicePrincipalIds []string
}
// NewRiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody instantiates a new RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody and sets the default values.
func NewRiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody()(*RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody) {
    m := &RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["servicePrincipalIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetServicePrincipalIds(res)
        }
        return nil
    }
    return res
}
// GetServicePrincipalIds gets the servicePrincipalIds property value. The servicePrincipalIds property
func (m *RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody) GetServicePrincipalIds()([]string) {
    return m.servicePrincipalIds
}
// Serialize serializes information the current object
func (m *RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetServicePrincipalIds() != nil {
        err := writer.WriteCollectionOfStringValues("servicePrincipalIds", m.GetServicePrincipalIds())
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
func (m *RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetServicePrincipalIds sets the servicePrincipalIds property value. The servicePrincipalIds property
func (m *RiskyServicePrincipalsMicrosoftGraphDismissDismissPostRequestBody) SetServicePrincipalIds(value []string)() {
    m.servicePrincipalIds = value
}
