package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EnterpriseModernAppManagementConfigurationCollectionResponseable 
type Windows10EnterpriseModernAppManagementConfigurationCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]Windows10EnterpriseModernAppManagementConfigurationable)
    SetValue(value []Windows10EnterpriseModernAppManagementConfigurationable)()
}
