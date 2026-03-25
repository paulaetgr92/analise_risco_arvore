CREATE TABLE weather_events (
                                id BIGSERIAL PRIMARY KEY,
                                region TEXT NOT NULL,
                                wind_speed DOUBLE PRECISION NOT NULL,      -- km/h
                                rainfall_mm DOUBLE PRECISION,              -- volume de chuva
                                severity TEXT NOT NULL CHECK (severity IN ('low','medium','high')),
                                occurred_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
                                created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
