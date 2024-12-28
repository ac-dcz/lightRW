drop table if exists reply;

create table if not exists reply(
    id bigint unsigned not null auto_increment comment "id",
    reply_id bigint unsigned not null comment "回复id",
    mid bigint unsigned not null comment "商家id",
    store_id bigint unsigned not null comment "店铺id",
    sku varchar(40) not null comment "sku",
    review_id bigint unsigned not null comment "评价id",

    reply_content varchar(512) not null comment "回复内容",
    has_image tinyint unsigned not null default 0 comment "0无/1有",
    image_json varchar(256) not null comment "image json",
    status tinyint unsigned not null default 20 comment '状态:10待审核；20审核通过；30审核不通过；40隐藏',
    `op_reason` varchar(512) NOT NULL DEFAULT '' COMMENT '运营审核拒绝原因',

    is_del tinyint unsigned not null default 0 comment "0否/1是",
    create_at datetime not null default now() comment "创建时间",
    update_at datetime not null default now() on update now() comment "更新时间",
    primary key (id),
    unique key uk_reply_id(reply_id),
    key ind_mid_sid_sku (store_id,sku),
    key ind_review_id(review_id)
)Engine = InnoDB default charset utf8mb4 collate utf8mb4_general_ci;

-- alter table reply add constraint fk_review_id foreign key (review_id) references review(id);

-- alter table reply modify column sku varchar(40) not null comment "sku";

-- create index ind_review_id on reply(review_id);