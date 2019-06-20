#create a keyspace article-db
create keyspace article_db with replication = {''class'': ''SimpleStrategy'', ''replication_factor'':1} and durable_writes = true;

#create a table news_headlines_by_country
create table news_headlines
(
	country text,
	date text,
	author text,
	title text,
	description text,
	url text,
	published_at timestamp,
	content text,
	Primary key (country,date,url)
)
with comment = 'table to hold the headlines data';

#select headlines where country = ?
select * from news_headlines where country = ?;