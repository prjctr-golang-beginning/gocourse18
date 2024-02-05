create table products
(
    id         varchar(36)  not null,
    brand_id   varchar(36)  not null,
    status     varchar(255) null,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    constraint brands_pk
        primary key (id)
);

