from selenium import webdriver
from selenium.webdriver.common.by import By
import time
import pandas as pd
import tqdm
import os
import pickle
from dotenv import load_dotenv
import psycopg2

login_url = 'https://account.nicovideo.jp/login'

def login(driver, login_url):
    driver.get(login_url)
    print(driver.current_url)

    load_dotenv()
    email = os.getenv('EMAIL')
    password = os.getenv('PASSWORD')
    print(email, password)

    driver.find_element(By.NAME, 'mail_tel').send_keys(email)
    driver.find_element(By.NAME, 'password').send_keys(password)
    driver.find_element(By.ID, 'login__submit').click()

    print(driver.current_url)
    if driver.current_url != 'https://www.nicovideo.jp/':
        print('Login may be failed.')
        exit(1)

    # save cookies
    pickle.dump(driver.get_cookies(), open('cookies.pkl', 'wb'))
    return driver


try:
    print('Start')
    options = webdriver.ChromeOptions()
    options.add_argument('--headless')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')
    # options.add_argument('--disable-gpu')
    driver = webdriver.Remote(
                command_executor='http://selenium:4444/wd/hub',
                options=options
                )

    driver.implicitly_wait(10)

    # driver.get('https://www.nicovideo.jp/my/history/video')
    # print(driver.current_url)

    db_host = os.getenv('DB_HOST')
    db_user = os.getenv('DB_USER')
    db_password = os.getenv('DB_PASSWORD')
    db_name = os.getenv('DB_NAME')
    db_port = os.getenv('DB_PORT')
    # print('db_host:', db_host, 'db_user:', db_user, 'db_password:', db_password, 'db_name:', db_name, 'db_port:', db_port)
    # conn = psycopg2.connect(f'sslmode=disable dbname={db_name} user={db_user} password={db_password} host={db_host} port={db_port}')
    # cur = conn.cursor()
    # cur.execute('SELECT * FROM history')
    # print(cur.fetchall())

    if os.path.exists('cookies.pkl'):
        print('Load cookies')
        cookies = pickle.load(open('cookies.pkl', 'rb'))
        driver.get('https://www.nicovideo.jp/')
        for cookie in cookies:
            driver.add_cookie(cookie)
        driver.get('https://www.nicovideo.jp/my/history/video')
        if driver.current_url != 'https://www.nicovideo.jp/my/history/video':
            print('Login may be failed.')
            driver = login(driver, login_url)
    else:
        print('Login')

        driver = login(driver, login_url)

    print('Login success')
    driver.get('https://www.nicovideo.jp/my/history/video')
    print(driver.current_url)

    buttons = driver.find_elements(By.CLASS_NAME, 'ShowMoreList-more')
    while len(buttons) > 0:
        print('click')
        buttons[0].click()
        buttons = driver.find_elements(By.CLASS_NAME, 'ShowMoreList-more')

    driver.implicitly_wait(1)

    max_date = pd.to_datetime('1970-01-01')
    if os.path.exists('history.csv'):
        history_df = pd.read_csv('history.csv', parse_dates=['watch_date'])
        max_date = history_df['watch_date'].max()
    else:
        history_df = pd.DataFrame(columns=['url', 'title', 'watch_date'])
    print('max_date:', max_date)

    media_objects = driver.find_elements(By.CLASS_NAME, 'NC-VideoMediaObject')
    print(len(media_objects))
    new_history = []
    bar = tqdm.tqdm(total=len(media_objects))
    for media_object in media_objects:
        video_url = media_object.find_element(By.TAG_NAME, 'a').get_attribute('href')
        video_title = media_object.find_element(By.CLASS_NAME, 'NC-VideoMediaObject-title').text
        video_watch_date = media_object.find_element(By.CLASS_NAME, 'VideoWatchHistoryItemAfter-meta').find_element(By.XPATH, 'span').text[:-3]
        video_watch_date = pd.to_datetime(video_watch_date)
        # print(video_watch_date, max_date, (video_watch_date - max_date).total_seconds())
        print(video_url, video_title, video_watch_date)
        if (video_watch_date - max_date).total_seconds() <= 0:
            bar.close()
            break
        new_history.append([video_url, video_title, video_watch_date])
        # sql = 'INSERT INTO history (video_url, title, watch_date) VALUES (%s, %s, %s)'
        # cur.execute(sql, (video_url, video_title, video_watch_date))
        bar.update(1)

    bar.close()
    print('number of new_history:', len(new_history))
    history_df = pd.concat([pd.DataFrame(new_history, columns=['url', 'title', 'watch_date']), history_df])
    # df = pd.DataFrame(history_data, columns=['url', 'title', 'watch_date'])
    history_df.to_csv('history.csv', index=False)
finally:
    driver.quit()
