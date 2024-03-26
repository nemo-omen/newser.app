-- +goose Up
CREATE TABLE IF NOT EXISTS
    users (
        id       TEXT PRIMARY KEY,
        email    TEXT UNIQUE NOT NULL,
        name     TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );

CREATE TABLE IF NOT EXISTS
    people (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL DEFAULT ''
    );

CREATE TABLE IF NOT EXISTS
    collections (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        slug TEXT NOT NULL,
        user_id TEXT NOT NULL,
        CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    images (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL DEFAULT '',
        url TEXT UNIQUE NOT NULL DEFAULT ''
    );

CREATE TABLE IF NOT EXISTS
    categories (
        id   TEXT PRIMARY KEY,
        term TEXT UNIQUE NOT NULL
    );

CREATE TABLE IF NOT EXISTS
    newsfeeds (
        id          TEXT PRIMARY KEY,
        title       TEXT NOT NULL,
        site_url    TEXT NOT NULL,
        feed_url    TEXT UNIQUE NOT NULL,
        description TEXT,
        copyright   TEXT,
        language    TEXT,
        feed_type   TEXT,
        slug        TEXT NOT NULL
    );

CREATE TABLE IF NOT EXISTS
    articles (
        id               TEXT PRIMARY KEY,
        title            TEXT NOT NULL,
        description      TEXT,
        content          TEXT,
        article_link     TEXT UNIQUE NOT NULL,
        published        TEXT NOT NULL,
        published_parsed TIMESTAMP NOT NULL,
        updated          TEXT NOT NULL,
        updated_parsed   TIMESTAMP NOT NULL,
        guid             TEXT UNIQUE NOT NULL,
        slug             TEXT NOT NULL,
        newsfeed_id      TEXT NOT NULL,
        CONSTRAINT       fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    subscriptions (
        user_id     TEXT NOT NULL,
        newsfeed_id TEXT NOT NULL,
        PRIMARY KEY (user_id, newsfeed_id),
        CONSTRAINT  fk_users FOREIGN KEY (user_id) REFERENCES users (id),
        CONSTRAINT  fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id)
    );

CREATE TABLE IF NOT EXISTS
    annotations (
        id         TEXT PRIMARY KEY,
        title      TEXT NOT NULL DEFAULT '',
        content    TEXT NOT NULL,
        user_id    TEXT NOT NULL,
        article_id TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    highlights (
        id TEXT PRIMARY KEY,
        start_offset INT NOT NULL,
        end_offset INT NOT NULL,
        user_id TEXT NOT NULL,
        annotation_id TEXT NOT NULL,
        article_id TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_annotations FOREIGN KEY (annotation_id) REFERENCES annotations (id) ON DELETE CASCADE
    );

-- JOIN TABLES
CREATE TABLE IF NOT EXISTS
    article_categories (
        article_id TEXT NOT NULL,
        category_id TEXT NOT NULL,
        PRIMARY KEY (article_id, category_id),
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_categories FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    article_images (
        article_id TEXT NOT NULL,
        image_id TEXT NOT NULL,
        PRIMARY KEY (article_id, image_id),
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_images FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    article_people (
        article_id TEXT NOT NULL,
        person_id TEXT NOT NULL,
        PRIMARY KEY (article_id, person_id),
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_people FOREIGN KEY (person_id) REFERENCES people (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    collection_articles (
        article_id TEXT NOT NULL,
        collection_id TEXT NOT NULL,
        PRIMARY KEY (article_id, collection_id),
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_collections FOREIGN KEY (collection_id) REFERENCES collections (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    newsfeed_people (
        newsfeed_id TEXT NOT NULL,
        person_id TEXT NOT NULL,
        PRIMARY KEY (newsfeed_id, person_id),
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE,
        CONSTRAINT fk_people FOREIGN KEY (person_id) REFERENCES people (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    newsfeed_categories (
        newsfeed_id TEXT NOT NULL,
        category_id TEXT NOT NULL,
        PRIMARY KEY (newsfeed_id, category_id),
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE,
        CONSTRAINT fk_categories FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    newsfeed_images (
        newsfeed_id TEXT NOT NULL,
        image_id TEXT NOT NULL,
        PRIMARY KEY (newsfeed_id, image_id),
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE,
        CONSTRAINT fk_images FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    collection_newsfeeds (
        collection_id TEXT NOT NULL,
        newsfeed_id TEXT NOT NULL,
        PRIMARY KEY (collection_id, newsfeed_id),
        CONSTRAINT fk_collections FOREIGN KEY (collection_id) REFERENCES collections (id) ON DELETE CASCADE,
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE
    );


-- +goose Down
DROP TABLE IF EXISTS newsfeed_images;
DROP TABLE IF EXISTS newsfeed_categories;
DROP TABLE IF EXISTS newsfeed_people;
DROP TABLE IF EXISTS newsfeed_subscriptions;
DROP TABLE IF EXISTS collection_newsfeeds;
DROP TABLE IF EXISTS collection_articles;
DROP TABLE IF EXISTS article_categories;
DROP TABLE IF EXISTS article_images;
DROP TABLE IF EXISTS article_people;

DROP TABLE IF EXISTS newsfeeds;
DROP TABLE IF EXISTS highlights;
DROP TABLE IF EXISTS annotations;
DROP TABLE IF EXISTS articles;

DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS images;
DROP TABLE IF EXISTS collections;
DROP TABLE IF EXISTS people;
DROP TABLE IF EXISTS users;