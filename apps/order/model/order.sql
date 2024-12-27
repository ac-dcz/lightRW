drop table if exists orders;

create table if not exists orders(
    id bigint unsigned not null auto_increment,
    order_id bigint unsigned not null comment "订单id",
    uid bigint unsigned not null comment "用户id",
    store_id bigint unsigned not null comment "店铺id",
    sku varchar(40) not null comment "sku",
    num int unsigned not null default 1 comment "购买数量",
    price decimal(10,2) not null comment "单价",
    status tinyint unsigned not null default 0 comment "0: 待支付;1: 已支付;2: 取消;3: 过期",
    create_at datetime not null default now(),
    update_at datetime not null default now() on update now(),
    primary key (id),
    unique key uk_oid_uid_sid_sku (order_id,uid,store_id,sku)
)Engine = InnoDB default charset utf8mb4 collate utf8mb4_general_ci;