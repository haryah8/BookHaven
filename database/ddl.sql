CREATE TABLE
  users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    name VARCHAR(255) NOT NULL,
    balance INT DEFAULT 0 -- User balance for borrowing books
  );

CREATE TABLE
  books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_year INT,
    isbn VARCHAR(20) UNIQUE,
    available_copies INT NOT NULL DEFAULT 1 -- Number of available copies for borrowing
  );

CREATE TABLE
  borrowings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL,
    borrowed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- When the book was borrowed
    returned_at TIMESTAMP NULL, -- When the book was returned (nullable if not returned yet)
    status ENUM ('borrowed', 'returned') DEFAULT 'borrowed', -- Borrowing status
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_book_id FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE
  );

CREATE TABLE
  transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    amount INT NOT NULL,
    transaction_id VARCHAR(255) UNIQUE NOT NULL,
    status ENUM ('pending', 'completed', 'failed') DEFAULT 'pending',
    description TEXT, -- Added description column
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_transactions_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
  );

-- UPDATE borrowings set borrowed_at = "2024-08-10 20:13:54" where id = 11;
INSERT INTO
  books (
    title,
    author,
    published_year,
    isbn,
    available_copies
  )
VALUES
  (
    'Journey to the Deserted Island',
    'Sophia Miller',
    1956,
    '0336133059168',
    1
  ),
  (
    'The Secrets of the Last Dragon',
    'David Brown',
    1954,
    '7781256385042',
    5
  ),
  (
    'The Legend of the Crystal Cave',
    'David Williams',
    1978,
    '8791211014979',
    1
  ),
  (
    'The Mystery of the Forgotten Kingdom',
    'Robert Miller',
    1964,
    '4809182530252',
    2
  ),
  (
    'The Legend of the Silver Sword',
    'Robert Williams',
    1981,
    '3788392494976',
    5
  ),
  (
    'The Quest for the Crystal Cave',
    'Sophia Jones',
    1978,
    '7945627090611',
    3
  ),
  (
    'Escape from the Ancient Ruins',
    'Sophia Brown',
    1962,
    '6714289942965',
    5
  ),
  (
    'The Secrets of the Crystal Cave',
    'Michael Garcia',
    1985,
    '3826830123342',
    4
  ),
  (
    'The Rise of the Hidden Treasure',
    'David Martinez',
    1983,
    '2340768015059',
    4
  ),
  (
    'The Quest for the Silver Sword',
    'Sophia Garcia',
    2018,
    '9337460817246',
    1
  ),
  (
    'The Lost City of Atlantis',
    'Emily Chen',
    1992,
    '6543210987654',
    2
  ),
  (
    'The Curse of the Haunted Mansion',
    'James Davis',
    2001,
    '9876543210987',
    3
  ),
  (
    'The Adventure of the Golden Amulet',
    'Sarah Lee',
    1995,
    '1234567890123',
    4
  ),
  (
    'The Secret of the Ancient Temple',
    'Kevin White',
    2005,
    '4567890123456',
    5
  ),
  (
    'The Mystery of the Missing Heirloom',
    'Rebecca Hall',
    1998,
    '7890123456789',
    2
  ),
  (
    'The Quest for the Golden Chalice',
    'Michael Kim',
    2003,
    '9012345678901',
    3
  ),
  (
    'The Legend of the Phoenix',
    'Emily Patel',
    2012,
    '2345678901234',
    4
  ),
  (
    'The Adventure of the Lost City',
    'David Kim',
    2008,
    '3456789012345',
    5
  ),
  (
    'The Secret of the Haunted Forest',
    'Sarah Taylor',
    2015,
    '4567890123453',
    2
  ),
  (
    'The Mystery of the Abandoned Mine',
    'James Brown',
    2010,
    '5678901234566',
    3
  ),
  (
    'The Quest for the Golden Sword',
    'Rebecca Lee',
    2016,
    '6789012345671',
    4
  ),
  (
    'The Legend of the Dragon Eye',
    'Michael Davis',
    2019,
    '7890123456788',
    5
  ),
  (
    'The Adventure of the Mysterious Island',
    'Emily Chen',
    2020,
    '8901234567890',
    2
  ),
  (
    'The Secret of the Ancient Ruins',
    'Kevin White',
    2017,
    '9012345678902',
    3
  ),
  (
    'The Mystery of the Haunted House',
    'Sarah Lee',
    2014,
    '2345678901237',
    4
  ),
  (
    'The Quest for the Golden Treasure',
    'David Kim',
    2013,
    '3456789012341',
    5
  ),
  (
    'The Legend of the Phoenix Rising',
    'Emily Patel',
    2021,
    '4567890123451',
    2
  ),
  (
    'The Adventure of the Lost Treasure',
    'Michael Kim',
    2019,
    '5678901234567',
    3
  ),
  (
    'The Secret of the Ancient Temple',
    'Rebecca Hall',
    2018,
    '6789012345678',
    4
  ),
  (
    'The Mystery of the Missing Person',
    'James Brown',
    2016,
    '7890123456787',
    5
  );