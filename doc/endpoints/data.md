# Data

## POST /data

Returns data from the given database.

| Parameter   | Type  | Optional | Description                                         |
| --------- | ---- | :----------------: | ------------------------------------------ |
| database  | string |        :x: | Database from which to extract the elements.  |
| quantity  | int |        :heavy_check_mark: | Number of items to be received. |

Response:

| Parameter   | Type   | Description                                         |
| --------- | ------------------------- | ----------- |
| status     | bool | Wether the operation was successful or not. |
| length     | int | Quantity of items returned. |
| items     | list[Item] | Items. |

Response codes:

| Code | Scenario   |
| ---- | -------- |
| 200  | The request was successful. |
| 422  | Some required fields are incorrect. |

## DELETE /data

Removes the given database.

Payload (JSON):

| Parameter   | Type | Optional  | Description                                         |
| --------- | ---- | :----------------: | ------------------------------------------ |
| database  | int |        :x:         | Database to be removed. |
| force  | bool |        :x:         | Force the database removal. |

Response codes:

| Code | Scenario   |
| ---- | -------- |
| 200  | The request was successful. |
| 422  | Some required fields are incorrect. |
