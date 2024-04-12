create table features (
    id bigserial primary key,
    name varchar(100) not null
);

create table tags (
    id bigserial primary key,
    name varchar(100) not null
);



create table banners (
    id bigserial primary key,
    content jsonb not null,
    is_active boolean,
    created_at timestamp,
    updated_at timestamp
);

create table banners_tags (
    banner_id bigint 
        references banners(id)
            on delete cascade
            on update cascade,
    feature_id bigint
                references features(id)
                on delete cascade
                on update cascade,
    tag_id bigint 
        references tags(id)
            on delete cascade
            on update cascade,
    unique(feature_id, tag_id),
    primary key (banner_id, tag_id)
);

create index ix_feature_id ON banners_tags using btree (feature_id);
create index ix_tag_id ON banners_tags using btree (tag_id);