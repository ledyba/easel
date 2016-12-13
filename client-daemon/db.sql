CREATE TABLE `resample_requests` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`src` varchar(1024) NOT NULL,
`dst` varchar(1024) NOT NULL,
`dst_width` int(10) unsigned NOT NULL,
`dst_height` int(10) unsigned NOT NULL,
`dst_quality` double unsigned NOT NULL,
`created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`uploaded` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`status` tinyint(1) unsigned NOT NULL COMMENT '0=enqueued, 1=in progress, 2=done, 3=error',
`message` text NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT
