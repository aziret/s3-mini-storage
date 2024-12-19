CREATE TABLE server_info (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    is_single_row BOOLEAN DEFAULT TRUE UNIQUE
);
