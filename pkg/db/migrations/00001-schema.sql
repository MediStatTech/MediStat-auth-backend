CREATE TABLE personal (
    personal_id UUID PRIMARY KEY,                 
    first_name VARCHAR(100) NOT NULL,             
    last_name VARCHAR(100) NOT NULL,             
    email VARCHAR(255) NOT NULL UNIQUE,     
    password_hash TEXT NOT NULL,     
    phone VARCHAR(20),                          
    status VARCHAR(50) NOT NULL,                  
    departure VARCHAR(100) NOT NULL,                       
    created_at TIMESTAMPTZ NOT NULL,              
    updated_at TIMESTAMPTZ                        
);

CREATE INDEX idx_personal_email ON personal(email);
