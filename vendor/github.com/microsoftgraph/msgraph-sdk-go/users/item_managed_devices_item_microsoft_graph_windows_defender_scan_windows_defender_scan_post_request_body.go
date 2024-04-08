package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody 
type ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The quickScan property
    quickScan *bool
}
// NewItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody instantiates a new ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody()(*ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody) {
    m := &ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["quickScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickScan(val)
        }
        return nil
    }
    return res
}
// GetQuickScan gets the quickScan property value. The quickScan property
func (m *ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody) GetQuickScan()(*bool) {
    return m.quickScan
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("quickScan", m.GetQuickScan())
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
func (m *ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetQuickScan sets the quickScan property value. The quickScan property
func (m *ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanPostRequestBody) SetQuickScan(value *bool)() {
    m.quickScan = value
}
