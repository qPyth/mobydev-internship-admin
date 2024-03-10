PRAGMA foreign_keys = ON;

CREATE TABLE users (
                      id INTEGER PRIMARY KEY,
                      email TEXT NOT NULL,
                      password TEXT NOT NULL,
                      role TEXT NOT NULL
);

CREATE TABLE projects (
                         id INTEGER PRIMARY KEY,
                         name TEXT NOT NULL,
                         category_id INTEGER NOT NULL ,
                         project_type_id INTEGER NOT NULL ,
                         age_category_id INTEGER NOT NULL ,
                         year TEXT,
                         duration TEXT,
                         key_words TEXT,
                         description TEXT,
                         director TEXT,
                         producer TEXT,
                         FOREIGN KEY(category_id) REFERENCES categories(id),
                         FOREIGN KEY(project_type_id) REFERENCES project_types(id),
                         FOREIGN KEY(age_category_id) REFERENCES age_categories(id)
);

CREATE TABLE categories (
                          id INTEGER PRIMARY KEY,
                          name TEXT NOT NULL
);

CREATE TABLE project_types (
                             id INTEGER PRIMARY KEY,
                             type TEXT NOT NULL
);

CREATE TABLE age_categories (
                             id INTEGER PRIMARY KEY,
                             category TEXT NOT NULL
);
