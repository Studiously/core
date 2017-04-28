-- +migrate Up
CREATE SEQUENCE global_id_sequence;

-- +migrate StatementBegin
CREATE FUNCTION generate_id(IN service_id int, OUT result bigint) AS $$
DECLARE
    -- TO START IDS SMALLER, YOU COULD CHANGE THIS TO A MORE RECENT UNIX TIMESTAMP
    our_epoch bigint := 1493337600;

    seq_id bigint;
    now_millis bigint;
    -- UNIQUE SERVICE IDENTIFIER
    -- CHANGE THIS FOR EACH SERVICE!!!
BEGIN
    SELECT nextval('global_id_sequence') % 1024 INTO seq_id;
    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp())) INTO now_millis;
    result := (now_millis - our_epoch) << 20;
    result := result | (service_id << 10);
    result := result | (seq_id);
END;
$$ LANGUAGE PLPGSQL;
-- +migrate StatementEnd