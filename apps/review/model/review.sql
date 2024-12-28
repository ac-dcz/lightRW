drop table if exists review;

create table if not exists review(
    id bigint unsigned not null auto_increment,
    review_id bigint unsigned not null comment  "评论id",
    uid bigint unsigned not null comment "用户id",
    order_id bigint unsigned not null comment "订单id",
    store_id bigint unsigned not null comment "店铺id",
    sku varchar(40) not null comment "sku",

    score tinyint unsigned not null comment "0差评/1中评/2好评",
    goods_desc varchar(256) not null default "" comment "商品描述",
    has_image tinyint unsigned not null default 0 comment "1有/0无",
    image_json varchar(256) not null default "" comment "image json",
    store_score tinyint unsigned not null comment "1-5星",
    is_reply tinyint unsigned not null default 0 comment "0否/1是",
    status tinyint unsigned not null default 10 comment '状态:10待审核；20审核通过；30审核
不通过；40隐藏',
    `op_reason` varchar(512) NOT NULL DEFAULT '' COMMENT '运营审核拒绝原因',
    `goods_snapshot` varchar(2048) NOT NULL DEFAULT '' COMMENT '商品快照信息',

    is_del tinyint unsigned not null default 0 comment "0否/1是",
    create_at datetime not null default now() comment "创建时间",
    update_at datetime not null default now() on update now() comment "更新时间",

    primary key (id),
    unique key uk_review_id (review_id),
    key ind_sid_sku(store_id,sku) comment "某一个商品有哪些评论",
    key ind_uid(uid) comment "用户进行了那些评论",
    check ( score in (0,1,2) ),
    check ( store_score>0 and store_score<=5 )
)Engine = InnoDB default charset utf8mb4 collate utf8mb4_general_ci;

create index ind_sid_sku on review(store_id,sku) comment "某一个商品有哪些评论";
create index ind_uid on review(uid) comment "用户进行了那些评论";

create unique index uk_review_id on review(review_id);