package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BitLockerRemovableDrivePolicy bitLocker Removable Drive Policies.
type BitLockerRemovableDrivePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This policy setting determines whether BitLocker protection is required for removable data drives to be writable on a computer.
    blockCrossOrganizationWriteAccess *bool
    // Select the encryption method for removable  drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
    encryptionMethod *BitLockerEncryptionMethod
    // The OdataType property
    odataType *string
    // Indicates whether to block write access to devices configured in another organization.  If requireEncryptionForWriteAccess is false, this value does not affect.
    requireEncryptionForWriteAccess *bool
}
// NewBitLockerRemovableDrivePolicy instantiates a new bitLockerRemovableDrivePolicy and sets the default values.
func NewBitLockerRemovableDrivePolicy()(*BitLockerRemovableDrivePolicy) {
    m := &BitLockerRemovableDrivePolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBitLockerRemovableDrivePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBitLockerRemovableDrivePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBitLockerRemovableDrivePolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BitLockerRemovableDrivePolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlockCrossOrganizationWriteAccess gets the blockCrossOrganizationWriteAccess property value. This policy setting determines whether BitLocker protection is required for removable data drives to be writable on a computer.
func (m *BitLockerRemovableDrivePolicy) GetBlockCrossOrganizationWriteAccess()(*bool) {
    return m.blockCrossOrganizationWriteAccess
}
// GetEncryptionMethod gets the encryptionMethod property value. Select the encryption method for removable  drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerRemovableDrivePolicy) GetEncryptionMethod()(*BitLockerEncryptionMethod) {
    return m.encryptionMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BitLockerRemovableDrivePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockCrossOrganizationWriteAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockCrossOrganizationWriteAccess(val)
        }
        return nil
    }
    res["encryptionMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBitLockerEncryptionMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionMethod(val.(*BitLockerEncryptionMethod))
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
    res["requireEncryptionForWriteAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireEncryptionForWriteAccess(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BitLockerRemovableDrivePolicy) GetOdataType()(*string) {
    return m.odataType
}
// GetRequireEncryptionForWriteAccess gets the requireEncryptionForWriteAccess property value. Indicates whether to block write access to devices configured in another organization.  If requireEncryptionForWriteAccess is false, this value does not affect.
func (m *BitLockerRemovableDrivePolicy) GetRequireEncryptionForWriteAccess()(*bool) {
    return m.requireEncryptionForWriteAccess
}
// Serialize serializes information the current object
func (m *BitLockerRemovableDrivePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("blockCrossOrganizationWriteAccess", m.GetBlockCrossOrganizationWriteAccess())
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionMethod() != nil {
        cast := (*m.GetEncryptionMethod()).String()
        err := writer.WriteStringValue("encryptionMethod", &cast)
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
        err := writer.WriteBoolValue("requireEncryptionForWriteAccess", m.GetRequireEncryptionForWriteAccess())
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
func (m *BitLockerRemovableDrivePolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlockCrossOrganizationWriteAccess sets the blockCrossOrganizationWriteAccess property value. This policy setting determines whether BitLocker protection is required for removable data drives to be writable on a computer.
func (m *BitLockerRemovableDrivePolicy) SetBlockCrossOrganizationWriteAccess(value *bool)() {
    m.blockCrossOrganizationWriteAccess = value
}
// SetEncryptionMethod sets the encryptionMethod property value. Select the encryption method for removable  drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerRemovableDrivePolicy) SetEncryptionMethod(value *BitLockerEncryptionMethod)() {
    m.encryptionMethod = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BitLockerRemovableDrivePolicy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRequireEncryptionForWriteAccess sets the requireEncryptionForWriteAccess property value. Indicates whether to block write access to devices configured in another organization.  If requireEncryptionForWriteAccess is false, this value does not affect.
func (m *BitLockerRemovableDrivePolicy) SetRequireEncryptionForWriteAccess(value *bool)() {
    m.requireEncryptionForWriteAccess = value
}
