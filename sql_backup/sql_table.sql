create database mscoin;
    /*
    用户表单
    */
CREATE TABLE `member`  (
                           `id` bigint(0) NOT NULL AUTO_INCREMENT,
                           `ali_no` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `qr_code_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `appeal_success_times` int(0) NOT NULL,
                           `appeal_times` int(0) NOT NULL,
                           `application_time` bigint(0) NOT NULL,
                           `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `bank` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `branch` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `card_no` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `certified_business_apply_time` bigint(0) NOT NULL,
                           `certified_business_check_time` bigint(0) NOT NULL,
                           `certified_business_status` int(0) NOT NULL,
                           `channel_id` int(0) NOT NULL DEFAULT 0,
                           `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `first_level` int(0) NOT NULL,
                           `google_date` bigint(0) NOT NULL,
                           `google_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `google_state` int(0) NOT NULL DEFAULT 0,
                           `id_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `inviter_id` bigint(0) NOT NULL,
                           `is_channel` int(0) NOT NULL DEFAULT 0,
                           `jy_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `last_login_time` bigint(0) NOT NULL,
                           `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `country` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `district` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `login_count` int(0) NOT NULL,
                           `login_lock` int(0) NOT NULL,
                           `margin` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `member_level` int(0) NOT NULL,
                           `mobile_phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `promotion_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `publish_advertise` int(0) NOT NULL,
                           `real_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `real_name_status` int(0) NOT NULL,
                           `registration_time` bigint(0) NOT NULL,
                           `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `second_level` int(0) NOT NULL,
                           `sign_in_ability` tinyint(4) NOT NULL DEFAULT b'1',
                           `status` int(0) NOT NULL,
                           `third_level` int(0) NOT NULL,
                           `token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `token_expire_time` bigint(0) NOT NULL,
                           `transaction_status` int(0) NOT NULL,
                           `transaction_time` bigint(0) NOT NULL,
                           `transactions` int(0) NOT NULL,
                           `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `qr_we_code_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `wechat` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `local` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `integration` bigint(0) NOT NULL DEFAULT 0,
                           `member_grade_id` bigint(0) NOT NULL DEFAULT 1 COMMENT '等级id',
                           `kyc_status` int(0) NOT NULL DEFAULT 0 COMMENT 'kyc等级',
                           `generalize_total` bigint(0) NOT NULL DEFAULT 0 COMMENT '注册赠送积分',
                           `inviter_parent_id` bigint(0) NOT NULL DEFAULT 0,
                           `super_partner` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `kick_fee` decimal(19, 2) NOT NULL,
                           `power` decimal(8, 4) NOT NULL DEFAULT 0.0000 COMMENT '个人矿机算力(每日维护)',
                           `team_level` int(0) NOT NULL DEFAULT 0 COMMENT '团队人数(每日维护)',
                           `team_power` decimal(8, 4) NOT NULL DEFAULT 0.0000 COMMENT '团队矿机算力(每日维护)',
                           `member_level_id` bigint(0) NOT NULL,
                           PRIMARY KEY (`id`) USING BTREE,
                           UNIQUE INDEX `UK_gc3jmn7c2abyo3wf6syln5t2i`(`username`) USING BTREE,
                           UNIQUE INDEX `UK_10ixebfiyeqolglpuye0qb49u`(`mobile_phone`) USING BTREE,
                           INDEX `FKbt72vgf5myy3uhygc90xna65j`(`local`) USING BTREE,
                           INDEX `FK8jlqfg5xqj5epm9fpke6iotfw`(`member_level_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

CREATE TABLE `coin`  (
                         `id` int(0) NOT NULL AUTO_INCREMENT,
                         `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '货币',
                         `can_auto_withdraw` int(0) NOT NULL COMMENT '是否能自动提币',
                         `can_recharge` int(0) NOT NULL COMMENT '是否能充币',
                         `can_transfer` int(0) NOT NULL COMMENT '是否能转账',
                         `can_withdraw` int(0) NOT NULL COMMENT '是否能提币',
                         `cny_rate` double NOT NULL COMMENT '对人民币汇率',
                         `enable_rpc` int(0) NOT NULL COMMENT '是否支持rpc接口',
                         `is_platform_coin` int(0) NOT NULL COMMENT '是否是平台币',
                         `max_tx_fee` double NOT NULL COMMENT '最大提币手续费',
                         `max_withdraw_amount` decimal(18, 8) NOT NULL COMMENT '最大提币数量',
                         `min_tx_fee` double NOT NULL COMMENT '最小提币手续费',
                         `min_withdraw_amount` decimal(18, 8) NOT NULL COMMENT '最小提币数量',
                         `name_cn` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '中文名称',
                         `sort` int(0) NOT NULL COMMENT '排序',
                         `status` tinyint(0) NOT NULL COMMENT '状态 0 正常 1非法',
                         `unit` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '单位',
                         `usd_rate` double NOT NULL COMMENT '对美元汇率',
                         `withdraw_threshold` decimal(18, 8) NOT NULL COMMENT '提现阈值',
                         `has_legal` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是合法币种',
                         `cold_wallet_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '冷钱包地址',
                         `miner_fee` decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '转账时付给矿工的手续费',
                         `withdraw_scale` int(0) NOT NULL DEFAULT 4 COMMENT '提币精度',
                         `account_type` int(0) NOT NULL DEFAULT 0 COMMENT '币种账户类型0：默认  1：EOS类型',
                         `deposit_address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '充值地址',
                         `infolink` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种资料链接',
                         `information` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '币种简介',
                         `min_recharge_amount` decimal(18, 8) NOT NULL COMMENT '最小充值数量',
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;
INSERT INTO `mscoin`.`coin`(`id`, `name`, `can_auto_withdraw`, `can_recharge`, `can_transfer`, `can_withdraw`, `cny_rate`, `enable_rpc`, `is_platform_coin`, `max_tx_fee`, `max_withdraw_amount`, `min_tx_fee`, `min_withdraw_amount`, `name_cn`, `sort`, `status`, `unit`, `usd_rate`, `withdraw_threshold`, `has_legal`, `cold_wallet_address`, `miner_fee`, `withdraw_scale`, `account_type`, `deposit_address`, `infolink`, `information`, `min_recharge_amount`) VALUES (1, 'Bitcoin', 0, 0, 1, 0, 0, 0, 0, 0.0002, 5.00000000, 0.0002, 0.00100000, '比特币', 1, 0, 'BTC', 0, 0.10000000, 0, '0', 0.00000000, 4, 0, '', '', '', 0.00000000);
INSERT INTO `mscoin`.`coin`(`id`, `name`, `can_auto_withdraw`, `can_recharge`, `can_transfer`, `can_withdraw`, `cny_rate`, `enable_rpc`, `is_platform_coin`, `max_tx_fee`, `max_withdraw_amount`, `min_tx_fee`, `min_withdraw_amount`, `name_cn`, `sort`, `status`, `unit`, `usd_rate`, `withdraw_threshold`, `has_legal`, `cold_wallet_address`, `miner_fee`, `withdraw_scale`, `account_type`, `deposit_address`, `infolink`, `information`, `min_recharge_amount`) VALUES (2, 'Bitcoincash', 1, 1, 1, 1, 0, 1, 0, 0.0002, 5.00000000, 0.0002, 0.00100000, '比特现金', 1, 0, 'BCH', 0, 0.10000000, 0, '0', 0.00000000, 4, 0, '', '', '', 0.00000000);
INSERT INTO `mscoin`.`coin`(`id`, `name`, `can_auto_withdraw`, `can_recharge`, `can_transfer`, `can_withdraw`, `cny_rate`, `enable_rpc`, `is_platform_coin`, `max_tx_fee`, `max_withdraw_amount`, `min_tx_fee`, `min_withdraw_amount`, `name_cn`, `sort`, `status`, `unit`, `usd_rate`, `withdraw_threshold`, `has_legal`, `cold_wallet_address`, `miner_fee`, `withdraw_scale`, `account_type`, `deposit_address`, `infolink`, `information`, `min_recharge_amount`) VALUES (3, 'DASH', 1, 1, 1, 1, 0, 1, 0, 0.0002, 5.00000000, 0.0002, 0.00100000, '达世币', 1, 0, 'DASH', 0, 0.10000000, 0, '0', 0.00000000, 4, 0, '', '', '', 0.00000000);
INSERT INTO `mscoin`.`coin`(`id`, `name`, `can_auto_withdraw`, `can_recharge`, `can_transfer`, `can_withdraw`, `cny_rate`, `enable_rpc`, `is_platform_coin`, `max_tx_fee`, `max_withdraw_amount`, `min_tx_fee`, `min_withdraw_amount`, `name_cn`, `sort`, `status`, `unit`, `usd_rate`, `withdraw_threshold`, `has_legal`, `cold_wallet_address`, `miner_fee`, `withdraw_scale`, `account_type`, `deposit_address`, `infolink`, `information`, `min_recharge_amount`) VALUES (4, 'Ethereum', 1, 1, 1, 1, 0, 1, 0, 0.0002, 5.00000000, 0.0002, 0.00100000, '以太坊', 1, 0, 'ETH', 0, 0.10000000, 0, '0', 0.00000000, 4, 0, '', '', '', 0.00000000);
INSERT INTO `mscoin`.`coin`(`id`, `name`, `can_auto_withdraw`, `can_recharge`, `can_transfer`, `can_withdraw`, `cny_rate`, `enable_rpc`, `is_platform_coin`, `max_tx_fee`, `max_withdraw_amount`, `min_tx_fee`, `min_withdraw_amount`, `name_cn`, `sort`, `status`, `unit`, `usd_rate`, `withdraw_threshold`, `has_legal`, `cold_wallet_address`, `miner_fee`, `withdraw_scale`, `account_type`, `deposit_address`, `infolink`, `information`, `min_recharge_amount`) VALUES (5, 'GalaxyChain', 1, 1, 1, 1, 1, 1, 0, 0.0002, 5.00000000, 0.0002, 0.00100000, '银河链', 1, 0, 'GCC', 0, 0.10000000, 0, '0', 0.00000000, 4, 0, '', '', '', 0.00000000);
INSERT INTO `mscoin`.`coin`(`id`, `name`, `can_auto_withdraw`, `can_recharge`, `can_transfer`, `can_withdraw`, `cny_rate`, `enable_rpc`, `is_platform_coin`, `max_tx_fee`, `max_withdraw_amount`, `min_tx_fee`, `min_withdraw_amount`, `name_cn`, `sort`, `status`, `unit`, `usd_rate`, `withdraw_threshold`, `has_legal`, `cold_wallet_address`, `miner_fee`, `withdraw_scale`, `account_type`, `deposit_address`, `infolink`, `information`, `min_recharge_amount`) VALUES (6, 'Litecoin', 1, 0, 1, 1, 1, 1, 0, 0.0002, 5.00000000, 0.0002, 0.00100000, '莱特币', 1, 0, 'LTC', 0, 0.10000000, 0, '0', 0.00000000, 4, 0, '', '', '', 0.00000000);
INSERT INTO `mscoin`.`coin`(`id`, `name`, `can_auto_withdraw`, `can_recharge`, `can_transfer`, `can_withdraw`, `cny_rate`, `enable_rpc`, `is_platform_coin`, `max_tx_fee`, `max_withdraw_amount`, `min_tx_fee`, `min_withdraw_amount`, `name_cn`, `sort`, `status`, `unit`, `usd_rate`, `withdraw_threshold`, `has_legal`, `cold_wallet_address`, `miner_fee`, `withdraw_scale`, `account_type`, `deposit_address`, `infolink`, `information`, `min_recharge_amount`) VALUES (7, 'SGD', 1, 1, 1, 1, 0, 1, 0, 0.0002, 500.00000000, 1, 1.00000000, '新币', 4, 0, 'SGD', 0, 0.10000000, 1, '0', 0.00000000, 4, 0, '', '', '', 0.00000000);
INSERT INTO `mscoin`.`coin`(`id`, `name`, `can_auto_withdraw`, `can_recharge`, `can_transfer`, `can_withdraw`, `cny_rate`, `enable_rpc`, `is_platform_coin`, `max_tx_fee`, `max_withdraw_amount`, `min_tx_fee`, `min_withdraw_amount`, `name_cn`, `sort`, `status`, `unit`, `usd_rate`, `withdraw_threshold`, `has_legal`, `cold_wallet_address`, `miner_fee`, `withdraw_scale`, `account_type`, `deposit_address`, `infolink`, `information`, `min_recharge_amount`) VALUES (8, 'USDT', 1, 1, 1, 1, 0, 1, 0, 0.0002, 5.00000000, 0.0002, 0.00100000, '泰达币T', 1, 0, 'USDT', 0, 0.10000000, 0, '0', 0.00000000, 4, 0, '', '', '', 0.00000000);

