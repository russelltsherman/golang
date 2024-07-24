
create database initdb;
create user inituser with encrypted password 'initpass';
grant all privileges on database initdb to inituser;