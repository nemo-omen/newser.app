-- database: ../newser.sqlite
CREATE TABLE IF NOT EXISTS
    sessions (
        token TEXT PRIMARY KEY,
        data BLOB NOT NULL,
        expiry REAL NOT NULL
    );

CREATE TABLE IF NOT EXISTS
    users (
        id TEXT PRIMARY KEY,
        email TEXT UNIQUE NOT NULL
        ON CONFLICT FAIL,
        name TEXT UNIQUE NOT NULL
        ON CONFLICT FAIL,
        password TEXT NOT NULL
    );

CREATE TABLE IF NOT EXISTS
    people (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL
        ON CONFLICT FAIL DEFAULT ''
    );

CREATE TABLE IF NOT EXISTS
    collections (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        slug TEXT NOT NULL,
        user_id INT NOT NULL,
        CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    images (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL DEFAULT '',
        url TEXT UNIQUE NOT NULL
        ON CONFLICT IGNORE DEFAULT ''
    );

CREATE TABLE IF NOT EXISTS
    categories (
        id TEXT PRIMARY KEY,
        term TEXT UNIQUE NOT NULL
        ON CONFLICT IGNORE
    );

CREATE TABLE IF NOT EXISTS
    subscriptions (
        user_id TEXT NOT NULL,
        newsfeed_id TEXT NOT NULL,
        CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (id),
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) PRIMARY KEY (user_id, newsfeed_id)
    );

CREATE TABLE IF NOT EXISTS
    articles (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        content TEXT,
        article_link TEXT UNIQUE NOT NULL
        ON CONFLICT IGNORE,
        published TEXT NOT NULL,
        published_parsed DATETIME NOT NULL,
        updated TEXT NOT NULL,
        updated_parsed DATETIME NOT NULL,
        guid TEXT UNIQUE NOT NULL
        ON CONFLICT IGNORE,
        slug TEXT NOT NULL,
        newsfeed_id int NOT NULL,
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS
    newsfeeds (
        id TEXT PRIMARY KEY
        ON CONFLICT IGNORE,
        title TEXT NOT NULL,
        site_url NOT NULL,
        feed_url TEXT UNIQUE NOT NULL
        ON CONFLICT IGNORE,
        description TEXT,
        copyright TEXT,
        language TEXT,
        feed_type TEXT,
        slug TEXT NOT NULL
    );

CREATE TABLE IF NOT EXISTS
    annotations (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL DEFAULT '',
        content TEXT NOT NULL,
        user_id TEXT NOT NULL,
        article_id TEXT NOT NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL,
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
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL,
        CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_annotations FOREIGN KEY (annotation_id) REFERENCES annotations (id) ON DELETE CASCADE
    );

-- JOIN TABLES
CREATE TABLE IF NOT EXISTS
    article_categories (
        article_id TEXT NOT NULL,
        category_id TEXT NOT NULL,
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_categories FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE PRIMARY KEY (article_id, category_id)
        ON CONFLICT IGNORE
    );

CREATE TABLE IF NOT EXISTS
    article_images (
        article_id TEXT NOT NULL,
        image_id TEXT NOT NULL,
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_images FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE CASCADE PRIMARY KEY (article_id, image_id)
        ON CONFLICT IGNORE
    );

CREATE TABLE IF NOT EXISTS
    article_people (
        article_id TEXT NOT NULL,
        person_id TEXT NOT NULL,
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_people FOREIGN KEY (person_id) REFERENCES people (id) ON DELETE CASCADE PRIMARY KEY (article_id, person_id)
        ON CONFLICT IGNORE
    );

CREATE TABLE IF NOT EXISTS
    collection_articles (
        article_id TEXT NOT NULL,
        collection_id TEXT NOT NULL,
        CONSTRAINT fk_articles FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
        CONSTRAINT fk_collections FOREIGN KEY (collection_id) REFERENCES collections (id) ON DELETE CASCADE PRIMARY KEY (article_id, collection_id)
        ON CONFLICT IGNORE
    );

CREATE TABLE IF NOT EXISTS
    newsfeed_people (
        newsfeed_id TEXT NOT NULL,
        person_id TEXT NOT NULL,
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE,
        CONSTRAINT fk_people FOREIGN KEY (person_id) REFERENCES people (id) ON DELETE CASCADE PRIMARY KEY (newsfeed_id, person_id)
        ON CONFLICT IGNORE
    );

CREATE TABLE IF NOT EXISTS
    newsfeed_categories (
        newsfeed_id TEXT NOT NULL,
        category_id TEXT NOT NULL,
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE,
        CONSTRAINT fk_categories FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE PRIMARY KEY (newsfeed_id, category_id)
        ON CONFLICT IGNORE
    );

CREATE TABLE IF NOT EXISTS
    newsfeed_images (
        newsfeed_id TEXT NOT NULL,
        image_id TEXT NOT NULL,
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE,
        CONSTRAINT fk_images FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE CASCADE PRIMARY KEY (newsfeed_id, image_id)
        ON CONFLICT IGNORE
    );

CREATE TABLE IF NOT EXISTS
    collection_newsfeeds (
        collection_id TEXT NOT NULL,
        newsfeed_id TEXT NOT NULL,
        CONSTRAINT fk_collections FOREIGN KEY (collection_id) REFERENCES collections (id) ON DELETE CASCADE,
        CONSTRAINT fk_newsfeeds FOREIGN KEY (newsfeed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE PRIMARY KEY (collection_id, newsfeed_id)
        ON CONFLICT IGNORE
    );
