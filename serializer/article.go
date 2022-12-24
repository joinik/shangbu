package serializer

import (
	"go_ctry/model"
)

// Article 文章序列化
type Article struct {
	ID        uint                `json:"id"`
	CreatedAt int64               `json:"created_at"`
	UpdatedAt int64               `json:"updated_at"`
	Title     string              `json:"title"`
	Cover     []map[string]string `json:"cover"`
	Comment   int64               `json:"commet"`
	Like      int64               `json:"like"`
	DisLike   int64               `json:"dislike"`
	AuthorID  uint                `json:"author_id"`
	Author    string              `json:"author"`
	CateID    uint                `json:"cate_id"`
	Cate      string              `json:"cate"`
	AreaID    uint                `json:"area_id"`
	Area      string              `json:"area"`
	Content   string              `json:"content"`
}

type ArtConet struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

func BuildArt(art *model.Article) Article {
	return Article{
		ID:        art.ID,
		CreatedAt: art.CreatedAt.Unix(),
		UpdatedAt: art.UpdatedAt.Unix(),
		Title:     art.Title,
		Cover:     art.Cover.V,
		Comment:   art.CommentCount,
		Like:      art.LikeCount,
		DisLike:   art.DisLikeCount,
		AuthorID:  art.AuthorID,
		Author:    art.Author.UserName,
		AreaID:    art.AreaID,
		Area:      art.Area.Name,
		CateID:    art.CateID,
		Cate:      art.Cate.CateName,
		Content:   art.Content.Content,
	}
}

func BuildArtContent(artContent *model.ArtContent) ArtConet {
	return ArtConet{
		ID:      artContent.ArticleID,
		Content: artContent.Content,
	}
}
