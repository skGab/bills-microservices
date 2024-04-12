- Application requiriments
  I'll start building a monolitch service and migrate to a microsservice at the end.

// EXTRACT SOME FUNCTIONAL REQUIRIMENTS FROM THE ORIGNAL EXCEL
// AFTER THE IMPLEMENTATION OF THEM, THINK IN MORE FUNCTIONALITYS

// GATWAY-SERVICE
[] - Create the gateway route
[] - Route incoming requests from client to appropriate service

//DATABASE SERVICE (gRPC)
[x] - Get Bills
[x] - Create Bill (TESTED)
[x] - Update Bill (after the first info about the bills)
[x] - Delete Bill
[x] - Delete Bills

//DOMAIN BUSINESS SERVICE
[] - Show the total amount spent on bills
[] - Show the total amout to pay of each user
[] - Show the individual amout of each bill for all users
[] - User can choose how many people will pay the bills
[] - The user can apply a discount on the total of his part of the bill (after the first info about the bill)

[] - Automatic send the value of each account to the inter account (after the user confirm that's everything fine)

// USER SERVICE (gRPC)
[] - User registration
[] - User login

- Required technologies

[] - Docker
[] - RabbitMQ
[] - gRPC

- Required Methodologies

[x] - Domain-driven design
[] - Test-driven Development
[] - Continuous Integration
