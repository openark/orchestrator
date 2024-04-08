package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StoragePlanInformation 
type StoragePlanInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Indicates whether there are higher storage quota plans available. Read-only.
    upgradeAvailable *bool
}
// NewStoragePlanInformation instantiates a new storagePlanInformation and sets the default values.
func NewStoragePlanInformation()(*StoragePlanInformation) {
    m := &StoragePlanInformation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStoragePlanInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStoragePlanInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStoragePlanInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StoragePlanInformation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StoragePlanInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["upgradeAvailable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAvailable(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *StoragePlanInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetUpgradeAvailable gets the upgradeAvailable property value. Indicates whether there are higher storage quota plans available. Read-only.
func (m *StoragePlanInformation) GetUpgradeAvailable()(*bool) {
    return m.upgradeAvailable
}
// Serialize serializes information the current object
func (m *StoragePlanInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("upgradeAvailable", m.GetUpgradeAvailable())
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
func (m *StoragePlanInformation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *StoragePlanInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUpgradeAvailable sets the upgradeAvailable property value. Indicates whether there are higher storage quota plans available. Read-only.
func (m *StoragePlanInformation) SetUpgradeAvailable(value *bool)() {
    m.upgradeAvailable = value
}
