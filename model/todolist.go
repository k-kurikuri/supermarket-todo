package model

// TodoList TodoList struct
type TodoList struct {
	UID  string
	Todo *[]Todo
}

// Todo Todo struct
type Todo struct {
	Title string
	Done  bool
}

// Indexes DELETE API Request struct
type Indexes struct {
	Indexes []string
}

// Result Response
type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
