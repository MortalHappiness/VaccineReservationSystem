from area_data import areaData
import csv
import random
choices = []
for county in areaData:
    for township in areaData[county]:
        if len(township) != 4 and township != "烏坵鄉":
            choices.append([county, township])


with open("hospital_search.csv", "w") as csvfile:
    writer = csv.writer(csvfile)
    writer.writerow(["county", "township"])
    writer.writerows(random.choices(choices, k=100000))
