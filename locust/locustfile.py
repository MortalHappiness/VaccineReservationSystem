from locust import HttpUser, between, task
import csv
import uuid
import random
import time
import json
users = []
hospitals = []
with open("/mnt/locust/data/user_10000.csv", "r", newline='') as csvfile:
    rows = csv.reader(csvfile)
    for row in rows:
        if row[0] == "ID":
            continue
        users.append(row)
        # ID,name,healthCardID,gender,birthDay,address,phone,vaccines
with open("/mnt/locust/data/hospitals_all_clean.csv", "r", newline='') as csvfile:
    rows = csv.reader(csvfile)
    for row in rows:
        if row[0] == "ID":
            continue
        hospitals.append(row)
        #ID,county,township,name,address,vaccineCnt




class WebsiteUser(HttpUser):
    wait_time = between(0.5, 10)
    user = None
    
    def on_start(self):
        self.user = random.choice(users)
        self.client.post("/api/session", json={
            "nationID": self.user[0],
            "healthCardID": self.user[2]
        })
    
    @task
    def index(self):
        self.client.get("/api/session")

    @task(2)
    def get_hospita(self):
        hospital = random.choice(hospitals)
        self.client.get("/api/hospitals", params={
            "county": hospital[1],
            "township": hospital[2]
        })

    
    @task(3)
    def get_hospital_and_post_reservation(self):
        hospital = random.choice(hospitals)
        # get hospital
        self.client.get("/api/hospitals", params={
            "county": hospital[1],
            "township": hospital[2]
        })
        
        # post reservation
        self.client.post(f"/api/reservations/users/{self.user[0]}", json={
            "id": str(uuid.uuid1()),
            "user": {
                "nationID": self.user[0]
            },
            "hospital": {
                "id": hospital[0],
                "county": hospital[1],
                "township": hospital[2]
            },
            "vaccineType": random.choice(list(json.loads(hospital[-1]).keys())),
            "completed": False,
            "date": int(time.time())
        })
