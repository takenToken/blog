CREATE DATABASE IF NOT EXISTS db_blog;
USE db_blog;

/**
*  tb_album       专辑
*/
-- drop table if exists tb_album;
create table tb_album
(
  id                        bigint(20)      AUTO_INCREMENT  not null        comment '自动增长id',
  name                      varchar(100)    default ''      not null        comment '名称',
  cover                	    varchar(70)     default ''      not null        comment '转换',
  posttime                  datetime                        not null        comment '提交日期',
  ishide                    tinyint(4)      default '0'     not null        comment '是否隐藏',
  rank                      tinyint(4)      default '0'     not null        comment '排名',
  photonum                  varchar(20)     default '0'     not null        comment '照片数量',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='专辑';
create index idx_tb_album_name on tb_album(name);
create index idx_tb_album_posttime on tb_album(posttime);
create index idx_tb_album_ishide on tb_album(ishide);
create index idx_tb_album_rank on tb_album(rank);
create index idx_tb_album_photonum on tb_album(photonum);


/**
*  tb_comments       评论
*/
-- drop table if exists tb_comments;
create table tb_comments
(
  id                        bigint(20)      AUTO_INCREMENT  not null        comment '自动增长id',
  obj_pk_id                 bigint(20)                      not null        comment '',
  reply_pk                	bigint(20)      default '0'     not null        comment '',
  reply_fk                  bigint(20)      default '0'     not null        comment '',
  user_id                   bigint(20)                      not null        comment '',
  comment                   longtext                        not null        comment '',
  submittime                datetime                        not null        comment '',
  ipaddress                 varchar(255)    default NULL    not null        comment '',
  is_removed                tinyint(4)      default '0'     not null        comment '',
  obj_pk_type               bigint(20)      default '0'     not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='评论';
create index idx_tb_comments_obj_pk_id on tb_album(obj_pk_id);
create index idx_tb_comments_reply_pk on tb_album(reply_pk);


/**
*  tb_link       链接
*/
-- drop table if exists tb_link;
create table tb_link
(
  id                        bigint(20)      AUTO_INCREMENT      not null        comment '自动增长id',
  sitename                  varchar(80)     default ''          not null        comment '',
  siteavator                varchar(200)    default '/static/upload/default/user-default-60x60.png'       not null        comment '',
  url                       varchar(200)    default ''          not null        comment '',
  sitedesc                  varchar(300)    default ''          not null        comment '',
  rank                      tinyint(4)      default '0'         not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='链接';
create index idx_tb_link_sitename on tb_album(sitename);



/**
*  tb_link       链接
*/
-- drop table if exists tb_mood;
create table tb_mood
(
  id                        bigint(20)      AUTO_INCREMENT      not null        comment '自动增长id',
  content                   longtext                            not null        comment '',
  cover                     varchar(70)     default '/static/upload/default/blog-default-0.png'       not null        comment '',
  posttime                  datetime                            not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='链接';
create index idx_tb_mood_posttime on tb_album(posttime);



/**
*  tb_link       链接
*/
-- drop table if exists tb_option;
create table tb_option
(
  id                        bigint(20)      AUTO_INCREMENT      not null        comment '自动增长id',
  name                      varchar(255)    DEFAULT ''          not null        comment '',
  value                     longtext                            not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='链接';
create index idx_tb_option_name on tb_album(name);



/**
*  tb_permission       权限
*/
-- drop table if exists tb_permission;
create table tb_permission
(
  id                        int(11)         AUTO_INCREMENT      not null        comment '自动增长id',
  name                      varchar(20)     DEFAULT ''          not null        comment '',
  value                     longtext                            not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='权限';


/**
*  tb_photo       照片管理
*/
-- drop table if exists tb_photo;
create table tb_photo
(
  id                        bigint(20)         AUTO_INCREMENT      not null        comment '自动增长id',
  albumid                   bigint(20)         DEFAULT '0'         not null        comment '',
  des                       varchar(100)       DEFAULT ''          not null        comment '',
  posttime                  datetime                               not null        comment '',
  url                       varchar(70)        DEFAULT ''          not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='照片管理';

/**
*  tb_post       照片管理
*/
-- DROP TABLE IF EXISTS `tb_post`;
CREATE TABLE `tb_post`
(
    `id`       bigint(20)   NOT NULL AUTO_INCREMENT,
    `user_id`  bigint(20)   NOT NULL,
    `title`    varchar(100) NOT NULL DEFAULT '',
    `color`    varchar(7)   NOT NULL DEFAULT '',
    `urlname`  varchar(100) NOT NULL DEFAULT '',
    `urltype`  tinyint(4)   NOT NULL DEFAULT '0',
    `content`  longtext     NOT NULL,
    `tags`     varchar(100) NOT NULL DEFAULT '',
    `posttime` datetime     NOT NULL,
    `views`    bigint(20)   NOT NULL DEFAULT '0',
    `status`   tinyint(4)   NOT NULL DEFAULT '0',
    `updated`  datetime     NOT NULL,
    `istop`    tinyint(4)   NOT NULL DEFAULT '0',
    `cover`    varchar(70)  NOT NULL DEFAULT '/static/upload/default/blog-default-0.png',
    PRIMARY KEY (`id`),
    KEY `tb_post_user_id` (`user_id`),
    KEY `tb_post_title` (`title`),
    KEY `tb_post_color` (`color`),
    KEY `tb_post_urlname` (`urlname`),
    KEY `tb_post_urltype` (`urltype`),
    KEY `tb_post_tags` (`tags`),
    KEY `tb_post_posttime` (`posttime`),
    KEY `tb_post_views` (`views`),
    KEY `tb_post_status` (`status`),
    KEY `tb_post_updated` (`updated`),
    KEY `tb_post_istop` (`istop`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


/**
*  tb_tag
*/
-- DROP TABLE IF EXISTS `tb_tag`;
CREATE TABLE `tb_tag`
(
    `id`    bigint(20)  NOT NULL AUTO_INCREMENT,
    `name`  varchar(20) NOT NULL DEFAULT '',
    `count` bigint(20)  NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `tb_tag_name` (`name`),
    KEY `tb_tag_count` (`count`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;




/**
*  tb_tag_post
*/
-- DROP TABLE IF EXISTS `tb_tag_post`;
CREATE TABLE `tb_tag_post`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `tag_id`     bigint(20) NOT NULL DEFAULT '0',
    `postid`     bigint(20) NOT NULL DEFAULT '0',
    `poststatus` tinyint(4) NOT NULL DEFAULT '0',
    `posttime`   datetime   NOT NULL,
    PRIMARY KEY (`id`),
    KEY `tb_tag_post_tag_id` (`tag_id`),
    KEY `tb_tag_post_postid` (`postid`),
    KEY `tb_tag_post_poststatus` (`poststatus`),
    KEY `tb_tag_post_posttime` (`posttime`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;


/**
*  tb_user       用户管理
*/
-- DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`
(
    `id`         bigint(20)   NOT NULL AUTO_INCREMENT,
    `username`   varchar(15)  NOT NULL DEFAULT '',
    `password`   varchar(32)  NOT NULL DEFAULT '密码',
    `nickname`   varchar(15)  NOT NULL DEFAULT '',
    `email`      varchar(50)  NOT NULL DEFAULT '',
    `salt`       varchar(20)  NOT NULL DEFAULT '密码盐',
    `lastlogin`  datetime     NOT NULL,
    `logincount` bigint(20)   NOT NULL DEFAULT '0',
    `lastip`     varchar(32)  NOT NULL DEFAULT '',
    `authkey`    varchar(10)  NOT NULL DEFAULT '',
    `active`     tinyint(4)   NOT NULL DEFAULT '0',
    `permission` varchar(100) NOT NULL DEFAULT '',
    `avator`     varchar(150) NOT NULL DEFAULT '/static/upload/default/user-default-60x60.png',
    `upcount`    bigint(20)   NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    KEY `tb_user_nickname` (`nickname`),
    KEY `tb_user_email` (`email`),
    KEY `tb_user_lastlogin` (`lastlogin`),
    KEY `tb_user_logincount` (`logincount`),
    KEY `tb_user_lastip` (`lastip`),
    KEY `tb_user_permission` (`permission`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8;


INSERT INTO `tb_link`
VALUES (1, '爱在发烧', '/static/upload/smallpic/20180619/1529369471932023700.jpg', 'http://azfashao.com/',
        '一个非常棒的站点，博主也很厉害', 99),
       (2, 'AN STUDIO', '/static/upload/smallpic/20180621/1529576340049076300.jpg', '//shop59002320.taobao.com',
        '外设韩国原单店铺', 100);


INSERT INTO `tb_option`
VALUES (1, 'sitename', 'inana用心交织的生活'),
       (2, 'siteurl', 'https://inana.top/'),
       (3, 'subtitle', '带着她和她去旅行'),
       (4, 'pagesize', '15'),
       (5, 'keywords', 'Python,MySQL,Golang,Windows,Linux'),
       (6, 'description', '来一场说走就走的旅行'),
       (7, 'theme', 'double'),
       (8, 'timezone', '8'),
       (9, 'stat', ''),
       (10, 'weibo', 'https://weibo.com/p/1005051484763434'),
       (11, 'github', 'https://github.com/griffin702'),
       (12, 'mybirth', '1987-09-30'),
       (13, 'albumsize', '9'),
       (14, 'nickname', '云丶先生'),
       (15, 'myoldcity', '湖北省 黄石市'),
       (16, 'mycity', '湖北省 武汉市'),
       (17, 'myprifessional', '游戏运维攻城师'),
       (18, 'myworkdesc', '1、Windows、Linux服务器运维，主要包括IIS、Apache、Nginx、Firewall、MySQL、SQLServer等常用服务。
2、日常备份与灾备恢复等确保数据安全，以及DBA相关其他职能。
3、负责公司内部网络运维，硬件维护、内外网分离以及常用第三方软件运维，主要包括SVN、FTP、BugFree、企业邮箱等服务。
4、业务线部分Golang服务端维护及二次开发。'),
       (19, 'mylang', 'Python、Golang、SQL、Shell'),
       (20, 'mylike', '旅行、游戏、技术');

INSERT INTO `tb_permission`
VALUES (1, 'user'),
      (2, 'article'),
      (3, 'album'),
      (4, 'link'),
      (5, 'mood'),
      (6, 'tag'),
      (7, 'system'),
      (8, 'fileupload');


INSERT INTO `tb_user`
VALUES (1, 'admin', 'b2cecd1598289f6fc5b5f4abe85ff159', '刘丶先生', '362310566@qq.com','abtfg', '2018-07-12 12:55:55', 0,
        '127.0.0.1', '', 1, '1|2|3|4|5|6|7|8', '/static/upload/default/user-default-60x60.png', 0);
