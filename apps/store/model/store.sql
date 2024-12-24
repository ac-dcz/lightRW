drop table if exists store;

create table if not exists store(
    id bigint unsigned not null auto_increment,
    store_id bigint unsigned not null comment "店铺id",
    `name` varchar(40) not null default "" comment "店铺名称",
    uid bigint unsigned not null comment "拥有者",
    creat_at datetime not null default now() comment "创建时间",
    update_at datetime not null default now() on update now() comment "更新时间",
    primary key (id),
    unique key (store_id),
    key ind_name (`name`(20))
)Engine = InnoDB default charset utf8mb4 collate utf8mb4_general_ci;


alter table store add constraint fk_uid foreign key(uid) references user(id);
