
/**
  (c) Code|ng, 2019
  MAINTAINER: Kevin A. Riedl
*/

CREATE TABLE User (
  usr_id VARCHAR(255) PRIMARY KEY,
  usr_forename VARCHAR(255),
  usr_surname VARCHAR(255),
  rsd_id VARCHAR(255),
  FOREIGN KEY (rsd_id) REFERENCES Residence(rsd_id)
);

CREATE TABLE Residence (
  rsd_id VARCHAR(255) PRIMARY KEY,
  rsd_state VARCHAR(255),
  rsd_province VARCHAR(255),
  cnt_id VARCHAR(255),
  rsd_street VARCHAR(255),
  FOREIGN KEY (cnt_id) REFERENCES Country(cnt_id)
);

CREATE TABLE Country (
  cnt_id VARCHAR(255) PRIMARY KEY,
  cnt_country VARCHAR(255),
  cnt_plz VARCHAR(255)
);
