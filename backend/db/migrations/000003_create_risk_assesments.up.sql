CREATE TABLE risk_assessments (
                                  id bigserial primary key  ,

                                  tree_id int NOT NULL,

                                  score INTEGER NOT NULL CHECK (score >= 0),

                                  level TEXT NOT NULL CHECK (level IN ('BAIXO', 'MEDIO', 'ALTO')),

                                  calculated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),

                                  CONSTRAINT fk_trees
                                      FOREIGN KEY (tree_id)
                                          REFERENCES trees(id)
                                          ON DELETE CASCADE


);
