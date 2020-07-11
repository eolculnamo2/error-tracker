SET FOREIGN_KEY_CHECKS=0;
CREATE TABLE `users` (
  `email` varchar(150) UNIQUE,
  `password` varchar(500) NOT NULL,
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) DEFAULT NULL,
  `organization_id` INT DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT NULL,
  `authority` varchar(45) DEFAULT NULL,

  PRIMARY KEY (`email`),
  FOREIGN KEY(`organization_id`)
  REFERENCES `organizations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `organizations` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `website` VARCHAR(200) DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT 1,

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
SET FOREIGN_KEY_CHECKS=1;