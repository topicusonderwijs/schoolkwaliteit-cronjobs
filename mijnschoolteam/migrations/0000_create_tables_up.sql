-- name: CreateMSTAccountTable :exec
create table accounts
(
    id             int  primary key,
    guid           varchar(36)          not null,
    type           varchar(191)         null,
    contact_id     int                  null,
    code           varchar(191)         null,
    name           varchar(191)         null,
    contact_name   varchar(191)         null,
    subdomain      varchar(191)         null,
    authority_name varchar(191)         null,
    authority_code varchar(191)         null,
    address1       varchar(191)         null,
    address2       varchar(191)         null,
    zipcode        varchar(191)         null,
    city           varchar(191)         null,
    country        varchar(191)         null,
    email          varchar(191)         null,
    phone          varchar(191)         null,
    website        varchar(191)         null,
    logo           varchar(191)         null,
    concept_id     int                  null,
    settings       json                 null,
    active         tinyint(1) default 1 null,
    customer_id    varchar(191)         null,
    expired_on     date                 null,
    expire_reason  varchar(191)         null,
    deleted_at     datetime             null,
    users_count    int        default 0 not null,
    people_count   int        default 0 not null,
    groups_count   int        default 0 not null,
    created_at     datetime             not null,
    updated_at     datetime             not null,
    parnassys      tinyint(1) default 0 not null,
    anonymized     tinyint(1) default 0 not null,
    constraint index_accounts_on_guid
        unique (guid),
    constraint index_accounts_on_subdomain
        unique (subdomain)
);