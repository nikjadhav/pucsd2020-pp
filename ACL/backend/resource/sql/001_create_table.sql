USE ACL;
/* DROP TABLE IF EXISTS */
DROP TABLE IF EXISTS groupfilesystem;
DROP TABLE IF EXISTS userfilesystem;
DROP TABLE IF EXISTS filesystem;
DROP TABLE IF EXISTS usergroup;
DROP TABLE IF EXISTS user_detail;
DROP TABLE IF EXISTS filetype;
DROP TABLE IF EXISTS user_rolw;
DROP TABLE IF EXISTS groups_;

/* CREATE TABLES*/

CREATE TABLE IF NOT EXISTS filetype(
	ftid int primary key,
	ftype char(20)   /*ftype can be file or dir*/
);

CREATE TABLE IF NOT EXISTS user_rolw(
	rid int AUTO_INCREMENT primary key,
	rtype char(10) /* rtype can be admin,read/write,read */

);

CREATE TABLE IF NOT EXISTS groups_(
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
    rtype INT NOT NULL DEFAULT 3
);


CREATE TABLE IF NOT EXISTS usergroup(
	id int,
	gid int,
	foreign key (id) references user_detail(id) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key (gid) references groups_(gid)
	ON DELETE CASCADE ON UPDATE CASCADE
	
);

CREATE TABLE IF NOT EXISTS filesystem(
	fid int AUTO_INCREMENT primary key,
	fname varchar(100),
	parent int,
	ftype int,
	unique(fname,parent),
	owner int,
	foreign key(owner) references user_detail(id) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key (parent) references filesystem(fid) ON DELETE CASCADE,
	foreign key (ftype) references filetype(ftid)
	ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS userfilesystem(
	id int ,
	fid int,
	ptype int,
	foreign key(id) references user_detail(id) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key (ptype) references user_rolw(rid) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(fid) references filesystem(fid)
	ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS groupfilesystem(
	gid int,
	fid int,
	ptype int,
	foreign key(gid) references groups_(gid) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(fid) references filesystem(fid) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key (ptype) references user_rolw(rid)
	ON DELETE CASCADE ON UPDATE CASCADE
	
);


/*INSERT VALUES*/
insert into user_rolw values(1,'Admin');
insert into user_rolw values(2,'Write');
insert into user_rolw values(3,'Read');

insert into filetype values(1,"file");
insert into filetype values(2,"dir");

insert into filesystem values(1,'home',NULL,2,NULL);
insert into filesystem values(2,'unknown',1,2,NULL);






