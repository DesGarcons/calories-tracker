create table if not exists users(
    id serial primary key,
    tid BIGINT,
    height NUMERIC(5,2),
    weight NUMERIC(5,2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
