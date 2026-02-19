create table books
(
    id          bigint auto_increment
        primary key,
    isbn        varchar(17)          not null,
    title       varchar(255)         not null,
    author      varchar(255)         not null,
    on_hold     tinyint(1) default 0 not null,
    checked_out tinyint(1) default 0 not null,
    late_fee    decimal    default 5 not null
);

create table users
(
    id             bigint auto_increment
        primary key,
    account_number bigint               not null,
    email          varchar(255)         not null,
    first_name     varchar(255)         not null,
    is_admin       tinyint(1) default 0 not null invisible,
    created_at     datetime             not null,
    constraint users_pk_2
        unique (account_number),
    constraint users_pk_3
        unique (email)
);

create table book_holds
(
    id            bigint auto_increment
        primary key,
    user_id       bigint                       not null,
    book_id       bigint                       not null,
    status        varchar(7) default 'Waiting' not null,
    hold_datetime datetime                     not null,
    pickup_due    datetime                     null,
    constraint book_holds_books_id_fk
        foreign key (book_id) references books (id),
    constraint book_holds_users_id_fk
        foreign key (user_id) references users (id)
);

create table borrowed_books
(
    id              bigint auto_increment
        primary key,
    user_id         bigint   not null,
    book_id         bigint   not null,
    borrow_datetime datetime not null,
    return_datetime datetime not null,
    constraint borrowed_books_books_id_fk
        foreign key (book_id) references books (id),
    constraint borrowed_books_users_id_fk
        foreign key (user_id) references users (id)
);

create table user_credentials
(
    id            bigint auto_increment
        primary key,
    user_id       bigint       not null,
    password_hash varchar(255) not null,
    constraint user_credentials_pk
        unique (user_id),
    constraint user_id
        foreign key (user_id) references users (id)
);

