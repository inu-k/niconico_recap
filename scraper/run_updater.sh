#!/bin/bash
echo "Running updater"
source /etc/environment
cd /app
/usr/local/bin/python video_info_updater.py
