# golangTask

# Goal

A project to read CSV files from a specified folder and save to Redis db.
Can get records from the database by ID

# More information

Golang version - 1.20.3.

Before running - need to change const folderPath in impl/main.go to real folder of .csv files.

If necessary, change the connection settings to the Redis db in impl/config/redis_conn.go.

For the test, you can use less time for the scheduled task (now 30 minutes) in impl/main.go.

After running the project it will run on port 1321

Endpoint to get record - http://localhost:1321/promotions/0000d8cf-25b0-4de7-8b80-ae4804b590bf
