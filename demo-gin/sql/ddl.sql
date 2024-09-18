CREATE TABLE `t_user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `user_name` varchar(255) NOT NULL COMMENT '用户名',
    `phone` varchar(255) DEFAULT NULL COMMENT '用户手机号',
    `password` varchar(255) NOT NULL COMMENT '用户密码sha1哈希',
    `status` int(1) DEFAULT NULL COMMENT '用户状态',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用户注册时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '用户注册时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_user_id` (`user_id`) COMMENT '用户Id唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;