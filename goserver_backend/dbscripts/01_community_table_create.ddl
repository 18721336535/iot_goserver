USE goweb_db;
CREATE TABLE IF NOT EXISTS community(  
    id INT NOT NULL AUTO_INCREMENT,  
    community_id VARCHAR(24),
    community_name VARCHAR(64),
    introduction VARCHAR(512),
    create_time DATE DEFAULT NOW(),  
    update_time DATE DEFAULT NOW(),  
    PRIMARY KEY ( id )
) engine=InnoDB default charset utf8;

insert into community(community_id,community_name,introduction)values("001","andy","trying new things")