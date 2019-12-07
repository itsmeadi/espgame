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
    answered int(2)#no of people answered
);

create table ESPGAME.answers(
      id int(5) primary key NOT NULL AUTO_INCREMENT,
      question_id int(5) references questions.id,
      answer_text varchar(500),
      media_url varchar(500)
      #answered int(2)
);

create table ESPGAME.user_questions(
    id int(5) primary key NOT NULL AUTO_INCREMENT,
    user_id varchar(50) references user.id,
    question_id int(5) references questions.id,
    answer_id int(5) references answers.id,
    correctness int(2) default 0
);