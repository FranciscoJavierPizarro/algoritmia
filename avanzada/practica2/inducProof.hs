{-# LANGUAGE BangPatterns #-}
import Data.Map
import Data.List
import LatinSquare
import Debug.Trace
import System.CPUTime
import System.IO.Unsafe (unsafePerformIO)
import Text.Printf

type Predicate = (Int -> Bool)

implies :: Bool -> Bool -> Bool
implies p q = (not p) || q

basecase :: Predicate -> Bool
basecase p = p 0

jump :: Predicate -> Int -> Bool
jump p n = implies (p n) (p (succ n))

indstep :: Predicate -> Bool
indstep p = forAllInt (jump p)

forAllInt :: Predicate -> Bool
forAllInt p = go 0
  where go n | n == 50 = True
             | otherwise = p n && go (succ n)

isLatinSquareSolved :: [LatinSquare] -> Int -> Bool
isLatinSquareSolved latinSquareSols n = 
    case latinSquareSols of
        [] -> False
        [h] -> all (isLatinSquareSolved' h n) [1..n]
        (h:t) -> all (isLatinSquareSolved' h n) [1..n]
    where isLatinSquareSolved' latinSquare n i = sort [latinSquare ! (i, j) | j <- [1..n]] == [1..n] && sort [latinSquare ! (j, i) | j <- [1..n]] == [1..n]

predicate :: Int -> Bool
predicate n = isLatinSquareSolved (resolver_latinSquare n (latinSquare_create n)) n

predicateDebug :: Int -> Bool
predicateDebug n = let
  (result, elapsedTime) = unsafePerformIO (timeIt $ return (isLatinSquareSolved (resolver_latinSquare n (latinSquare_create n)) n))
  in trace ("Execution time for n = " ++ show n ++ " is " ++ formatTime elapsedTime) result

-- Function to measure execution time
timeIt :: IO a -> IO (a, Integer)
timeIt action = do
  start <- getCPUTime
  result <- action
  end <- getCPUTime
  let !elapsedTime = end - start
  return (result, elapsedTime)

formatTime :: Integer -> String
formatTime picoseconds = printf "%.11f" (fromIntegral picoseconds / 10^12 :: Double)

main :: IO ()
main = do
    if basecase predicate && indstep predicateDebug then putStrLn "El código funciona para cualquier N menor que el límite establecido"
    else putStrLn "El código no funciona para todos los N menores que el límite establecido"