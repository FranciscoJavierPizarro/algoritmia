{-# LANGUAGE BangPatterns #-}
import BinaryTree
import System.Environment
import System.CPUTime
import Control.Exception
import Text.Printf
import Data.Time.Clock
import Data.Sequence
import Data.Bits
-- compilar ghc Simulation.hs -threaded
-- ejecutar time ./Simulation 10000000 10000 +RTS -s -N12

import Control.Concurrent

main :: IO ()
main = do
  args <- getArgs
  let n = read (head args) :: Integer
      m = read (args !! 1) :: Integer
  
  -- Create an MVar for each function
  mvars <- mapM (\_ -> newEmptyMVar) [1..2]

  forkIO $ do
    start <- getCurrentTime
    let !resultado1 = simularConRecorridoParcial n m BinaryTree.Empty
    end <- getCurrentTime
    let diff1 = (realToFrac (diffUTCTime end start) :: Double) * (10^3)
    putMVar (mvars !! 0) (diff1, resultado1)

  forkIO $ do
    start <- getCurrentTime
    let !resultado2 = simulacionDirecta n m
    end <- getCurrentTime
    let diff2 = (realToFrac (diffUTCTime end start) :: Double) * (10^3)
    putMVar (mvars !! 1) (diff2, resultado2)

  -- Wait for all threads to finish and print the results
  (diff1, resultado1) <- takeMVar (mvars !! 0)
  (diff2, resultado2) <- takeMVar (mvars !! 1)

  printf "resultado recorrido:%i\nresultado directo:%i\n" resultado1 resultado2
  printf "Execution time normal: \t\t%0.9f ms\n" diff1
  printf "Execution time direct(no tree): %0.9f ms\n" diff2

