create table users(
    username varchar(255) primary key comment '用户名',
    password varchar(510) not null comment '密码',
    email varchar(50) not null comment '邮箱',
    phone varchar(30) comment '手机号'
);

create table files(
    filehash varchar(510) primary key comment '文件hash',
    filepath varchar(510) not null comment '存储位置',
    filestatus tinyint default 1 comment '文件状态'
)

create table userfiles(
    username varchar(255) not null unique,
    filename varchar(255) unique not null ,
    filehash varchar(510) not null unique ,
    foreign key (username) REFERENCES users(username),
    foreign key (filehash) REFERENCES files(filehash)
);