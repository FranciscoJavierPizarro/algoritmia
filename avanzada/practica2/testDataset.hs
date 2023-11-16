--------------------------------------------------------------------------------
--                                                                            --
--    Archivo: testDataset.hs                                                 --
--    Fecha de última revisión: 16/11/2023                                    --
--    Autores: Francisco Javier Pizarro 821259                                --
--             Jorge Solán Morote   	816259                                  --
--    Comms:                                                                  --
--          Este archivo carga latinSquares desde un fichero y los resuelve   --
--          Adicionalmente devuelve el tiempo total empleado                  --
--                                                                            --
--------------------------------------------------------------------------------
import LatinSquare
import Data.List.Split (splitOn)
import System.IO
import System.CPUTime
--ghc testDataset.hs -package minisat-solver -package split -O2
-- -- | The main function.
processPart :: String -> IO ()
processPart part = do
  let s = replaceAsterisks part -- Modify as needed
  print s
  let num = getLatinSquareDimensions s
  putStrLn $ show num
--   num <- getInt
--   let s = latinSquare_create num
  case resolver_latinSquare num s of
    [] -> do
      putStrLn "No solution."
    [h] -> do
      putStrLn "Unique solution:"
      putStr (mostrar_latinSquare h num)
    h:t -> do
      putStrLn "Non-unique solution:"
      putStr (mostrar_latinSquare h num)

filename = "latin_squares.txt"
main :: IO ()
main = do
  contents <- readFile filename
  let contents_parts = splitOn "#" contents
  start <- getCPUTime
  mapM_ processPart contents_parts
  end <- getCPUTime
  let diff = fromIntegral(end - start) / (10^12) :: Float
  putStrLn (show diff)