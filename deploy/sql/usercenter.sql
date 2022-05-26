SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
# '1970-00-00 00:00:00'
DROP TABLE IF EXISTS `user`;
create table `user` (
    `id` bigint not null auto_increment,
    `create_time` datetime not null default current_timestamp,
    `update_time` datetime not null default current_timestamp on update current_timestamp,
    `delete_time` datetime not null default '1970-01-01 08:00:00',
    `del_state` tinyint not null default '0',
    `version` bigint not null default '0' comment '版本号',

    `mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `username` varchar(255) character set utf8mb4 collate utf8mb4_general_ci  not null default '' comment '用户名',
    `password` varchar(255) character set utf8mb4 collate utf8mb4_general_ci  not null default '' comment '密码',

    primary key (`id`),
    unique key `idx_mobile` (`mobile`),
    unique key `idx_username` (`username`)
) engine = InnoDB  default charset = utf8mb4 collate = utf8mb4_general_ci comment = '用户表';

# 密码加密过　明文密码是password
insert into  `user` (mobile, username, password) value ('123456781', 'fzj', '5f4dcc3b5aa765d61d8327deb882cf99');
insert into  `user` (mobile, username, password) value ('123456782','hyb', '5f4dcc3b5aa765d61d8327deb882cf99');
insert into  `user` (mobile, username, password) value ('123456783','wsj', '5f4dcc3b5aa765d61d8327deb882cf99');
insert into  `user` (mobile, username, password) value ('123456785','dh', '5f4dcc3b5aa765d61d8327deb882cf99');
