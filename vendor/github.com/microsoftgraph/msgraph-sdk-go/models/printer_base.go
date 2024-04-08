package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterBase 
type PrinterBase struct {
    Entity
    // The capabilities of the printer/printerShare.
    capabilities PrinterCapabilitiesable
    // The default print settings of printer/printerShare.
    defaults PrinterDefaultsable
    // The name of the printer/printerShare.
    displayName *string
    // Whether the printer/printerShare is currently accepting new print jobs.
    isAcceptingJobs *bool
    // The list of jobs that are queued for printing by the printer/printerShare.
    jobs []PrintJobable
    // The physical and/or organizational location of the printer/printerShare.
    location PrinterLocationable
    // The manufacturer of the printer/printerShare.
    manufacturer *string
    // The model name of the printer/printerShare.
    model *string
    // The status property
    status PrinterStatusable
}
// NewPrinterBase instantiates a new printerBase and sets the default values.
func NewPrinterBase()(*PrinterBase) {
    m := &PrinterBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrinterBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.printer":
                        return NewPrinter(), nil
                    case "#microsoft.graph.printerShare":
                        return NewPrinterShare(), nil
                }
            }
        }
    }
    return NewPrinterBase(), nil
}
// GetCapabilities gets the capabilities property value. The capabilities of the printer/printerShare.
func (m *PrinterBase) GetCapabilities()(PrinterCapabilitiesable) {
    return m.capabilities
}
// GetDefaults gets the defaults property value. The default print settings of printer/printerShare.
func (m *PrinterBase) GetDefaults()(PrinterDefaultsable) {
    return m.defaults
}
// GetDisplayName gets the displayName property value. The name of the printer/printerShare.
func (m *PrinterBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["capabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterCapabilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilities(val.(PrinterCapabilitiesable))
        }
        return nil
    }
    res["defaults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterDefaultsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaults(val.(PrinterDefaultsable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isAcceptingJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAcceptingJobs(val)
        }
        return nil
    }
    res["jobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintJobable, len(val))
            for i, v := range val {
                res[i] = v.(PrintJobable)
            }
            m.SetJobs(res)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(PrinterLocationable))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(PrinterStatusable))
        }
        return nil
    }
    return res
}
// GetIsAcceptingJobs gets the isAcceptingJobs property value. Whether the printer/printerShare is currently accepting new print jobs.
func (m *PrinterBase) GetIsAcceptingJobs()(*bool) {
    return m.isAcceptingJobs
}
// GetJobs gets the jobs property value. The list of jobs that are queued for printing by the printer/printerShare.
func (m *PrinterBase) GetJobs()([]PrintJobable) {
    return m.jobs
}
// GetLocation gets the location property value. The physical and/or organizational location of the printer/printerShare.
func (m *PrinterBase) GetLocation()(PrinterLocationable) {
    return m.location
}
// GetManufacturer gets the manufacturer property value. The manufacturer of the printer/printerShare.
func (m *PrinterBase) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The model name of the printer/printerShare.
func (m *PrinterBase) GetModel()(*string) {
    return m.model
}
// GetStatus gets the status property value. The status property
func (m *PrinterBase) GetStatus()(PrinterStatusable) {
    return m.status
}
// Serialize serializes information the current object
func (m *PrinterBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("capabilities", m.GetCapabilities())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaults", m.GetDefaults())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAcceptingJobs", m.GetIsAcceptingJobs())
        if err != nil {
            return err
        }
    }
    if m.GetJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCapabilities sets the capabilities property value. The capabilities of the printer/printerShare.
func (m *PrinterBase) SetCapabilities(value PrinterCapabilitiesable)() {
    m.capabilities = value
}
// SetDefaults sets the defaults property value. The default print settings of printer/printerShare.
func (m *PrinterBase) SetDefaults(value PrinterDefaultsable)() {
    m.defaults = value
}
// SetDisplayName sets the displayName property value. The name of the printer/printerShare.
func (m *PrinterBase) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsAcceptingJobs sets the isAcceptingJobs property value. Whether the printer/printerShare is currently accepting new print jobs.
func (m *PrinterBase) SetIsAcceptingJobs(value *bool)() {
    m.isAcceptingJobs = value
}
// SetJobs sets the jobs property value. The list of jobs that are queued for printing by the printer/printerShare.
func (m *PrinterBase) SetJobs(value []PrintJobable)() {
    m.jobs = value
}
// SetLocation sets the location property value. The physical and/or organizational location of the printer/printerShare.
func (m *PrinterBase) SetLocation(value PrinterLocationable)() {
    m.location = value
}
// SetManufacturer sets the manufacturer property value. The manufacturer of the printer/printerShare.
func (m *PrinterBase) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The model name of the printer/printerShare.
func (m *PrinterBase) SetModel(value *string)() {
    m.model = value
}
// SetStatus sets the status property value. The status property
func (m *PrinterBase) SetStatus(value PrinterStatusable)() {
    m.status = value
}
