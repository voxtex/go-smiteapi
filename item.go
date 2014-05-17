package smiteapi

type MenuItem struct {
	Description string
	Value       string
}

type ItemDescription struct {
	Description          string
	MenuItems            []MenuItem `json:"Menuitems"`
	SecondaryDescription string
}

type Item struct {
	ChildItemId      int
	Name             string `json:"DeviceName"`
	IconId           int
	ItemDescription  ItemDescription
	ItemId           int
	ItemTier         int
	Price            int
	RootItemId       int
	ShortDescription string `json:"ShortDesc"`
	StartingItem     bool
	Type             string
}
