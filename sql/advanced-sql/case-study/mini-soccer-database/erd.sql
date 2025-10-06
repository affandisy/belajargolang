Table Fields {
  FieldID INT [primary key]
  FieldName varchar
  FieldSize varchar
  Location varchar
  HourlyRate decimal
}

Table Customers {
  CustomerID INT [pk]
  FirstName varchar
  LastName varchar
  PhoneNumber varchar
}

Table Bookings {
  BookingID INT [pk]
  CustomerID INT [ref: > Customers.CustomerID]
  FieldID INT [ref: > Fields.FieldID]
  BookingDate date
  StartTime timestamp
  EndTime timestamp
  TotalAmount float
}

Table Payments {
  PaymentID INT [pk]
  BookingID INT [ref: > Bookings.BookingID]
  PaymentDate Float64
  PaymentMethod Float64
}