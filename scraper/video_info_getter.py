import requests

ENDPOINT_URL = 'https://snapshot.search.nicovideo.jp/api/v2/snapshot/video/contents/search'
# payload = {
#     'q': '',
#     'targets': 'tags',
#     'filters[contentId][0]': 'sm42690205',
#     '_sort': '-viewCounter',
#     'fields': 'contentId,title,tags,viewCounter,startTime,thumbnailUrl'
# }

# response = requests.get(ENDPOINT_URL, params=payload)
# response.raise_for_status()
# res = response.json()
# for res in response.json()['data']:
#     print(res['contentId'])
#     print(res['title'])
#     print(res['tags'].split(' '))
#     print(res['viewCounter'])
#     print(res['startTime'])
#     print(res['thumbnailUrl'])
#     print('-----------------')
# print(response.json())

def upsert_video_basic_info(conn, cur, video_id):
    payload = {
        'q': '',
        'targets': 'tags',
        'filters[contentId][0]': video_id,
        '_sort': '-viewCounter',
        'fields': 'contentId,title,thumbnailUrl'
    }

    try:
        response = requests.get(ENDPOINT_URL, params=payload)
        response.raise_for_status()
        res = response.json()
        if len(res['data']) == 0:
            print('No data for', video_id)
            return
        for res in response.json()['data']:
            sql = 'INSERT INTO video_basic_info (video_id, title, thumbnail_url) VALUES (%s, %s, %s) ON CONFLICT (video_id) DO UPDATE SET title = %s, thumbnail_url = %s'
            cur.execute(sql, (res['contentId'], res['title'], res['thumbnailUrl'], res['title'], res['thumbnailUrl']))
            conn.commit()
    except Exception as e:
        print(f'Failed to get video info for {video_id}: {e}',)
        return

def insert_video_tag_info(conn, cur, video_id):
    payload = {
        'q': '',
        'targets': 'tags',
        'filters[contentId][0]': video_id,
        '_sort': '-viewCounter',
        'fields': 'contentId,tags'
    }

    try:
        response = requests.get(ENDPOINT_URL, params=payload)
        response.raise_for_status()
        res = response.json()
        if len(res['data']) == 0:
            print('No data for', video_id)
            return
        for res in response.json()['data']:
            tags = res['tags'].split(' ')
            for tag in tags:
                sql = 'INSERT INTO video_tag_info (video_id, tag) VALUES (%s, %s)'
                cur.execute(sql, (res['contentId'], tag))
                conn.commit()
    except Exception as e:
        print(f'Failed to get tag info for {video_id}: {e}',)
        return

def upsert_video_basic_info_and_video_tag_info(conn, cur, video_id):
    upsert_video_basic_info(conn, cur, video_id)
    insert_video_tag_info(conn, cur, video_id)
    return

