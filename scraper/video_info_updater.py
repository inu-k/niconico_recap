import pandas as pd
import psycopg2
import os
import datetime
from zoneinfo import ZoneInfo

from video_info_getter import upsert_video_basic_info_and_video_tag_info

db_host = os.getenv('DB_HOST')
db_user = os.getenv('DB_USER')
db_password = os.getenv('DB_PASSWORD')
db_name = os.getenv('DB_NAME')
db_port = os.getenv('DB_PORT')
print('db_host:', db_host, 'db_user:', db_user, 'db_password:', db_password, 'db_name:', db_name, 'db_port:', db_port)
conn = psycopg2.connect(f'sslmode=disable dbname={db_name} user={db_user} password={db_password} host={db_host} port={db_port}')
cur = conn.cursor()

now_jst = datetime.datetime.now(ZoneInfo('Asia/Tokyo'))

senti_time = pd.to_datetime(f'{(now_jst - datetime.timedelta(days=1)).strftime("%Y-%m-%d")} 05:00:00')

cur.execute('SELECT video_id, watch_date FROM history WHERE watch_date >= %s', (senti_time,))

for row in cur.fetchall():
    video_id = row[0]
    watch_date = row[1]
    print(video_id, watch_date)
    upsert_video_basic_info_and_video_tag_info(conn, cur, video_id)
