
-- Table Definition
CREATE TABLE IF NOT EXISTS expenses (
    id SERIAL PRIMARY KEY,
    title TEXT,
    amount FLOAT,
    note TEXT,
    tags TEXT[]
);

INSERT INTO "expenses" ("title", "amount", "note", "tags") VALUES ('test-title', 100, 'test-note', '{test-tag1}');
