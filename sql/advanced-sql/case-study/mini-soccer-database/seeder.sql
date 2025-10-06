-- Customers Seeders
INSERT INTO Customers (FirstName, LastName, Email, PhoneNumber) VALUES
mini_soccer_sl-# ('Ahmad', 'Setiawan', 'ahmad.set@example.com', '+628123456789'),
mini_soccer_sl-# ('Siti', 'Fatimah', 'siti.fat@example.com', '+628987654321'),
mini_soccer_sl-# ('Budi', 'Harsono', 'budi.h@example.com', '+628456789012');

-- Fields Seeders
INSERT INTO Fields (FieldName, FieldSize, Location, HourlyRate) VALUES
mini_soccer_sl-# ('Ronaldo Arena', '5v5', 'Jakarta', 50.00),
mini_soccer_sl-# ('Messi Park', '7v7', 'Bali', 70.00),
mini_soccer_sl-# ('Neymar Ground', '5v5', 'Surabaya', 45.00);

-- Bookings Seeders
INSERT INTO Bookings (CustomerID, FieldID, BookingDate, StartTime, EndTime) VALUES
mini_soccer_sl-# (1, 2, '2023-08-10', '16:00:00', '18:00:00'),
mini_soccer_sl-# (3, 1, '2023-08-11', '10:00:00', '12:00:00'),
mini_soccer_sl-# (2, 3, '2023-08-12', '14:00:00', '15:00:00');

-- Payments Seeders
INSERT INTO Payments (BookingID, PaymentDate, PaymentAmount, PaymentMethod) VALUES
mini_soccer_sl-# (2, '2023-08-09', 140.00, 'Credit Card'),
mini_soccer_sl-# (3, '2023-08-10', 100.00, 'Bank Transfer'),
mini_soccer_sl-# (4, '2023-08-11', 45.00, 'Cash');