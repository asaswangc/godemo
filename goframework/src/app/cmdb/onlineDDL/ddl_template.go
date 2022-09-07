package onlineDDL

var CreateTableSql = `
create table %s
(
    id int(11) not null auto_increment comment '自增id',
    created_at timestamp not null default current_timestamp comment '创建时间',
    updated_at timestamp not null default current_timestamp on update current_timestamp comment '最后更新时间',
    primary key (id)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_unicode_ci;
`
