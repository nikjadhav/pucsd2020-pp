USE restapi;
/* DROP TABLE IF EXISTS */
DROP TABLE IF EXISTS groupfilesystem;
DROP TABLE IF EXISTS userfilesystem;
DROP TABLE IF EXISTS filesystem;
DROP TABLE IF EXISTS usergroup;
DROP TABLE IF EXISTS user_detail;
DROP TABLE IF EXISTS filetype;
DROP TABLE IF EXISTS user_rolw;
DROP TABLE IF EXISTS groups;

/* CREATE TABLES*/

CREATE TABLE IF NOT EXISTS filetype(
	ftype char(20) primary key   /*ftype can be file or dir*/
);

CREATE TABLE IF NOT EXISTS user_rolw(
	rid int AUTO_INCREMENT primary key,
	rtype char(10) /* rtype can be admin,read/write,read */

);

CREATE TABLE IF NOT EXISTS groups(
	gid int AUTO_INCREMENT primary key,
	gname varchar(20)
);

CREATE TABLE IF NOT EXISTS user_detail (
    id                  INT         AUTO_INCREMENT      PRIMARY KEY,
    first_name          CHAR(25)    NOT NULL,
    last_name           CHAR(25)    NOT NULL,
    email               CHAR(64)    NOT NULL UNIQUE,
    password            VARBINARY(128)    NOT NULL,
    contact_number      CHAR(15)    NOT NULL,
    updated_by          INT         NOT NULL DEFAULT 0,
    deleted             TINYINT(1)  NOT NULL DEFAULT 0,
    creation_date       DATETIME    DEFAULT CURRENT_TIMESTAMP,
    last_update         DATETIME    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    rtype		int,
    foreign key (rtype) references user_rolw(rid)
);


CREATE TABLE IF NOT EXISTS usergroup(
	id int,
	gid int,
	foreign key (id) references user_detail(id),
	foreign key (gid) references groups(gid),
	primary key(id,gid)
);

CREATE TABLE IF NOT EXISTS filesystem(
	fid int AUTO_INCREMENT primary key,
	fname char(20),
	parent int,
	ftype char(10),
	foreign key (parent) references filesystem(fid),
	foreign key (ftype) references filetype(ftype)

);


CREATE TABLE IF NOT EXISTS userfilesystem(
	id int ,
	fid int,
	ptype int,
	foreign key(id) references user_detail(id),
	foreign key(fid) references filesystem(fid),
	foreign key (ptype) references user_rolw(rid));

CREATE TABLE IF NOT EXISTS groupfilesystem(
	gid int,
	fid int,
	ptype int,
	foreign key(gid) references groups(gid),
	foreign key(fid) references filesystem(fid),
	foreign key (ptype) references user_rolw(rid)
);


