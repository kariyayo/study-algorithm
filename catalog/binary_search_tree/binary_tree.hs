f >: g = g f


data Tree a = Empty | Node a (Tree a) (Tree a) deriving (Show)

insert :: (Ord a) => a -> Tree a -> Tree a
insert x Empty = Node x Empty Empty
insert x (Node a left right)
    | x == a = Node x left right
    | x < a  = Node a (insert x left) right
    | x > a  = Node a left (insert x right)

contains :: (Ord a) => a -> Tree a -> Bool
contains _ Empty = False
contains x (Node a left right)
    | x == a  = True
    | x < a   = contains x left
    | x > a   = contains x right

remove :: (Ord a) => a -> Tree a -> Tree a
remove x (Node a Empty Empty)
    | x == a    = Empty
    | otherwise = Node a Empty Empty
remove x (Node a left Empty)
    | x == a    = left
    | otherwise = Node a (remove x left) Empty
remove x (Node a Empty right)
    | x == a    = right
    | otherwise = Node a Empty (remove x right)
remove x (Node a left right)
    | x == a  = Node maxFromRight left (remove maxFromRight right)
    | x < a   = Node a (remove x left) right
    | x > a   = Node a left (remove x right)
    where maxFromRight = maxValue right

maxValue :: (Ord a) => Tree a -> a
maxValue (Node a Empty Empty) = a
maxValue (Node a left Empty)  = max a (maxValue left)
maxValue (Node a Empty right) = max a (maxValue right)
maxValue (Node a left right)  = maximum [a, (maxValue left), (maxValue right)]

