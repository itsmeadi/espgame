#drop database ESPGAME;
create database ESPGAME;

#use database ESPGAME;
create table ESPGAME.user(
    id varchar(50) primary key,
    name varchar(50),
    password varchar(50),
    usertype varchar(10)
);

create table ESPGAME.questions(
    id int(5) primary key NOT NULL AUTO_INCREMENT,
    question_text varchar(500),
    media_url varchar(500),
    answered_by_users int(2) default 0
);

create table ESPGAME.answers(
      id int(5) primary key NOT NULL AUTO_INCREMENT,
      question_id int(5) references questions.id,
      answer_text varchar(500),
      media_url varchar(500)
);

create table ESPGAME.user_questions_answers(
    id int(5) primary key NOT NULL AUTO_INCREMENT,
    user_id varchar(50) references user.id,
    question_id int(5) references questions.id,
    answer_id int(5) references answers.id,
    correctness int(2) default 0,
    UNIQUE `unique_index`(`user_id`, `question_id`, `answer_id`)
);

CREATE INDEX q_idx
    ON ESPGAME.user_questions_answers (question_id);

CREATE INDEX uid_idx
    ON ESPGAME.user_questions_answers (user_id);

create table ESPGAME.group(
    id int(5) primary key NOT NULL AUTO_INCREMENT,
    user_id varchar(50) default 0,
    answered_by_users int(2) default 0
);

insert into ESPGAME.user (id,name,password, usertype) values('admin','admin','21232f297a57a5a743894a0e4a801fc3','admin'); -- insert admin user