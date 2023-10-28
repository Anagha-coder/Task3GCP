# Employee Management System on Google Cloud

This project showcases the implementation of an Employee Management System on Google Cloud Platform (GCP). It utilizes the Go programming language for creating APIs and Firestore for data storage. The project encompasses basic CRUD (Create, Read, Update, Delete) functionalities, which are made accessible through REST APIs. These APIs are deployed as serverless microservices using Google Cloud Functions.

## Project Overview - check id somethis is missing
1. GCP Services: Google Cloud Functions, Firestore, IAM/Service Accounts
2. Programming Language: Go (Golang)
3. API Documentation: Swagger
4. Data Storage: Firestore NoSQL database

## Features -- change this
1. CRUD Operations: Create, Read, Update, and Delete employee records.
2. Serverless Deployment: Utilizes Google Cloud Functions for serverless execution.
3. Swagger Documentation: API endpoints are documented using Swagger definitions.
4. Firestore Integration: Employee data is stored in Google Cloud Firestore collections.

## Getting Started

## Prerequisites

* Refer the links mentioned in the documentation for a detailed overview


1. Go installed on your local machine.
[Golang Documentation](https://go.dev/doc/)

2. A Google Cloud Platform account with a project created.
[Google Cloud Platform](https://cloud.google.com/)

3. Google Cloud SDK installed and configured.
[Install the gcloud CLI](https://cloud.google.com/sdk/docs/install)

4. Basic understanding of Google Cloud Functions and Firestore
- [Google Cloud Functions](https://cloud.google.com/functions#)
- [Google Firestore](https://cloud.google.com/firestore/docs/)

- Check Google Cloud's Identity and Access Management to understand more about Access Control
[Identity and Access Management | IAM](https://cloud.google.com/iam/docs)

5. Swagger UI - REST API Documentation Tool
[Swagger Documenatation](https://swagger.io/docs/)

6. [Terraform by HashiCorp](https://www.terraform.io/)
- [Teraform Documentation](https://developer.hashicorp.com/terraform?product_intent=terraform)
- [Terraform on Google Cloud documentation](https://cloud.google.com/docs/terraform)




## Project Setup - not fully setup- do changes accordingly

### 1. Clone the Repository:

```bash
git clone <https://github.com/Anagha-coder/Task3GCP.git>
cd employee-management-system --configure folder name etc
```

### 2. Install Dependencies

```bash
# Install necessary Go packages and dependencies
go mod tidy
```

### 3. Environment Variables:

- Configure necessary environment variables for Firestore authentication and API endpoints.

### 4. Run Locally

```bash
go run main.go
```

## Deployment to Google Cloud - do chnages - untouched
Firestore Setup:

Create a Firestore database on Google Cloud Console.
Configure necessary security rules for Firestore.
Cloud Functions Deployment:

Deploy individual Cloud Functions for each API endpoint using the Google Cloud Console or Terraform.
Swagger Documentation:

Deploy Swagger UI on a web server to provide interactive API documentation.
API Documentation
Swagger documentation is available at <swagger-url> after deployment.
Explore and test the API endpoints using Swagger's interactive interface.

## more sections to come in between
- keep adding as you build project



## Security Considerations

- Ensure that appropriate IAM roles and permissions are set for Cloud Functions to access Firestore.

- Implement secure coding practices to prevent common vulnerabilities like SQL injection and cross-site scripting. -change this

## Contributing

- Contributions are welcome! Feel free to open issues or submit pull requests.
