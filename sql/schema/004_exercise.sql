-- +goose Up
CREATE TABLE exercise (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,  
    description TEXT,            
    duration INT,               
    repetitions INT,             
    sets INT,                    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);


-- +goose Down
DROP TABLE exercise;