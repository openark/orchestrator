package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody 
type ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The managedEBookAssignments property
    managedEBookAssignments []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookAssignmentable
}
// NewManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody instantiates a new ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody and sets the default values.
func NewManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody()(*ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody) {
    m := &ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedEBookAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedEBookAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookAssignmentable)
            }
            m.SetManagedEBookAssignments(res)
        }
        return nil
    }
    return res
}
// GetManagedEBookAssignments gets the managedEBookAssignments property value. The managedEBookAssignments property
func (m *ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody) GetManagedEBookAssignments()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookAssignmentable) {
    return m.managedEBookAssignments
}
// Serialize serializes information the current object
func (m *ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedEBookAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedEBookAssignments()))
        for i, v := range m.GetManagedEBookAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("managedEBookAssignments", cast)
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
func (m *ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManagedEBookAssignments sets the managedEBookAssignments property value. The managedEBookAssignments property
func (m *ManagedEBooksItemMicrosoftGraphAssignAssignPostRequestBody) SetManagedEBookAssignments(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedEBookAssignmentable)() {
    m.managedEBookAssignments = value
}
