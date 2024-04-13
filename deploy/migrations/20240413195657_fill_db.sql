-- +goose Up
-- +goose StatementBegin
insert into features (name) values ('Список товаров'),
('Отзывы на товары'),
('Страница сведений о товаре'),
('Оформление заказа'),
('Интеграция с платежной ситсемой НСПК'),
('Варианты доставки'),
('Отслеживание заказа');

insert into tags (name) values ('Одежда'), 
('Обувь'), 
('Электроника'), 
('Мебель'), 
('Техника'), 
('СпортТовары'), 
('ДетскиеИгрушки'), 
('Книги'), 
('Искусство'), 
('Коллекционные'), 
('Вещи'), 
('БытоваяТехника'), 
('СадовыйИнвентарь'), 
('Инструменты');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM features WHERE TRUE;
DELETE FROM tags WHERE TRUE;
-- +goose StatementEnd
