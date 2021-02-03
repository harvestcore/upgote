import time
import warnings
from locust import HttpUser, task, between

warnings.filterwarnings("ignore")

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def log(self):
        self.client.get("/api/log", verify=False)

    @task
    def log_with_params(self):
        self.client.post("/api/log", verify=False, json={
            "quantity": 1
        })

    @task
    def log_data(self):
        self.client.post("/api/data", verify=False, json={
            "database": "log"
        })

    @task
    def log_data_with_quantity(self):
        self.client.post("/api/data", verify=False, json={
            "database": "log",
            "quantity": 50
        })


