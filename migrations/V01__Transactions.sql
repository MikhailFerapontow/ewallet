CREATE TABLE IF NOT EXISTS Transactions (
    "time" VARCHAR(100) NOT NULL,
    "from" VARCHAR(32) REFERENCES Wallet(id) NOT NULL,
    "to" VARCHAR(32) REFERENCES Wallet(id) NOT NULL,
    amount FLOAT CHECK (amount >= 0) NOT NULL
)
