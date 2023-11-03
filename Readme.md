# Employee Management System on Google Cloud

This project showcases the implementation of an Employee Management System on Google Cloud Platform (GCP). It utilizes the Go programming language for creating APIs and Firestore for data storage. The project encompasses basic CRUD (Create, Read, Update, Delete) functionalities, which are made accessible through REST APIs. These APIs are deployed as serverless microservices using Google Cloud Functions.

## Project Overview - 
1. GCP Services: Google Cloud Functions, Firestore, IAM/Service Accounts

2. Programming Language: Go (Golang)

3. API Documentation: SwaggerUI

4. Data Storage: Firestore NoSQL database

## Key Functionalities
1. Employee Management: Enables the creation, retrieval, updating, and deletion of employee records.

2. Serverless Architecture: Leverages Google Cloud Functions for seamless and scalable deployment.

3. Interactive API Documentation: Provides comprehensive Swagger documentation for all API endpoints.
Seamless Integration with Google 

4. Cloud Firestore: Stores and manages employee data efficiently using Google Cloud Firestore collections.


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

- Check Google Cloud's [Identity and Access Management](https://cloud.google.com/iam/docs)  to understand more about Access Control


5. Swagger UI - REST API Documentation Tool
[Swagger Documenatation](https://swagger.io/docs/)



## Project Setup - not fully setup- do changes accordingly

### 1. Clone the Repository:

```bash
git clone https://github.com/Anagha-coder/Task3GCP.git
cd employee-management-system --configure folder name etc
```

### 2. Install Dependencies

```bash
# Install necessary Go packages and dependencies
go mod tidy
```

### 3. Environment Variables:

- Configure necessary environment variables for Firestore authentication and API endpoints.
  - Application Default Credentials Using Cloud SDK

### 4. Run Locally

```bash
go run main.go
```

## Configuring Google Cloud Environment

## Sevice Account on GCP

[Service accounts overview](https://cloud.google.com/iam/docs/service-account-overview)

- Create a New User Service account for your GCP Project

- Generate a Security key and
[Set up Application Default Credentials](https://cloud.google.com/docs/authentication/provide-credentials-adc) using [Cloud SDK](https://cloud.google.com/sdk/docs/install)

## Firestore Setup

[Create a collection of Firestore documents](https://cloud.google.com/firestore/docs/samples/firestore-query-collection-group-dataset#firestore_query_collection_group_dataset-go)



[Configure Security Rules for Firestore](https://cloud.google.com/firestore/docs/security/get-started)

## Set Up IAM

[IAM basic and predefined roles reference](https://cloud.google.com/iam/docs/understanding-roles)

## Cloud Functions 

[Basics Of Cloud Functions](https://cloud.google.com/functions#)

Deploy individual Cloud Functions for each API endpoint.

## Swagger

- Deploy Swagger UI on a web server to provide interactive API documentation.
- [Swagger Static Files](https://swagger.io/docs/open-source-tools/swagger-ui/usage/installation/) 




## Cloud Logging

[Setting Up Cloud Logging for Go](https://cloud.google.com/logging/docs/setup/go)



## Security Considerations

- Ensure that appropriate IAM roles and permissions are set for Cloud Functions to access Firestore.

- Check Security Rules of Firestore go get access to the database.



## Contributing

- Contributions are welcome! Feel free to open issues or submit pull requests.
