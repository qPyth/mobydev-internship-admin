INSERT INTO projects(name, category_id,  project_type_id, age_category_id, year, duration, key_words, description, director, producer) values ('film1', 1, 1, 1, 2018, 90, 'key1, key2, key3', 'description', 'director', 'producer');

INSERT INTO categories(name) values ('history');
INSERT INTO categories(name) values ('comedy');

INSERT INTO project_types(type) values ('full-length');
INSERT INTO project_types(type) values ('short');

INSERT INTO age_categories(category) values ('0+');
INSERT INTO age_categories(category) values ('6+');
INSERT INTO age_categories(category) values ('12+');
INSERT INTO age_categories(category) values ('16+');
INSERT INTO age_categories(category) values ('18+');