create table `register` (
    `id` bigint(20) not null auto_increment ,
    `name` varchar(64) collate utf8mb4_general_ci not null ,
    `gender` tinyint(4) not null default '0' ,
    `phone` varchar(64) collate utf8mb4_general_ci not null ,
    `qq` varchar(64) collate utf8mb4_general_ci not null ,
    `specialty` varchar(64) collate utf8mb4_general_ci not null ,
    `self_introduction` varchar(64) collate utf8mb4_general_ci not null ,
    primary key ('id')
)engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci;