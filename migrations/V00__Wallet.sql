CREATE TABLE IF NOT EXISTS Wallet (
    id VARCHAR(32) PRIMARY KEY DEFAULT md5(random()::text),
    balance FLOAT CHECK (balance >= 0)
)