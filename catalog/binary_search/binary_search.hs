contains :: Ord a => a -> [a] -> Bool
contains _ [] = False
contains x [y]
    | x == y    = True
    | otherwise = False
contains x ys
    | x == m = True
    | x < m  = contains x (take p ys)
    | x > m  = contains x (drop p ys)
    where p = (length ys) `div` 2
          m = ys !! p

