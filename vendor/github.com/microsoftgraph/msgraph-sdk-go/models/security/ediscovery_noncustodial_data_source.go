package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryNoncustodialDataSource 
type EdiscoveryNoncustodialDataSource struct {
    DataSourceContainer
    // User source or SharePoint site data source as non-custodial data source.
    dataSource DataSourceable
    // Operation entity that represents the latest indexing for the non-custodial data source.
    lastIndexOperation EdiscoveryIndexOperationable
}
// NewEdiscoveryNoncustodialDataSource instantiates a new ediscoveryNoncustodialDataSource and sets the default values.
func NewEdiscoveryNoncustodialDataSource()(*EdiscoveryNoncustodialDataSource) {
    m := &EdiscoveryNoncustodialDataSource{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    odataTypeValue := "#microsoft.graph.security.ediscoveryNoncustodialDataSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryNoncustodialDataSource(), nil
}
// GetDataSource gets the dataSource property value. User source or SharePoint site data source as non-custodial data source.
func (m *EdiscoveryNoncustodialDataSource) GetDataSource()(DataSourceable) {
    return m.dataSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryNoncustodialDataSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["dataSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSource(val.(DataSourceable))
        }
        return nil
    }
    res["lastIndexOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryIndexOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastIndexOperation(val.(EdiscoveryIndexOperationable))
        }
        return nil
    }
    return res
}
// GetLastIndexOperation gets the lastIndexOperation property value. Operation entity that represents the latest indexing for the non-custodial data source.
func (m *EdiscoveryNoncustodialDataSource) GetLastIndexOperation()(EdiscoveryIndexOperationable) {
    return m.lastIndexOperation
}
// Serialize serializes information the current object
func (m *EdiscoveryNoncustodialDataSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("dataSource", m.GetDataSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastIndexOperation", m.GetLastIndexOperation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDataSource sets the dataSource property value. User source or SharePoint site data source as non-custodial data source.
func (m *EdiscoveryNoncustodialDataSource) SetDataSource(value DataSourceable)() {
    m.dataSource = value
}
// SetLastIndexOperation sets the lastIndexOperation property value. Operation entity that represents the latest indexing for the non-custodial data source.
func (m *EdiscoveryNoncustodialDataSource) SetLastIndexOperation(value EdiscoveryIndexOperationable)() {
    m.lastIndexOperation = value
}
