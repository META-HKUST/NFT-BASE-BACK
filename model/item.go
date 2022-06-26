package model

const (
	insertItem      = string("insert into items(item_id, item_name, collection_id, item_data, description, owner_id, creater_id, category) values(?,?,?,?,?,?,?,?);")
	queryItem       = string("select item_id, item_name, collection_id, item_data, description, owner_id, creater_id, category, created_at from items where item_id=?;")
	insertItemLable = string("insert into item_label(item_id, label) values(?,?);")
	queryItemLable  = string("select item_id, label from item_label where item_id=?")
	updateItemOwner = string("update items set owner_id=? where item_id=?;")
	searchLabel     = string("select label from item_label where item_id=?;")
)

type Item struct {
	ItemID       string `db:"item_id"`
	ItemName     string `db:"item_name"`
	CollectionID int    `db:"collection_id"`
	ItemData     string `db:"item_data"`
	Description  string `db:"description"`
	OwnerID      string `db:"owner_id"`
	CreaterID    string `db:"creater_id"`
	Category     string `db:"category"`
	CreatedAt    string `db:"created_at"`
}

func CreateItem(item Item) (Item, error) {
	_, err := db.Exec(insertItem,
		item.ItemID,
		item.ItemName,
		item.CollectionID,
		item.ItemData,
		item.Description,
		item.OwnerID,
		item.CreaterID,
		item.Category,
	)
	if err != nil {
		return Item{}, err
	}
	var ret Item
	err = db.Get(&ret, queryItem, item.ItemID)
	if err != nil {
		return Item{}, err
	}
	return ret, nil
}

type ItemLable struct {
	ItemID    string `db:"item_id"`
	ItemLabel string `db:"label"`
}

func CreateItemLabel(label ItemLable) (ItemLable, error) {
	_, err := db.Exec(insertItemLable,
		label.ItemID,
		label.ItemLabel,
	)
	if err != nil {
		return ItemLable{}, err
	}
	var ret ItemLable
	err = db.Get(&ret, queryItemLable, label.ItemID)
	if err != nil {
		return ItemLable{}, err
	}
	return ret, nil
}

func UpdateItemOwner(itemId, toUserId string) (Item, error) {
	_, err := db.Exec(updateItemOwner,
		toUserId,
		itemId,
	)
	if err != nil {
		return Item{}, err
	}
	var ret Item
	err = db.Get(&ret, queryItem, itemId)
	if err != nil {
		return Item{}, err
	}
	return ret, nil
}

type Label struct {
	Label string `db:"label"`
}
type LabelSlice struct {
	Label []Label
}

func SearchLable(itemID string) ([]string, error) {
	var LableSlice []string
	err := db.Select(&LableSlice, searchLabel, itemID)
	if err != nil {
		return []string{}, err
	}
	return LableSlice, nil
}
