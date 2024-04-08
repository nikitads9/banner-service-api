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
    tag_id bigint,
    is_active boolean,
    created_at timestamp,
    updated_at timestamp,
    version integer check (version < 3)
    unique(feature_id, tag_id, version),
    unique(id, feature_id, version),
    constraint fk_feature 
        foreign key (feature_id) 
            references features(id)
                on delete cascade
                on update cascade,
    constraint fk_tag 
        foreign key (tag_id) 
            references tags(id)
                on delete cascade
                on update cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table features;
drop table tags;
drop table banners;
-- +goose StatementEnd
