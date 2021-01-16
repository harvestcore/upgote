# Log

> Log item

| Parameter   | Type   | I/O  | Description                                         |
| ----------- | ------ | :--: | --------------------------------------------------- |
| connotation | string | O | Connotation of the log message. [INFO, WARNING, ERROR] |
| datetime        | string/datetime | O | Date and time when the log message was recorded. |
| from        | string/uuid | O | From whom is this message. |
| to        | string/uuid | O | To whom this log message is addressed. |
| id         | string | O | Log message ID. |
| message   | string | O | Log message. |

## GET /log

Returns the logfile in plain text.

> Content-Type: text/plain

## POST /log

Returns log messages.

Payload (JSON):

| Parameter   | Type   | Optional | Description                                         |
| --------- | ---- | :----------------: | ------------------------------------------ |
| quantity  | int |        :heavy_check_mark:         | Number of items to be received. |

Response:

| Parameter   | Type   | Description                                         |
| --------- | ------------------------- | ----------- |
| status     | bool | Wether the operation was successful or not. |
| length     | int | Quantity of log messages returned. |
| items     | list[[Log](#Log)] | Log messages. |
