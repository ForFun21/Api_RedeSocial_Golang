insert into usuarios (nome, nick, email, senha)
values
("usuario1", "usuario_1", "usuario1@gmail.com", "$2a$10$p5/aAQsukYGjj3Yag672M.zk.sgurKnn.QnOruHYCUy2o7NEKJBOq"),
("usuario2", "usuario_2", "usuario2@gmail.com", "$2a$10$p5/aAQsukYGjj3Yag672M.zk.sgurKnn.QnOruHYCUy2o7NEKJBOq"),
("usuario3", "usuario_3", "usuario3@gmail.com", "$2a$10$p5/aAQsukYGjj3Yag672M.zk.sgurKnn.QnOruHYCUy2o7NEKJBOq");


insert into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);
