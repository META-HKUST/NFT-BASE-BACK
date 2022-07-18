#!/bin/bash
mysql -ufabric_release_admin -pfabric_release2022 -e "
drop database if exists fabric_release_db;
create database fabric_release_db;
use fabric_release_db;
source mysql.sql;
show tables;
quit"
rm -rf /home/fabric_release/03_End/NFT-BASE-BACK/wallet/*.id

