--------------------------------------------------------------------------------
--                                                                            --
--    Archivo: LatinSquare.hs                                                 --
--    Fecha de última revisión: 16/11/2023                                    --
--    Autores: Francisco Javier Pizarro 821259                                --
--             Jorge Solán Morote   	816259                                  --
--    Comms:                                                                  --
--          Este archivo contiene la implementación del LatinSquare empleando --
--          MiniSat para resolverlo, ofrece funciones para cargarlo desde     --
--          un string, para crearlo sabiendo su N y para resolverlo           --
--                                                                            --
--------------------------------------------------------------------------------
module LatinSquare where
import SAT.MiniSat
-- cabal install minisat-solver
import qualified Data.Map as Map
import Data.Map (Map)
import Data.List
import Data.Char
import System.IO

import Data.List (genericLength)
import Data.Maybe (listToMaybe)
import Data.Function (on)
import Data.Ord (comparing)
-- ghc latinSquare.hs -package minisat-solver -O2 && ./latinSquare

data Celda = Celda Int Int Int
            deriving (Eq, Ord, Show)

type LatinSquare = Map (Int, Int) Int

getLatinSquareDimensions :: LatinSquare -> Int
getLatinSquareDimensions square = maximum [maxRow, maxCol, maxVal]
  where
    rows = map fst (Map.keys square)
    cols = map snd (Map.keys square)
    values = Map.elems square
    maxRow = maximum rows
    maxCol = maximum cols
    maxVal = maximum values


celda :: Int -> Int -> Int -> Formula Celda
celda i j n = Var (Celda i j n)

reglas :: Int -> Formula Celda
reglas m = celdas :&&: rows :&&: columns
  where
    celdas = All [ ExactlyOne [ celda i j n | n <- [1..m] ] | i <- [1..m], j <- [1..m] ]
    rows = All [ ExactlyOne [ celda i j n | j <- [1..m] ] | i <- [1..m], n <- [1..m] ]
    columns = All [ ExactlyOne [ celda i j n | i <- [1..m] ] | j <- [1..m], n <- [1..m] ]

formula_del_latinSquare :: LatinSquare -> Formula Celda
formula_del_latinSquare s = All [ celda i j n | ((i,j),n) <- Map.toList s ]

latinSquare_de_la_solucion :: Map Celda Bool -> LatinSquare
latinSquare_de_la_solucion m = Map.fromList [ ((i,j),n) | (Celda i j n, True) <- Map.toList m ] 

resolver_latinSquare :: Int -> LatinSquare -> [LatinSquare]
resolver_latinSquare n s = map latinSquare_de_la_solucion (solve_all ((reglas n):&&: formula_del_latinSquare s))

mostrar_latinSquare :: LatinSquare -> Int -> String
mostrar_latinSquare s n =
  concat [entry i j n | i <- [1..n], j <- [1..n]]
  where
    entry i j nFila = case Map.lookup (i,j) s of
      Nothing -> "$"
      Just n -> (if (j == 1) then "| " else "") ++ show n ++ " " ++ (if ((mod j nFila) == 0) then "|\n" else "")
    
latinSquare_create :: Int -> LatinSquare
latinSquare_create n = Map.fromList [ ((i,j),0) | i <- [0..(n-1)], j <- [0..(n-1)] ]

latinSquare_de_lista :: [Int] -> LatinSquare
latinSquare_de_lista xs = Map.fromList [ ((i, j), n) | ((i, j), n) <- zip coords xs, n >= 1, n <= len ]
  where
    len = floor (sqrt (fromIntegral (length xs)))
    coords = [(i, j) | i <- [1..len], j <- [1..len]]

replaceAsterisks :: String -> LatinSquare
replaceAsterisks input = latinSquare_de_lista $ map replaceChar (filter (not . isSpace) input)
  where
    replaceChar c
      | isDigit c = digitToInt c 
      | c == '*'  = 0
      | otherwise = error "Invalid character in the input"

getInt :: IO Int
getInt = do
    input <- getLine
    return (read input)