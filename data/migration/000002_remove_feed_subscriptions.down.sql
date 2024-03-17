CREATE TABLE IF NOT EXISTS
    feed_subscriptions (
        feed_id TEXT NOT NULL,
        subscription_id TEXT NOT NULL,
        CONSTRAINT fk_newsfeeds FOREIGN KEY (feed_id) REFERENCES newsfeeds (id) ON DELETE CASCADE,
        CONSTRAINT fk_subscriptions FOREIGN KEY (subscription_id) REFERENCES subscriptions (id) ON DELETE CASCADE
    );