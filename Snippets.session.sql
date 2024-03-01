












-- WITH RECURSIVE temp as (
--     select *, 1 as level from People
--     where father_id = NULL || mother_id = NULL

--     UNION 

--     SELECT ppl.ID, ppl.Name, ppl.DateBirth, t.level + 1 as level
--     from temp t, People ppl
--     where ppl.father_id = t.ID || ppl.mother_id = t.ID
-- )
-- select max(level) from temp;

--   ('Grandpa1', 'Smith', 'Male', NULL, NULL),
--   ('Grandma1', 'Smith', 'Female', NULL, NULL),
--   ('Dad1', 'Smith', 'Male', 1, 2),
--   ('Mom1', 'Smith', 'Female', 1, 2),
--   ('Son1', 'Smith', 'Male', 3, 4),
--   ('Daughter1', 'Smith', 'Female', 3, 4),
--   ('Grandson1', 'Smith', 'Male', 5, 6),
--   ('Granddaughter1', 'Smith', 'Female', 5, 6),
--   ('Grandpa2', 'Johnson', 'Male', NULL, NULL),
--   ('Grandma2', 'Johnson', 'Female', NULL, NULL),
--   ('Dad2', 'Johnson', 'Male', 9, 10),
--   ('Mom2', 'Johnson', 'Female', 9, 10),
--   ('Son2', 'Johnson', 'Male', 11, 12),
--   ('Daughter2', 'Johnson', 'Female', 11, 12),
--   ('Grandson2', 'Johnson', 'Male', 13, 14),
--   ('Granddaughter2', 'Johnson', 'Female', 13, 14),
--   ('Grandpa3', 'Black', 'Male', NULL, NULL),
--   ('Grandma3', 'Black', 'Female', NULL, NULL),
--   ('Dad3', 'Black', 'Male', 17, 18),
--   ('Mom3', 'Black', 'Female', 17, 18),
--   ('Son3', 'Black', 'Male', 19, 20),
--   ('Daughter3', 'Black', 'Female', 19, 20),
--   ('Grandson3', 'Black', 'Male', 21, 22),
--   ('Granddaughter3', 'Black', 'Female', 21, 22),
--   ('Son1', 'Smith-Johnson', 'Male', 3, 12),
--   ('Daughter3', 'Black-Smith', 'Female', 19, 4);
