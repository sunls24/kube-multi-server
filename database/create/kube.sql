create table if not exists kube
(
    id        integer not null
        constraint kube_pk
            primary key autoincrement,
    name      text    not null,
    config    text    not null,
    namespace text default 'default' not null
);

create unique index if not exists kube_id_uindex
    on kube (id);

create unique index if not exists kube_name_uindex
    on kube (name);