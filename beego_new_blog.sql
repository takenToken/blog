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
  ranking                   tinyint(4)      default '0'     not null        comment '排名',
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
  ipaddress                 varchar(255)    default ''      not null        comment '',
  is_removed                tinyint(4)      default '0'     not null        comment '',
  obj_pk_type               bigint(20)      default '0'     not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='评论';
create index idx_tb_comments_obj_pk_id on tb_comments(obj_pk_id);
create index idx_tb_comments_reply_pk on tb_comments(reply_pk);


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
  ranking                   tinyint(4)      default '0'         not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='链接';
create index idx_tb_link_sitename on tb_link(sitename);



/**
*  tb_mood       心情
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
create index idx_tb_mood_posttime on tb_mood(posttime);



/**
*  tb_option       配置
*/
-- drop table if exists tb_option;
create table tb_option
(
  id                        bigint(20)      AUTO_INCREMENT      not null        comment '自动增长id',
  name                      varchar(255)    DEFAULT ''          not null        comment '链接名',
  value                     longtext                            not null        comment '链接值',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='链接';
create index idx_tb_option_name on tb_option(name);



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
*  tb_post       文章管理
*/
-- drop table if exists tb_post;
create table tb_post
(
  id                        bigint(20)         AUTO_INCREMENT      not null        comment '自动增长id',
  user_id                   bigint(20)         DEFAULT '0'         not null        comment '',
  title                     varchar(100)       DEFAULT ''          not null        comment '',
  color                     varchar(7)         DEFAULT ''          not null        comment '',
  urlname                   varchar(100)       DEFAULT ''          not null        comment '',
  urltype                   tinyint(4)         DEFAULT '0'         not null        comment '',
  content                   longtext                               not null        comment '',
  tags                      varchar(100)       DEFAULT ''          not null        comment '',
  posttime                  datetime                               not null        comment '',
  views                     bigint(20)         DEFAULT '0'         not null        comment '',
  status                    tinyint(4)         DEFAULT '0'         not null        comment '',
  updated                   datetime                               not null        comment '',
  istop                     tinyint(4)         DEFAULT '0'         not null        comment '',
  cover                     varchar(70)        DEFAULT '/static/upload/default/blog-default-0.png'          not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='照片管理';
create index idx_tb_post_title on tb_post(title);

/**
*  tb_tag       标签管理
*/
-- drop table if exists tb_tag;
create table tb_tag
(
  id                        bigint(20)         AUTO_INCREMENT      not null        comment '自动增长id',
  name                      varchar(20)        DEFAULT ''         not null         comment '',
  count                     varchar(20)        DEFAULT '0'          not null       comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='标签管理';
create index idx_tb_tag_name on tb_tag(name);


/**
*  tb_tag_post       标签文章管理
*/
-- drop table if exists tb_tag_post;
create table tb_tag_post
(
  id                        bigint(20)         AUTO_INCREMENT      not null        comment '自动增长id',
  tag_id                    bigint(20)         DEFAULT '0'         not null        comment '',
  postid                    bigint(20)         DEFAULT '0'         not null        comment '',
  poststatus                tinyint(4)         DEFAULT '0'         not null        comment '',
  posttime                  datetime                               not null        comment '',
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='标签文章管理';


/**
*  tb_user       用户管理
*/
-- drop table if exists tb_user;
create table tb_user
(
  id                        bigint(20)         AUTO_INCREMENT      not null        comment '自动增长id',
  username                  varchar(32)                            not null        comment '账号',
  password                  varchar(32)                            not null        comment '密码',
  nickname                  varchar(15)                            not null        comment '昵称',
  email                     varchar(50)                            not null        comment '邮箱',
  salt                      varchar(20)                            not null        comment '密码盐',
  lastlogin                 datetime                               not null        comment '最后登录时间',
  logincount                bigint(20)         DEFAULT '0'         not null        comment '登录次数',
  lastip                    varchar(32)        DEFAULT ''          not null        comment '最后登录ip',
  authkey                   varchar(10)        DEFAULT ''          not null        comment '',
  active                    tinyint(4)         DEFAULT '0'         not null        comment '是否激活',
  permission                varchar(100)       DEFAULT ''          not null        comment '权限',
  avator                    varchar(150)       DEFAULT '/static/upload/default/user-default-60x60.png'         not null        comment '',
  upcount                   bigint(20)         DEFAULT '0'         not null        comment '',
  primary key(id),
  constraint unique_tb_user_username unique(username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='用户管理';
create index idx_tb_user_email on tb_user(email);


INSERT INTO tb_link
VALUES (1, '爱在发烧', '/static/upload/smallpic/20180619/1529369471932023700.jpg', 'http://azfashao.com/',
        '一个非常棒的站点，博主也很厉害', 99);


INSERT INTO tb_option
VALUES (1, 'sitename', 'pack'),
       (2, 'siteurl', 'https://qlblog.cn'),
       (3, 'subtitle', '创造价值,更好生活'),
       (4, 'pagesize', '15'),
       (5, 'keywords', 'Java,MySQL,Golang,Redis,Kafka,Zookeeper,JVM,Docker,K8s,Windows,Linux'),
       (6, 'description', '学无止境'),
       (7, 'theme', 'double'),
       (8, 'timezone', '8'),
       (9, 'stat', ''),
       (11, 'github', 'https://github.com/takenToken'),
       (12, 'mybirth', '1989-06-11'),
       (13, 'albumsize', '9'),
       (14, 'nickname', '刘丶先生'),
       (15, 'myoldcity', '广东省 深圳市'),
       (16, 'mycity', '广东省 深圳市'),
       (17, 'myprifessional', 'Java/Go工程师'),
       (18, 'myworkdesc', '1、构建企业级微服务平台,熟悉Docker。
2、负责架构系统调优,线上问题跟进。
3、负责公司供应链业务功能开发。
4、希望云平台Golang、零售方向前进'),
       (19, 'mylang', 'Java、Golang、Java Framework'),
       (20, 'mylike', '技术、运动、旅行');

INSERT INTO tb_permission VALUES (1, 'user',''),
      (2, 'article',''),
      (3, 'album',''),
      (4, 'link',''),
      (5, 'mood',''),
      (6, 'tag',''),
      (7, 'system',''),
      (8, 'fileupload','');


INSERT INTO `tb_user`
VALUES (1, 'admin', 'b2cecd1598289f6fc5b5f4abe85ff159', '刘丶先生', '362310566@qq.com','abtfg', '2018-07-12 12:55:55', 0,
        '127.0.0.1', '', 1, '1|2|3|4|5|6|7|8', '/static/upload/default/user-default-60x60.png', 0);
