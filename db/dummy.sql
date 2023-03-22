INSERT INTO "loan_applications" ("personal_id", "name", "amount", "term", "created_at")
VALUES
    (1, "John Smith", 5700, 24, '2022-03-15 15:30:00'),
    (2, "Michael Scott ", 3570, 24, '2022-03-12 15:30:00'),
    (2, "Michael Scott ", 1500, 12, '2022-03-15 15:30:00'),
    (3, "Christian Bale", 30000, 36, CURRENT_TIMESTAMP),
    (4, "The Rock", 10000, 48, CURRENT_TIMESTAMP);

INSERT INTO "blacklist" ("personal_id")
VALUES
    (2),
    (4);