## Features

- Upload images through a REST API (`/upload/screenshot`).
- Customizable API key for secure uploads.
- Rate limiting to control upload frequency.
- Supports dynamic image file names for security.
- Docker container support for easy deployment.

## Prerequisites

Before running the project, make sure you have:

- [Go](https://golang.org/dl/) installed (version 1.23.0 or higher).
- [Docker](https://www.docker.com/products/docker-desktop) installed (for running in a container).
- A `.env` file with the necessary environment variables.

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/mundaelol/ShareX-Screenshot-Uploader.git
    cd ShareX-Screenshot-Uploader
    ```

2. Create a .env file in the root directory and add the following:
    ```env
    UPLOAD_API_KEY=YOUR_PRIVATE_API_KEY
    ```
    Replace `YOUR_PRIVATE_API_KEY` with your actual API key.

3. Install dependencies and build the project:
    ```bash
    go mod download
    ```

4. Run the server:
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

5. To use Docker for containerized deployment, build and run the container:
    ```bash
    docker-compose up --build
    ```

## API Endpoints
- `POST /upload/screenshot`: Upload a screenshot to the server. Requires a valid `UPLOAD_API_KEY` in the `.env` file.

- `GET /images/{filename}`: Retrieve an uploaded image by its filename.

## Docker Deployment
You can also run the project in a Docker container by using the included `docker-compose.yml` file. The configuration will build the Go application inside a container and expose port `8080`.

## File Upload Process
The application limits file size to 500 MB for each upload. The uploaded file names are randomized for security and uniqueness.

- File names are sanitized and randomized to avoid conflicts and ensure unique names.
- The `/images/` directory is used to store the uploaded files.

## Configuration
You can customize the following settings in the `utils` package:

- `UploadFileSize`: Maximum allowed file size (500 MB by default).
- `FileNameLength`: Length of the randomized file name (12 characters by default).
- `RateLimit`: Maximum allowed requests per second (10 by default).
- `Directory`: Directory where uploaded files are stored (default: `./images/`).

## Rate Limiting
The server uses rate limiting to control how many requests can be made per second. The default is 10 requests per second. If the rate limit is exceeded, requests will be delayed until allowed again.

## Contributing
Feel free to fork the repository and submit pull requests with improvements or bug fixes.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/mundaelol/ShareX-Screenshot-Uploader/blob/main/LICENSE) file for details.
