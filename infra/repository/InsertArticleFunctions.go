package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

func InsertArticle(db *sqlx.DB, a *model.Article) (*model.Article, error) {
	q := `
	INSERT INTO articles(
		title,
		description,
		content,
		article_link,
		published,
		published_parsed,
		updated,
		updated_parsed,
		guid,
		slug,
		feed_id,
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`
	res, err := db.Exec(
		q,
		a.Title,
		a.Description,
		a.Content,
		a.ArticleLink,
		a.Published,
		a.PublishedParsed,
		a.Updated,
		a.UpdatedParsed,
		a.GUID,
		a.Slug,
		a.FeedId,
	)
	if err != nil {
		fmt.Println("article insert err: ", err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	a.ID = id

	if a.Image != nil {
		_ = InsertArticleImageTx(db, a.Image, a.ID)
	}

	if a.Author != nil {
		_ = InsertArticlePersonTx(db, a.Author, a.ID)
	}
	return a, nil
}

func InsertArticleImageTx(db *sqlx.DB, i *model.Image, aId int64) error {
	tx := db.MustBegin()
	defer tx.Rollback()
	qi := `
	INSERT INTO images(title, url)
		VALUES(?,?);
	`
	res, err := tx.Exec(qi, i.Title, i.URL)
	if err != nil {
		return err
	}

	imgId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	qai := `
	INSERT INTO article_images(article_id, image_id)
		VALUES(?,?);
	`
	_, err = tx.Exec(qai, aId, imgId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func InsertArticlePersonTx(db *sqlx.DB, p *model.Person, aId int64) error {
	tx := db.MustBegin()
	defer tx.Rollback()
	qp := `
		INSERT INTO people(name, email)
		VALUES(?,?);
		`
	res, err := tx.Exec(qp, p.Name, p.Email)

	if err != nil {
		return err
	}

	pId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	qap := `
	INSERT INTO article_people(article_id, person_id)
		VALUES(?,?);
	`
	_, err = tx.Exec(qap, aId, pId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
