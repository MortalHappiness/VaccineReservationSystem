import csv
from area_data import areaData
data = []
# 開啟 CSV 檔案
# with open('全台醫療院所清單.csv', newline='') as csvfile:

#     # 讀取 CSV 檔案內容
#     rows = csv.reader(csvfile)

#     # 以迴圈輸出每一列
#     i = 0
#     for row in rows:
#         name = row[1]
#         address = row[2]
#         county = address[:3]
#         township = address[3:6]
#         if township[-1] not in ["區", "市", "鄉", "鎮"]:
#             township = township[:-1]
#         data.append([i, county, township, name, address, ""])
#         i += 1

# with open("hospitals_all.csv", "w", newline='') as csvfile:
#     writer = csv.writer(csvfile)
#     writer.writerow(["ID", "county", "township",
#                     "name", "address", "vaccineCnt"])
#     writer.writerows(data)

data = []
with open("hospitals_all.csv", "r", newline='') as csvfile:
    rows = csv.reader(csvfile)
    for row in rows:
        if row[1] in areaData and row[2] in areaData[row[1]]:
            row[-1] = '{"BNT": 500, "AZ": 500}'
            data.append(row)

with open("hospitals_all_clean.csv", "w", newline='') as csvfile:
    writer = csv.writer(csvfile)
    writer.writerow(["ID", "county", "township",
                    "name", "address", "vaccineCnt"])
    writer.writerows(data)
