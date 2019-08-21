
 /*drop table sso_city_info;*/

	create table sso_city_info(
		id bigint  primary key auto_increment      comment '' ,
		company_name varchar(32) not null     comment '公司名称' ,
		account varchar(16) not null    comment '帐号' ,
		pwd varchar(8)  not null    comment '密码' 
				
  ) COMMENT='中石化公司账号密码表';

 




