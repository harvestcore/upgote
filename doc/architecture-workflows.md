# Architecture workflows

> This document defines all the different modules, workflows and scenarios between all the parts of this software.
> Please note that this is a living document and it will be updated along the development of HarvestCCode.

## API

The landing place for all user requests. It communicates with the core to handles all those requests and communicates with the core.

Possible scenarios:

- **User performs a request to HarvestCCode**: The API handles that request.

## Core

Main core of the software. It contains all the logic related to Updater creation and data storing.

Possible scenarios:

- **The core creates a new updater**.
- **The core updates an existing updater**.
- **The core removes an existing updater**.

## Updater

Background process that fetches the configured data.

Possible scenarios:

- **The Core creates an Updater**: The background process starts.
- **The Updater time interval is met**: The system fetches the data and then stores it.

## Log

Logs all operations performed across all the system.

Possible scenarios:

- **The system logs an event**: The system logs all the related information.
