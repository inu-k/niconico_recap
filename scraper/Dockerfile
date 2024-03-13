FROM python:3.11

COPY . /app
WORKDIR /app
RUN apt-get update
RUN apt-get install -y cron
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

RUN chmod +x /app/run.sh
RUN chmod +x /app/scraper.py
RUN echo "*/30 * * * * bash /app/run.sh >>/var/tmp/log 2>>/var/tmp/log2\n* * * * * echo \"test\" >>/var/tmp/test" > /etc/cron.d/crontab
RUN chmod 0644 /etc/cron.d/crontab
RUN crontab /etc/cron.d/crontab
CMD ["cron", "-f"]
