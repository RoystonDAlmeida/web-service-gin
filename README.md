# Simple Album API with Gin

Welcome to the Simple Album API! This API is built using the Gin web framework in Go and provides endpoints to manage albums. You can create new albums, retrieve all albums, and fetch a specific album by its ID.

## Features

- **Create an Album**: `POST /albums`
- **Get All Albums**: `GET /albums`
- **Get Album by ID**: `GET /albums/:id`

## Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or later)
- Git (for cloning the repository)

### Installation

1. Clone the repository:
	```bash
	git clone https://github.com/RoystonDAlmeida/web-service-gin.git
	cd web-service-gin
	```

2. Install the dependencies:

	```bash
	go mod tidy
	```

3. Run the application:

	```bash
	go run main.go
	```

The server will start on `http://localhost:8080`.

## API Endpoints

### 1. Create an Album

**Endpoint**: `POST /albums`

**Request Body**:
	```json
	{
		"title": "Album Title",
		"artist": "Artist Name",
		"year": 2021
	}
	```

**Response**:
- **201 Created**
	```json
	{
		"id": 1,
		"title": "Album Title",
		"artist": "Artist Name",
		"year": 2021
	}
	```

### 2. Get All Albums

**Endpoint**: `GET /albums`

**Response**:
- **200 OK**
	```json
	[
		{
			"id": 1,
			"title": "Album Title",
			"artist": "Artist Name",
			"year": 2021
		},
		{
			"id": 2,
			"title": "Another Album",
			"artist": "Another Artist",
			"year": 2020
		}
	]
	```


### 3. Get Album by ID

**Endpoint**: `GET /albums/:id`

**Response**:
- **200 OK**
	```json
	{
		"id": 1,
		"title": "Album Title",
		"artist": "Artist Name",
		"year": 2021
	}
	```

- **404 Not Found**

	```json
	{
		"error": "Album not found"
	}

## Example Usage

You can use tools like [Postman](https://www.postman.com/) or [cURL](https://curl.se/) to test the API endpoints.

### Create an Album Example with cURL

	```json
		curl -X POST http://localhost:8080/albums
		-H 'Content-Type: application/json'
		-d '{"title":"My First Album","artist":"John Doe","year":2022}'
	```

### Get All Albums Example with cURL

	```json
		curl http://localhost:8080/albums
	```

### Get Album by ID Example with cURL

	```json
		curl http://localhost:8080/albums/1
	```

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Happy coding! 🚀
