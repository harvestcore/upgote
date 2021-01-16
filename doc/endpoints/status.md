# Status and healthcheck

## GET /status

Returns the actual state of the backend.

Response:

| Parameter | Type | Description                       |
| --------- | ---- | --------------------------------- |
| status    | bool | If the backend is working or not. |
| updaters  | int | Number of existing (running or not) updaters. |

## GET /healthcheck

Returns the actual state of the backend in a simplified JSON.

Response:

| Parameter | Type | Description                       |
| --------- | ---- | --------------------------------- |
| status    | bool | If the backend is working or not. |
