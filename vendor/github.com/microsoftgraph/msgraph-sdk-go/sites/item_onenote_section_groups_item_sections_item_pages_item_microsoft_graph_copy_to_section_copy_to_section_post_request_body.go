package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody 
type ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The groupId property
    groupId *string
    // The id property
    id *string
    // The siteCollectionId property
    siteCollectionId *string
    // The siteId property
    siteId *string
}
// NewItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody instantiates a new ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) {
    m := &ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["siteCollectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCollectionId(val)
        }
        return nil
    }
    res["siteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The groupId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) GetGroupId()(*string) {
    return m.groupId
}
// GetId gets the id property value. The id property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) GetId()(*string) {
    return m.id
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) GetSiteCollectionId()(*string) {
    return m.siteCollectionId
}
// GetSiteId gets the siteId property value. The siteId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) GetSiteId()(*string) {
    return m.siteId
}
// Serialize serializes information the current object
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteCollectionId", m.GetSiteCollectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteId", m.GetSiteId())
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
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) SetGroupId(value *string)() {
    m.groupId = value
}
// SetId sets the id property value. The id property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) SetId(value *string)() {
    m.id = value
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) SetSiteCollectionId(value *string)() {
    m.siteCollectionId = value
}
// SetSiteId sets the siteId property value. The siteId property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemMicrosoftGraphCopyToSectionCopyToSectionPostRequestBody) SetSiteId(value *string)() {
    m.siteId = value
}
