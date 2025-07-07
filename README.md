# triple-s

# Learning Objectives

- HTTP
- Basic networking concepts
- REST API
- Operations with CSV, XML formats
- Operations with os, io, path libraries

# Abstract

In this project, I developed a tool called `triple-s`, designed to implement a simplified version of S3 (Simple Storage Service) object storage. This tool will provide a REST API that allows clients to interact with the storage system, offering core functionalities such as creating and managing buckets, uploading, retrieving, and deleting files, as well as handling object metadata. The project aims to demonstrate key concepts of RESTful API design, basic networking, and data management, providing a practical foundation for understanding cloud storage solutions.

## Bucket Management

You will implement core functionalities to manage storage containers, known as "buckets" in the S3 paradigm. Buckets are fundamental units in the storage system where files (objects) are stored. This phase involves creating, listing, and deleting buckets via REST API endpoints, ensuring all responses conform to the XML format required by Amazon S3 specifications.

Note: Authentication and authorization are outside the scope of this project. All operations can be performed without credentials.

### API Endpoints for Bucket Management

#### 1. Create a Bucket:
- **HTTP Method:** `PUT`
- **Endpoint:** `/{BucketName}`
- **Request Body:** Empty. Additional parameters can be passed in the request headers.
- **Behavior:**
  - Validate the bucket name to ensure it meets Amazon S3 naming requirements (3-63 characters, only lowercase letters, numbers, hyphens, and periods).
  - Ensure the bucket name is unique across the entire storage system.
  - If the bucket name is valid and unique, create a new entry in the bucket metadata storage.
  - Return a `200 OK` status code and details of the created bucket, or an appropriate error message if the creation fails (e.g., `400 Bad Request` for invalid names, `409 Conflict` for duplicate names).

#### 2. List All Buckets:
- **HTTP Method:** `GET`
- **Endpoint:** `/`
- **Behavior:**
  - Read the bucket metadata from the storage (e.g., a CSV file).
  - Return an XML response containing a list of all matching buckets, including metadata like creation time, last modified time, etc.
  - Respond with a `200 OK` status code and the [XML list of buckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html#API_ListBuckets_Examples).

#### 3. Delete a Bucket:
- **HTTP Method:** `DELETE`
- **Endpoint:** `/{BucketName}`
- **Behavior:**
  - Check if the specified bucket exists by looking it up in the bucket metadata storage.
  - Ensure the bucket is empty (no objects are stored in it) before deletion.
  - If the bucket exists and is empty, remove it from the metadata storage.
  - Return a `204 No Content` status code if the deletion is successful, or an error message in XML format if the bucket does not exist or is not empty (e.g., `404 Not Found` for a non-existent bucket, 409 Conflict for a non-empty bucket).

### Ensuring Unique and Valid Bucket Names:

#### Bucket Naming Conventions:

- Bucket names must be unique across the system.
- Names should be between 3 and 63 characters long.
- Only lowercase letters, numbers, hyphens (`-`), and dots (`.`) are allowed.
- Must not be formatted as an IP address (e.g., 192.168.0.1).
- Must not begin or end with a hyphen and must not contain two consecutive periods or dashes.

#### Validation Implementation

- Use regular expressions to enforce naming rules.
- Check the uniqueness of a bucket name by reading the existing entries from the CSV metadata file.
- If the bucket name does not meet the rules, return a `400 Bad Request` response with a relevant error message.

### Example:

>##### Scenario 1: Bucket Creation
>- A client sends a `PUT` request to `/{BucketName}` with the name `my-bucket`.
>- The server checks for the validity and uniqueness of the bucket name, then creates an entry in the bucket metadata storage (e.g., `buckets.csv`).
>- The server responds with `200 OK` and the details of the new bucket or an appropriate error message if the creation fails.

>##### Scenario 2: Listing Buckets
>- A client sends a `GET` request to `/`.
>- The server reads the bucket metadata storage (e.g., `buckets.csv`) and returns an XML list of all bucket names and metadata.
>- The server responds with a `200 OK` status code.

>##### Scenario 3: Deleting a Bucket
>- A client sends a `DELETE` request to `/{BucketName}` for the bucket `my-bucket`.
>- The server checks if `my-bucket` exists and is empty.
>- If the conditions are met, the bucket is removed from the bucket metadata storage (e.g., `buckets.csv`).

## Object Operations

### API Endpoints for Object Operations

You will implement three main API endpoints to handle object operations:

#### 1. Upload a New Object:
- **HTTP Method:** `PUT`
- **Endpoint:** `/{BucketName}/{ObjectKey}`
- **Request Body:** Binary data of the object (file content).
- **Headers:**
  - `Content-Type`: The object's data type.
  - `Content-Length`: The length of the content in bytes.
- **Behavior:**
  - Verify if the specified bucket `{BucketName}` exists by reading from the bucket metadata.
  - Validate the object key `{ObjectKey}`.
  - Save the object content to a file in a directory named after the bucket (`data/{BucketName}/`).
  - Store object metadata in a CSV file (`data/{BucketName}/objects.csv`).
  - Respond with a 200 status code or an appropriate error message if the upload fails.

#### 2. Retrieve an Object:
- **HTTP Method:** `GET`
- **Endpoint:** `/{BucketName}/{ObjectKey}`
- **Behavior:**
  - Verify if the bucket `{BucketName}` exists.
  - Check if the object `{ObjectKey}` exists.
  - Return the object data or an error.

#### 3. Delete an Object:
- **HTTP Method:** `DELETE`
- **Endpoint:** `/{BucketName}/{ObjectKey}`
- **Behavior:**
  - Verify if the bucket and object exist.
  - Delete the object and update metadata.
  - Respond with a `204 No Content` status code or an appropriate error message.

### Example Scenarios

>- **Scenario 1: Object Upload**
   >  - A client sends a `PUT` request to `/photos/sunset.png` with the binary content of an image.
>  - The server checks if the `photos` bucket exists, validates the object key `sunset.png`, and saves the file to `data/photos/sunset.png`.
>  - The server updates `data/photos/objects.csv` with metadata for `sunset.png` and responds with `200 OK`.

>- **Scenario 2: Object Retrieval**
   >  - A client sends a `GET` request to `/photos/sunset.png`.
>  - The server checks if the `photos` bucket exists and if `sunset.png` exists within the bucket.
>  - The server reads the file from `data/photos/sunset.png` and returns the binary content with the `Content-Type` header set to `image/png`.

>- **Scenario 3: Object Deletion**
   >  - A client sends a `DELETE` request to `/photos/sunset.png`.
>  - The server checks if the `photos` bucket exists and if `sunset.png` exists within the bucket.
>  - The server deletes `data/photos/sunset.png` and removes the corresponding entry from `data/photos/objects.csv`.
>  - The server responds with `204 No Content`.

## Storing Metadata in a CSV File:

### Bucket CSV File Structure:
- Each line in the CSV file represents a bucket's metadata.
- The columns could include:
  - `Name`: The unique name of the bucket.
  - `CreationTime`: The timestamp when the bucket was created.
  - `LastModifiedTime`: The last time any modification was made to the bucket.
  - `Status`: Indicates whether the bucket is active or marked for deletion.

### Storing Object Metadata in a CSV File:

- **CSV File Structure for Object Metadata:**
  - Each bucket will have its own object metadata CSV file (e.g., `data/{bucket-name}/objects.csv`).
  - The columns could include:
    - `ObjectKey`: The unique key or identifier of the object within the bucket.
    - `Size`: The size of the object in bytes.
    - `ContentType`: The MIME type of the object (e.g., `image/png`, `application/pdf`).
    - `LastModified`: The timestamp when the object was last uploaded or modified.

## Usage

```
$ ./triple-s --help  
Simple Storage Service.

**Usage:**
    triple-s [-port <N>] [-dir <S>]  
    triple-s --help

**Options:**
- --help     Show this screen.
- --port N   Port number
- --dir S    Path to the directory
```

## Author

This project has been created by:

Savva Savostyanov
