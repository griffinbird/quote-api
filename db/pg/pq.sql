create table quotes(
  id        int primary key not null,
  quote  varchar(200)
);

drop table quotes;

insert into quotes (id, quote) values (1, 'Life is either a daring adventure or nothing at all.');
insert into quotes (id, quote) values (2, 'Push yourself, because no one else is going to do it for you.');
insert into quotes (id, quote) values (3, 'Days of efficient actions builds up to achieving goals. String these days together and success will happen to you');


select quote from quotes order by random() limit 1;

select * from quotes where id = 1;

curl -X POST -H "Content-Type: application/json" -d '{"quote": "Days of efficient actions builds up to achieving goals. String these days together and success will happen to you"}' localhost:8080/quotes
curl -d '{"quote": "Life is either a daring adventure or nothing at all."}' -X POST -H "Content-Type: application/json" localhost:8080/quotes
curl -d '{"key1":"value1"}' -H "Content-Type: application/json" -X POST localhost:8080/quotes

curl -X POST -H "Content-Type: application/json" -d '{"quote": "Life is either a daring adventure or nothing at all."}' localhost:8080/quotes

curl -X POST -H "Content-Type: application/json" -d '{"quote":"asdfasdfsadf"}' localhost:8080/quotes

curl -X POST -H "Content-Type: application/json" \
 -d '{"quote":"asdfasdfsadf"}' \
 http://localhost:8080/quotes

curl -X GET localhost:8080/quotes

fmt.Printf("%+v\n", quote)