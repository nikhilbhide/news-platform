#create a keyspace article-db
create keyspace article_db with replication = {''class'': ''SimpleStrategy'', ''replication_factor'':1} and durable_writes = true;

#create a table news_headlines_by_country
create table table_name
(
	country text,
	date timestamp,
	author text,
	title text,
	description text,
	url text,
	publishedAt timestamp,
	content text,
	Primary key (country)

)
with comment = 'table to hold the headlines data';