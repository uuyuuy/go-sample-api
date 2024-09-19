import random
from faker import Faker
import datetime

fake = Faker('ja_JP')

sql_template = "INSERT INTO users (last_name, first_name, tel, age) VALUES ('{last_name}', '{first_name}', '{tel}', {age});"

def generate_phone_number():
    return f"{random.randint(1000000000, 9999999999)}"

def generate_datetime():
    return fake.date_time_between(start_date='-10y', end_date='now').strftime('%Y-%m-%d %H:%M:%S.%f')[:-3]

data = []
for _ in range(50):
    last_name = fake.last_name()  # 日本語の苗字を生成
    first_name = fake.first_name()  # 日本語の名前を生成
    tel = generate_phone_number()  # ランダムな電話番号を生成
    age = random.randint(18, 80)  # 18歳から80歳までの年齢をランダムに生成
    created_at = generate_datetime()  # ランダムな作成日時を生成
    updated_at = generate_datetime()  # ランダムな更新日時を生成

    data.append(sql_template.format(last_name=last_name, first_name=first_name, tel=tel, age=age, created_at=created_at, updated_at=updated_at))

file_name = "dummy_data/dummy_data.sql"
with open(file_name, "w", encoding="utf-8") as f:
    for row in data:
        f.write(row + "\n")

print("DONE")
