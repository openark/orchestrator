package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBodyable 
type ItemItemsItemWorkbookWorksheetsItemNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetFormula()(*string)
    GetName()(*string)
    SetComment(value *string)()
    SetFormula(value *string)()
    SetName(value *string)()
}
