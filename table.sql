create table users
(
  id           serial                not null
    constraint users_pk
      primary key,
  real_name    text,
  student_id   text,
  username     text                  not null,
  password     integer               not null,
  created_at   timestamp,
  modified_at  timestamp,
  is_admin     boolean default false not null,
  is_certified boolean default false not null,
  is_blocked   boolean default false not null,
  email        text                  not null
);

alter table users
  owner to improcuss;

create unique index users_email_uindex
  on users (email);

create unique index users_student_id_uindex
  on users (student_id);

create unique index users_username_uindex
  on users (username);

create table sessions
(
  id         serial      not null
    constraint sessions_pk
      primary key,
  uuid       varchar(64) not null,
  email      text        not null,
  user_id    integer
    constraint sessions_users__fk
      references users,
  created_at timestamp
);

alter table sessions
  owner to improcuss;

create unique index sessions_uuid_uindex
  on sessions (uuid);

