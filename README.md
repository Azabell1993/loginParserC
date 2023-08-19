# loginParserC

C Parser and MySQL Login Integration Example
=============

This repository contains an example implementation of a C parser and MySQL login integration. The program demonstrates how to connect to a MySQL database, authenticate users, and perform queries using C programming language.

Prerequisites
-------------
Before you begin, ensure you have the following requirements in place:

MySQL Server installed and running.
GCC compiler for compiling C code.
MySQL development libraries for linking the MySQL client.

Setup
-------------
Clone the repository to your local machine

bash
-------------
git clone https://github.com/Azabell1993/loginParserC.git
cd ClangLoginParser
make

Usage
-------------
Update the connect.json file with your MySQL database connection details:

json
-------------
```  
{
  "db_host": "localhost",
  "db_user": "your_db_user",
  "db_pass": "your_db_password",
  "db_name": "your_db_name"
}
```  


Run the compiled executable:
-------------

./loginSecurithLib

Follow the prompts to enter a username and password for authentication.

Structure
-------------
connector.h: Header file containing JSON parsing and database connection functions.
login.c: Main program file containing the login and MySQL interaction logic.
connect.json: JSON configuration file for MySQL connection details.
Makefile: Makefile for compiling the program.

Notes
-------------
The program establishes a connection to the MySQL database using the details provided in connect.json.
It performs user authentication and executes sample queries using MySQL.
The connector.h header provides JSON parsing utilities to read the connect.json file.
Contributing
Contributions are welcome! If you find any issues or want to enhance the example, feel free to create a pull request.

License
-------------
This project is licensed under the MIT License - see the LICENSE file for details.

Build Example
-------------
```
azabell@azabellUbuntu:~/Desktop/workspace/ClangLoginParser$ ./loginSecurithLib 
=======================================
db_host localhost
db_user test_user
db_pass 1234
db_name test_db

Connect Sucess!!
아이디 :azabell1993
비밀번호 :azabell1
 !! RUN QUERY !!
=====================================================================================
	 SELECT DISTINCT                            IFNULL(                             CAST(                                   (                                       SELECT CASE                                         WHEN COUNT(*) > 0 THEN 1                                            ELSE 0                                      END                                     FROM (                                          SELECT A1.USER_INFONUM                                          FROM ADMIN AS A1                                            INNER JOIN ADMIN_SECURITY AS S ON ADMIN.USERNAME = S.USERNAME                                           WHERE ADMIN.USER_INFONUM = S.USER_SECURITY_INFO                                          AND   ADMIN.USERNAME = S.USERNAME                                        ) A                                 ) AS CHAR                               ), 0                            ) AS LOGIN                      FROM ADMIN                  WHERE                     USERNAME = 'azabell1993'                     AND PASSWORD = 'azabell1'; 
=====================================================================================
 	 query_stat_chk : 0 
 !! RUN QUERY !!
=====================================================================================
	 SELECT DISTINCT ad.USERNAME                            FROM ADMIN ad                                                   WHERE USERNAME 		= 'azabell1993'                           AND PASSWORD 	= 'azabell1'; 
=====================================================================================
 	 query_stat : 0
  !! RUN QUERY !!
=====================================================================================
	 select now() 
=====================================================================================
 	 query_stat_date : 0
로그인 실패!!
azabell@azabellUbuntu:~/Desktop/workspace/ClangLoginParser$ ./loginSecurithLib 
=======================================
db_host localhost
db_user test_user
db_pass 1234
db_name test_db

Connect Sucess!!
아이디 :azabell
비밀번호 :azabell!@
 !! RUN QUERY !!
=====================================================================================
	 SELECT DISTINCT                            IFNULL(                             CAST(                                   (                                       SELECT CASE                                         WHEN COUNT(*) > 0 THEN 1                                            ELSE 0                                      END                                     FROM (                                          SELECT A1.USER_INFONUM                                          FROM ADMIN AS A1                                            INNER JOIN ADMIN_SECURITY AS S ON ADMIN.USERNAME = S.USERNAME                                           WHERE ADMIN.USER_INFONUM = S.USER_SECURITY_INFO                                          AND   ADMIN.USERNAME = S.USERNAME                                        ) A                                 ) AS CHAR                               ), 0                            ) AS LOGIN                      FROM ADMIN                  WHERE                     USERNAME = 'azabell'                     AND PASSWORD = 'azabell!@'; 
=====================================================================================
 	 query_stat_chk : 0 
 !! RUN QUERY !!
=====================================================================================
	 SELECT DISTINCT ad.USERNAME                            FROM ADMIN ad                                                   WHERE USERNAME 		= 'azabell'                           AND PASSWORD 	= 'azabell!@'; 
=====================================================================================
 	 query_stat : 0
  !! RUN QUERY !!
=====================================================================================
	 select now() 
=====================================================================================
 	 query_stat_date : 0
==============================로그인 성공!!!===================================
회원 azabell님 환영합니다. 
##############################################
로그인 시각 :   2023-08-19 22:54:19

```  

Replace your-username and your-repo-name with your GitHub username and repository name respectively. Make sure to update the MySQL connection details in connect.json before running the program.

Feel free to customize this README to include any additional information or instructions specific to your project.