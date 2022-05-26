#!/bin/bash
mysql -u $1 -p -e "
drop database if exists $1_db;
create database $1_db;
use $1_db;
source mysql.sql;
show tables;
quit"