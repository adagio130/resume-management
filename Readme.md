
# Resume Management System

A simple and scalable resume management system built with Go, Docker, and MySQL. This system allows users to create, manage, and retrieve resumes, including their experiences, skills, and education.

## Features

- CRUD operations for resumes.
- Association with experiences and education.
- Dockerized environment for easy setup and deployment.
- MySQL as the database backend.
- RESTful APIs built using the Gin framework.

---

## API Endpoints

### User APIs

| Method | Endpoint   | Description               | Request Body            | Response         |
|--------|------------|---------------------------|-------------------------|------------------|
| POST   | `/user`    | Create a new user         | JSON with user details  | Created user info|

**Example Request**:

```bash
curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{
    "name": "John Doe",
    "account": "john",
    "gender": "male",
    "location": "New York"
}'
```

**Example Response**:

```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "John Doe",
  "account": "john",
  "gender": "male",
  "location": "New York"
}
```

---

### Resume APIs

| Method | Endpoint               | Description                                  | Request Body                  | Response        |
|--------|------------------------|----------------------------------------------|-------------------------------|-----------------|
| GET    | `/resume/:id`          | Get a specific resume by its ID             | None                          | Resume details  |
| POST   | `/resume`              | Create a new resume                         | JSON with resume details      | Created resume  |
| PUT    | `/resume/:id`          | Update a specific resume by its ID          | JSON with updated details     | Updated resume  |
| DELETE | `/resume/:id`          | Delete a specific resume by its ID          | None                          | Success message |
| GET    | `/resumes/:id`         | Get all resumes for a specific user by ID   | None                          | List of resumes |

**Example: Fetch a Resume**

```bash
curl -X GET http://localhost:8080/resume/123e4567-e89b-12d3-a456-426614174000
```

**Example Response**:

```json
{
  "id": "d8279226-508f-4dc0-98eb-143786906a5b",
  "user_id": "362bb265-8976-49c5-8d4c-283bebce33ed",
  "title": "First Resume",
  "email": "johndoe@example.com",
  "phone": "123-456-7890",
  "experience": {
    "74635212-8cdc-45e9-ac1f-6097097be84c": {
      "id": "74635212-8cdc-45e9-ac1f-6097097be84c",
      "company": "Web Solutions",
      "position": "Developer",
      "is_present": false,
      "start_date": "2018-01",
      "end_date": "2019-12",
      "description": "Worked on various web development projects."
    },
    "e7880175-b0de-4f13-8ff3-7e2b229a7791": {
      "id": "e7880175-b0de-4f13-8ff3-7e2b229a7791",
      "company": "Tech Corp",
      "position": "Senior Developer",
      "is_present": true,
      "start_date": "2020-01",
      "end_date": "",
      "description": "Developing and maintaining web applications."
    }
  },
  "skills": [
    "Go",
    "JavaScript",
    "React"
  ],
  "education": {
    "6adb570c-1c58-4cb0-b8b8-e8ab931b3781": {
      "id": "6adb570c-1c58-4cb0-b8b8-e8ab931b3781",
      "school": "University of Technology",
      "major": "Computer Science",
      "degree": "Bachelor's",
      "start_date": "2015-09",
      "end_date": "2019-06"
    }
  }
}
```

**Example: Create a Resume**

```bash
curl -X POST http://localhost:8080/resume
```

**Example Response**:

```json
{
  "user_id": "362bb265-8976-49c5-8d4c-283bebce33ed",
  "title": "First Resume",
  "email": "johndoe@example.com",
  "phone": "123-456-7890",
  "experience": [
    {
      "company": "Tech Corp",
      "position": "Senior Developer",
      "is_present": true,
      "start_date": "2020-01",
      "end_date": "",
      "description": "Developing and maintaining web applications."
    },
    {
      "company": "Web Solutions",
      "position": "Developer",
      "is_present": false,
      "start_date": "2018-01",
      "end_date": "2019-12",
      "description": "Worked on various web development projects."
    }
  ],
  "skills": [
    "Go",
    "JavaScript",
    "React"
  ],
  "education": [
    {
      "school": "University of Technology",
      "major": "Computer Science",
      "degree": "Bachelor",
      "start_date": "2015-09",
      "end_date": "2019-06"
    }
  ]
}
```

**Example: Create a Resume**

```bash
curl -X PUT http://localhost:8080/resume
```

**Example Response**:

```json
{
  "user_id": "332efd8c-4be5-4af6-af3b-0cce177d8a2a",
  "title": "First Resume 123",
  "email": "johndoe123@example.com",
  "phone": "123-456-7890",
  "experience": [
    {
      "id": "038954bc-06b9-4e20-9276-0889b4a36d31",
      "company": "Web Solutions1",
      "position": "Developer1",
      "is_present": false,
      "start_date": "2018-01",
      "end_date": "2019-12",
      "description": "Worked on various web development projects."
    },
    {
      "id": "0f42110e-4c5a-41e0-86e0-1a5602a1b410",
      "company": "Tech Corp",
      "position": "Senior Developer",
      "is_present": true,
      "start_date": "2020-02",
      "end_date": "",
      "description": "Developing and maintaining web applications."
    }
  ],
  "skills": [
    "Go",
    "JavaScript",
    "React"
  ],
  "education": [
    {
      "id": "da82a08b-53a7-4aa7-951f-8f070c2d780d",
      "school": "University of Technology",
      "major": "Computer Science",
      "degree": "Bachelor",
      "start_date": "2015-09",
      "end_date": "2019-06"
    }
  ]
}
```

---

## Requirements

- Docker (v20+)
- Docker Compose (v2+)
- Go (v1.22+)
- MySQL (v8.0+)

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/resume-management.git
   cd resume-management
   ```

2. Set up the `.env` file:

   Create a `.env` file in the project root and provide the following environment variables:

   ```env
   DB_HOST=db
   DB_PORT=3306
   DB_USER=hank
   DB_PASSWORD=hanktest
   DB_NAME=resume
   APP_PORT=8888
   LOG_ENV=dev
   ```

3. Build and start the services:

   ```bash
   make build
   make run
   ```

4. Verify the services are running:

   ```bash
   docker ps
   ```

---

## Docker Commands

- **Build the project**:

  ```bash
  docker-compose build
  ```

- **Start the project**:

  ```bash
  docker-compose up -d
  ```

- **Stop the project**:

  ```bash
  docker-compose down
  ```

- **View logs**:

  ```bash
  docker-compose logs -f
  ```

---

## Development

1. Run the application locally without Docker:

   ```bash
   go run cmd/main.go
   ```

2. Run database migrations:

   ```bash
   make migrate
   ```

---

## Technologies Used

- **Backend**: Go (Gin framework)
- **Database**: MySQL
- **Containerization**: Docker, Docker Compose
- **Configuration**: Environment variables, YAML
- **Logging**: Configurable logging environment (`dev`, `test`, `prod`)

---

## Future Improvements

- Add user authentication and authorization.
- Create a web interface for managing resumes.
- Add support for exporting resumes as PDF.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
