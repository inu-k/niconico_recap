!/bin/bash
psql -f ./setup.sh -U user -d history
echo "Database setup complete."