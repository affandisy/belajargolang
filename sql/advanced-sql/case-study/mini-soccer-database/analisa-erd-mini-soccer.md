## ERD Title: Mini Soccer Rental System

Entity: Fields
## Attributes:
- FieldID INT PRIMARY KEY
- FieldName VARCHAR(100)
- FieldSize VARCHAR(100)
- Location VARCHAR(50)
- HourlyRate DECIMAL

Entity: Customers
## Attributes:
- CustomerID INT PRIMARY KEY
- FirstName VARCHAR(100)
- LastName VARCHAR(100)
- PhoneNumber VARCHAR(100)

Entity: Bookings
## Attributes: 
- BookingID INT PRIMARY KEY
- CustomerID INT
- FieldINT INT
- BookingDate DATE
- StartTime TIMESTAMP
- EndTime TIMESTAMP
- TotalAmount FLOAT

Entity: Payments
## Attributes: 
- PaymentID INT PRIMARY KEY
- BookingID INT
- PaymentDate DATE
- PaymentMethod FLOAT64

Description: Satu Customers bisa booking banyak, tapi satu booking hanya bisa satu customer
Table_Name to Table_Name: Customers - Bookings
Type: One to Many

Description: One field can place many orders, but each order is linked to only one field
Table_Name to Table_Name: Fields - Bookings
Type: One to Many

Description: One Booking have only one payment
Table_Name to Table_Name: Payments - Bookings
Type: One to One

Integrity Constraints:
- TotalAmount cannot be 0
- We can select BookingDate manually
- PaymentMethod only accept 3 values: Credit Card, Bank Transfer, Cash
- Customer's Email and PhoneNumber is unique
- All numeric type cannot below 0 and positive only



Additional Notes:


