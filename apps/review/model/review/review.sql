drop table if exists review;

create table if not exists review(
    id bigint unsigned not null auto_increment,
    uid bigint unsigned not null comment "用户id",
    store_id bigint unsigned not null comment "店铺id",
    sku varchar(40) not null comment "sku",
    order_id bigint unsigned not null comment "订单id",

    score tinyint unsigned not null comment "0差评/1中评/2好评",
    goods_desc varchar(256) not null default "" comment "商品描述",
    has_image bit(1) not null default 0 comment "0有/1无",
    image_json varchar(256) not null default "" comment "image json",
    store_score tinyint unsigned not null comment "1-5星",
    is_reply bit(1) not null default 0 comment "0否/1是",
    status tinyint unsigned not null default 10 comment '状态:10待审核；20审核通过；30审核
不通过；40隐藏',
    `op_reason` varchar(512) NOT NULL DEFAULT '' COMMENT '运营审核拒绝原因',
    `goods_snapshot` varchar(2048) NOT NULL DEFAULT '' COMMENT '商品快照信息',

    is_del bit(1) not null default 0 comment "0否/1是",
    creat_at datetime not null default now() comment "创建时间",
    update_at datetime not null default now() on update now() comment "更新时间",

    primary key (id),
    unique key uk_uid_sid_sku_oid(uid,store_id,sku,order_id),
    check ( score in (0,1,2) ),
    check ( store_score>0 and store_score<=5 )
)Engine = InnoDB default charset utf8mb4 collate utf8mb4_general_ci;