package onlineDDL

type BaseModel struct {
	Id        int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:自增id" json:"id"`
	CreatedAt string `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:最后更新时间;NOT NULL" json:"updated_at"`
}

var CreateTableSql = `
create table {{.TableName}}
(
    id int(11) not null auto_increment comment '自增id',
    {{ range .TableFields }}{{ . }}{{- end }}
    created_at timestamp not null default current_timestamp comment '创建时间',
    updated_at timestamp not null default current_timestamp on update current_timestamp comment '最后更新时间',
    primary key (id)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_unicode_ci;
`
