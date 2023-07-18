-- {
--   "couriers": [
--     {
--       "courier_type": "BIKE",
--       "regions": [
--         2,3
--       ],
--       "working_hours": [
--         "12:00-20:00"
--       ]
--     },
--      {
--       "courier_type": "CAR",
--       "regions": [
--         3,4
--       ],
--       "working_hours": [
--         "20:00-23:00"
--       ]
--     }
--   ]
-- }

INSERT INTO couriers (courier_type, regions, working_hours)
VALUES
    ('BIKE', ARRAY[1, 2], ARRAY[
        TIME '08:00:00', TIME '12:00:00',
     TIME '13:00:00', TIME '17:00:00',
     TIME '18:00:00', TIME '22:00:00'
         ]),
    ('CAR', ARRAY[1, 2, 3], ARRAY[
        TIME '09:00:00', TIME '18:00:00',
     TIME '19:00:00', TIME '22:00:00'
         ]),
    ('FOOT', ARRAY[1], ARRAY[
        TIME '10:00:00', TIME '14:00:00',
     TIME '15:00:00', TIME '19:00:00'
         ]);
