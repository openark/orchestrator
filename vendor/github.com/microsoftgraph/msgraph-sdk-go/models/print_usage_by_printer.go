package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintUsageByPrinter 
type PrintUsageByPrinter struct {
    PrintUsage
    // The printerId property
    printerId *string
}
// NewPrintUsageByPrinter instantiates a new PrintUsageByPrinter and sets the default values.
func NewPrintUsageByPrinter()(*PrintUsageByPrinter) {
    m := &PrintUsageByPrinter{
        PrintUsage: *NewPrintUsage(),
    }
    odataTypeValue := "#microsoft.graph.printUsageByPrinter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePrintUsageByPrinterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintUsageByPrinterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintUsageByPrinter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintUsageByPrinter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrintUsage.GetFieldDeserializers()
    res["printerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterId(val)
        }
        return nil
    }
    return res
}
// GetPrinterId gets the printerId property value. The printerId property
func (m *PrintUsageByPrinter) GetPrinterId()(*string) {
    return m.printerId
}
// Serialize serializes information the current object
func (m *PrintUsageByPrinter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrintUsage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("printerId", m.GetPrinterId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPrinterId sets the printerId property value. The printerId property
func (m *PrintUsageByPrinter) SetPrinterId(value *string)() {
    m.printerId = value
}
