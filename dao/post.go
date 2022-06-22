package dao

import (
	"errors"
	"goblog/models"
	"log"
)

func CountGetAllPost() (pageCount int) {
	row := DB.QueryRow("select  count(1) from blog_post")
	if row.Err() != nil {
		log.Println("获取row出错")
		panic(row.Err())
	}
	row.Scan(&pageCount)
	return
}

func CountGetAllPostBySlug(slug string) (pageCount int) {
	row := DB.QueryRow("select  count(1) from blog_post where slug=?", slug)
	if row.Err() != nil {
		log.Println("获取row出错")
		panic(row.Err())
	}
	row.Scan(&pageCount)
	return
}

func CountGetAllPostByCategoryId(cId int) (pageCount int) {
	row := DB.QueryRow("select  count(1) from blog_post where category_id=?", cId)
	if row.Err() != nil {
		log.Println("获取row出错")
		panic(row.Err())
	}
	row.Scan(&pageCount)
	return
}

func GetPostSearch(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select  * from blog_post where  title like ?", "%"+condition+"%")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	log.Println(len(posts))
	return posts, nil
}

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select  * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select  * from blog_post where slug=? limit ?,?", slug, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageByCategoryId(cId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select  * from blog_post where category_id=? limit ?,?", cId, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostById(pId int) (models.Post, error) {
	row := DB.QueryRow("select  * from blog_post where pid=?", pId)
	var post models.Post
	if row.Err() != nil {
		return post, errors.New("没有此文章")
	}
	var err = row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return post, err
	}
	return post, nil
}

func SavePost(post *models.Post) {
	exec, err := DB.Exec("insert into blog_post (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at)"+
		"values(?,?,?,?,?,?,?,?,?,?)", post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
		return
	}
	pid, err := exec.LastInsertId()
	post.Pid = int(pid)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post title=?,content=?,markdown=?,category_id=?,type =?,slug=?,update_at=? pid=?",
		post.Title, post.Content, post.Markdown, post.CategoryId, post.Type, post.Slug, post.UpdateAt, post.Pid)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetAllPost() ([]models.Post, error) {
	rows, err := DB.Query("select  * from blog_post")
	if err != nil {
		return nil, err
	}
	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil

}
