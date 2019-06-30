#metadata store table
create table news_headlines_web_country_metadata
(
    country   text primary key,
    shortname text
)
    with comment = 'This table stores the metadata for google web site'
        and caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
        and compaction = {'max_threshold': '32', 'min_threshold': '4', 'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
        and compression = {'class': 'org.apache.cassandra.io.compress.LZ4Compressor', 'chunk_length_in_kb': '64'}
        and dclocal_read_repair_chance = 0.1;