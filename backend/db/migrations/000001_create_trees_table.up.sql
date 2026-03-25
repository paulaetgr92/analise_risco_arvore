
CREATE TABLE trees (
                       id BIGSERIAL PRIMARY KEY ,
                       latitude DOUBLE PRECISION NOT NULL,
                       longitude DOUBLE PRECISION NOT NULL,
                       species TEXT,
                       height DOUBLE PRECISION,
                       diameter DOUBLE PRECISION,
                       age INTEGER,
                       health TEXT,
                       status text,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
