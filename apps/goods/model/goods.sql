drop table if exists goods;

create table if not exists goods(
    id bigint unsigned not null auto_increment comment "id",
    sku varchar(40) not null comment "sku",
    `name` varchar(40) not null default "" comment "名称",
    create_at datetime not null default now() comment "创建时间",
    update_at datetime not null default now() on update now() comment "更新时间",
    primary key (id),
    unique key uk_sku(sku),
    key ind_name(`name`(10))
)Engine = InnoDB default charset utf8mb4 collate utf8mb4_general_ci;
