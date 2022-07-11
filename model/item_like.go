package model

var (
	updateLikeCount = string("update items set like_count = ? where item_id = ?")
	addItemLike     = string("insert into item_like(item_id,user_id) values(?,?)")
	deleteItemLike  = string("delete from item_like where item_id = ?")
	getLikeCount    = string("select count(*) from item_like where item_id = ?")
	doesLike        = string("SELECT item_id FROM item_like WHERE item_id = ? and user_id = ?")
)

func Like(itemId string, UserId string) error {
	c, err := GetLikeCount(itemId)
	if err != nil {
		return err
	}

	_, e := db.Exec(updateLikeCount, c+1, itemId)
	if e != nil {
		return e
	}

	_, e = db.Exec(addItemLike, itemId, UserId)
	if e != nil {
		return e
	}

	return nil
}

func UnLike(itemId string, UserId string) error {
	c, err := GetLikeCount(itemId)
	if err != nil {
		return err
	}

	_, e := db.Exec(updateLikeCount, c-1, itemId)
	if e != nil {
		return e
	}

	_, e = db.Exec(deleteItemLike, itemId)
	if e != nil {
		return e
	}

	return nil
}

func DoesLike(itemId string, UserId string) (bool, error) {
	var g string
	e := db.Get(&g, doesLike, itemId, UserId)
	if e != nil {
		return false, e
	}
	if g == "" {
		return false, nil
	} else {
		return true, nil
	}
}

func GetLikeCount(itemId string) (int, error) {
	var a int
	err := db.Get(&a, getLikeCount, itemId)
	if err != nil {
		return 0, err
	}
	return a, nil
}
