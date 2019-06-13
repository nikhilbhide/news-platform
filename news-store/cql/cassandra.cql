#create a keyspace article-db
create keyspace article_db with replication = {''class'': ''SimpleStrategy'', ''replication_factor'':1} and durable_writes = true;

#create a table news_headlines_by_country
create table news_headlines_by_country
(
	country text,
	date timestamp,
	author text,
	title text,
	description text,
	url text,
	published_at timestamp,
	content text,
	Primary key (country)
)
with comment = 'table to hold the headlines data';