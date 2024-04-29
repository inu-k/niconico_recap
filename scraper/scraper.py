from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.common.exceptions import TimeoutException
import time
import pandas as pd
import tqdm
import os
import pickle
from dotenv import load_dotenv
import psycopg2

from video_info_getter import upsert_video_basic_info_and_video_tag_info

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
    if driver.current_url != 'https://www.nicovideo.jp':
        print('Login may be failed.')
        exit(1)

    # save cookies
    pickle.dump(driver.get_cookies(), open('cookies.pkl', 'wb'))
    return driver

def page_load_with_retry(driver, url, max_retry=2):
    retry = 0
    ret = None
    print('try to load:', url)
    while retry < max_retry:
        try:
            driver.get(url)
            ret = driver
            break
        except TimeoutException:
            retry += 1
            print('Retry:', retry)
            continue
    
    if ret is None:
        print('Failed to load:', url)
        raise TimeoutException('Failed to load:', url)

    print('success to load:', url)
    return ret


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
    print('Driver start')

    driver.implicitly_wait(60)
    driver.set_page_load_timeout(60)

    # driver.get('https://www.nicovideo.jp/my/history/video')
    # print(driver.current_url)

    db_host = os.getenv('DB_HOST')
    db_user = os.getenv('DB_USER')
    db_password = os.getenv('DB_PASSWORD')
    db_name = os.getenv('DB_NAME')
    db_port = os.getenv('DB_PORT')
    print('db_host:', db_host, 'db_user:', db_user, 'db_password:', db_password, 'db_name:', db_name, 'db_port:', db_port)
    conn = psycopg2.connect(f'sslmode=disable dbname={db_name} user={db_user} password={db_password} host={db_host} port={db_port}')
    cur = conn.cursor()

    if os.path.exists('cookies.pkl'):
        print('Loading cookies')
        cookies = pickle.load(open('cookies.pkl', 'rb'))
        driver = page_load_with_retry(driver, 'https://www.nicovideo.jp/ranking')
        for cookie in cookies:
            driver.add_cookie(cookie)
        print('Cookies loaded')
        # driver = page_load_with_retry(driver, 'https://www.nicovideo.jp/video_top')
        driver = page_load_with_retry(driver, 'https://www.nicovideo.jp/my/history/video')
        if driver.current_url != 'https://www.nicovideo.jp/my/history/video':
            print('Login may be failed.')
            driver = login(driver, login_url)
    else:
        print('Login')

        driver = login(driver, login_url)

    print('Login success')
    driver = page_load_with_retry(driver, 'https://www.nicovideo.jp/my/history/video')
    print(driver.current_url)

    buttons = driver.find_elements(By.CLASS_NAME, 'ShowMoreList-more')
    while len(buttons) > 0:
        print('click')
        buttons[0].click()
        buttons = driver.find_elements(By.CLASS_NAME, 'ShowMoreList-more')

    media_objects = driver.find_elements(By.CLASS_NAME, 'NC-VideoMediaObject')
    print(len(media_objects))
    new_history_num = 0
    bar = tqdm.tqdm(total=len(media_objects))
    for media_object in media_objects:
        video_url = media_object.find_element(By.TAG_NAME, 'a').get_attribute('href')
        video_title = media_object.find_element(By.CLASS_NAME, 'NC-VideoMediaObject-title').text
        video_watch_date = media_object.find_element(By.CLASS_NAME, 'VideoWatchHistoryItemAfter-meta').find_element(By.XPATH, 'span').text[:-3]
        video_watch_date = pd.to_datetime(video_watch_date)
        
        # print(video_watch_date, max_date, (video_watch_date - max_date).total_seconds())
        video_id = video_url.split('/')[-1]
        print(video_id, video_title, video_watch_date)

        sql = 'INSERT INTO history (video_id, watch_date) VALUES (%s, %s) ON CONFLICT (video_id, watch_date) DO NOTHING'
        cur.execute(sql, (video_id, video_watch_date))
        conn.commit()

        row_count = cur.rowcount
        if row_count == 0:
            break
        new_history_num += row_count
        print(new_history_num)
        # print(cur.statusmessage)
        # print(cur.query)
        upsert_video_basic_info_and_video_tag_info(conn, cur, video_id)
        bar.update(1)

    bar.close()
    print('number of new_history:', new_history_num)
finally:
    driver.quit()
    cur.close()
    conn.close()
