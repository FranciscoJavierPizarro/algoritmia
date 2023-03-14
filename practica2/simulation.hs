{-# LANGUAGE BangPatterns #-}
import BinaryTree
import System.Environment
import System.CPUTime
import Control.Exception
import Text.Printf
import Data.Time.Clock
import Data.Sequence
import Data.Bits
-- compilar ghc -O2 simulation.hs 
-- ejecutar time ./simulation 10000000 10000 [0]

main :: IO ()
main = do
  args <- getArgs
  let n = read (head args) :: Integer
      m = read (args !! 1) :: Integer
      algoritmo = if Prelude.length args == 3 then read (args !! 2) :: Int else -1

  start <- getCurrentTime
  let !resultado = if algoritmo == 0 then simularConRecorridoParcial n m BinaryTree.Empty else simulacionDirecta n m
  end <- getCurrentTime
  let diff = (realToFrac (diffUTCTime end start) :: Double) * (10^3)

  --printf "resultado recorrido:%i\n" resultado
  printf "%0.9f" diff --tiempo en MS