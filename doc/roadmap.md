# HarvestCCode Roadmap

> This is a living document and it will be updated with more details in the future. Learn more about the different parts and modules of the project [here](architecture-workflows.md).

The development roadmap is the following:

## Project and architecture definition

Define the basis of the project, its architecture and create the first user stories.

Milestone:

- [Architecture design](https://github.com/harvestcore/HarvestCCode/milestone/2)

## Development

Main development of the project. This part is divided in different milestones that are cronologically ordered in the following list.

### Phase #1

This first phase aims to create the Updater, which is the part that fetches the data.

- [Milestone: Updater](https://github.com/harvestcore/HarvestCCode/milestone/7)
  - [HU - Fetch data from time to time](https://github.com/harvestcore/HarvestCCode/issues/15)

### Phase #2

In this phase the Core of the software will be created, along with the Handler, which will handle all the events between the different parts of the software.

- [Milestone: Core](https://github.com/harvestcore/HarvestCCode/milestone/6)
  - [HU - Store the fetched data](https://github.com/harvestcore/HarvestCCode/issues/16)
  - [HU - Create Updaters](https://github.com/harvestcore/HarvestCCode/issues/31)
  - [HU - Change Updater config](https://github.com/harvestcore/HarvestCCode/issues/17)
  - [HU - Stop Updaters](https://github.com/harvestcore/HarvestCCode/issues/32)
- [Milestone: Handler](https://github.com/harvestcore/HarvestCCode/milestone/9)

### Phase #3

In this phase the software will be capable of communicating with "the exterior" via a REST-API.

- [Milestone: REST-API](https://github.com/harvestcore/HarvestCCode/milestone/8)
  - [HU - Manage Updaters](https://github.com/harvestcore/HarvestCCode/issues/12)
  - [HU - Fetch data from the system](https://github.com/harvestcore/HarvestCCode/issues/13)

### Phase #0

The logging module will be developed in the early stages of the project, since all of the modules mentioned above make use of this one.

- [Milestone: Logs](https://github.com/harvestcore/HarvestCCode/milestone/10)
  - [HU - Logs](https://github.com/harvestcore/HarvestCCode/issues/14)
  - [HU - Logs download](https://github.com/harvestcore/HarvestCCode/issues/18)

## Container deployment

Deploy the project to containers.

More info TBD.

## Cloud deployment

Deploy the project to the cloud.

More info TBD.
