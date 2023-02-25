# go-todo
Simple To-Do List Rest API

## Models

### List

| Field | Description |
|-------|-------------|
|`id` | ToDo list id |
|`name`| ToDo list name |

### Task

| Field | Description |
|-------|-------------|
|`id` | Task list id |
|`title`| Task title |
|`status`| Task status |

## Endpoints

- `POST|GET /lists` - create and get todo lists
- `POST|GET|PUT /lists/:id/tasks` - create, update and get todo list tasks
