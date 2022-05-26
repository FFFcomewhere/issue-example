set names utf8mb4;
set FOREIGN_KEY_CHECKS = 0;

# comment
drop table if exists `comment`;
create table `comment` (
    `id` bigint not null auto_increment,
    `create_time` datetime not null default current_timestamp,
    `update_time` datetime not null default current_timestamp on update current_timestamp,
    `delete_time` datetime not null default  '1970-01-01 08:00:00',
    `del_state` tinyint not null default '0',
    `version` bigint not null default '0' comment '版本号',

    `issueid` bigint not null  default 0 comment '提案id',
    `userid` bigint not null  default 0 comment '评论者id',
    `content` varchar(255) not null default '' comment '评论内容',

    primary key (`id`),
    key `idx_issueid` (`issueid`)
) engine = InnoDB  default charset = utf8mb4 collate = utf8mb4_general_ci comment = '评论表';

insert into  `comment` (issueid, userid, content) value (1, 1, 'comment1');
insert into  `comment` (issueid, userid, content) value (1, 2, 'comment2');
insert into  `comment` (issueid, userid, content) value (2, 1, 'comment3');

select *
from comment
where `issueid` > 1;

# issue
drop table if exists `issue`;
create table `issue` (
    `id` bigint not null auto_increment,
    `create_time` datetime not null default current_timestamp,
    `update_time` datetime not null default current_timestamp on update current_timestamp,
    `delete_time` datetime not null default  '1970-01-01 08:00:00',
    `del_state` tinyint not null default '0',
    `version` bigint not null default '0' comment '版本号',

    `name` varchar(255) not null default '' comment '提案名称',
    `userid` bigint not null  default 0 comment '提案发起者id',
    `tagid` bigint not null default 0 comment '标签id',
    `milestoneid` bigint not null default 0 comment '里程碑id',

    primary key (`id`),
    key `idx_tagid` (`tagid`),
    key `idx_milestone` (`milestoneid`)
) engine = InnoDB  default charset = utf8mb4 collate = utf8mb4_general_ci comment = '提案表';

insert into  `issue` (name, userid) value ('issue1', 1);
insert into  `issue` (name, userid) value ('issue2', 1);
insert into  `issue` (name, userid) value ('issue3', 1);
insert into  `issue` (name, userid) value ('issue4', 2);
insert into  `issue` (name, userid) value ('issue5', 2);
insert into  `issue` (name, userid) value ('issue6', 2);
insert into  `issue` (name, userid) value ('issue7', 3);
insert into  `issue` (name, userid) value ('issue8', 3);
insert into  `issue` (name, userid) value ('issue9', 3);
insert into  `issue` (name, userid) value ('issue10', 4);
insert into  `issue` (name, userid) value ('issue11', 4);
insert into  `issue` (name, userid) value ('issue12', 4);
insert into  `issue` (name, userid) value ('issue13', 4);


# milestone
drop table if exists `milestone`;
create table `milestone` (
    `id` bigint not null auto_increment,
    `create_time` datetime not null default current_timestamp,
    `update_time` datetime not null default current_timestamp on update current_timestamp,
    `delete_time` datetime not null default  '1970-01-01 08:00:00',
    `del_state` tinyint not null default '0',
    `version` bigint not null default '0' comment '版本号',

    `name` varchar(255)  not null default '' comment '里程碑名称',

    primary key (`id`),
    unique key `idx_name` (`name`)
) engine = InnoDB  default charset = utf8mb4 collate = utf8mb4_general_ci comment = '里程碑表';

insert into  `milestone` (name) value ('milestone1');

drop table if exists `tag`;
create table `tag` (
    `id` bigint not null auto_increment,
    `create_time` datetime not null default current_timestamp,
    `update_time` datetime not null default current_timestamp on update current_timestamp,
    `delete_time` datetime not null default  '1970-01-01 08:00:00',
    `del_state` tinyint not null default '0',
    `version` bigint not null default '0' comment '版本号',

    `name` varchar(255) not null default '' comment '标签名称',

    primary key (`id`),
    unique key `idx_name` (`name`)
) engine = InnoDB  default charset = utf8mb4 collate = utf8mb4_general_ci comment = '标签表';

insert into  `tag` (name) value ('bug');
insert into `tag` (name) value ('ask');
