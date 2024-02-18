CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" text,
  "password" text
);

CREATE TABLE "wallets" (
  "id" serial PRIMARY KEY,
  "userId" int,
  "balance" numeric,
  "lastFunded" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "wallet_transaction" (
  "id" serial PRIMARY KEY,
  "walletId" int,
  "amount" serial,
  "type" text
);

CREATE TABLE "drivers" (
  "id" serial PRIMARY KEY,
  "email" text,
  "password" text,
  "lastLocation" point
);

CREATE TABLE "carDetails" (
  "id" serial PRIMARY KEY,
  "driverId" int,
  "make" text,
  "model" text,
  "plateNumber" text,
  "vin" text,
  "imageUrl" text
);

CREATE TABLE "trip" (
  "id" serial PRIMARY KEY,
  "userId" int,
  "driverId" int,
  "transactionId" int,
  "pickUpLocation" point,
  "destination" point,
  "currentTripLocation" point,
  "tripStartedAt" timestampz NOT NULL DEFAULT (now()),
  "tripEndedAt" timestampz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "wallet" ("userId");

ALTER TABLE "wallet" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id");

ALTER TABLE "wallet_transaction" ADD FOREIGN KEY ("walletId") REFERENCES "wallet" ("id");

ALTER TABLE "carDetails" ADD FOREIGN KEY ("driverId") REFERENCES "driver" ("id");

ALTER TABLE "trip" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id");

ALTER TABLE "trip" ADD FOREIGN KEY ("driverId") REFERENCES "driver" ("id");

ALTER TABLE "trip" ADD FOREIGN KEY ("transactionId") REFERENCES "trip" ("id");
