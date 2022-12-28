package models

type Node struct {
	NodeId            string  `grom:"primaryKey;size:128" json:"id" `
	NodeName          string  `grom:"size:128" json:"name"`
	NodeSymbolizeSize float32 `grom:"" json:"symbolSize"`
	NodeValue         int     `grom:"" json:"value"`
	NodeCategory      int     `grom:"" json:"category"`
}

// Link n--n的关系
type Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Value  int    `json:"value"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Links []Link `json:"links"`
}

type OneUserPatents struct {
	Patentsid []int
}

type InventorPatent struct {
	UserId   int
	PatentId int
}
