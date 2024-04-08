package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody 
type ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The keepEnrollmentData property
    keepEnrollmentData *bool
    // The keepUserData property
    keepUserData *bool
    // The macOsUnlockCode property
    macOsUnlockCode *string
    // The persistEsimDataPlan property
    persistEsimDataPlan *bool
}
// NewItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody instantiates a new ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody()(*ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) {
    m := &ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keepEnrollmentData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepEnrollmentData(val)
        }
        return nil
    }
    res["keepUserData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    res["macOsUnlockCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOsUnlockCode(val)
        }
        return nil
    }
    res["persistEsimDataPlan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistEsimDataPlan(val)
        }
        return nil
    }
    return res
}
// GetKeepEnrollmentData gets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) GetKeepEnrollmentData()(*bool) {
    return m.keepEnrollmentData
}
// GetKeepUserData gets the keepUserData property value. The keepUserData property
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) GetKeepUserData()(*bool) {
    return m.keepUserData
}
// GetMacOsUnlockCode gets the macOsUnlockCode property value. The macOsUnlockCode property
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) GetMacOsUnlockCode()(*string) {
    return m.macOsUnlockCode
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) GetPersistEsimDataPlan()(*bool) {
    return m.persistEsimDataPlan
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("keepEnrollmentData", m.GetKeepEnrollmentData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macOsUnlockCode", m.GetMacOsUnlockCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("persistEsimDataPlan", m.GetPersistEsimDataPlan())
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
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. The keepEnrollmentData property
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) SetKeepEnrollmentData(value *bool)() {
    m.keepEnrollmentData = value
}
// SetKeepUserData sets the keepUserData property value. The keepUserData property
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
// SetMacOsUnlockCode sets the macOsUnlockCode property value. The macOsUnlockCode property
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) SetMacOsUnlockCode(value *string)() {
    m.macOsUnlockCode = value
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. The persistEsimDataPlan property
func (m *ItemManagedDevicesItemMicrosoftGraphWipeWipePostRequestBody) SetPersistEsimDataPlan(value *bool)() {
    m.persistEsimDataPlan = value
}
