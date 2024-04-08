package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportRoot 
type ReportRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dailyPrintUsageByPrinter property
    dailyPrintUsageByPrinter []PrintUsageByPrinterable
    // The dailyPrintUsageByUser property
    dailyPrintUsageByUser []PrintUsageByUserable
    // The monthlyPrintUsageByPrinter property
    monthlyPrintUsageByPrinter []PrintUsageByPrinterable
    // The monthlyPrintUsageByUser property
    monthlyPrintUsageByUser []PrintUsageByUserable
    // The OdataType property
    odataType *string
    // The security property
    security SecurityReportsRootable
}
// NewReportRoot instantiates a new ReportRoot and sets the default values.
func NewReportRoot()(*ReportRoot) {
    m := &ReportRoot{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReportRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportRoot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportRoot) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDailyPrintUsageByPrinter gets the dailyPrintUsageByPrinter property value. The dailyPrintUsageByPrinter property
func (m *ReportRoot) GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable) {
    return m.dailyPrintUsageByPrinter
}
// GetDailyPrintUsageByUser gets the dailyPrintUsageByUser property value. The dailyPrintUsageByUser property
func (m *ReportRoot) GetDailyPrintUsageByUser()([]PrintUsageByUserable) {
    return m.dailyPrintUsageByUser
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dailyPrintUsageByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByPrinterable)
            }
            m.SetDailyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["dailyPrintUsageByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByUserable)
            }
            m.SetDailyPrintUsageByUser(res)
        }
        return nil
    }
    res["monthlyPrintUsageByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByPrinterable)
            }
            m.SetMonthlyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["monthlyPrintUsageByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByUserable)
            }
            m.SetMonthlyPrintUsageByUser(res)
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
    res["security"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityReportsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurity(val.(SecurityReportsRootable))
        }
        return nil
    }
    return res
}
// GetMonthlyPrintUsageByPrinter gets the monthlyPrintUsageByPrinter property value. The monthlyPrintUsageByPrinter property
func (m *ReportRoot) GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable) {
    return m.monthlyPrintUsageByPrinter
}
// GetMonthlyPrintUsageByUser gets the monthlyPrintUsageByUser property value. The monthlyPrintUsageByUser property
func (m *ReportRoot) GetMonthlyPrintUsageByUser()([]PrintUsageByUserable) {
    return m.monthlyPrintUsageByUser
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ReportRoot) GetOdataType()(*string) {
    return m.odataType
}
// GetSecurity gets the security property value. The security property
func (m *ReportRoot) GetSecurity()(SecurityReportsRootable) {
    return m.security
}
// Serialize serializes information the current object
func (m *ReportRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDailyPrintUsageByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageByPrinter()))
        for i, v := range m.GetDailyPrintUsageByPrinter() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("dailyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageByUser()))
        for i, v := range m.GetDailyPrintUsageByUser() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("dailyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageByPrinter()))
        for i, v := range m.GetMonthlyPrintUsageByPrinter() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("monthlyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageByUser()))
        for i, v := range m.GetMonthlyPrintUsageByUser() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("monthlyPrintUsageByUser", cast)
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
        err := writer.WriteObjectValue("security", m.GetSecurity())
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
func (m *ReportRoot) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDailyPrintUsageByPrinter sets the dailyPrintUsageByPrinter property value. The dailyPrintUsageByPrinter property
func (m *ReportRoot) SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)() {
    m.dailyPrintUsageByPrinter = value
}
// SetDailyPrintUsageByUser sets the dailyPrintUsageByUser property value. The dailyPrintUsageByUser property
func (m *ReportRoot) SetDailyPrintUsageByUser(value []PrintUsageByUserable)() {
    m.dailyPrintUsageByUser = value
}
// SetMonthlyPrintUsageByPrinter sets the monthlyPrintUsageByPrinter property value. The monthlyPrintUsageByPrinter property
func (m *ReportRoot) SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)() {
    m.monthlyPrintUsageByPrinter = value
}
// SetMonthlyPrintUsageByUser sets the monthlyPrintUsageByUser property value. The monthlyPrintUsageByUser property
func (m *ReportRoot) SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)() {
    m.monthlyPrintUsageByUser = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ReportRoot) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSecurity sets the security property value. The security property
func (m *ReportRoot) SetSecurity(value SecurityReportsRootable)() {
    m.security = value
}
