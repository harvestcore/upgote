# Architecture workflows

> This document defines all the workflows and scenarios between all the different parts of this software.

> Please note that this is a living document and it will be updated along the development of HarvestCCode.

## API

The landing place for all user requests. It handles all those requests and communicates with the Handler by sending events with all the necessary information.

Possible scenarios:

- **User performs a request to HarvestCCode**: The API handles that request and issues an event to the Handler.

## Handler

Central part that handles all the events (and only events) sent and received by the rest of the parts. It contains different queues, one for each type of event.

Possible scenarios:

- **The handler receives an API event**: The handler forwards the event to the core. Issues back an event to the Updater if needed. Finally issues a log event.
- **The handler receives a Updater event**: The handler forwards the event to the core. Issues back an event to the Updater if needed. Finally issues a log event.

## Core

Main core of the software. It contains all the logic related to Updater creation and data storing.

Possible scenarios:

- **The core receives create Updater event**: The core checks the data from the event and creates a background process that will handle that data update. Finally issues a log event.
- **The core receives data storing event**: The core checks the data from the event and stores the data. Finally issues a log event.

## Updater

Background process that fetches the configured data.

Possible scenarios:

- **The Core creates an Updater**: The background process starts.
- **The Updater time interval is met**: It fetches the data. Finally issues an event to the Handler.

## Log

Logs all operations performed across all the system.

Possible scenarios:

- **The Log receives an event**: It logs all the related information.
