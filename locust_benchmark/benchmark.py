import time
import warnings
from locust import HttpUser, task, between

warnings.filterwarnings("ignore")

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def status(self):
        self.client.get("/api/status", verify=False)

    @task
    def healthcheck(self):
        self.client.get("/api/healthcheck", verify=False)

    @task
    def updater(self):
        self.client.get("/api/updater", verify=False)

    @task
    def updater_wrong_data(self):
        self.client.post("/api/updater", verify=False, json={})

    @task
    def updater_correct_data(self):
        self.client.post("/api/updater", verify=False, json={
            "database": "locust",
            "schema": {
                "my": "schema"
            },
            "interval": 60,
            "source": "https://ipinfo.io/json",
            "method": "GET",
            "timeout": 30
        })

    @task
    def data(self):
        self.client.post("/api/data", verify=False, json={
            "database": "locust"
        })

    @task
    def data_with_data(self):
        self.client.post("/api/data", verify=False, json={
            "database": "locust",
            "quantity": 50
        })
