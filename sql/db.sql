
/**
  (c) Code|ng, 2019
  MAINTAINER: Kevin A. Riedl
*/

CREATE TABLE User (
  usr_id VARCHAR(255) PRIMARY KEY, /* public key!! */
  usr_forename VARCHAR(255),
  usr_surname VARCHAR(255),
  usr_email VARCHAR(255),
  usr_facebook VARCHAR(255),
  usr_instagram VARCHAR(255),
  usr_linkedin VARCHAR(255),
  usr_google VARCHAR(255),
  usr_github VARCHAR(255),
  usr_dropbox VARCHAR(255),
  usr_amazon VARCHAR(255)
);