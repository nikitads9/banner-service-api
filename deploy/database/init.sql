create table features (
    id bigserial primary key,
    name varchar(100) not null
);

create table tags (
    id bigbigserial primary key,
    name varchar(100) not null
);

create table banners (
    id bigserial primary key,
    data text not null,
    feature_id bigint,
    is_active boolean,
    unique(id, feature_id),
    constraint fk_feature 
        foreign key (feature_id) 
            references features(id)
                on delete cascade
                on update cascade
);

create table banners_tags (
    banner_id bigint 
        references banners(id)
            on delete cascade
            on update cascade,
    tag_id bigint 
        references tags(id)
            on delete cascade
            on update cascade,
    primary key (banner_id, tag_id)
);