FROM python:3.11

COPY . /app
WORKDIR /app
RUN apt-get update
RUN apt-get install -y cron
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

RUN env > /etc/environment
RUN chmod +x /app/run.sh
RUN chmod +x /app/scraper.py
RUN echo "*/30 * * * * bash /app/run.sh >>/var/tmp/log 2>>/var/tmp/log2\n30 5 * * * bash /app/run_updater.sh >>/var/tmp/log_updater 2>>/var/tmp/log2_updater" > /etc/cron.d/crontab
RUN chmod 0644 /etc/cron.d/crontab
RUN crontab /etc/cron.d/crontab
RUN chmod +x /app/entrypoint.sh
ENTRYPOINT ["bash", "/app/entrypoint.sh"]
