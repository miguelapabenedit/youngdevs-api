package data

type Condition struct {
	ID      int
	Display string
	Type    int
	Cell    int
}

type Action struct {
	ID      int
	Display string
	Type    int
}

type Command struct {
	ID        int
	Display   string
	Type      int
	Condition Condition
	Action    Action
}
