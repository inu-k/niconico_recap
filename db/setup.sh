#!/bin/bash
ls /tmp
psql -f /tmp/setup.sql -U user -d db_history
echo "Database setup complete."