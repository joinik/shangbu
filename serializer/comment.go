package serializer

import "go_ctry/model"

type Comment struct {
	ID        uint   `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	User      string `json:"user"`
}

func BuildComment(com *model.Comment) Comment {
	return Comment{
		ID:        com.ID,
		CreatedAt: com.CreatedAt.Unix(),
		UpdatedAt: com.UpdatedAt.Unix(),
		Content:   com.Ccoment,
		UserID:    com.UserID,
		User:      com.User.UserName,
	}
}

func BuildListComment(items []*model.Comment) (rest []Comment) {
	for _, item := range items {
		rest = append(rest, Comment{
			ID:        item.ID,
			CreatedAt: item.CreatedAt.Unix(),
			UpdatedAt: item.UpdatedAt.Unix(),
			Content:   item.Ccoment,
			UserID:    item.UserID,
			User:      item.User.UserName,
		})
	}
    return 

    
}


