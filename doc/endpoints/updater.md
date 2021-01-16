# Updater

> Updater item

| Parameter   | Type   | I/O  | Description                                         |
| ----------- | ------ | :--: | --------------------------------------------------- |
| schema | string | I/O | Schema to match the fetched data against. |
| interval        | int | I/O | Time interval between data fetch. |
| source        | string | I/O | Source from where to fetch the data. |
| method        | string | I/O | Method used to fetch the data. [GET, POST] |
| id         | string/uuid | I/O | Updater ID. |
| requestBody   | dict or null | I/O | Request body, if existant. |
| timeout   | int | I/O | Maximun request timeout. |
| collection   | string | I/O | Collection or database where the data is stored. |

## GET /updater

Returns all the existant updaters.

Response:

| Parameter   | Type   | Description                                         |
| --------- | ------------------------- | ----------- |
| status     | bool | Wether the operation was successful or not. |
| length     | int | Quantity of items returned. |
| items     | list[[Updater](#Updater)] | Items. |

## POST /updater

Creates a new updater.

Payload (JSON): See [Updater](#Updater).

Response:

| Parameter   | Type   | Description                                         |
| --------- | ------------------------- | ----------- |
| status     | bool | Wether the operation was successful or not. |
| id     | string/uuid | ID of the created updater. |

## PUT /updater

Updates an existant updater.

Payload (JSON): See [Updater](#Updater).

Response:

| Parameter   | Type   | Description |
| --------- | ------------------------- | ----------- |
| status     | bool | Wether the operation was successful or not. |

## DELETE /data

Removes the given updater.

Payload (JSON):

| Parameter   | Type | Optional  | Description                                         |
| --------- | ---- | :----------------: | ------------------------------------------ |
| id  | string/uuid |        :x:         | Updater to be stopped. |
| force  | bool |        :x:         | Force the updater removal. |

Response:

| Parameter   | Type   | Description |
| --------- | ------------------------- | ----------- |
| status     | bool | Wether the operation was successful or not. |

## POST /updater/start

Stars an existant updater.

Payload (JSON):

| Parameter   | Type | Optional  | Description       |
| --------- | ---- | :----------------: | --------------- |
| id  | string/uuid |        :x:         | Updater to be started. |

Response:

| Parameter   | Type   | Description |
| --------- | ------------------------- | ----------- |
| status     | bool | Wether the operation was successful or not. |

## POST /updater/stop

Stops an existant updater.

Payload (JSON):

| Parameter   | Type | Optional  | Description       |
| --------- | ---- | :----------------: | --------------- |
| id  | string/uuid |        :x:         | Updater to be stopped. |

Response:

| Parameter   | Type   | Description |
| --------- | ------------------------- | ----------- |
| status     | bool | Wether the operation was successful or not. |
