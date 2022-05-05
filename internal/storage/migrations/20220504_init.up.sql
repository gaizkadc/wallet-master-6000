create table if not exists customers
(
    id       uuid           primary key,
    password varchar        not null,
    balance  numeric(20,10)
);

insert into customers(id, password, balance) values ('273e4de9-a5ff-42c2-bdf2-54884c1d19cc', 'abc', 100);
insert into customers(id, password, balance) values ('12e1c07f-5305-476e-b0bb-079dd60bd1dc', 'def', 1000);
insert into customers(id, password, balance) values ('450a7d47-874a-4064-8178-10b3bffc6e2e', 'ghi', 10000);
insert into customers(id, password, balance) values ('60306e90-c9d1-4923-9e9a-0db9aade3dbe', 'jkl', 10);

create table if not exists transactions
(
    id             uuid           primary key,
    from_customer  uuid           not null,
    to_customer    uuid           not null,
    amount         numeric(20,10) not null,
    datetime       timestamp      not null
);

insert into transactions(id, from_customer, to_customer, amount, datetime)
    values ('93065ceb-f26c-4c18-b2d5-27ce2f9adb3b', '450a7d47-874a-4064-8178-10b3bffc6e2e', '273e4de9-a5ff-42c2-bdf2-54884c1d19cc', 5, '2016-06-22 19:10:25-07');
insert into transactions(id, from_customer, to_customer, amount, datetime)
    values ('734a4661-84ee-4085-a69c-c2c9276fe645', '273e4de9-a5ff-42c2-bdf2-54884c1d19cc', '450a7d47-874a-4064-8178-10b3bffc6e2e', 10, '2016-06-22 19:10:26-07');
insert into transactions(id, from_customer, to_customer, amount, datetime)
    values ('ff064151-3f1c-4509-8be1-e273b2b0c42b', '60306e90-c9d1-4923-9e9a-0db9aade3dbe', '12e1c07f-5305-476e-b0bb-079dd60bd1dc', 15, '2016-06-22 19:10:27-07');
insert into transactions(id, from_customer, to_customer, amount, datetime)
    values ('227c9e03-9fa5-40b5-9164-b20cd171ef74', '12e1c07f-5305-476e-b0bb-079dd60bd1dc', '60306e90-c9d1-4923-9e9a-0db9aade3dbe', 55, '2016-06-22 19:10:28-07');
insert into transactions(id, from_customer, to_customer, amount, datetime)
    values ('176545af-1ab6-48d2-8035-b93bdaae2ca2', '450a7d47-874a-4064-8178-10b3bffc6e2e', '273e4de9-a5ff-42c2-bdf2-54884c1d19cc', 55, '2016-06-22 19:10:29-07');
