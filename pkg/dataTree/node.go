package datatree

type NodeType int

const (
	RootType = iota
	CountryType
	DeviceType
)

type Node struct {
	Type       NodeType         `json:"type"`
	Name       string           `json:"name"`
	Children   map[string]*Node `json:"children"`
	TimeSpent  int64            `json:"time_spent" default:"0"`
	WebRequest int64            `json:"web_request" default:"0"`
}

func NewNode(nodeType NodeType) *Node {
	return &Node{
		Type:       nodeType,
		Children:   make(map[string]*Node),
		WebRequest: 0,
		TimeSpent:  0,
	}
}
