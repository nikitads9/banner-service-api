-- +goose Up
-- +goose StatementBegin
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
    content jsonb not null,
    feature_id bigint,
    is_active boolean,
    created_at timestamp,
    updated_at timestamp,
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

create index ix_feature_id ON banners using btree (feature_id);
create index ix_tag_id ON banners_tags using btree (tag_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table features;
drop table tags;
drop table banners_tags;
drop table banners;
-- +goose StatementEnd
