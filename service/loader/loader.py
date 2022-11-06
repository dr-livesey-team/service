import json
import socket
import time

import pandas as pd

df = pd.read_csv("results.csv", delimiter='$')

offset = 0
size = 1000
timeout = 0.5

for i in range(offset, offset + size):
    info = {
        "request_root_identifier": int(df["Корневой ИД заявки"][i]),
        "opening_date": str(df["Дата создания заявки в формате Timezone"][i]),
        "closing_date": str(df["Дата закрытия"][i]),
        "district_name": str(df["Наименование района"][i]),
        "address": str(df["Адрес проблемы"][i]),
        "fault_name": str(df["Наименование дефекта"][i]),
        "management_company_name": str(df["Наименование управляющей компании"][i]),
        "service_organization_name": str(df["Наименование обслуживавшей организации (исполнителя)"][i]),
        "urgency_category_name": str(df["Наименование категории срочности: Аварийная, Обычная"][i]),
        "feedback": str(df["Отзыв"][i]),
        "group_id": int(df["ИД ситуации"][i])
    }

    buffer = json.dumps(info)

    request = {
        "func": "InsertRequest",
        "data": buffer
    }

    buffer = json.dumps(request)
    buffer += '\n'

    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect(("0.0.0.0", 50052))
    s.send(str.encode(buffer))
    s.close()
    time.sleep(timeout)

for i in range(offset, offset + size):
    if df["Корневой ИД заявки"][i] == df["ИД ситуации"][i] and int(df["Категории аномальности"][i]) != 0:
        info = {
            "id": int(df["ИД ситуации"][i]),
            "opening_date": str(df["Начало ситуации"][i]),
            "closing_date": str(df["Окончание ситуации"][i]),
            "district_name": str(df["Наименование района"][i]),
            "address": str(df["Адрес проблемы"][i]),
            "fault_name": str(df["Наименование дефекта"][i]),
            "management_company_name": str(df["Наименование управляющей компании"][i]),
            "service_organization_name": str(df["Наименование обслуживавшей организации (исполнителя)"][i]),
            "urgency_category_name": str(df["Наименование категории срочности: Аварийная, Обычная"][i]),
            "anomaly_category": str(df["Категории аномальности"][i]),
            "latitude": float(0),
            "longitude": float(0),
        }

        buffer = json.dumps(info)

        request = {
            "func": "InsertAnomaly",
            "data": buffer
        }

        buffer = json.dumps(request)
        buffer += '\n'

        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect(("0.0.0.0", 50052))
        s.send(str.encode(buffer))
        s.close()
        time.sleep(timeout)

for i in range(offset, offset + size):
    if df["Корневой ИД заявки"][i] == df["ИД ситуации"][i] and int(df["Категории аномальности"][i]) == 0:
        info = {
            "id": int(df["ИД ситуации"][i]),
            "opening_date": str(df["Начало ситуации"][i]),
            "closing_date": str(df["Окончание ситуации"][i]),
            "district_name": str(df["Наименование района"][i]),
            "address": str(df["Адрес проблемы"][i]),
            "fault_name": str(df["Наименование дефекта"][i]),
            "management_company_name": str(df["Наименование управляющей компании"][i]),
            "service_organization_name": str(df["Наименование обслуживавшей организации (исполнителя)"][i]),
            "urgency_category_name": str(df["Наименование категории срочности: Аварийная, Обычная"][i]),
        }

        buffer = json.dumps(info)

        request = {
            "func": "InsertNormal",
            "data": buffer
        }

        buffer = json.dumps(request)
        buffer += '\n'

        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect(("0.0.0.0", 50052))
        s.send(str.encode(buffer))
        s.close()
        time.sleep(timeout)
