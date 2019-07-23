# 数据库表设计

## 0 创建mysql库

- 容器化：

```
MYPWD=`pwd`
docker run \
-p 3306:3306 \
--name c1mysql \
--privileged=true   \
-v $MYPWD/data:/mysql_data \
-e MYSQL_ROOT_PASSWORD=password \
-d mysql:latest
```

```
docker cp ddb816eca33e:/etc/mysql /root/db/mysql/conf
docker cp  /root/db/mysql/conf/my.cnf ddb816eca33e:/etc/mysql/my.cnf
```

- 解决：Authentication plugin 'caching_sha2_password' cannot be loaded

```
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'password'; 
```

## 1 数据库建库

```
CREATE USER "mysql_db"@"localhost" IDENTIFIED BY "password"; 
CREATE USER "mysql_db"@"%" IDENTIFIED BY "password"; 
DROP database mysql_dbDB;
create database mysql_dbDB default charset utf8 collate utf8_general_ci;
GRANT ALL PRIVILEGES ON mysql_db.* TO 'mysql_db'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON mysql_db.* TO 'mysql_db'@'%' WITH GRANT OPTION;
flush privileges; 

```


## 2 数据库建表


```$xslt
DROP TABLE IF EXISTS `demo_user`;
CREATE TABLE `demo_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  `pass` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
);
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pass` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `created_at` int(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(20) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for users_accounts
-- ----------------------------
DROP TABLE IF EXISTS `users_accounts`;
CREATE TABLE `users_accounts` (
  `uid` int(11) unsigned NOT NULL,
  `account_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账户',
  `account_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '账户类型: 1->手机号码, 2 -> 邮箱, 3 -> 微信',
  `created_at` int(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(20) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  UNIQUE KEY `uniq_uaa` (`uid`,`account_name`,`account_type`,`deleted_at`) USING BTREE COMMENT '唯一索引',
  UNIQUE KEY `uniq_index` (`account_name`,`account_type`,`deleted_at`) USING BTREE COMMENT '唯一索引',
  KEY `idx_uid` (`uid`) USING BTREE COMMENT '用户uid'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

```


## 3 数据初始化



