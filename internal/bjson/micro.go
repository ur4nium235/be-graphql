package bjson

type UserInfo struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Type        int    `json:"type"`
	Role        string `json:"role"`
}
