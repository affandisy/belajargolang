CREATE TABLE "Fields" (
  "FieldID" INT PRIMARY KEY,
  "FieldName" varchar,
  "FieldSize" varchar,
  "Location" varchar,
  "HourlyRate" decimal
);

CREATE TABLE "Customers" (
  "CustomerID" INT PRIMARY KEY,
  "FirstName" varchar,
  "LastName" varchar,
  "PhoneNumber" varchar
);

CREATE TABLE "Bookings" (
  "BookingID" INT PRIMARY KEY,
  "CustomerID" INT,
  "FieldID" INT,
  "BookingDate" date,
  "StartTime" timestamp,
  "EndTime" timestamp,
  "TotalAmount" float
);

CREATE TABLE "Payments" (
  "PaymentID" INT PRIMARY KEY,
  "BookingID" INT,
  "PaymentDate" "Float64",
  "PaymentMethod" "Float64"
);

ALTER TABLE "Bookings" ADD FOREIGN KEY ("CustomerID") REFERENCES "Customers" ("CustomerID");

ALTER TABLE "Bookings" ADD FOREIGN KEY ("FieldID") REFERENCES "Fields" ("FieldID");

ALTER TABLE "Payments" ADD FOREIGN KEY ("BookingID") REFERENCES "Bookings" ("BookingID");