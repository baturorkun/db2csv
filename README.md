### Export table from database to csv file

Export table from database to csv file.


##### Supported Databases

- SQlite
- MySQL
- PosgreSQL

##### Install (For Linux and OS X)

> sh install.sh


##### Usage

At first, you must fill in the "db2csv.ini". The INI file includes application and database settings. 
You can find a detailed explanation below. If INI file is in the working directory and named "db2scv.ini", you can run the binary directly.

> ./bin/db2scv-darwin-amd64

If not, you must say where it is by "-conf" parameter.

> ./db2scv-darwin-amd64 -conf ~/batur/settings/my-dv2csv-1.ini

##### Setting INI File

```
[app]

Sql = "select * from attack_result  where vector = 1"
Filename = "result"
FilesPath = ./files
OverwriteFile = false
#false: Adds date suffix to the filename

[database]

Type = "sqlite3"
#Types: sqlite3, mysql, postgres

# Sqlite Settings

SqliteFile = "/Users/batur/go/db/demo_picus_db"

# Postgres setting

#User = postgres
#Password = postgres
#Host = localhost
#DbName = test
#Port = 5432
#SslMode = disable

# Mysql setting

#User = root
#Password = root
#Host = localhost
#DbName = test
#Port = 3306
```
[app]

* Sql : Raw SQL command
* Filename : CSV Filename
* FilesPath : Where csv files will save
* OverwriteFile : True > Always overwrite same file,  False > Adds date suffix to the filename, there will be a different file saved every run.

[database]

* Type : Database type (Values:  sqlite3, mysql, postgres)

The other parameters change depend on the database type. You should check example above.