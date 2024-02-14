
create TABLE sessions (
    token char(43) primary key,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

create INDEX sessions_expiry_idx on sessions (expiry);