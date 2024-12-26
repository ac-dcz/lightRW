drop table if exists user;

create table if not exists user(
    id bigint unsigned not null auto_increment comment "id",
    nick_name varchar(40) not null default "" comment "昵称",
    tel char(11) not null comment "手机号",
    password varchar(128) not null comment "密码",
    level bit(4) not null default 1 comment "1普通用户/2商家/4管理员",
    status tinyint unsigned not null default 0 comment "0有效/1注销",
    create_at datetime not null default now() comment "创建时间",
    update_at datetime not null default now() on update now() comment "更新时间",
    primary key(id),
    unique key uk_tel(tel)
)Engine = InnoDB default charset utf8mb4 collate utf8mb4_general_ci;