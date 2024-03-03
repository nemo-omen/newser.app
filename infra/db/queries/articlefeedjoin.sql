-- database: ../../../data/newser.sqliteSELECT
    newsfeeds.id as feed_id,
    newsfeeds.title as feed_title,
    newsfeeds.feed_url as feed_url,
    newsfeeds.site_url as feed_site_url,
    newsfeeds.slug as feed_slug,
    articles.*,
		images.title as feed_image_title,
    images.url as feed_image_url
FROM
    newsfeeds
    LEFT JOIN articles ON newsfeeds.id = articles.feed_id
    LEFT JOIN newsfeed_images ON newsfeeds.id = newsfeed_images.newsfeed_id
    LEFT JOIN images ON newsfeed_images.image_id = images.id
WHERE
    newsfeeds.id = 1;
