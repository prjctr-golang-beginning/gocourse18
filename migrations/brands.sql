create table brands
(
    id         varchar(36)  not null,
    name       varchar(255) null,
    code       varchar(255) null,
    alias      varchar(255) null,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    constraint brands_pk
        primary key (id)
);

