INSERT INTO "loan_applications" ("personal_id", "name", "amount", "term")
VALUES
    (2, "John Smith", 5700, 24),
    (4, "Michael Scott ", 1500, 12),
    (5, "Christian Bale", 30000, 36),
    (9, "The Rock", 10000, 48);

INSERT INTO "blacklist" ("personal_id")
VALUES
    (1),
    (4),
    (9);