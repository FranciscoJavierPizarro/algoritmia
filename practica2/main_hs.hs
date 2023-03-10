{-# LANGUAGE BangPatterns #-}
import BinaryTree
import System.Environment
import System.CPUTime
import Control.Exception
import Text.Printf
import Data.Time.Clock
import Data.Sequence
import Data.Bits
-- compilar ghc main_hs.hs -threaded
-- ejecutar time ./main_hs 10000000 10000 +RTS -s -N12

-- main :: IO ()
-- main = do
--   let binarySeq = toBinarySeq 40
--   print (binarySeq)
--   args <- getArgs
--   let n = read (head args) :: Integer
--       m = read (args !! 1) :: Integer

--   start <- getCurrentTime
--   let !resultado = simularConRecorridoParcial n m Empty
--   end <- getCurrentTime
--   let diff1 = (realToFrac (diffUTCTime end start) :: Double) * (10^3)

--   start <- getCurrentTime
--   let !resultado1 = simulacionDirecta n m
--   end <- getCurrentTime
--   let diff2 = (realToFrac (diffUTCTime end start) :: Double) * (10^3)

--   start <- getCurrentTime
--   let !resultado2 = simulacionOptimizada n m Empty
--   end <- getCurrentTime
--   let diff3 = (realToFrac (diffUTCTime end start) :: Double) * (10^3)


--   print (resultado2 - resultado)
--   printf "Execution time normal: \t\t%0.9f ms\n" diff1
--   printf "Execution time direct(no tree): %0.9f ms\n" diff2
--   printf "Execution time optimized:\t %0.9f ms\n" diff3

import Control.Concurrent

main :: IO ()
main = do
  args <- getArgs
  let n = read (head args) :: Integer
      m = read (args !! 1) :: Integer
  
  -- Create an MVar for each function
  mvars <- mapM (\_ -> newEmptyMVar) [1..3]

  -- Spawn a new thread for each function and store the result in the corresponding MVar
  -- forkIO $ do
  --   start <- getCurrentTime
  --   let !resultado = simulacionOptimizada n m Empty
  --   end <- getCurrentTime
  --   let diff1 = (realToFrac (diffUTCTime end start) :: Double) * (10^3)
  --   putMVar (mvars !! 0) (diff1, resultado)

  forkIO $ do
    start <- getCurrentTime
    let !resultado1 = simularConRecorridoParcial n m BinaryTree.Empty
    end <- getCurrentTime
    let diff2 = (realToFrac (diffUTCTime end start) :: Double) * (10^3)
    putMVar (mvars !! 1) (diff2, resultado1)

  forkIO $ do
    start <- getCurrentTime
    let !resultado2 = simulacionDirecta n m
    end <- getCurrentTime
    let diff3 = (realToFrac (diffUTCTime end start) :: Double) * (10^3)
    putMVar (mvars !! 2) (diff3, resultado2)

  -- Wait for all threads to finish and print the results
  -- (diff1, resultado) <- takeMVar (mvars !! 0)
  (diff2, resultado1) <- takeMVar (mvars !! 1)
  (diff3, resultado2) <- takeMVar (mvars !! 2)

  print (resultado1 - resultado2)
  -- printf "Execution time optimized: \t%0.9f ms\n" diff1
  printf "Execution time normal: \t\t%0.9f ms\n" diff2
  printf "Execution time direct(no tree): %0.9f ms\n" diff3

