#!/bin/bash
echo "Running scraper"
source /etc/environment
cd /app
echo "cd"
ls -la 
/usr/local/bin/python scraper.py
