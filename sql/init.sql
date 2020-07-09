DROP TABLE IF EXISTS `id_space`;
CREATE TABLE `id_space` (
  `space_name` varchar(45) NOT NULL DEFAULT '',
  `prefix` varchar(45) NOT NULL DEFAULT '',
  `suffix` varchar(45) NOT NULL DEFAULT '',
  `seed` bigint(20) unsigned NOT NULL,
  `batch_size` bigint(11) unsigned NOT NULL,
  PRIMARY KEY (`space_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `id_space` (`space_name`, `prefix`, `suffix`, `seed`, `batch_size`) VALUES ('label', 'label', '', '8541000', '1000');
INSERT INTO `id_space` (`space_name`, `prefix`, `suffix`, `seed`, `batch_size`) VALUES ('order', 'order', 'WD0ZV6AFGSCU9MBKN24HJL578XIOP31QERTY', '78436785196604', '1000');
INSERT INTO `id_space` (`space_name`, `prefix`, `suffix`, `seed`, `batch_size`) VALUES ('recommendcode', '', '0ZV6AWDFGSCU9HJL578X1MBKN24QERTYIOP3', '78436785196604', '1000');
INSERT INTO `id_space` (`space_name`, `prefix`, `suffix`, `seed`, `batch_size`) VALUES ('user', 'user', '', '8541000', '1000');


INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_ADMIN', 'sitea', '*https://res.cloudinary.com/*', '*', 'allow', NULL);
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_ADMIN', 'sitea', '/acc/*', '*', 'allow', NULL);
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '1', 'ROLE_ADMIN', 'sitea', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_BASIC', 'sitea', '/acc/paychannels', 'GET', 'allow', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_BASIC', 'sitea', '/acc/systemmessage', 'GET', 'allow', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_BASIC', 'sitea', '/acc/approval/alert', 'GET', 'allow', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_BASIC', 'sitea', '/acc/roles', 'GET', 'allow', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_BASIC', 'sitea', '/acc/menus', 'GET', 'allow', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_BASIC', 'sitea', '/acc/login*', 'POST', 'allow', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'ROLE_BASIC', 'sitea', '/acc/report/data', 'GET', 'allow', '');
