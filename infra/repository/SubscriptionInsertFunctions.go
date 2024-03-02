package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

func InsertAggregateSubscriptionTx(db *sqlx.DB, n *model.Newsfeed, userId int64) error {
	fmt.Println("STARTING TRANSACTION")
	tx := db.MustBegin()
	defer tx.Rollback()

	persistedNf, err := InsertNewsfeedWithTx(tx, n)
	if err != nil {
		return err
	}

	for _, a := range n.Articles {
		a.FeedId = persistedNf.ID
		persistedArticle, err := InsertArticleWithTx(tx, a)
		if err != nil {
			return err
		}

		err = InsertUnreadWithTx(tx, persistedArticle.ID)
		if err != nil {
			return err
		}
	}

	err = InsertSubscriptionWithTx(tx, userId, persistedNf.ID)
	if err != nil {
		return err
	}

	fmt.Println("COMMITTING TX")
	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error: ", err)
		return ErrTransactionError
	}

	return nil
}

func InsertNewsfeedWithTx(tx *sqlx.Tx, n *model.Newsfeed) (*model.Newsfeed, error) {
	q1 := `
	INSERT INTO newsfeeds(
		title,
		site_url,
		feed_url,
		description,
		updated,
		updated_parsed,
		copyright,
		language,
		feed_type,
		slug
	)
		VALUES(?,?,?,?,?,?,?,?,?,?)
		ON CONFLICT(feed_url) do nothing;
	`
	res, err := tx.Exec(
		q1,
		n.Title,
		n.SiteUrl,
		n.FeedUrl,
		n.Description,
		n.Updated,
		n.UpdatedParsed,
		n.Copyright,
		n.Language,
		n.FeedType,
		n.Slug,
	)
	if err != nil {
		fmt.Println("newsfeed insertion error: ", err)
		return nil, ErrInsertError
	}
	feedId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("lastInserted err: ", err)
		return nil, ErrNotFound
	}
	n.ID = feedId

	if n.Image != nil {
		_ = InsertNewsfeedImageWithTx(tx, n.Image, n.ID)
	}

	if n.Author != nil {
		_ = InsertNewsfeedPersonWithTx(tx, n.Author, n.ID)
	}
	return n, nil
}

func InsertArticleWithTx(tx *sqlx.Tx, a *model.Article) (*model.Article, error) {
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
		feed_id
	) VALUES(?,?,?,?,?,?,?,?,?,?,?);
	`
	res, err := tx.Exec(
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
		return nil, ErrInsertError
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("error getting article id: ", err.Error())
		return nil, ErrNotFound
	}
	a.ID = id

	if a.Image != nil {
		_ = InsertArticleImageWithTx(tx, a.Image, a.ID)
	}

	if a.Author != nil {
		_ = InsertArticlePersonWithTx(tx, a.Author, a.ID)
	}
	return a, nil
}

func InsertArticleImageWithTx(tx *sqlx.Tx, i *model.Image, aId int64) error {
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
		fmt.Println("error getting image id: ", err.Error())
		return ErrNotFound
	}

	qai := `
	INSERT INTO article_images(article_id, image_id)
		VALUES(?,?);
	`
	_, err = tx.Exec(qai, aId, imgId)
	if err != nil {
		fmt.Println("error inserting article_images: ", err.Error())
		return ErrInsertError
	}
	return nil
}

func InsertNewsfeedImageWithTx(tx *sqlx.Tx, i *model.Image, nId int64) error {
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
		fmt.Println("error getting image id: ", err.Error())
		return ErrNotFound
	}

	qai := `
	INSERT INTO newsfeed_images(newsfeed_id, image_id)
		VALUES(?,?);
	`
	_, err = tx.Exec(qai, nId, imgId)
	if err != nil {
		fmt.Println("error inserting newsfeed_images: ", err.Error())
		return ErrInsertError
	}
	return nil
}

func InsertArticlePersonWithTx(tx *sqlx.Tx, p *model.Person, aId int64) error {
	qp := `
		INSERT INTO people(name, email)
		VALUES(?,?);
		`
	res, err := tx.Exec(qp, p.Name, p.Email)

	if err != nil {
		fmt.Println("error inserting into people: ", err.Error())
		return ErrInsertError
	}

	pId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("error getting person_id: ", err.Error())
		return ErrNotFound
	}

	qap := `
	INSERT INTO article_people(article_id, person_id)
		VALUES(?,?);
	`
	_, err = tx.Exec(qap, aId, pId)
	if err != nil {
		fmt.Println("error inserting article_people: ", err.Error())
		return ErrInsertError
	}
	return nil
}

func InsertNewsfeedPersonWithTx(tx *sqlx.Tx, p *model.Person, nId int64) error {
	qp := `
		INSERT INTO people(name, email)
		VALUES(?,?);
		`
	res, err := tx.Exec(qp, p.Name, p.Email)

	if err != nil {
		fmt.Println("error inserting into people: ", err.Error())
		return ErrInsertError
	}

	pId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("error getting person_id: ", err.Error())
		return ErrNotFound
	}

	qap := `
	INSERT INTO newsfeed_people(newsfeed_id, person_id)
		VALUES(?,?);
	`
	_, err = tx.Exec(qap, nId, pId)
	if err != nil {
		fmt.Println("error inserting newsfeed_people: ", err.Error())
		return ErrInsertError
	}
	return nil
}

func InsertUnreadWithTx(tx *sqlx.Tx, aId int64) error {
	var collId int64
	err := tx.Get(&collId, "SELECT id FROM collections WHERE title=?", "unread")

	if err != nil {
		fmt.Println("error finding collection: ", err.Error())
		return ErrNotFound
	}

	q3 := `
	INSERT INTO collection_articles(article_id, collection_id)
		VALUES(?, ?)
	`
	_, err = tx.Exec(q3, aId, collId)
	if err != nil {
		fmt.Println("insertion error: ", err)
		return ErrInsertError
	}
	return nil
}

func InsertSubscriptionWithTx(tx *sqlx.Tx, userId, nId int64) error {
	q4 := `
	INSERT INTO subscriptions(user_id, newsfeed_id)
		VALUES(?, ?);
	`
	_, err := tx.Exec(q4, userId, nId)
	if err != nil {
		fmt.Println("insertion error: ", err)
		return ErrInsertError
	}
	return nil
}