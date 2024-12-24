drop table if exists goods_store;

create table if not exists goods_store(
    id bigint unsigned not null auto_increment,
    store_id bigint unsigned not null,
    sku varchar(40) not null comment "sku",
    stock int unsigned not null default 0 comment "库存",
    create_at datetime not null default now() comment "创建时间",
    update_at datetime not null default now() on update now() comment "更新时间",
    primary key (id),
    unique key (store_id,sku)
)Engine = InnoDB default charset utf8mb4 collate utf8mb4_general_ci;

alter table goods_store add constraint fk_sku foreign key (sku) references goods(sku);

alter table goods_store add constraint fk_store_id foreign key (store_id) references store(store_id);