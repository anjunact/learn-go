CREATE TABLE `gratisography` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `img_url` varchar(255) DEFAULT NULL,
  `type_name` varchar(50) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `width` int(11) DEFAULT NULL,
  `height` int(11) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=388 DEFAULT CHARSET=utf8;
