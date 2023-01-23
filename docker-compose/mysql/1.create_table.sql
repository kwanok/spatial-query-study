create table if not exists db_live.locations
(
    id          int auto_increment
        primary key,
    name        varchar(255) not null,
    coordinates point        not null srid 0
);

create spatial index geo_index
    on db_live.locations (coordinates);